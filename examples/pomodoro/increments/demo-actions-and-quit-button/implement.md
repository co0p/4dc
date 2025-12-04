# Implementation Plan: Demo Tray Actions, Quit Button, and Minimal Icon

## Context and Inputs

- Increment goal: add two visible tray actions (`Pomodoro`, `Break`), a visible `Quit` action that cleanly exits the app, and a minimal tray icon for the `examples/pomodoro` demo. See `increment.md` for acceptance clues.
- Design approach: small Go demo app with `internal/app` (timer logic), `internal/tray` (systray wrapper), `cmd/pomodoro` entry, and embedded icon assets. See `design.md`.
- Key inputs assumed: `increment.md`, `design.md`, and `CONSTITUTION.md` in `examples/pomodoro`.
- Important assumptions: one small systray dependency is acceptable; assets will be embedded with `//go:embed`; CI can run a smoke command.

## High-Level Approach

- Keep domain logic (timer) pure and testable under `internal/app`.
- Isolate platform code in `internal/tray` and use a small systray package for macOS tray/menu integration.
- Provide `cmd/pomodoro` that wires app + tray, supports `--smoke` and `--version`, and performs graceful shutdown on Quit.
- Embed a single, minimal icon in `assets/` to preserve a single-binary distribution.

## Workstreams and Steps

Workstreams are ordered to minimize risk: Setup → App → Tray → CLI/Wiring → Assets → Tests/CI → Docs/PR.

- Workstream: Setup (repo-local module setup)
  1. Initialize a Go module for the example: run `go mod init github.com/co0p/4dc/examples/pomodoro` inside `examples/pomodoro`.
     - Why: Keep the example compilable and dependency-managed independently of the repo root; required before adding `go` code and external systray dep.
     - Validation: `go env GOMOD` prints `examples/pomodoro/go.mod` and `go list` succeeds.

- Workstream: App (domain) — `examples/pomodoro/internal/app`
  1. Add package and interface `App` with methods: `StartPomodoro()`, `StartBreak()`, `Shutdown(ctx context.Context) error`, and `OnStateChange(func(State))`.
     - Why: Clear contract for UI and tests.
     - Validation: package compiles.
  2. Implement minimal state machine (`Idle`, `PomodoroRunning`, `BreakRunning`) with cancellation semantics.
     - Validation: unit tests for transitions and idempotence.
  3. Add unit tests for transitions, repeated calls, and shutdown behavior.
     - Validation: `go test ./...` for the package passes and provides coverage for core logic.

- Workstream: Tray (platform) — `examples/pomodoro/internal/tray`
  1. Add `Tray` interface (`Run(ctx) error`, `Close() error`) and a systray-backed implementation using `github.com/getlantern/systray` (or similar approved package).
     - Why: Isolate platform code, allow mocking, and provide the visible menu.
     - Validation: package compiles; manual run shows menu created.
  2. Create menu items (`Pomodoro`, `Break`, `Quit`) and wire clicks to app methods; on `Quit` call `App.Shutdown(ctx)` then exit cleanly.
     - Validation: manual verification on macOS: menu items visible, log output on actions, `Quit` exits without crash.
  3. Add a mock tray implementation for unit tests to verify wiring without launching a UI.
     - Validation: `go test` for tray package passes.

- Workstream: CLI/Wiring — `examples/pomodoro/cmd/pomodoro`
  1. Add `main.go` that parses `--version` and `--smoke` flags, sets up logging, constructs `App` and `Tray`, and runs the tray loop.
     - Why: Single entrypoint for local demos and CI smoke checks.
     - Validation: `go build ./cmd/pomodoro` succeeds.
  2. Implement `--smoke` to initialize subsystems then immediately request shutdown and exit 0; use this in CI to validate non-crashing startup.
     - Validation: running `./pomodoro --smoke` exits 0 quickly.

- Workstream: Assets & Embedding — `examples/pomodoro/assets`
  1. Add a minimal icon file (e.g., `icon.png`, small size) to `assets/` and embed it using `//go:embed` in the tray package.
     - Why: Keep packaging single-binary and assets tiny to respect constitution.
     - Validation: `go build` succeeds and local run displays the embedded icon.

- Workstream: Tests & CI
  1. Add unit tests as described (app + tray mocks). Keep tests fast (<60s).
  2. Add CI smoke step to build `cmd/pomodoro` and run `./pomodoro --smoke` (exit 0 expected). Also run `go test ./...`.
     - Validation: CI passes for PRs with new code.

- Workstream: Docs & PR
  1. Add/update `examples/pomodoro/README.md` with instructions to build and run the demo and acceptance checks (menu labels, Quit behavior, icon visibility).
  2. In the PR description, include acceptance clues, mention added dependency, and a short test plan.

## Testing Plan

- Unit tests:
  - `internal/app`: test transitions, idempotence, and Shutdown.
  - `internal/tray`: mock-based tests to assert wiring.
- Smoke test:
  - CI builds binary and runs `./pomodoro --smoke` to confirm non-crashing startup and exit.
- Local manual validation:
  - Build: `go build -o ./bin/pomodoro ./cmd/pomodoro`
  - Run: `./bin/pomodoro` — verify tray icon, menu items, action logs, and Quit behavior.

## CI/CD and Rollout

- CI: run `go test ./...`, build `cmd/pomodoro`, and run smoke command. Optionally measure binary size; fail if above threshold.
- Rollout: additive example-only change — merge to `main`. No feature flags required.
- Rollback: revert the PR to remove the example if issues appear.

## Observability and Post-Deployment Checks

- Logging: log `action=<name>` for user actions and lifecycle events (`status=shutdown-started`, `status=shutdown-complete`). Keep logs human-friendly.
- Post-merge checks (manual/staging): verify icon visible, both actions present and produce logs, and Quit exits cleanly.

## Risks, Dependencies, and Coordination

- Dependency: `github.com/getlantern/systray` (or similar). Risk: binary size increase. Mitigation: choose a small package, measure size in CI, document trade-offs in PR.
- Risk: macOS tray differences. Mitigation: keep interactions minimal and document tested macOS versions.
- Coordination: notify reviewers in PR; no cross-team coordination expected.

## Follow-up and Cleanup

- Optional follow-ups: keyboard shortcuts, improved icons, packaging notes for retina assets, or adding lightweight UI smoke tests.
- Cleanup: consider whether `--smoke` remains useful; remove or keep based on team preference.

## Acceptance Criteria Mapping

- The work is acceptable when:
  - Tray shows two functioning actions and a Quit action that cleanly exits.
  - Minimal icon is visible in the tray.
  - Unit tests for `internal/app` and mock tests for `internal/tray` pass.
  - CI smoke step builds the binary and `--smoke` exits 0.
