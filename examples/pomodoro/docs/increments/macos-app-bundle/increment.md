# Add macOS .app bundle via Makefile

## Job Story
**When** I use the app on macOS,
**I want to** build a native `.app` bundle with a single command,
**So I can** launch it from Finder like any other menu bar app without using the terminal.

**Assumption Being Tested:** Bundling as a macOS `.app` makes the systray app easier to install, launch, and adopt.

## Acceptance Criteria
- **Given** a macOS host
  **When** I run `make bundle`
  **Then** a bundle exists at `examples/pomodoro/dist/Pomodoro.app` containing:
  - `Contents/MacOS/pomodoro` (binary built from `./app`)
  - `Contents/Info.plist` including keys:
    - `CFBundleName`: `Pomodoro`
    - `CFBundleIdentifier`: `co0p.pomodoro`
    - `CFBundleVersion`: from `VERSION` file or `git describe --tags --always`
    - `LSUIElement`: `1` (hide Dock icon; menu-bar app)
  - `Contents/Resources/AppIcon.icns` present (placeholder or converted from existing glyph)

- **Given** the bundle exists
  **When** I run `open examples/pomodoro/dist/Pomodoro.app`
  **Then** the tray icon appears (icon-only) with the items “Start Sample”, “Pause Sample”, and “Quit”, and no Dock icon is shown.

- **Given** a non‑macOS host
  **When** I run `make bundle`
  **Then** the task prints a clear message (e.g., "macOS only") and exits without creating artifacts.

- **Given** bundle artifacts exist
  **When** I run `make clean`
  **Then** the `examples/pomodoro/dist/` directory is removed and the source tree remains untouched.

## Success Signal
A teammate on macOS runs `make bundle` and launches `Pomodoro.app` from Finder. The menu bar icon appears within 5 seconds, the menu contains the expected items, and no terminal is required.

## Out of Scope
- Code signing or notarization.
- DMG packaging or auto‑updates.
- CI/CD integration for release artifacts.
