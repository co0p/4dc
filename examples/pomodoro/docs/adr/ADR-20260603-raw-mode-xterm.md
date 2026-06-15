# ADR-20260603: golang.org/x/term for raw-mode terminal input

## Context

The Pomodoro binary requires:
- Putting the terminal into raw mode (so single keypresses are read without waiting for Enter).
- Restoring the terminal to its original state on exit.

These operations differ between macOS and Linux at the syscall level (`termios` structures and `ioctl` constants vary). Implementing this directly via `syscall` or `golang.org/x/sys/unix` requires platform-specific build tags and non-trivial boilerplate.

## Alternatives Considered

| Option | Pros | Cons |
|--------|------|------|
| `golang.org/x/term` | Abstracts macOS/Linux; `MakeRaw`, `Restore`, `IsTerminal` built-in; only `golang.org/x/sys` as transitive dep | Not stdlib |
| `syscall` / `unix` directly | Zero extra dependency | Platform-specific build tags; more code; easy to get wrong |
| Third-party TUI library (e.g. bubbletea) | Full UI framework | Far exceeds scope; large dependency tree; conflicts with "simplest thing" principle |

## Decision

Use `golang.org/x/term`. It is maintained by the Go team, has a single transitive dependency (`golang.org/x/sys`), and provides exactly the three functions needed: `IsTerminal`, `MakeRaw`, `Restore`.

## Consequences

- Cross-platform raw-mode handling with minimal code.
- Dependency is locked in `go.sum`; no transitive dependency surprises.
- If stdlib ever absorbs `golang.org/x/term`, migration is a one-line import change.
