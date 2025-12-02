# Design: Add macOS .app bundle via Makefile
**Date:** 2025-12-02  
**Status:** Initial Technical Design

## Design Summary
Introduce a `Makefile`-driven bundle step that creates a macOS `.app` for the Pomodoro tray app. The build will compile the Go binary, stage a standard macOS bundle layout, generate `Info.plist` with `LSUIElement=1` (menu-bar app), and include an `AppIcon.icns` generated during the build. This aligns with the constitution’s Small, Safe Iterations and Dependency Discipline: build logic stays outside app code and behind simple scripts.

## Initial Approach
- **Make targets (bundle/clean/version)**
  - **Approach:** Add `make bundle`, `make clean`, and a helper `make version`. Bundle compiles the binary and assembles `dist/Pomodoro.app`.
  - **Rationale:** Backgroundable build steps (1C) and clear developer ergonomics.
  - **Trade-offs:** Requires macOS/Xcode CLT for `iconutil`/`plutil`.
  - **Alternatives:** Pre-commit generated assets; or use Go-based packager. Prefer simple shell-first.
- **Info.plist with LSUIElement**
  - **Approach:** Template minimal plist (Name/Identifier/Version/LSUIElement) and fill version from `VERSION` or `git describe`.
  - **Rationale:** Meets acceptance; hides Dock; keeps logic out of core code (2B).
  - **Trade-offs:** Requires `plutil` or envsubst; simple sed is acceptable.
  - **Alternatives:** Generate via small Go tool, but that couples build and app.
- **Icon pipeline (.icns)**
  - **Approach:** Generate a simple monochrome glyph PNG and convert to `.iconset` → `.icns` using `iconutil`. If tools missing, fall back to placeholder.
  - **Rationale:** High-contrast template icon matches menu-bar expectations; background processing friendly (1C).
  - **Trade-offs:** Tooling dependency; placeholder if unavailable.
  - **Alternatives:** Commit a static `AppIcon.icns` asset.
- **Platform gating**
  - **Approach:** Guard `make bundle` with `uname` check; macOS prints friendly no-op message elsewhere.
  - **Rationale:** Satisfies cross-platform no-op criterion.
  - **Trade-offs:** Extra conditionals in Makefile.
  - **Alternatives:** Separate platform-specific Makefiles.

## Architecture Overview
- **Components:**
  - `Makefile`: orchestration (bundle/version/clean).
  - `dist/` bundle: `Pomodoro.app/Contents/{MacOS,Resources,Info.plist}`.
  - Build helpers (shell snippets) to produce `AppIcon.icns` and plist.
- **Data Flow:**
  - Source → `go build ./app` → stage `dist/Pomodoro.app/Contents/MacOS/pomodoro` → write `Info.plist` (inject version) → create `Resources/AppIcon.icns` → `open dist/Pomodoro.app` for manual check.
- **Integration Points:**
  - macOS tools: `iconutil`, `plutil`, `defaults`/`sed`; Go toolchain.
  - Existing tray adapter remains unchanged (3B).
- **State Management:**
  - No runtime state changes; build artifacts live under `examples/pomodoro/dist/`.

## Implementation Constraints
- Runs only on macOS; requires Xcode Command Line Tools for icon/plist utilities.
- Output path fixed: `examples/pomodoro/dist/Pomodoro.app`.
- Bundle identifier stable: `co0p.pomodoro`.

## Open Questions
- Universal binary support now or later? (arm64 + amd64 via `lipo`).
- Should we commit a static `.icns` to avoid `iconutil` dependency?
- Exact versioning scheme (strict `VERSION` file vs. `git describe` fallback)?

## Save Location
`examples/pomodoro/docs/increments/macos-app-bundle/design.md`
