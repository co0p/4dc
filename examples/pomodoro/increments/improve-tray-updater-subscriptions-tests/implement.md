Implementation Plan: Improve Tray Updater, Subscriptions & Tests

## Context and Inputs

- Increment: `increment.md` in this folder — improve tray updater lifecycle, add unsubscribeable subscriptions, add timing constants, and make tests deterministic.
- Design: `design.md` in this folder — introduces `TitleUpdater`, additive `SubscribeStateChange` API, named constants, and deterministic test strategy.
- Assumptions: local dev environment supports `go test ./...` and building the example; no feature-flag system is required for this small, internal change.

Date: 2025-12-05

## High-Level Approach

Break work into four small workstreams to keep changes reviewable and reversible:
- Workstream A: Add an additive subscribe/unsubscribe API to `internal/app` and tests.
- Workstream B: Implement `TitleUpdater` in `internal/tray` and add deterministic unit tests.
- Workstream C: Wire `TitleUpdater` in `systray_impl.go`, add `titleUpdateInterval` constant and thread docs, update README.
- Workstream D: Test hygiene and small naming/doc passes.

Each workstream is a small PR that includes focused tests and a short acceptance checklist. Run `go test ./...` after each PR.

## Workstreams and Steps

Workstream A — App subscription API
1. Add `SubscribeStateChange(fn func(State)) func()` to the `App` interface (additive).
   - Files: `internal/app/app.go`.
   - Why: Provide a clean unsubscribe mechanism without breaking callers.
   - Validation: Add unit tests verifying multiple subscribers are called and that the returned unsubscribe stops notifications. Run `go test ./internal/app`.

2. Implement subscriber list in `timerApp` and adapt `OnStateChange(fn)` to call `SubscribeStateChange(fn)` and discard the unsubscribe.
   - Files: `internal/app/app.go`, add `internal/app/app_subscribe_test.go`.
   - Validation: `go test ./internal/app` passes and existing `remaining_test.go` still passes.

Workstream B — TitleUpdater implementation & tests
3. Add named constant `titleUpdateInterval = 10 * time.Second` in `internal/tray` (top of updater file) with a short comment.
   - Files: `internal/tray/updater.go` (or a small constants section)
   - Validation: code compiles.

4. Create `TitleUpdater` type with `NewTitleUpdater(...)`, `Run(ctx)` and `Stop()` methods. Move and adapt logic from `ManageTitleUpdates` into these methods. Keep injected `tickerFactory` dependency.
   - Files: `internal/tray/updater.go` (replace the function-style helper with the typed implementation).
   - Validation: `go test ./internal/tray` builds; add unit tests in the next step.

5. Implement deterministic unit tests for the updater:
   - `TestTitleUpdater_ImmediateUpdateOnStart` — immediate update when transitioning to running.
   - `TestTitleUpdater_UpdatesOnTick` — tick-driven updates via fake tick channel.
   - `TestTitleUpdater_StopUnsubscribes` — calling `Stop()` prevents further updates.
   - Use explicit channels to detect wiring and injected `tickerFactory` to control ticks.
   - Files: `internal/tray/updater_test.go`.
   - Validation: `go test ./internal/tray` passes consistently.

Workstream C — Wiring & docs
6. Wire `TitleUpdater` into `internal/tray/systray_impl.go`:
   - Create setter wrappers inside `systray.Run` and pass them to `NewTitleUpdater`.
   - Start `TitleUpdater.Run(ctx)` and ensure `Stop()` is called on shutdown.
   - Add a brief comment in `systray_impl.go` explaining the macOS main-thread requirement and the setter-wrapping pattern.
   - Validation: `go build ./cmd/pomodoro` succeeds and a manual smoke of the binary shows unchanged tray behavior.

7. Update `examples/pomodoro/README.md` with a short note referencing `titleUpdateInterval` and the main-thread setter pattern.
   - Validation: README updated; no code impact.

Workstream D — Test hygiene & small renames
8. Replace `time.Sleep` and wiring waits in `internal/tray/updater_test.go` and `internal/app/remaining_test.go` with explicit synchronization (channels/wired signals) and use intention-revealing test names.
   - Files: `internal/tray/updater_test.go`, `internal/app/remaining_test.go`.
   - Validation: `go test ./...` passes reliably and faster.

9. Small hygiene pass: add doc comments for exported items touched and rename ambiguous locals (e.g., `a` -> `appSvc`) in changed files. Keep renames limited to files changed above.
   - Validation: Build and tests pass; diffs are reviewable.

## Testing Plan

- Unit tests: add/modify tests for `internal/app` and `internal/tray` as described above.
- Integration/smoke: run `go build -o bin/pomodoro ./cmd/pomodoro` and `./bin/pomodoro --smoke` locally to confirm startup behaviour.
- Test commands (run from `examples/pomodoro`):

```bash
go test ./internal/app
go test ./internal/tray
go test ./...
``` 

- Test determinism notes:
  - Use channels for wiring confirmation; avoid arbitrary sleeps.
  - Keep test durations short by using `New(...)` with small durations when needed.

## CI/CD and Rollout Plan

- CI: no changes to pipeline required. Each PR must pass existing `go test ./...` checks.
- Rollout: safe to merge directly to `main` after review since changes are internal and additive. If issues appear, revert the PR.
- Optional: dogfood via local builds before merging.

## Observability and Post-Deployment Checks

- Add minimal info logs in `TitleUpdater.Run` and `Stop` to aid debugging (stdout/stderr). Keep logs concise.
- Post-merge checks:
  - Ensure `go test ./...` passes on CI.
  - Manual smoke: run `./bin/pomodoro` and confirm tray label shows `Xm` while running and clears when idle.

## Risks, Dependencies, and Coordination

- Risk: concurrency bugs in subscription handling. Mitigation: add focused unit tests for subscribe/unsubscribe behavior, and a test where unsubscribe is called from a callback.
- Risk: test flakiness during first iteration; mitigate with explicit channel synchronization and short watchdogs.
- Dependency: none external. Keep all changes within `examples/pomodoro`.

## Follow-up and Cleanup

- Optional follow-up: add a small integration test exercising repeated start/stop cycles to ensure no leaks.
- Consider extracting tray lifecycle helper if the pattern is reused across examples.

## Acceptance Checklist (for PRs)

- All new and updated unit tests pass locally and in CI (`go test ./...`).
- `go build ./cmd/pomodoro` succeeds.
- The PR description includes: what changed, how to run tests, and a short manual smoke-check.
- Changes are small and reversible; maintainers can review diffs quickly.
