Design: TitleUpdater, Unsubscribeable Subscriptions & Deterministic Tests

## Context and Problem

This increment improves the maintenance and testability of the tray-title updater introduced earlier. The product goal is to keep the tray showing a concise remaining-time label while a Pomodoro or Break is running, but implement that behavior with clearer lifecycle control, unsubscribeable app subscriptions, named timing constants, and deterministic tests.

Why now
- The current helper is a function-style updater with an implicit lifecycle and the `App` subscription surface supports a single `OnStateChange` setter without an unsubscribe mechanism. Tests rely on brief sleeps and wiring loops which makes them less deterministic. These issues are localized, small, and straightforward to fix while preserving user-visible behavior.

Relevant existing code
- `internal/tray/updater.go` — function `ManageTitleUpdates` implements the ticker and state-handling loop.
- `internal/tray/updater_test.go` — unit test using a fake app and a fake ticker channel; relies on `time.Sleep` and a wiring wait loop.
- `internal/tray/systray_impl.go` — systray wiring where the updater is started; macOS requires systray calls on the main thread.
- `internal/app/app.go` — `App` interface includes `OnStateChange(fn func(State))` and `Remaining()`; implementation stores a single callback and has no unsubscribe mechanism.

Guardrails and constraints
- Respect the project's `CONSTITUTION.md`: keep changes small, avoid new runtime dependencies, preserve single-binary buildability, and keep tests fast and reliable.

## Proposed Solution (Technical Overview)

Introduce a `TitleUpdater` type and an unsubscribe-capable subscription API while preserving backwards compatibility:

- New `TitleUpdater` type (in `internal/tray/updater.go`):
  - Purpose: make updater lifecycle explicit and testable.
  - Fields: `app app.App`, `setTitle func(string)`, `clearTitle func()`, `tickerFactory func(d time.Duration) (<-chan time.Time, func())`, internal state for the current ticker/stopper and the unsubscribe function returned by the app subscription.
  - API:
    - `NewTitleUpdater(app app.App, setTitle func(string), clearTitle func(), tickerFactory func(d time.Duration) (<-chan time.Time, func())) *TitleUpdater` — constructor.
    - `Run(ctx context.Context)` — subscribes to app state, begins the management loop and returns when ctx is cancelled.
    - `Stop()` — detaches subscription and stops any running ticker; idempotent and safe to call from tests and wiring.

- App subscription API (in `internal/app/app.go`):
  - Add `SubscribeStateChange(fn func(State)) func()` which returns an unsubscribe function. Implement `timerApp` to maintain a small slice (or map) of subscribers and call them while holding the lock.
  - Keep `OnStateChange(fn func(State))` as a compatibility adapter that calls `SubscribeStateChange(fn)` and discards the unsubscribe. This avoids breaking existing callers and makes migration incremental.

- Named timing constants:
  - Add `const titleUpdateInterval = 10 * time.Second` in `internal/tray` and use it instead of inline literals. Optionally document default session durations in `internal/app` (e.g., `defaultPomodoro` / `defaultBreak`).

- Tests & determinism:
  - In `internal/tray/updater_test.go`, replace sleep-based wiring waits with explicit channels. TitleUpdater tests will expose a `wired` channel or the test will wait for the subscription to be set via the fake app/unsubscribe handshake.
  - Continue using injected `tickerFactory` that returns a deterministic tick channel and stopper; tests will drive ticks by sending into that channel.
  - Update `internal/app/remaining_test.go` where appropriate to avoid long sleeps by using short durations already supported by `New(...)` and asserting state changes via observable conditions rather than blind sleeps.

How the flow works after change
- `systray_impl.go` will create title setter wrappers inside `systray.Run` (respecting macOS main-thread constraint), construct the `TitleUpdater` with `time.NewTicker` based factory, and call `Run(ctx)`.
- `TitleUpdater.Run` subscribes to app state (receives an unsubscribe func), and on state transitions to running starts or restarts the ticker and performs an immediate update using `app.Remaining()`. On transition to idle it clears the title and stops the ticker. `Stop()` will call unsubscribe and stop the ticker.

## Scope and Non-Scope (Technical)

In-scope
- Add `TitleUpdater` and tests in `internal/tray`.
- Add `SubscribeStateChange` to `App` and `timerApp` implementation and test subscription/unsubscribe behavior.
- Update tests to use explicit synchronization and deterministic tick channels.
- Add named timing constants and short doc comments (tray threading guidance) and a `README.md` note.

Out-of-scope
- No changes to user-visible behavior or UI beyond implementation hygiene.
- No new dependencies or packaging changes.
- No cross-repo extraction or broad API versioning beyond the small `SubscribeStateChange` addition and backward-compatible adapter.

## Architecture and Boundaries

- Components involved:
  - Domain: `internal/app` (timer state and Remaining calculation).
  - Platform/UI: `internal/tray` (TitleUpdater, systray wiring).
  - Entry: `cmd/pomodoro` (bootstraps systray; unchanged).
- Boundaries:
  - Domain remains responsible for timer state and remaining-time computation. The updater queries only read-only `Remaining()` and subscribes to state change events.
  - The tray updater is a platform concern: it controls presentation and must obey the systray main-thread constraint (setters created inside `systray.Run`).

Constitution alignment
- No new runtime deps, changes are small and reversible, tests remain fast and localized, and documentation clarifies threading assumptions.

## Contracts and Data

- `App` interface change (addition):
  - New method signature: `SubscribeStateChange(fn func(State)) func()` — returns `unsubscribe` function. This is additive; `OnStateChange(fn func(State))` remains as a compatibility adapter.
- `TitleUpdater` contract:
  - `Run(ctx)` should be invoked in the systray-run context or with setter wrappers that are safe to call; `Stop()` is safe to call concurrently.
- No persistent data or schema changes are required.

Compatibility and migration
- Existing callers of `OnStateChange` continue to work. New code should prefer `SubscribeStateChange` so it can unsubscribe cleanly. Migration is local to the example; if other packages in the repo depend on the old `OnStateChange`, they will continue to operate.

## Testing and Safety Net

Unit tests
- `internal/tray`:
  - `TestTitleUpdater_ImmediateUpdateOnStart` — verify immediate setTitle when state goes to running.
  - `TestTitleUpdater_UpdatesOnTick` — simulate ticks through fake tick channel and verify periodic updates.
  - `TestTitleUpdater_StopUnsubscribes` — ensure `Stop()` prevents further calls after unsubscribe.
- `internal/app`:
  - `TestSubscribeStateChange_MultipleSubscribers` — verify multiple subscribers are called and unsubscribe works.
  - `TestRemainingBehavior` — keep existing tests but use shorter durations and observable conditions where possible.

Integration / smoke
- `go test ./...` remains the primary verification. No end-to-end GUI tests required for this increment; manual smoke of the binary (`./bin/pomodoro --smoke` and running the app locally) is recommended before release.

Flakiness mitigations
- Use explicit channels for wiring confirmation in tests instead of `time.Sleep`.
- Use injected ticker factory for deterministic tick control.

## CI/CD and Rollout

- CI impact
  - No new CI jobs required. Changes should pass the existing build + `go test ./...` pipeline.
  - Keep commits small; the PR should include the updated tests and acceptance notes.

- Rollout and rollback
  - Safe to land in main as the change is internal and additive. If any regression appears, revert the small PR.
  - No feature flags required; if desired for dogfooding, the systray wiring can gate the `TitleUpdater` behind a boolean toggle in `cmd/pomodoro` (not necessary here).

## Observability and Operations

- Logging:
  - Add brief info-level logs when `TitleUpdater` starts, stops, subscribes, and unsubscribes. Keep messaging concise and human-friendly per constitution.
  - Log unexpected errors (none expected here) at warn/error level.

- Metrics (optional):
  - An in-memory counter for title updates can be added for debugging; not required for this increment.

- Operational checks:
  - Manual smoke: verify tray title updates and that no stray goroutines remain after stop.

## Risks, Trade-offs, and Alternatives

- Risks:
  - API surface increase (additional subscription method) requires small callers update if they want unsubscribe; mitigated by preserving `OnStateChange` as an adapter.
  - Tests must be carefully written to avoid reintroducing sleeps; initial iteration may require small tweaks.

- Trade-offs:
  - `TitleUpdater` struct provides better lifecycle clarity at the cost of a slightly larger API surface (one new type). This pays off in readability and testability.

- Alternatives considered:
  - Mutating `OnStateChange` to return an unsubscribe: more breaking — rejected in favor of an additive `SubscribeStateChange` method.
  - Keep function-style helper but add explicit unsubscribe token: this reduces clarity compared to a small typed `TitleUpdater` and was not chosen.

## Follow-up Work

- Create a small integration test exercising repeated start/stop cycles to ensure no leaks under repeated use.
- If the pattern repeats across examples, consider extracting a tiny helper package to standardize tray lifecycle wiring.

## References

- `increment.md` (this folder)
- `CONSTITUTION.md` (project root: `examples/pomodoro/CONSTITUTION.md`)
- Relevant code: `internal/tray/updater.go`, `internal/tray/updater_test.go`, `internal/app/app.go`, `internal/app/remaining_test.go`, `internal/tray/systray_impl.go`
