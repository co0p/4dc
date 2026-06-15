package timer_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"pomodoro/internal/timer"
)

// TestSessionRun_Countdown verifies that Run emits the correct sequence of
// remaining durations using a fast-clock shim (no real time.Sleep).
func TestSessionRun_Countdown(t *testing.T) {
	const dur = 3 * time.Second

	// Pre-fill a buffered tick channel with exactly dur/time.Second ticks.
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var got []time.Duration
	for remaining := range sess.Run(ctx) {
		got = append(got, remaining)
	}

	want := []time.Duration{
		3 * time.Second,
		2 * time.Second,
		1 * time.Second,
		0,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Run() emitted %v, want %v", got, want)
	}
}

// TestNewCycle verifies the 8-interval standard Pomodoro sequence produced by
// NewCycle: correct kind, number, and duration at each position.
func TestNewCycle(t *testing.T) {
	work := 25 * time.Minute
	short := 5 * time.Minute
	long := 15 * time.Minute

	intervals := timer.NewCycle(work, short, long)

	if len(intervals) != 8 {
		t.Fatalf("NewCycle returned %d intervals, want 8", len(intervals))
	}

	want := []struct {
		kind   timer.IntervalKind
		number int
		dur    time.Duration
	}{
		{timer.Pomodoro, 1, work},
		{timer.ShortBreak, 0, short},
		{timer.Pomodoro, 2, work},
		{timer.ShortBreak, 0, short},
		{timer.Pomodoro, 3, work},
		{timer.ShortBreak, 0, short},
		{timer.Pomodoro, 4, work},
		{timer.LongBreak, 0, long},
	}

	for i, w := range want {
		got := intervals[i]
		if got.Kind != w.kind {
			t.Errorf("interval[%d].Kind = %v, want %v", i, got.Kind, w.kind)
		}
		if got.Number != w.number {
			t.Errorf("interval[%d].Number = %d, want %d", i, got.Number, w.number)
		}
		if got.Session.Duration != w.dur {
			t.Errorf("interval[%d].Session.Duration = %v, want %v", i, got.Session.Duration, w.dur)
		}
	}
}

// TestSessionRun_CancelledByContext verifies that Run closes its channel when
// the context is cancelled before the countdown finishes.
func TestSessionRun_CancelledByContext(t *testing.T) {
	ticks := make(chan time.Time) // never ticks

	sess := timer.Session{
		Duration: 10 * time.Second,
		NewTicker: func(d time.Duration) (<-chan time.Time, func()) {
			return ticks, func() {}
		},
	}

	ctx, cancel := context.WithCancel(context.Background())

	ch := sess.Run(ctx)

	// Drain initial emission then cancel.
	<-ch     // should receive 10s immediately
	cancel() // cancel before any tick

	var count int
	for range ch {
		count++
	}
	// Channel must close after cancellation; we may receive 0 or more extra
	// ticks that raced but must not block forever.
	_ = count
}
