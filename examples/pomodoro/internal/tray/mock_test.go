package tray

import (
	"context"
	"testing"
	"time"

	"github.com/co0p/4dc/examples/pomodoro/internal/app"
)

func TestMockTrayTriggersAppActions(t *testing.T) {
	a := app.New(50*time.Millisecond, 20*time.Millisecond)
	mt := NewMockTray(a)

	// Start pomodoro via mock trigger
	mt.Trigger("Pomodoro")
	if a.State() != app.StatePomodoroRunning {
		t.Fatalf("expected pomodoro running, got %s", a.State())
	}

	// Trigger break; should switch state
	mt.Trigger("Break")
	if a.State() != app.StateBreakRunning {
		t.Fatalf("expected break running, got %s", a.State())
	}

	// Trigger quit; Shutdown should set idle
	mt.Trigger("Quit")
	// allow a tiny moment for shutdown hooks
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(10 * time.Millisecond):
	case <-ctx.Done():
		t.Fatalf("shutdown did not complete in time")
	}
	if a.State() != app.StateIdle {
		t.Fatalf("expected idle after quit, got %s", a.State())
	}
}
