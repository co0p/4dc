Design: Short and Long Break Actions

## 1. Context and Problem

The product increment requires exposing explicit Short Break (5m) and Long Break (25m) actions in the tray UI so users can start either break length directly. Currently the demo exposes a single `Break` action (5m), `Pomodoro` (25m), and `Quit`.

Why now: this is a small, user-visible UX improvement with low technical risk and immediate benefit: users can choose a longer break without changing pomodoro durations or adding automatic scheduling.

Relevant existing behavior and components:
- `internal/app` (core domain): exposes `StartPomodoro()` and `StartBreak()`, `Remaining()` and state subscription APIs. `New()` sets default durations (25m pomodoro, 5m break).
- `internal/tray/systray_impl.go`: constructs tray menu and wires click handlers to `app` methods; logs menu actions.
- `internal/tray/mock.go`: in-process mock tray used in tests.
- `internal/tray/updater.go`: subscribes to app state changes and updates tray title based on `Remaining()`; no API changes required here.

Links:
- increment.md: ../short-and-long-break-actions/increment.md
- constitution: ../../CONSTITUTION.md

## 2. Proposed Solution (Technical Overview)

High-level approach (small, backward-compatible):

- Add two explicit methods to the `App` interface and the `timerApp` implementation:
  - `StartShortBreak()` — starts a break with the existing short-break duration (5m default).
  - `StartLongBreak()` — starts a break with a long-break duration (25m default).

- Keep `StartBreak()` present for compatibility (calls `StartShortBreak()` internally or preserved for callers). This keeps the public package surface backward-compatible for any existing internal callers.

- Update `internal/tray/systray_impl.go` to add two menu items: `Short Break` and `Long Break`, and wire their `ClickedCh` handlers to call `a.StartShortBreak()` and `a.StartLongBreak()` respectively. Add corresponding log lines for each action.

- Update `internal/tray/mock.go` `Trigger` to support the new names (`"Short Break"`, `"Long Break"`) so tests can simulate clicks.

- Add unit tests in `internal/app` validating that `StartShortBreak()` and `StartLongBreak()` set state to break-running and set `Remaining()` to approximately the expected durations.

- Add tests in `internal/tray` for the mock tray and integration-style test(s) that exercise the wiring between tray and app using the `MockTray` (or existing test helpers). Update existing tests that referenced the old single `Break` action.

Rationale: adding small explicit methods keeps the change minimal and easy to reason about. It avoids changing method signatures or core timer internals and preserves backward compatibility.

Typical flow after change:
1. User clicks `Short Break` in the tray.
2. `systray_impl` handler calls `app.StartShortBreak()`.
3. `timerApp` starts a break session with `end = now + shortBreakDuration`, sets state to `StateBreakRunning`, and notifies subscribers.
4. `TitleUpdater` (already subscribing) updates the tray title using `Remaining()`; tests and manual observation validate behavior.

## 3. Scope and Non-Scope (Technical)

In-scope:
- Adding `StartShortBreak()` and `StartLongBreak()` to the `App` interface and implementing them on `timerApp`.
- Updating `systray_impl.go` to add two menu items and wire handlers.
- Updating `internal/tray/mock.go` and tests to support new menu names.
- Adding unit tests and small integration tests validating the observable behavior described in the increment.
- Adding a short README/CHANGELOG note (increment task).

Out-of-scope (explicit):
- Adding user-configurable durations or persistent settings.
- Adding automatic scheduling of long breaks after N pomodoros.
- Introducing new external dependencies or telemetry.

## 4. Architecture and Boundaries

Components involved and responsibilities:
- `internal/app` (domain)
  - Responsibilities: manage session lifecycle, durations, subscribers, and `Remaining()`.
  - Change: add two convenience methods `StartShortBreak()` and `StartLongBreak()` that reuse existing timer-start logic.
- `internal/tray/systray_impl.go` (infrastructure/UI)
  - Responsibilities: build tray menu, handle clicks, invoke domain API, log actions.
  - Change: add two menu items and wire them to new domain methods.
- `internal/tray/mock.go` and tests (test infra)
  - Responsibilities: simulate tray events in tests; updated to include new triggers.
- `internal/tray/updater.go` (infrastructure)
  - No change required; continues to read `app.Remaining()` and update tray title.

Guardrails and constraints:
- No new packages or external services.
- Keep changes confined to `examples/pomodoro` to honor single-binary constraint.
- Follow existing naming and documentation conventions (receiver names, godoc, logs).

## 5. Contracts and Data

API changes (internal package `app`):
- Add to `App` interface:
  - `StartShortBreak()`
  - `StartLongBreak()`
- Implementation contract: invoking either method starts a break session and sets the app state to `StateBreakRunning`. `Remaining()` must report the correct remaining duration immediately after start.

Compatibility:
- Existing callers of `StartBreak()` remain functional; consider implementing `StartBreak()` to call `StartShortBreak()` to preserve behavior.
- No storage, events, or external contracts are changed.

No migration or versioning strategies are required — this is an internal API change within the example project.

## 6. Testing and Safety Net

Unit tests:
- `internal/app` tests:
  - Test that `StartShortBreak()` sets state to `StateBreakRunning` and `Remaining()` is between 4m50s and 5m10s (allow a small tolerance for test execution delay). Use test-friendly durations (constructor injection) to make the test fast: e.g., create `New(25*time.Millisecond, 5*time.Millisecond)` if such constructor is available; prefer small, deterministic durations in tests.
  - Test that `StartLongBreak()` sets state to `StateBreakRunning` and `Remaining()` is approximately 25m (or short test duration variant).
  - Test that `StartPomodoro()` behavior is unchanged.

Integration tests (small, fast):
- `internal/tray` tests:
  - Using `MockTray`, simulate `Trigger("Short Break")` and assert `app.State()` == `StateBreakRunning` and `Remaining()` ≈ 5m (using injected test durations).
  - Similarly for `Trigger("Long Break")`.
  - Update or add a test that the tray menu includes the four items (where possible in unit tests that construct the systray implementation with test helpers or by asserting the mock exposes names).

Regression tests:
- Update any test that assumed the single `Break` item.

Safety notes:
- Use constructor injection of short durations in tests to keep unit/integration tests fast (< 60s) per constitution.
- Keep tests deterministic by avoiding time.Sleep where possible; use injected ticker factories or small durations.

## 7. CI/CD and Rollout

CI implications:
- No new pipeline steps required. Normal `go test ./...` should cover updated tests.
- Ensure unit tests remain fast by using short durations in test constructors.

Rollout plan:
- Deliver as a small PR that includes code changes and tests.
- Merge after CI passes.
- No staged rollout required; behavior is local to the binary and has no external dependencies.

Rollback:
- Revert the PR if behavioral regression appears; as a small change the revert is low-risk.
- `StartBreak()` remains available as compatibility, so callers will keep working (minimizes regression risk).

## 8. Observability and Operations

Logging:
- Add log lines in `systray_impl.go` click handlers for `Short Break` and `Long Break` similar to existing `StartPomodoro` logs, e.g. `log.Println("action=StartShortBreak")` and `log.Println("action=StartLongBreak")`.
- Domain-level logs: `timerApp` may log state transitions in existing levels (info) when starting/ending sessions — avoid noisy logs.

Metrics (optional):
- A simple counter of how often Short vs Long break is used could be added later as a follow-up. For now, logs provide sufficient operational signal.

Signals to observe:
- Success: user-initiated break actions present in logs and `Remaining()` behavior matches expectations.
- Trouble: unexpected panics, missing menu items (regression in systray wiring), or test failures in CI.

## 9. Risks, Trade-offs, and Alternatives

Risks:
- Slight API surface increase in `app` (two new methods) — small and internal to the example project.
- Tests must be updated; small risk of introducing flakiness if not using constructor injection.

Trade-offs and alternatives considered:
- Change `StartBreak()` to accept a duration or enum: this reduces the number of methods but requires changing the method signature (potentially breaking callers) and slightly more refactoring. Rejected for this increment in favor of minimal, explicit methods.
- Keep only `StartBreak()` and add a parameterless `StartLongBreak()` helper elsewhere: rejected because adding clear `StartShortBreak()`/`StartLongBreak()` keeps intent explicit and simple.

## 10. Follow-up Work

- Add configurable durations via a small settings mechanism (file or UI) as a separate increment.
- Implement optional automatic long-break scheduling after N pomodoros as a separate feature (follow-up increment).
- Add small usage metrics to measure adoption of long vs short break.

## 11. References

- increment.md: ../short-and-long-break-actions/increment.md
- CONSTITUTION.md: ../../CONSTITUTION.md
- ADR: docs/ADR-2025-12-05-receiver-naming-and-docs.md
