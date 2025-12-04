package app

import (
	"context"
	"testing"
	"time"
)

func TestPomodoroTransitionAndIdle(t *testing.T) {
	// very short durations to keep tests fast
	a := New(10*time.Millisecond, 5*time.Millisecond)

	stateCh := make(chan State, 4)
	a.OnStateChange(func(s State) { stateCh <- s })

	a.StartPomodoro()

	// expect PomodoroRunning then Idle shortly after
	var seenPomodoro, seenIdle bool
	timeout := time.After(200 * time.Millisecond)
	for !seenIdle {
		select {
		case s := <-stateCh:
			if s == StatePomodoroRunning {
				seenPomodoro = true
			}
			if s == StateIdle && seenPomodoro {
				seenIdle = true
			}
		case <-timeout:
			t.Fatalf("timeout waiting for state transitions: seenPomodoro=%v", seenPomodoro)
		}
	}
}

func TestIdempotentStarts(t *testing.T) {
	a := New(30*time.Millisecond, 10*time.Millisecond)
	a.StartPomodoro()
	// second start should be no-op and not panic
	a.StartPomodoro()
	// then start break while pomodoro running should cancel previous timer and switch
	a.StartBreak()
	if a.State() != StateBreakRunning {
		t.Fatalf("expected break running, got %s", a.State())
	}
}

func TestShutdownCancelsTimer(t *testing.T) {
	a := New(1*time.Second, 1*time.Second)
	a.StartPomodoro()

	// ensure we're in running state
	if a.State() != StatePomodoroRunning {
		t.Fatalf("expected pomodoro running")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := a.Shutdown(ctx); err != nil {
		t.Fatalf("shutdown failed: %v", err)
	}
	if a.State() != StateIdle {
		t.Fatalf("expected idle after shutdown, got %s", a.State())
	}
}
