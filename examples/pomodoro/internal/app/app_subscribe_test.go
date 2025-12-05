package app

import (
	"testing"
	"time"
)

func TestSubscribeStateChange_MultipleAndUnsubscribe(t *testing.T) {
	// use short durations for test
	a := New(50*time.Millisecond, 20*time.Millisecond)

	ch1 := make(chan State, 4)
	ch2 := make(chan State, 4)

	unsub1 := a.SubscribeStateChange(func(s State) { ch1 <- s })
	_ = a.SubscribeStateChange(func(s State) { ch2 <- s })

	// Start a pomodoro: both should receive the running state
	a.StartPomodoro()

	select {
	case s := <-ch1:
		if s != StatePomodoroRunning {
			t.Fatalf("unexpected state for ch1: %v", s)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("timeout waiting for ch1 notification")
	}
	select {
	case s := <-ch2:
		if s != StatePomodoroRunning {
			t.Fatalf("unexpected state for ch2: %v", s)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("timeout waiting for ch2 notification")
	}

	// Unsubscribe first listener
	unsub1()

	// Start a break: only ch2 should receive it
	a.StartBreak()

	select {
	case s := <-ch2:
		if s != StateBreakRunning {
			t.Fatalf("unexpected state for ch2 after break: %v", s)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("timeout waiting for ch2 after break")
	}

	select {
	case s := <-ch1:
		t.Fatalf("did not expect ch1 to receive after unsubscribe, got %v", s)
	case <-time.After(30 * time.Millisecond):
		// expected: no message
	}
}
