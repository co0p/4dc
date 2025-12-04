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

func NewMockTray(a app.App) *MockTray { return &MockTray{App: a} }

func (m *MockTray) Run(ctx context.Context) error {
	m.started = true
	<-ctx.Done()
	m.started = false
	return ctx.Err()
}

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
