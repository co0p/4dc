package timer

import (
	"context"
	"time"
)

// Session holds the configuration for a single countdown interval.
// NewTicker is a pluggable clock: in production it wraps time.NewTicker;
// in tests it returns a pre-filled channel for fast, deterministic runs.
type Session struct {
	Duration  time.Duration
	NewTicker func(d time.Duration) (<-chan time.Time, func())
}

// NewSession returns a Session configured for wall-clock use.
func NewSession(d time.Duration) Session {
	return Session{
		Duration: d,
		NewTicker: func(d time.Duration) (<-chan time.Time, func()) {
			t := time.NewTicker(d)
			return t.C, t.Stop
		},
	}
}

// Run starts the countdown and emits remaining time after each tick, starting
// with the full duration before the first tick. The returned channel is closed
// when the countdown reaches zero or ctx is cancelled.
func (s Session) Run(ctx context.Context) <-chan time.Duration {
	out := make(chan time.Duration)
	go func() {
		defer close(out)

		remaining := s.Duration

		// Emit the initial (full) remaining time before the first tick.
		select {
		case out <- remaining:
		case <-ctx.Done():
			return
		}

		tick, stop := s.NewTicker(time.Second)
		defer stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick:
				remaining -= time.Second
				select {
				case out <- remaining:
				case <-ctx.Done():
					return
				}
				if remaining <= 0 {
					return
				}
			}
		}
	}()
	return out
}
