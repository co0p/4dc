package ui

import (
	"context"
	"fmt"
	"io"
	"os"
	"syscall"
	"time"

	"pomodoro/internal/timer"
)

// WriteIdleMessage writes the idle screen line to w.
func WriteIdleMessage(w io.Writer) {
	fmt.Fprintln(w, "No session in progress. Press Enter or Space to start.")
}

// WriteCountdown writes the remaining duration as \r[label] MM:SS remaining,
// overwriting the current terminal line.
func WriteCountdown(w io.Writer, label string, remaining time.Duration) {
	if remaining < 0 {
		remaining = 0
	}
	total := int(remaining.Seconds())
	mins := total / 60
	secs := total % 60
	fmt.Fprintf(w, "\r[%s] %02d:%02d remaining", label, mins, secs)
}

// WriteBell writes a single ASCII bell character (0x07) to w.
func WriteBell(w io.Writer) {
	fmt.Fprint(w, "\a")
}

// WriteSessionComplete writes a session-complete message to w.
func WriteSessionComplete(w io.Writer) {
	fmt.Fprintln(w, "\nSession complete!")
}

// intervalLabel returns the display label for an Interval.
func intervalLabel(iv timer.Interval) string {
	switch iv.Kind {
	case timer.Pomodoro:
		return fmt.Sprintf("Pomodoro %d/4", iv.Number)
	case timer.ShortBreak:
		return "Short break"
	default:
		return "Long break"
	}
}

// RunCycle shows the idle screen, waits for Enter or Space on in, then
// automatically runs all intervals in sequence — writing labeled countdown
// ticks to out and ringing the bell at the end of each interval — before
// returning to the idle screen.
// It returns when all intervals complete or ctx is cancelled.
func RunCycle(ctx context.Context, in io.Reader, out io.Writer, intervals []timer.Interval) {
	WriteIdleMessage(out)

	startCh := make(chan struct{}, 1)

	go func() {
		buf := make([]byte, 1)
		for {
			n, err := in.Read(buf)
			if err != nil || n == 0 {
				return
			}
			switch buf[0] {
			case '\n', '\r', ' ':
				select {
				case startCh <- struct{}{}:
				default:
				}
			case 0x03, 0x04:
				syscall.Kill(os.Getpid(), syscall.SIGINT)
				return
			}
		}
	}()

	select {
	case <-startCh:
	case <-ctx.Done():
		return
	}

	for _, iv := range intervals {
		label := intervalLabel(iv)
		for remaining := range iv.Session.Run(ctx) {
			WriteCountdown(out, label, remaining)
		}
		if ctx.Err() != nil {
			return
		}
		WriteBell(out)
	}

	WriteIdleMessage(out)
}

// to out. label is shown on every tick in the format [label] MM:SS remaining.
// It returns when sess completes or ctx is cancelled.
//
// A single goroutine owns all reads from in for the lifetime of RunSession.
// Byte 0x03 (Ctrl+C swallowed by raw mode) re-raises SIGINT so the process
// exits cleanly even when ISIG is disabled.
func RunSession(ctx context.Context, in io.Reader, out io.Writer, label string, sess timer.Session) {
	WriteIdleMessage(out)

	startCh := make(chan struct{}, 1)

	go func() {
		buf := make([]byte, 1)
		for {
			n, err := in.Read(buf)
			if err != nil || n == 0 {
				return
			}
			switch buf[0] {
			case '\n', '\r', ' ':
				select {
				case startCh <- struct{}{}:
				default:
				}
			case 0x03, 0x04: // Ctrl+C / Ctrl+D: re-raise SIGINT (raw mode suppressed it)
				syscall.Kill(os.Getpid(), syscall.SIGINT)
				return
			}
		}
	}()

	select {
	case <-startCh:
	case <-ctx.Done():
		return
	}

	for remaining := range sess.Run(ctx) {
		WriteCountdown(out, label, remaining)
	}

	if ctx.Err() != nil {
		return
	}

	WriteBell(out)
	WriteSessionComplete(out)
	WriteIdleMessage(out)
}
