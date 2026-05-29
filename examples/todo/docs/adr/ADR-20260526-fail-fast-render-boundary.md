# ADR-20260526: Fail-Fast Validation at Render Boundary

- Status: Accepted
- Date: 2026-05-26

## Context

The todo example receives a list of items before rendering. Invalid state values can silently corrupt UI behavior unless validated immediately.

Domain rule for this increment: allowed states are exactly `open`, `active`, `completed`.

## Decision

Validate input todo items at the render boundary and fail fast with explicit errors when constraints are violated.

Validation includes:

- Input is an array.
- Each todo is an object with string `id`.
- `title` is non-empty and maximum 256 characters.
- `state` must be one of `open`, `active`, `completed`.

On violation, throw an error and stop rendering flow.

## Consequences

Positive:

- Prevents silent data drift.
- Makes domain violations immediate and visible.
- Keeps tests deterministic by enforcing strict preconditions.

Trade-off:

- Rendering is intentionally strict and will not attempt recovery from malformed data in this increment.

## Evidence

- Implementation: `app.js` validation and throw path.
- Tests: `tests.js` includes explicit invalid-state fail-fast check.
- Increment evidence: `.4dc/implementation.md` reports passing verification.
