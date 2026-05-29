# Testing Decisions

## Decision Summary

Use a browser-native test harness (`tests.html` + `tests.js`) with no external test framework for early frontend increments.

## Why

- Zero dependency overhead.
- Easy local execution by opening one HTML file.
- Direct verification against real DOM rendering behavior.

## Minimum Checks for This Example

1. Renders all sample todos.
2. Renders correct state labels for allowed states.
3. Shows empty state when data is empty.
4. Fails fast on invalid state.
5. Writes a `list-load` log entry with item count.
6. Storage module saves and loads todos from localStorage.
7. addTodo rejects empty titles.
8. addTodo rejects titles over 256 characters.
9. addTodo creates todos with `open` state.
10. Added todos persist across page reload.

## Execution

- Open `tests.html` in a browser.
- Confirm summary reports all checks passing.
- Investigate failures by reading line-level assertion messages in the test output list.

## Test Isolation Strategy

To prevent state pollution between tests:

- Clear localStorage in `resetFixture()` before each test
- Ensures deterministic test execution regardless of test order
- Each test starts with clean global state

Pattern:

```javascript
function resetFixture() {
  localStorage.clear();
  // ... other cleanup
}
```

## Forward Usage

- Keep this pattern for small static increments.
- Reassess and formalize a larger test stack only when complexity or cross-module integration requires it.
