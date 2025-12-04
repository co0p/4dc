# Increment: Minimal Demo UI — Start/Stop + Quit + Icon

## Context

- Brief description: The `examples/pomodoro` app is a small macOS menu-bar Pomodoro timer implemented as a single Go binary. The demo currently lacks a complete, user-facing tray menu and a minimal automated check that the binary builds correctly.
- Important background:
  - Existing behavior: The project is a lightweight tray app with focus on simplicity and local operation (see `CONSTITUTION.md`).
  - Constraints: Must remain a single distributable binary, avoid scope-expanding dependencies, and keep changes small and reversible.
  - Links: `README.md` and `CONSTITUTION.md` in the project root describe purpose and constraints.
- Key assumptions: Use a simple monochrome icon for the tray; add one fast unit test for timer core; add a CI smoke check that validates the built binary responds to `--version`/`--help`.

## Goal

- Outcome / hypothesis:
  - Users can interact with a simple tray menu that provides Start/Stop and Quit actions and see a clear icon. Maintainers have a small unit test protecting core timer logic and a CI smoke check that confirms the binary builds and runs `--version`/`--help`.
- Scope:
  - Add Start/Stop switch and a Quit action to the tray menu.
  - Provide a simple monochrome icon used in the tray.
  - Add one unit test for timer start/stop behavior and a CI smoke check that validates the binary.
- Non-goals:
  - Not adding analytics, crash-reporting, cloud sync, or full multi-resolution `icns` packaging.
  - Not achieving full project-wide test coverage in this increment.
  - Not introducing large new dependencies or external services.
- Why this is a good, small increment:
  - It completes a minimal user flow (start/stop/quit) that makes the demo usable and testable.
  - It is small, reversible, and verifiable with a quick manual check and simple automated tests, fitting the project's constitution.

## Tasks

- Task: Tray menu exposes two actions: Start and Stop (toggle behavior)
  - User/Stakeholder Impact: End users can start and stop the timer from the menu bar without opening additional windows.
  - Acceptance Clues: Selecting Start changes the menu to show Stop; selecting Stop halts the timer; observable timer state is consistent with menu label.

- Task: Tray menu includes a Quit action that cleanly exits the app
  - User/Stakeholder Impact: Users can reliably quit the app and not leave background processes running.
  - Acceptance Clues: Selecting Quit terminates the process and no background process remains for the app; on next launch the app starts fresh or restores documented minimal state.

- Task: Add a simple monochrome icon for the tray
  - User/Stakeholder Impact: Users can identify the app in the menu bar; improves discoverability and polish of the demo.
  - Acceptance Clues: When the built binary runs on macOS (or a dev environment), the tray shows the provided icon or an acceptable placeholder in development.

- Task: Add one unit test covering core timer start/stop behavior
  - User/Stakeholder Impact: Maintainers gain automated protection against regressions in core timer logic.
  - Acceptance Clues: The test passes locally and in CI; it runs quickly (< 1s) and verifies that starting advances state and stopping halts time progression in the core timer abstraction.

- Task: CI smoke check validates build and that the binary responds to `--version` or `--help`
  - User/Stakeholder Impact: Maintainers get immediate build feedback and confidence the artifact runs.
  - Acceptance Clues: CI job completes the build step and the smoke step successfully prints version/help output and exits with status 0.

## Risks & Assumptions

- Known risks:
  - UI semantics may vary across macOS versions; tray APIs can behave differently in older OS versions.
  - Adding icon asset may require minor layout/testing adjustments on retina displays; however multi-resolution packaging is out of scope for this increment.
- Key assumptions:
  - A simple monochrome PNG/SVG is sufficient for demonstration and development.
  - The unit test will target a timer abstraction (not OS-specific UI code) so it remains fast and stable.
- Mitigations:
  - Keep UI changes minimal and rely on a well-scoped timer abstraction for tests.
  - Document any platform notes and test Quit behavior manually on a representative macOS runner.

## Success Criteria & Observability

- How we will know this increment is successful:
  - Manual checks: On a dev macOS machine the tray menu shows the icon, Start/Stop toggles the timer state, and Quit exits cleanly.
  - Automated checks: Unit test passes and CI smoke check (build + `--version`/`--help`) succeeds.
- What we will observe after release:
  - CI logs show successful build and smoke step.
  - Test results show green for the new unit test.
  - No immediate crash reports or widespread regressions from manual testers.

## Process Notes

- Implementation should be done as small, focused changes that are easy to review and revert if necessary.
- Changes should pass the local test and smoke checks before pushing; CI will validate the build artifact and test.
- Rollback: keep builds and tags for quick rollback if a regression is discovered.

## Follow-up Increments (Optional)

- Add a lightweight packaging workflow to generate an `icns` and `dmg` for macOS, optionally with notarization steps.
- Extend automated tests to include UI screenshot tests or platform-specific smoke tests in a macOS runner.
- Improve test coverage across the codebase to meet the constitution's configurable minimum for core functionality.
