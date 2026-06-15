package ui_test

import (
	"context"
	"io"
	"strings"
	"testing"
	"time"

	"pomodoro/internal/timer"
	"pomodoro/internal/ui"
)

func TestWriteIdleMessage(t *testing.T) {
	var buf strings.Builder
	ui.WriteIdleMessage(&buf)
	got := buf.String()
	want := "No session in progress. Press Enter or Space to start."
	if !strings.Contains(got, want) {
		t.Errorf("WriteIdleMessage output %q does not contain %q", got, want)
	}
}

func TestWriteCountdown(t *testing.T) {
	cases := []struct {
		label     string
		remaining time.Duration
		want      string
	}{
		{"Pomodoro 1/4", 25 * time.Minute, "\r[Pomodoro 1/4] 25:00 remaining"},
		{"Pomodoro 1/4", 25*time.Minute - time.Second, "\r[Pomodoro 1/4] 24:59 remaining"},
		{"Short break", 0, "\r[Short break] 00:00 remaining"},
		{"Long break", 90 * time.Second, "\r[Long break] 01:30 remaining"},
	}
	for _, c := range cases {
		var buf strings.Builder
		ui.WriteCountdown(&buf, c.label, c.remaining)
		if buf.String() != c.want {
			t.Errorf("WriteCountdown(%q, %v) = %q, want %q", c.label, c.remaining, buf.String(), c.want)
		}
	}
}

func TestWriteBell(t *testing.T) {
	var buf strings.Builder
	ui.WriteBell(&buf)
	if buf.String() != "\a" {
		t.Errorf("WriteBell() = %q, want %q", buf.String(), "\a")
	}
}

func TestWriteSessionComplete(t *testing.T) {
	var buf strings.Builder
	ui.WriteSessionComplete(&buf)
	if !strings.Contains(buf.String(), "Session complete") {
		t.Errorf("WriteSessionComplete() = %q, want it to contain %q", buf.String(), "Session complete")
	}
}

// TestRunSession_FullCycle drives the entire user-visible flow:
// idle → Space key → countdown → bell → "Session complete" → idle.
func TestRunSession_FullCycle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use an io.Pipe so we can feed a single key byte then leave the reader open.
	r, w := io.Pipe()
	defer r.Close()

	// Send Space to trigger session start.
	go func() { w.Write([]byte{' '}) }()

	// 2-second fast-clock session.
	const dur = 2 * time.Second
	ticks := make(chan time.Time, int(dur/time.Second))
	for range int(dur / time.Second) {
		ticks <- time.Time{}
	}
	sess := timer.Session{
		Duration: dur,
		NewTicker: func(d time.Duration) (<-chan time.Time, func()) {
			return ticks, func() {}
		},
	}

	var out strings.Builder
	ui.RunSession(ctx, r, &out, "Pomodoro 1/4", sess)

	output := out.String()

	// AC1: idle message shown before countdown
	idleIdx := strings.Index(output, "No session in progress")
	if idleIdx < 0 {
		t.Fatalf("idle message not found in output: %q", output)
	}

	// AC2: countdown lines present with label
	if !strings.Contains(output, "\r[Pomodoro 1/4] 00:02 remaining") {
		t.Errorf("expected countdown start in output: %q", output)
	}
	if !strings.Contains(output, "\r[Pomodoro 1/4] 00:00 remaining") {
		t.Errorf("expected countdown end in output: %q", output)
	}

	// AC3: exactly one bell
	bells := strings.Count(output, "\a")
	if bells != 1 {
		t.Errorf("expected exactly 1 bell, got %d in output: %q", bells, output)
	}

	// AC4: session complete message followed by idle message again
	completeIdx := strings.Index(output, "Session complete")
	if completeIdx < 0 {
		t.Fatalf("'Session complete' not found in output: %q", output)
	}
	finalIdleIdx := strings.LastIndex(output, "No session in progress")
	if finalIdleIdx <= idleIdx {
		t.Errorf("expected idle message to reappear after session complete")
	}
}

// makeFastInterval returns an Interval of the given kind/number backed by a
// fast-clock session of dur seconds.
func makeFastInterval(kind timer.IntervalKind, number int, dur time.Duration) timer.Interval {
	ticks := make(chan time.Time, int(dur/time.Second))
	for range int(dur / time.Second) {
		ticks <- time.Time{}
	}
	return timer.Interval{
		Kind:   kind,
		Number: number,
		Session: timer.Session{
			Duration: dur,
			NewTicker: func(d time.Duration) (<-chan time.Time, func()) {
				return ticks, func() {}
			},
		},
	}
}

// TestRunCycle_FullCycle drives the full 8-interval Pomodoro cycle:
// idle → Space key → all intervals auto-advance → idle.
func TestRunCycle_FullCycle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, w := io.Pipe()
	defer r.Close()

	// Single Space to start the cycle; no further input needed.
	go func() { w.Write([]byte{' '}) }()

	const dur = 2 * time.Second

	// Build 8 fast-clock intervals matching the standard cycle.
	intervals := []timer.Interval{
		makeFastInterval(timer.Pomodoro, 1, dur),
		makeFastInterval(timer.ShortBreak, 0, dur),
		makeFastInterval(timer.Pomodoro, 2, dur),
		makeFastInterval(timer.ShortBreak, 0, dur),
		makeFastInterval(timer.Pomodoro, 3, dur),
		makeFastInterval(timer.ShortBreak, 0, dur),
		makeFastInterval(timer.Pomodoro, 4, dur),
		makeFastInterval(timer.LongBreak, 0, dur),
	}

	var out strings.Builder
	ui.RunCycle(ctx, r, &out, intervals)

	output := out.String()

	// Idle message at start.
	firstIdleIdx := strings.Index(output, "No session in progress")
	if firstIdleIdx < 0 {
		t.Fatalf("idle message not found in output: %q", output)
	}

	// All four Pomodoro labels present.
	for _, label := range []string{
		"[Pomodoro 1/4]",
		"[Pomodoro 2/4]",
		"[Pomodoro 3/4]",
		"[Pomodoro 4/4]",
	} {
		if !strings.Contains(output, label) {
			t.Errorf("expected label %q in output", label)
		}
	}

	// Break labels present.
	if !strings.Contains(output, "[Short break]") {
		t.Errorf("expected [Short break] in output: %q", output)
	}
	if !strings.Contains(output, "[Long break]") {
		t.Errorf("expected [Long break] in output: %q", output)
	}

	// Exactly 8 bells — one per interval end.
	bells := strings.Count(output, "\a")
	if bells != 8 {
		t.Errorf("expected 8 bells, got %d", bells)
	}

	// Idle message at end (after the last interval).
	finalIdleIdx := strings.LastIndex(output, "No session in progress")
	if finalIdleIdx <= firstIdleIdx {
		t.Errorf("expected idle message to reappear after cycle completes")
	}
}
