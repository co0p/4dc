// Package tray contains the platform-specific tray (systray) wiring used by
// the pomodoro demo. The package provides a `TitleUpdater` responsible for
// subscribing to app state changes and updating the tray title with a
// concise remaining-time label.
//
// Systray threading note:
// The underlying `systray` library requires `systray.Run` to be called on
// the main OS thread on macOS. The demo calls `systray.Run` from `main()` to
// satisfy this constraint. Methods that call into OS UI APIs (for example,
// `systray.SetTitle`) should be reviewed for thread-safety if you move them
// to other goroutines or refactor the wiring.
//
// See `../../docs/ADR-2025-12-05-receiver-naming-and-docs.md` for documentation
// and receiver-naming conventions applied in this example.
package tray
