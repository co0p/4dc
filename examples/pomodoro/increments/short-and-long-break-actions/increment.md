Increment: Short and Long Break Actions

## Context

The pomodoro demo is a small, single-binary macOS menu-bar app that currently exposes a single "Break" action (5 minutes) alongside "Pomodoro" (25 minutes) and "Quit". Users sometimes want an extended break option without changing the Pomodoro duration or introducing automatic scheduling rules. The project's constitution emphasizes small, safe, and observable changes with fast tests and minimal operational risk.

This increment focuses on a narrow, user-visible UX change: expose explicit Short Break and Long Break actions so users can choose a 5-minute or 25-minute break directly from the tray.

Related notes within the project: the tray already offers menu items and wiring for start/quit actions; there are small unit tests and a mock tray used for testing UI wiring.

## Goal

Users can start a Pomodoro (25m), a Short Break (5m), or a Long Break (25m) directly from the tray menu. The tray explicitly shows four actions: `Pomodoro`, `Short Break`, `Long Break`, and `Quit`.

Scope:
- Adds explicit Short Break and Long Break actions to the tray UI and wiring.
- Keeps Pomodoro behavior unchanged (25 minutes).

Non-goals:
- This increment will not add automatic long-break scheduling (for example, after N pomodoros).
- It will not add a UI for configuring durations or persistent user preferences.
- It will not introduce telemetry or external services.

Why this is a good increment:
- Small, self-contained, and low-risk change to the tray UI.
- Easy to validate with quick tests and manual checks.
- Aligns with the project's constitution: small changes, fast feedback, and observable behavior.

## Tasks

- Task: Tray exposes four menu actions: `Pomodoro`, `Short Break`, `Long Break`, `Quit`.
  - User/Stakeholder Impact: Users will explicitly see both short and long break choices in the tray and can choose the desired break length directly.
  - Acceptance Clues: Running the demo shows the four items in the tray menu with those labels.

- Task: Starting `Short Break` transitions the app to a break-running state with approximately 5 minutes remaining.
  - User/Stakeholder Impact: Users selecting Short Break get a 5-minute break and immediate visual feedback that a break is active.
  - Acceptance Clues: After selecting Short Break, the app reports a break-running state and Remaining returns a value close to 5 minutes.

- Task: Starting `Long Break` transitions the app to a break-running state with approximately 25 minutes remaining.
  - User/Stakeholder Impact: Users selecting Long Break get a 25-minute break and immediate visual feedback that a long break is active.
  - Acceptance Clues: After selecting Long Break, the app reports a break-running state and Remaining returns a value close to 25 minutes.

- Task: Starting `Pomodoro` preserves current behavior (25 minutes).
  - User/Stakeholder Impact: Existing users experience no regression; Pomodoro remains a 25-minute focused session.
  - Acceptance Clues: Selecting Pomodoro starts a 25-minute session as before and Remaining returns a value close to 25 minutes.

- Task: Add or update tests to cover the new explicit break actions and verify observable state transitions and remaining-time behavior.
  - User/Stakeholder Impact: Maintainers get machine-checkable, fast feedback preventing regressions.
  - Acceptance Clues: New/updated tests run in CI and assert that invoking each break action results in the expected state and approximate remaining durations.

- Task: Add a short note to the README or CHANGELOG describing the new tray actions so users and maintainers are aware of the change.
  - User/Stakeholder Impact: Users and contributors can discover the change without inspecting code or PRs.
  - Acceptance Clues: README or CHANGELOG includes a short entry stating the tray now offers Short Break and Long Break actions.

## Risks and Assumptions

Risks:
- Adding another menu item may slightly increase tray UI complexity; clear labels mitigate confusion.
- Existing tests that assume a single break action may need small updates.

Assumptions:
- Default durations remain unchanged: Pomodoro 25m, Short Break 5m, Long Break 25m.
- No user-configurable durations or persistence are included in this increment.
- The project’s current test and CI setup will validate the change.

Mitigations (high level):
- Use clear labels for menu items (e.g., "Short Break (5m)") and update tests accordingly.

## Success Criteria and Observability

- The running app's tray contains four menu items: `Pomodoro`, `Short Break`, `Long Break`, `Quit`.
- Selecting `Short Break` results in a break-running state and Remaining ≈ 5m.
- Selecting `Long Break` results in a break-running state and Remaining ≈ 25m.
- Selecting `Pomodoro` continues to start a 25-minute session.
- Automated tests that assert these behaviors pass in CI.
- README/CHANGELOG includes a short note describing the user-visible change.

Checks to perform after merge:
- Manual: run the app and verify the tray displays the four items and each action behaves as expected.
- Automated: run the test suite and ensure new tests pass quickly.

## Process Notes

- Implement the change as a small, focused PR that updates the tray labels and wiring and adds tests.
- Keep changes minimal and reversible; include a test that demonstrates the new menu items and their behavior.
- Follow the normal review and CI process; no special rollout or coordination is required.

## Follow-up Increments (Optional)

- Add a simple settings mechanism to let users configure break durations.
- Introduce an optional automatic long-break scheduler (for example, long break after N pomodoros) as an opt-in feature.
- Improve tray title formatting to show remaining time for active sessions.
