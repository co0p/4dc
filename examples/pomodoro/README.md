# Pomodoro (Tray App)

A lightweight Pomodoro timer that lives in the macOS menu bar.

## Build & Run

- Build the bundled app (macOS only):
	```zsh
	make bundle
	```
- Launch the app:
	```zsh
	make open
	```

This produces `dist/Pomodoro.app` and opens it. The app is a menu‑bar app (no Dock icon) with an icon‑only glyph. Menu entries: Start Sample, Pause Sample, Quit.

## Versioning

- The bundle version is sourced from `VERSION` (fallback to `git describe`).
- Update `examples/pomodoro/VERSION` to bump the app version.

## Icon

- The runtime tray icon is a high‑contrast template glyph.
- The app bundle icon (`AppIcon.icns`) can be generated procedurally:
	```zsh
	make icon
	```
	This creates PNGs at required sizes and builds `assets/AppIcon.icns` via `iconutil`.

## Clean

- Remove build artifacts:
	```zsh
	make clean
	```

## Notes

- macOS only tasks (`icon`, `bundle`, `open`) no‑op on other platforms.
- Requires Go toolchain and Xcode Command Line Tools for `iconutil`.