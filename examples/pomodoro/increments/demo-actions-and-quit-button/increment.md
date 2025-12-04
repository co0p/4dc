Increment: Demo tray actions, quit button, and minimal icon

## Context

The `pomodoro` example is a small, macOS menu-bar Pomodoro timer that prioritizes minimalism and a single-binary distribution. The app currently provides the core timer behavior but lacks a small, discoverable tray surface suitable for demos and manual testing.

This increment adds a tiny, demo-focused UI surface: two explicit tray/menu actions (buttons), a visible Quit action that cleanly exits the app, and a minimal project-related icon. The intent is to provide a boilerplate that is easy to review, run locally, and iterate on, while respecting the project's constitution (keep changes small, avoid heavy dependencies, and keep the single-binary constraint in mind).

Key constraints and background:
- Keep the change minimal and reversible; no telemetry, no cloud services, and no heavy GUI frameworks.
- The icon should be simple and minimal (a small glyph or basic tomato motif) to avoid large asset or packaging changes.
- Changes should pass the project's normal build and smoke checks and be safe to roll back.

## Goal

Users can interact with two clearly labeled tray actions (buttons) and a Quit action from the app’s tray/menu; the Quit action exits the app gracefully. A simple, minimal icon representing the project appears in the tray so the app is recognizable during demos.

Scope (what this increment will do):
- Add two visible tray/menu actions (implemented as buttons) that perform their labelled action when clicked.
- Add a Quit action that cleanly terminates the application.
- Bundle a minimal, project-related icon for the tray UI.
- Add short, discoverable labels or tooltips for each action so demo users understand their purpose.

Non-goals (explicitly out of scope):
- This increment will not redesign the overall UI, add telemetry, introduce external services, or add complex animations or polished artwork.
- It will not change the packaging model or introduce heavy native bindings that materially increase binary size without an ADR.

Why this is a good increment:
- Small and focused: delivers a visible, testable UI surface in one delivery.
- Low risk: UI affordances are straightforward and reversible.
- Fast feedback: maintainers can validate behavior manually in staging or locally and iterate quickly.

## Tasks

- Task: Provide two visible tray/menu action buttons that perform their labeled action when clicked.
  - User/Stakeholder Impact: End users and demo participants have clear, clickable controls for common timer behaviors during demos and manual testing.
  - Acceptance Clues: Both action labels are visible in the tray/menu and clicking each produces the expected, observable effect in the UI (for example, the timer starts or toggles visibly); there are no hidden keyboard-only actions required to exercise them.

- Task: Provide a visible Quit action that gracefully shuts down the app.
  - User/Stakeholder Impact: Demo participants and maintainers can exit the app reliably without needing to kill the process manually.
  - Acceptance Clues: Selecting Quit exits the app cleanly (the process terminates, the tray icon disappears, and no crash/error logs are produced during normal shutdown); local state is left in a stable state.

- Task: Include a minimal, project-related icon visible in the tray.
  - User/Stakeholder Impact: The app is recognizable in the menu bar during demos and testing; maintainers have a simple starting asset to iterate on later.
  - Acceptance Clues: The tray shows a minimal icon (for example, a simple tomato/glyph) instead of a generic placeholder when the app is running.

- Task: Add short, discoverable labels or tooltips for the actions and Quit so demo users understand their purpose.
  - User/Stakeholder Impact: Demo participants and developers can quickly understand available actions without reading docs.
  - Acceptance Clues: Opening the tray menu or hovering reveals labels/tooltips for each action and Quit.

- Task: Document the change briefly in the project README or an increment note so reviewers know this is a demo boilerplate intended for iteration.
  - User/Stakeholder Impact: Maintainers and reviewers understand the intent and scope, reducing scope creep during reviews.
  - Acceptance Clues: A short one-paragraph note is present in the README or increment documentation describing what to test and the demo nature of the change.

## Risks and Assumptions

- Risks:
  - Introducing new UI code or icon assets could slightly increase binary size or require small platform-specific packaging changes; keep assets minimal to mitigate this.
  - Incorrectly handled shutdown could leave state inconsistent; acceptance must confirm clean termination.

- Assumptions:
  - The two actions are simple, immediate UI actions (buttons) and do not require long-running background workflows.
  - The work can be accomplished without adding heavy external dependencies that violate the constitution.

Mitigation (high-level): If any binary-size or dependency concerns appear during review, revert the asset or defer to a follow-up increment and document the trade-off in the PR description.

## Success Criteria and Observability

- Success Criteria:
  - The tray UI shows two functioning action buttons, a Quit action that exits cleanly, and a minimal icon.
  - The behavior can be validated by a quick local demo run and manual checks in staging.

- What to observe after release:
  - Manual verification in staging or locally: check the tray icon is visible, both action labels are present, clicking actions triggers visible behavior, and Quit terminates the process with no crash logs.
  - A short note in the README points to how to exercise these flows for reviewers and testers.

## Process Notes

- Keep the change small and go through the normal build + lint + smoke verification flow.
- Make a single, focused PR with a short description and the acceptance clues above to keep review time low.
- Rollback is straightforward: revert the demo UI changes if issues appear.

## Follow-up Increments (optional)

- Add keyboard shortcuts for the two actions and document them.
- Replace the minimal icon with a polished, retina-ready asset set and include packaging notes if needed.
- Add lightweight automated UI smoke tests that exercise the tray actions and Quit behavior if feasible without adding heavy test infra.
