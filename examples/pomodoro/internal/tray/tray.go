package tray

import (
	"context"

	"github.com/co0p/4dc/examples/pomodoro/internal/app"
)

// Tray abstracts a platform-specific tray/menu implementation.
type Tray interface {
	// Run starts the tray UI and blocks until it exits or context is canceled.
	Run(ctx context.Context) error
	// Close requests the tray to shut down.
	Close() error
}

// NewSystray is implemented in systray_impl.go and returns a Tray backed by a systray package.
func NewSystray(a app.App, icon []byte) Tray {
	return newSystrayImpl(a, icon)
}
