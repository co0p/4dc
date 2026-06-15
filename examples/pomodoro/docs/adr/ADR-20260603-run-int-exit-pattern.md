# ADR-20260603: run() int exit pattern

## Context

`cmd/pomodoro/main.go` needs to:
1. Restore terminal state (via `defer term.Restore`) before the process exits.
2. Exit with a controlled exit code (0 on clean quit).

`os.Exit` terminates the process immediately — it does **not** run any deferred functions in the calling frame. Placing `defer term.Restore(...)` in `main` and then calling `os.Exit(0)` at the end would silently skip the restore, leaving the terminal in raw mode.

## Decision

Wrap all application logic in a `run() int` function that returns an exit code. `main` contains only:

```go
func main() {
    os.Exit(run())
}
```

All deferred calls (e.g. `term.Restore`) are registered inside `run()`. When `run()` returns, its deferred calls execute in the normal Go defer order before control returns to `main`, which then calls `os.Exit` with the returned code.

## Consequences

- Terminal state is always restored on clean exit paths.
- The pattern is explicit: anyone reading `main` immediately sees that `run()` owns the lifecycle.
- Any future cleanup (e.g. saving session state) follows the same pattern — add a `defer` inside `run()`.
- `run()` is unit-testable independently of `os.Exit`.
