# CONSTITUTION — Pomodoro

> Permanent engineering guardrails. Updated only on explicit revision via the constitution skill.

---

## Engineering Principles

- **Small, verified steps.** No increment ships without a passing test suite. Every PR is a complete, green slice — never a half-built feature.
- **Simplest thing that works.** Prefer stdlib over dependencies. Add a dependency only when the alternative is materially more complex or fragile.
- **Explicit over implicit.** No global state beyond the top-level timer loop. Interval state is always passed; never inferred from side channels.
- **Observable behaviour, not internal structure.** Tests assert what the user sees (countdown output, interval transitions, key handling), not implementation internals.

---

## Architecture & Dependency Direction

| Path | Responsibility | Allowed dependencies |
|------|---------------|----------------------|
| `internal/timer` | Interval logic: countdown, transition, cycle tracking | stdlib only |
| `internal/ui` | Terminal rendering, raw-mode input, bell | `internal/timer` types; `golang.org/x/term` |
| `cmd/pomodoro` | Wires timer + ui; handles OS signals | `internal/timer`, `internal/ui` |

Dependency arrows point inward: `cmd/pomodoro → internal/ui → internal/timer`.  
The `timer` package must never import `ui`.

---

## Testing Strategy & Quality Gates

- **Red → Green → Refactor.** Write the failing test first. No production code before a failing test exists. No refactor before green.
- **Structural changes (tidying) are committed separately from behavioural changes.** A single commit must not mix both.
- **Unit tests** cover `internal/timer`: countdown arithmetic, interval transitions, long-break trigger at 4th Pomodoro, session summary calculation.
- **Integration tests** simulate a full cycle with a fast-clock shim. Assert: correct interval sequence, correct total elapsed time, correct summary line, terminal restore on quit.
- **Signal-handling integration tests** must synchronize on observable output (e.g. read the expected stdout line) before sending the signal. Never use `time.Sleep` as a synchronization mechanism — it is flaky and masks real races.
- **Raw-byte accumulator for newline-less output.** When asserting on terminal output that uses carriage-return (`\r`) in-place updates (no trailing `\n`), do NOT use `bufio.Scanner` — it blocks waiting for a newline that never arrives. Instead, read bytes into a goroutine-fed channel and accumulate them into a string before asserting.
- **Cycle-level tests use a `makeFastInterval` helper.** When testing `RunCycle` or any function that takes `[]timer.Interval`, construct intervals with a pre-filled tick channel via a test-local helper rather than inline. This keeps tests readable without adding production abstractions.
- **Gate:** `go test ./...` must be green before any increment is marked complete.

---

## Tidy First Policy

- Tidy before a feature increment when the existing code makes the increment unclear or error-prone.
- Valid tidying: rename for clarity, extract a function with no behaviour change, reorder declarations, remove dead code.
- Invalid tidying: changing logic under the guise of cleanup, adding generality "for later".
- Every tidy must have a test that was green before and remains green after.

---

## Documentation Rules & ADR Policy

- `README.md` is the sole user-facing doc. Keep it current with every increment.
- Create an ADR under `docs/adr/` for any decision that is: (a) hard to reverse, or (b) non-obvious given the constraints. Template: context → decision → consequences.
- ADR trigger examples: choice of raw-mode library (`golang.org/x/term`), signal handling approach, test clock strategy.
- No ADR needed for routine implementation choices.

---

## SDLC Artifact Expectations

| File | Created by | Lifecycle |
|------|-----------|-----------|
| `CONSTITUTION.md` | constitution skill | Permanent; updated only on explicit revision |
| `.agent/increment.md` | increment skill | Per cycle; deleted after promote |
| `.agent/plan.md` | plan skill | Per cycle; deleted after promote |
| `.agent/implementation.md` | implement skill | Per cycle; deleted after promote |
| `.agent/learnings.md` | implement skill | Per cycle; appended, reviewed at promote |

The `.agent/` directory is gitignored. No `.agent/` file is ever committed.

---

## AI Collaboration Rules

- **One task per context.** Each AI prompt addresses a single, scoped task. Do not batch unrelated changes into one session.
- **Human owns architecture.** The AI may propose; the human decides. Package boundaries, dependency choices, and test strategy are human-approved via the 4dc stop gates.
- **No speculative work.** The AI must not add features, error handling, or abstractions beyond what the current increment requires.
- **Evidence required.** The AI must never claim an increment is complete without showing `go test ./...` output.
- **Preserve optionality.** Avoid coupling decisions early. Defer platform-specific work (e.g., Windows support) until it is in scope.

---

## Platform & Build Constraints

- Targets: macOS (arm64, amd64) and Linux (amd64). Windows is out of scope for v1.
- Build: `go build -o pomodoro ./cmd/pomodoro` produces a single static binary.
- Go version: latest stable, pinned in `go.mod`.
- Third-party dependencies: `golang.org/x/term` is permitted for raw-mode terminal input. Must be locked via `go.sum`. No transitive dependency trees.
