# ADR: Receiver naming and function documentation conventions (Ross Cox style)

## Context

Go code style benefits from small, consistent choices about naming and documentation. Two recurring questions in reviews are:

- How should pointer versus value receivers be chosen and named?
- How should exported functions, types, and methods be documented so package docs are useful and idiomatic?

Ross Cox (and others in the Go community) have advocated minimal, consistent receiver names and clear godoc-style comments that start with the symbol name and are full sentences. This ADR captures those recommendations for this repository so reviewers and contributors have a shared, lightweight rule-set.

## Decision

Adopt the following conventions for Go code in this repository:

- Receiver choice
  - Use a pointer receiver when the method needs to modify the receiver's state, when the type contains a mutex or other synchronization primitives, or when the type is large enough that copying is undesirable.
  - Use a value receiver when the type is small, immutable, or when copying is cheap and the method does not modify state.
  - Be consistent within a named type: prefer all methods on a type to use the same receiver kind (pointer vs value) unless there is a strong reason to mix.

- Receiver naming
  - Use short, typically one-letter receiver names for exported types (e.g., `s *Server`, `r *Runner`, `t *Timer`).
  - Prefer the first letter of the type name (e.g., `t` for `Timer`, `s` for `Server`). If the first letter is ambiguous or commonly used, use a short, mnemonic two-letter name (e.g., `sr` for `sessionRunner`) rather than a long name.
  - Avoid stuttering: do not repeat the type name in method names to make up for a long receiver name (e.g., prefer `(s *Server) Start` over `(server *Server) ServerStart`).
  - When the receiver represents a conceptual role, choose a meaningful short name (e.g., `db *DB`, `cfg *Config`). The goal is local clarity in method bodies, not long descriptive names.

- Documentation style (godoc)
  - All exported functions, methods, and types must have a comment immediately preceding the declaration.
  - The comment must be a full sentence that begins with the name of the symbol being documented (e.g., `Start starts the server...`).
  - Keep comments concise and focused: document the behavior, important side effects, concurrency/goroutine-safety expectations, and error conditions. Don’t restate obvious parameter names or return types.
  - For complex behaviors or invariants, include short examples or cross-references to package-level documentation or an ADR.
  - Prefer documenting the contract (what the caller can expect) over implementation details. Implementation notes may go into a `// NOTE:` comment in the body if necessary.

## Consequences

- Benefits
  - Reviews focus on behavior rather than stylistic nitpicks; small receiver names are familiar to Go programmers and reduce visual noise.
  - Consistent documentation improves the generated package godoc and makes APIs easier to understand for new contributors.
  - Using pointer receivers where appropriate avoids subtle bugs (missing mutations) and performance surprises when types are large.

- Drawbacks / Trade-offs
  - Single-letter receiver names are less descriptive in very large methods; choose clarity over brevity when a method's logic benefits from a slightly longer but still short name.
  - Contributors unfamiliar with these conventions may need a short orientation, but the rules are small and straightforward.

## Alternatives Considered

- Always use descriptive receiver names (e.g., `server *Server`) — rejected because Go idiom favors short receiver names and descriptive names add noise in many method bodies.
- Always use pointer receivers — rejected because small, copyable value types (e.g., small structs) are safe and sometimes more convenient as value receivers.
- Omit godoc comments for exported symbols — rejected because package-level documentation is valuable for users and maintainers.

## Examples

```go
// Good: short receiver, pointer used because Start modifies state.
func (s *Server) Start() error {
    s.running = true
    // ...
}

// Good: comment starts with the symbol name.
// Start starts the server and returns an error if the server cannot be started.
func (s *Server) Start() error { ... }

// Good: value receiver for a small immutable type.
func (t Timer) Duration() time.Duration { return t.d }
```

## References

- Ross Cox — "Go comment and naming guidance" (community posts and reviews)
- Effective Go and the Go code review comments: https://golang.org/doc/effective_go.html
- Go project code-review conventions and examples in the standard library


---

This ADR is intentionally short and pragmatic. If maintainers want to tighten or relax any rule (for example, by standardizing two-letter receivers for certain subsystems) we can revise this ADR in a follow-up.