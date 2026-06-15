# ADR-20260604 — RunCycle does not delegate to RunSession

## Context

The `internal/ui` package contains two functions for running timer intervals:

- `RunSession(ctx, in, out, label, sess)` — a standalone, single-interval runner. It owns the full user flow: shows the idle screen, waits for Enter or Space, runs the countdown, rings the bell, and returns to the idle screen. It is designed to be used directly (e.g. for a one-shot timer invocation) and has dedicated unit test coverage.

- `RunCycle(ctx, in, out, intervals)` — added in this increment to drive a full 8-interval Pomodoro cycle with automatic advancement between intervals.

During planning, `RunCycle` was expected to call `RunSession` as a per-interval helper. During implementation this proved incorrect.

## Decision

`RunCycle` does **not** call `RunSession`. It drives the countdown loop directly:

```go
for _, iv := range intervals {
    label := intervalLabel(iv)
    for remaining := range iv.Session.Run(ctx) {
        WriteCountdown(out, label, remaining)
    }
    if ctx.Err() != nil {
        return
    }
    WriteBell(out)
}
```

`RunSession` is retained for its standalone use case and existing unit test coverage. The two functions share no code path.

## Rationale

`RunSession` contains idle-message and key-wait logic at its entry point. Calling it from `RunCycle` for each interval would show the idle screen and block for a keypress between every interval — breaking auto-advance entirely. Extracting just the countdown loop from `RunSession` into a shared helper was not done because it would add an abstraction with a single current caller (`RunCycle`), which the Constitution prohibits ("do not create helpers or abstractions for one-time operations").

## Consequences

- Changes to the standalone session flow (e.g. adding a "skip" key or a per-session summary line) only touch `RunSession`.
- Changes to cycle-level behaviour (e.g. adding a skip-interval key during a cycle) only touch `RunCycle`.
- Future contributors must not attempt to unify the two functions without a new increment that explicitly scopes the refactor.
