# Observability Notes

## Scope

This example uses lightweight local logging to verify key UI flow events during development.

## Event Contract

Current required events:

- `list-load`
  - Emitted after render path completes.
  - Payload includes `count` with rendered item total.

- `todo-add`
  - Emitted after a new todo is successfully added.
  - Payload includes:
    - `id`: Generated unique ID
    - `title`: Validated todo title
    - `state`: Always "open" for new todos

## Storage Model

- Logs are stored in `window.__todoLog` as an array.
- Each entry contains:
  - `type`
  - `timestamp` (ISO string)
  - `payload`

## Why This Is Durable

- Supports quick behavioral checks during manual and browser-test runs.
- Keeps observability simple while app scope is read-only.

## Guardrails

- Keep payloads compact and non-sensitive.
- Add new event types only when tied to explicit acceptance criteria.
- Treat log shape as a test-visible contract once asserted in tests.
