Increment: Improve Tray Updater, Subscriptions & Tests

## Context

The Pomodoro demo shows session state and a concise remaining-time label in the macOS tray. Recent work added a small updater to surface remaining minutes in the tray, but a few maintainability and testability issues remain:

- The tray updater is packaged as a function-style helper with an implicit lifecycle, making it easy to leak background work when wiring/unwiring UI.
- The app's state subscription pattern does not provide a simple, reliable way to unsubscribe listeners; this increases the risk of stray callbacks and test interference.
- Tests for the tray updater and the timer domain use short sleeps and implicit wiring, which can make them less deterministic and slower to diagnose when they fail.
- Threading expectations for the tray API (main-thread requirement on macOS) and key timing values are under-documented, which increases onboarding friction for new contributors.

These issues are small, localized, and suitable for a focused, low-risk improvement that preserves the project's single-binary, minimal-dependency goals.

## Goal

Outcome / Hypothesis
- Maintainability and reliability will improve: maintainers can start and stop the tray updater explicitly, attach and detach listeners without leaks, and run deterministic tests that reveal intent quickly.
- User-visible behavior will remain the same (the tray continues to show remaining minutes while a session runs), but the implementation will be clearer and safer for future changes.

Scope
- This increment focuses on the tray-updater surface, subscription lifecycle, and tests. It is intentionally small and limited to improving readability, lifecycle control, and test determinism.

Non-Goals
- This increment will not change the app's user-facing behavior or UX beyond reduced flakiness in tests and clearer code.
- It will not introduce new runtime dependencies, external services, or packaging changes.
- It will not perform large-scale refactors outside the tray, subscription, and test surface.

Why this is a good small increment
- The change is narrowly scoped, reversible, and testable. It reduces the maintenance cost of an area that directly affects reliability (background updaters and subscriptions) and improves CI feedback by making tests deterministic.

## Tasks

- Task: Provide an explicit updater abstraction with a clear start/stop lifecycle so title updates can be started and stopped deterministically.
  - User/Stakeholder Impact: Maintainers can reason about when background updates run and ensure they stop during teardown or shutdown; reviewers see explicit lifecycle use in wiring.
  - Acceptance Clues: The codebase contains an updater abstraction with externally-callable start/stop methods (or equivalent lifecycle controls) used by the tray wiring; reviewers can point to explicit stop/teardown usage and confirm no background updater remains after stop.

- Task: Make state subscriptions unsubscribeable so listeners can be detached reliably.
  - User/Stakeholder Impact: Tests and runtime wiring can attach listeners temporarily and remove them without leaving stray callbacks, reducing unexpected updates and test interference.
  - Acceptance Clues: The subscription API returns a callable unsubscribe (or equivalent) and the updater and tests call unsubscribe as part of teardown.

- Task: Document the tray threading requirement and the canonical setter pattern for contributors and reviewers.
  - User/Stakeholder Impact: New contributors understand macOS main-thread constraints and know the intended pattern for safely updating tray UI.
  - Acceptance Clues: The repository includes a short comment in the tray-related code describing the main-thread constraint and a short README note describing the setter-wrapping pattern.

- Task: Introduce named constants for timing values (for example, the updater cadence and session-duration defaults) and use them consistently.
  - User/Stakeholder Impact: Timing values are easier to find, adjust, and reference in tests and docs.
  - Acceptance Clues: Named constants exist and are referenced where appropriate rather than scattering inline literals.

- Task: Improve test determinism and naming for affected tests by replacing sleep-based synchronization with explicit synchronization (channels/waits) and using intention-revealing test names.
  - User/Stakeholder Impact: CI becomes more reliable; maintainers spend less time debugging flaky tests and can see test intent at a glance.
  - Acceptance Clues: Tests have descriptive names indicating intent (for example, showing update-on-tick behavior), use explicit synchronization instead of arbitrary sleeps, and pass locally in a stable way.

- Task: Apply a small hygiene pass: add doc comments on exported items and rename ambiguous local variables to clearer names.
  - User/Stakeholder Impact: Code is easier to read in reviews and new contributors ramp faster.
  - Acceptance Clues: Exported behaviors touched by this increment have short doc comments; variable names are improved in diffs and reduce cognitive load.

## Risks and Assumptions

- Risk: Changing the subscription API surface may require small adaptations where subscriptions are consumed. Assumption: this demo is the primary consumer and any adapters will be straightforward.
- Risk: If lifecycle stop/teardown is not called correctly in all paths, tests or runtime code may still observe stray updates. Mitigation: require explicit stop/unsubscribe calls in tests and wiring, and include teardown assertions in modified tests.
- Assumption: No new runtime dependencies will be added; all work uses the standard library and existing project patterns.

## Success Criteria and Observability

- Tests: The affected test packages run deterministically and `go test ./...` succeeds without introducing flakiness.
- Code Review: Diffs show an updater abstraction with explicit lifecycle control, an unsubscribe-capable subscription API, clearer timing constants, and renamed tests with explicit synchronization.
- Manual Validation: Running the demo locally still shows the remaining-time label while sessions run and it stops showing after sessions finish or cancellation; starting and stopping the updater does not leave background activity in logs.

## Process Notes

- Implement as a focused, single change set touching only the updater, subscription surface, documentation, and tests. Keep commits small and focused so the change is reversible if necessary.
- Include acceptance notes and a short test plan in the PR: how to run the updated tests and a quick smoke-check to validate the tray label behavior.
- Prefer minimal, well-documented code changes that reviewers can verify quickly; avoid large cross-cutting refactors in this increment.

## Follow-up Increments (optional)

- Add a small integration test that exercises start/stop lifecycle and verifies no stray callbacks on repeated start/stop cycles.
- If this pattern is reused across examples, extract a lightweight, documented helper package for tray lifecycle wiring.

Date: 2025-12-05
