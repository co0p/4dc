Implementation Plan: Short and Long Break Actions

## 1. Context and Inputs

- Increment goal: expose explicit `Short Break` (5m) and `Long Break` (25m) actions in the tray so users can start either break length directly. See `increment.md`.
- Agreed design: add two explicit domain methods `StartShortBreak()` and `StartLongBreak()` implemented by `timerApp`; keep `StartBreak()` for compatibility; add tray menu items and wiring; add tests and a short README/CHANGELOG note. See `design.md`.
- Key assumptions (from constitution and design): small, reversible changes; no new dependencies; tests must be fast (use small durations in tests by constructing or adjusting internal structs); change is local to `examples/pomodoro`.

## 2. High-Level Approach

Work will be organized into four workstreams:

- Domain (app) changes — add API surface and implement behavior inside `internal/app`.
- Tray UI wiring — add menu items and wire to domain methods in `internal/tray/systray_impl.go`.
- Tests — unit tests for domain and small integration tests for tray wiring using `MockTray`.
- Docs & housekeeping — README/CHANGELOG entry and formatting.

Changes are additive and backward-compatible: `StartBreak()` remains available and will delegate to `StartShortBreak()` so existing callers still work.

Constructor decision (confirmed): Option B — keep the existing `New(...)` semantics (no new required parameters). `longBreakDuration` will default to 25 minutes inside the implementation. Tests will set test-friendly values by constructing `timerApp` directly (tests run in the same package) or by mutating the instance after `New()` where appropriate to keep tests fast.

## 3. Workstreams and Steps

Workstream A — Domain (app) changes

Purpose: Add two explicit domain APIs and implement them on `timerApp`.

Steps:
1. Add method signatures to `App` interface (file: `internal/app/app.go`):
   - `StartShortBreak()`
   - `StartLongBreak()`
   - Why: define contract for tray and tests.
   - Validation: ensure project compiles (`go test ./internal/app`).

2. Add `longBreakDuration time.Duration` to `timerApp` struct and initialize it with a 25-minute default in `New()` (no signature change to `New`).
   - Why: keep default long-break behavior as requested.
   - Validation: compile and run domain tests.

3. Implement `StartShortBreak()` and `StartLongBreak()` on `timerApp`.
   - Reuse the existing timer-start logic (cancel existing timer, set `end`, set state, notify subscribers, start goroutine to clear state when elapsed).
   - Implement `StartBreak()` to call `StartShortBreak()` to preserve compatibility.
   - Validation: compile and run unit tests (fast tests; see Workstream C).

4. Update documentation comments/godoc on new exported methods (follow ADR conventions).
   - Why: maintain codebase style and generated docs.
   - Validation: `gofmt` and `go vet` (if used).

Workstream B — Tray UI wiring

Purpose: Add menu items and wire actions to new domain methods.

Steps:
1. Update `internal/tray/systray_impl.go` to add two menu items when constructing menus:
   - `mShort := systray.AddMenuItem("Short Break", "Start Short Break")`
   - `mLong := systray.AddMenuItem("Long Break", "Start Long Break")`
   - Why: user-visible labels.
   - Validation: compile and optionally run the app to inspect the menu.

2. Add goroutine handlers matching existing style: for each menu item's `ClickedCh` start a goroutine that logs the action and calls the corresponding domain method:
   - `StartShortBreak()` and `StartLongBreak()`.
   - Add log lines: `log.Println("action=StartShortBreak")` and `log.Println("action=StartLongBreak")`.
   - Validation: run `go test ./internal/tray` (after tests updated) and manual run.

3. Keep existing `Pomodoro` and `Quit` flows unchanged.

Workstream C — Tests

Purpose: Provide a safety net and fast checks.

Steps:
1. Domain unit tests (in `internal/app`):
   - Add `TestStartShortBreakSetsRemaining`:
     - Construct a `timerApp` instance with test-friendly durations by either:
       - Creating `t := &timerApp{pomodoroDuration: 25*time.Millisecond, breakDuration: 5*time.Millisecond, longBreakDuration: 25*time.Millisecond}` and any necessary fields initialized, or
       - Call `New()` and mutate `longBreakDuration` directly in the test since tests live in the same package.
     - Call `t.StartShortBreak()` and assert `t.State()` == `StateBreakRunning` and `t.Remaining()` is approximately the configured test duration (allow a small tolerance).
   - Add `TestStartLongBreakSetsRemaining` similarly for `StartLongBreak()`.
   - Add `TestStartBreakIsCompatible` to assert `StartBreak()` triggers short break behavior.
   - Validation: `go test ./internal/app` should execute quickly because tests use small durations.

2. Tray integration tests (in `internal/tray`):
   - Update `internal/tray/mock.go:Trigger` to handle `"Short Break"` and `"Long Break"` names.
   - Add tests that construct `MockTray` with a test `app` instance (fast durations) and call `Trigger("Short Break")` / `Trigger("Long Break")` and assert `app.State()` and `Remaining()`.
   - Validation: `go test ./internal/tray` passes.

3. Update any existing tests that referenced the old single `Break` menu entry to use the new triggers or compatibility path.

4. Run full test suite locally:
   - `go test ./...` and ensure runtime is fast. If any test is slow, revise to use smaller durations or mocking.

Workstream D — Docs & housekeeping

Purpose: Surface the change to users and maintainers and keep code style consistent.

Steps:
1. Add a short note to `README.md` or `CHANGELOG.md` under `examples/pomodoro` describing the new menu actions and default durations.
   - Validation: file added and spelled correctly.

2. Run `gofmt -w` on changed files and ensure godoc comments for new exported methods follow ADR conventions.

3. Update `internal/tray/mock.go` documentation comments to reflect new triggers.

## 4. Testing Plan (when to run which tests)

- After Domain step 1–3: run `go test ./internal/app`.
- After Tray wiring changes: run `go test ./internal/tray`.
- After all changes: run `go test ./...` to validate full project.
- Use short durations in tests to keep full-suite run times low.

## 5. CI/CD and Rollout Plan

- CI: No new pipeline steps required. Changes will be validated by the existing `go test ./...` in CI.
- Rollout: Merge the small PR after CI passes. No feature flags required because change is local and reversible.
- Rollback: Revert the PR if any regression appears. `StartBreak()` compatibility reduces immediate client breakage risk.

## 6. Observability and Post-Deployment Checks

- Ensure new log lines are present when actions are exercised locally: `action=StartShortBreak` and `action=StartLongBreak`.
- Manual post-merge check: run the binary locally and verify tray contains `Pomodoro`, `Short Break`, `Long Break`, `Quit` and each action starts the correct session.
- CI check: new tests must pass. If tests fail, investigate failing assertion details and times.

## 7. Risks, Dependencies, and Coordination

- Risk: Tests may be slow if not using test durations — mitigated by constructing timerApp with short durations inside package tests.
- Risk: Minor API surface increase — acceptable for example project.
- Decision point: constructor remains unchanged (Option B). Tests will set `longBreakDuration` directly when needed.

## 8. Follow-up and Cleanup

- After rollout, optionally remove `StartBreak()` (if desired) in a follow-up increment once callers are updated and compatibility is no longer needed.
- Consider adding an optional constructor overload or helper for tests if future tests need easier injection of `longBreakDuration` (follow-up).

## Quick commands

- Format and lint changed files:

```bash
gofmt -w internal/app/*.go internal/tray/*.go
```

- Run tests locally:

```bash
cd examples/pomodoro
go test ./...
```

---

If this outline looks good, I will write the `implement.md` file accordingly (or save it here). Reply with `yes` to proceed and I will create `implement.md` under `examples/pomodoro/increments/short-and-long-break-actions/implement.md`.