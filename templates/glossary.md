# Glossary

Living dictionary of domain terms for this project. Updated during the `promote` phase each cycle.

> Source of truth for shared language. If a term is not here, its meaning is ambiguous — add it.

---

## Format

Each entry follows this pattern:

**Term** — One-sentence definition in domain context.
- *Example:* How it appears in code, CLI output, or tests.
- *Related:* Other terms this one depends on or contrasts with.
- *Added:* Increment slug (e.g. `increment-2026-06-01`)

---

## Terms

### [Term A]

**[Term A]** — [One-sentence definition grounded in this project's domain, not a generic definition].
- *Example:* `[concrete example from code, CLI, or test]`
- *Related:* [Term B], [Term C]
- *Added:* `[increment slug]`

---

### [Term B]

**[Term B]** — [One-sentence definition].
- *Example:* `[example]`
- *Related:* [Term A]
- *Added:* `[increment slug]`

---

## Rules

- Definitions are written from the perspective of this project's domain, not general software engineering.
- When a term's meaning changes, update the definition in place — do not add a new entry.
- If two terms turn out to mean the same thing, pick one and add a redirect: `**Alias** — See [Canonical Term].`
- Remove terms that are no longer part of the codebase.
