# Todo Example

This example demonstrates the first shippable slice of the mobile todo app:

- Render todos from bundled sample data.
- Enforce strict state values: `open`, `active`, `completed`.
- Fail fast on invalid state values.
- Provide browser-native behavior tests without external dependencies.

## Files

- `index.html`: app page shell and DOM anchors.
- `sample-data.js`: deterministic seed dataset.
- `app.js`: validation, rendering, and lightweight logging.
- `styles.css`: brutalist visual system for mobile UI.
- `tests.html`: browser test runner page.
- `tests.js`: test suite for rendering, validation, and logging behavior.

## Run Locally

1. Open `index.html` in a browser.
2. Verify the list renders three seeded items.
3. Open `tests.html` in a browser.
4. Confirm all tests pass.

## Guaranteed Behavior in this Increment

- Only states `open`, `active`, and `completed` are accepted.
- Any invalid state throws an explicit error and stops rendering.
- Empty dataset shows an empty-state message.
- A `list-load` log entry is written with the rendered item count.
- Users can add new todo items via form input above the list.
- New todos always start with state `open`.
- New todos persist across page reloads via localStorage.
- Empty or >256 character titles are rejected with error message.
- Input field is cleared after successful add.

## Scope Boundaries

Out of scope for this increment:

- Edit, delete, or toggle todo items.
- Filtering.

These behaviors are intentionally deferred to later increments.
