Design: Demo tray actions, quit button, and minimal icon

## Context and Problem

This increment implements the product-level WHAT described in `increment.md`: add two visible tray actions (`Pomodoro`, `Break`), a Quit action that cleanly exits the app, and a minimal tray icon so the example is demo-ready. The change should be small, testable, and respect `CONSTITUTION.md` (single-binary, minimal dependencies, no telemetry by default).

Current state:
- The `examples/pomodoro` folder contains documentation but no source code in this repository. This design assumes adding a minimal Go app under `examples/pomodoro` to serve as a demo binary.

Links:
- increment: `examples/pomodoro/increments/demo-actions-and-quit-button/increment.md`
- constitution: `examples/pomodoro/CONSTITUTION.md`

## Proposed Solution (Technical Overview)

High-level approach
- Add a small, self-contained demo app under `examples/pomodoro` that implements:
  - an application core (`internal/app`) containing a minimal timer state machine and lifecycle methods,
  - a thin tray-wrapper (`internal/tray`) that integrates with a lightweight systray package to render the tray icon and menu items,
  - a `cmd/pomodoro` entry that wires the two and runs the app.
- Use one small Go dependency for tray integration (example: `github.com/getlantern/systray`) to avoid writing low-level native bindings. Keep usage minimal and provide an interface to allow mocking in tests.
- Include a minimal icon asset (placeholder PNG) in `examples/pomodoro/assets/` and embed it into the binary via Go's `//go:embed` (optional, but recommended to keep single-binary packaging simple).

Component responsibilities
- `internal/app` (pure, domain logic)
  - Maintain a simple state machine representing timer states: `Idle`, `PomodoroRunning`, `BreakRunning`.
  - Expose methods: `StartPomodoro()`, `StartBreak()`, and `Shutdown(ctx)`.
  - Provide an event/callback registration API (a channel or callback hook) so the UI can observe state changes for future updates.

- `internal/tray` (platform interaction)
  - Implement a `Tray` interface that can be constructed with an `App` instance and an icon; create menu items: `Pomodoro`, `Break`, and `Quit`.
  - Wire menu item clicks to call `App.StartPomodoro`, `App.StartBreak`, and `App.Shutdown` respectively.
  - Surface errors (initialization failure, icon load failure) via returned errors and logged messages.

- `cmd/pomodoro` (wiring)
  - Initialize logging, parse lightweight flags (`--version`, `--smoke`), instantiate `App` and `Tray`, and run the tray loop.
  - On Quit, request `App.Shutdown(ctx)` and ensure process exit after cleanup.

Why this split
- Keeps domain logic testable and independent of UI/frameworks.
- Keeps platform-specific code isolated and mockable.
- Matches constitutional guardrails: small, reversible, and minimizes new dependencies.

## Scope and Non-Scope (Technical)

In-scope:
- New minimal app skeleton under `examples/pomodoro` with `cmd/pomodoro`, `internal/app`, `internal/tray`, `assets/` and tests.
- One small external dependency for tray/menu support.
- Placeholder icon(s) embedded in the binary.
- Unit tests for the app state machine and tray wrapper using mocks.

Out-of-scope:
- Full-featured UI redesign, telemetry, cloud services, or retina-ready icon sets.
- Cross-platform polishing beyond the minimal macOS tray behavior.

## Architecture and Boundaries

- Domain vs. UI: Timer logic lives in `internal/app` (pure Go). UI and platform integration live in `internal/tray`.
- Dependency guardrail: Only one small systray package; avoid larger GUI frameworks.
- Packaging: Embed small assets into the binary (via `embed`) to keep packaging simple and single-binary friendly.

## Contracts and Data

Tray interface (conceptual)
```
type Tray interface {
  Run(ctx context.Context) error
  Close() error
}

type MenuItem interface {
  OnClick(fn func())
  SetTitle(title string)
}
```

App interface (conceptual)
```
type App interface {
  StartPomodoro()
  StartBreak()
  Shutdown(ctx context.Context) error
  OnStateChange(func(State)) // optional for UI updates
}
```

Notes:
- No persistent data schema changes are needed. If local state is later required, keep it simple (JSON file) and document location (per constitution).

## Testing and Safety Net

Unit tests
- `internal/app`:
  - Verify state transitions: Idle -> PomodoroRunning -> BreakRunning and back to Idle where appropriate.
  - Verify repeated actions behave sensibly (multiple Start calls should be idempotent or documented).
  - Verify Shutdown leaves the app in a stable state and cancels any timers.

- `internal/tray` (using a mock Tray implementation):
  - Assert that menu items are created with correct labels and that OnClick callbacks call the correct `app` methods.

Integration / smoke tests
- Add a CI smoke job that builds the binary and runs it with `--smoke` or `--version` to confirm a non-crashing startup. If `--smoke` is provided, the binary should initialize subsystems and exit quickly to allow automated validation.

Avoid UI end-to-end automation in this increment to keep scope small and tests fast.

## CI/CD and Rollout

CI changes
- Ensure `examples/pomodoro` builds in the repository's normal CI pipeline.
- Add unit tests to the default test suite; add a smoke step that runs the built binary with a flag to ensure it starts and exits cleanly.

Rollout
- This change is small and should follow the normal release process. No feature flags are required.

Rollback
- Revert the PR to remove the new code and assets. Keep the PR small to make rollbacks simple.

## Observability and Operations

Logging
- Log info-level messages for user actions (e.g., `action=StartPomodoro`), and lifecycle events (`action=Quit`, `status=shutdown-started`, `status=shutdown-complete`).
- Log initialization errors (icon load, tray init) at error level with context.

Metrics (optional, non-telemetry):
- Keep local counters for action invocations (e.g., in-memory counters logged on shutdown). Do not send telemetry off-host by default.

Diagnostics
- Provide clear stdout/stderr logs for smoke checks and local debugging per the constitution's human-centered logs guideline.

## Risks, Trade-offs, and Alternatives

Risks
- Adding a tray dependency increases binary size slightly; mitigation: pick a small, well-maintained package and measure binary size in CI.
- Platform differences in tray behavior across macOS versions; mitigation: keep interactions minimal and document any OS constraints.

Trade-offs
- Chosen approach uses a small external package to reduce implementation complexity and risk versus writing custom native bindings. This increases dependency surface slightly but keeps the change deliverable and reversible.

Alternatives
- Implement native tray integration directly (higher effort, larger scope) — candidate for follow-up if dependency is unacceptable.

## Follow-up Work

- Add keyboard shortcuts and more polished UI affordances.
- Add retina/hd icon assets and packaging notes for releases.
- Add lightweight UI smoke/e2e tests when a stable cross-platform harness is available.

## References
- `examples/pomodoro/increments/demo-actions-and-quit-button/increment.md`
- `examples/pomodoro/CONSTITUTION.md`
