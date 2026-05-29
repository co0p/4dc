# CONSTITUTION.md

## Engineering Principles
- Ship in small increments with user-visible outcomes.
- Keep implementation simple and explicit.
- Prefer behavior-first tests for critical user flows.
- Use retrospectives to promote durable decisions.

## Architecture Boundaries
- Stack: HTML, CSS, vanilla JavaScript.
- Runtime dependencies: none external.
- Storage boundary: localStorage adapter module only.
- UI logic is separated from persistence logic.

## Quality and Testing
- Minimum checks before merge:
  - Add todo
  - Toggle completion
  - Edit todo
  - Delete todo
  - Filter todo list
  - Reload persistence
- Browser-based testing is required to support a TDD workflow.
- Manual mobile viewport verification required.
- Accessibility baseline: semantic controls, focus visibility, tap targets.

## Delivery Rules
- One increment should stay small and shippable.
- Every increment must define acceptance criteria and completion signals.
- Promote learnings to permanent docs before cleanup.

## Documentation and ADR Policy
- ADRs live in docs/adr/ using ADR-YYYYMMDD-title.md.
- Keep README updated with user-facing behavior changes.
- Keep design notes updated when structure or behavior patterns change.

## Visual Design Prompt
- Visual direction: brutalism-like mobile web UI.
- Look and feel: intense contrast, sharp edges, strong lines, and a muted color palette.
- Layout: high legibility, bold section separation, rectangular surfaces, and minimal rounded corners.
- Typography: clear hierarchy with strong weight changes for headings and labels.
- Components: explicit borders, visible focus states, and high-contrast action elements.
- Motion: minimal and purposeful; prefer instant feedback over decorative transitions.
- Consistency rule: all new screens and components should follow this visual direction unless an approved ADR states otherwise.

## Domain Constraints
- Todo title maximum length is 256 characters.
- Empty titles are rejected.
- Allowed todo states are exactly: `open`, `active`, `completed`.
- Any other state value must fail fast and stop rendering flow until corrected.

## Observability
- Simple logging is sufficient for core activities: add, toggle, edit, delete, filter, and load.
- Keep logs lightweight and local.

## SDLC Artifacts
The project should keep these in sync as needed:
- OpenAPI specification (if public API appears)
- ADRs
- Visual design guide
- Personas
- Deployment strategy
- Testing decisions
- Observability notes
- C4 diagrams (system, container, component)

## Artifact Map for this Example
- Permanent:
  - examples/todo/CONSTITUTION.md
  - examples/todo/docs/
- Transient:
  - examples/todo/.4dc/increment.md
  - examples/todo/.4dc/plan.md
  - examples/todo/.4dc/implementation.md
  - examples/todo/.4dc/promote.md
  - examples/todo/.4dc/*-review.html
