package tray

import (
	"context"

	"github.com/co0p/4dc/examples/pomodoro/internal/app"
)

// MockTray is a simple in-process mock that calls App methods directly.
type MockTray struct {
	App     app.App
	started bool
}

// NewMockTray constructs a new in-process mock tray that calls the App
// methods directly. Useful for tests that exercise UI wiring without a
// real OS tray.
func NewMockTray(a app.App) *MockTray { return &MockTray{App: a} }

// Run starts the mock tray and blocks until the context is cancelled.
func (m *MockTray) Run(ctx context.Context) error {
	m.started = true
	<-ctx.Done()
	m.started = false
	return ctx.Err()
}

// Close requests the mock tray to stop. It is idempotent.
func (m *MockTray) Close() error {
	m.started = false
	return nil
}

// Trigger simulates a user clicking a menu item by name.
// Supported names: "Pomodoro", "Break", "Quit".
func (m *MockTray) Trigger(name string) {
	switch name {
	case "Pomodoro":
		m.App.StartPomodoro()
	case "Break":
		m.App.StartBreak()
	case "Quit":
		_ = m.App.Shutdown(context.Background())
	}
}
