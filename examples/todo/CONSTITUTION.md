# CONSTITUTION.md

> Architectural decisions for the `todo` CLI tool.
> Generated during the constitution phase. All other prompts defer to this.

---

## Architectural Decisions

### Simplicity Over Flexibility
- Pure bash. No dependencies beyond standard POSIX tools (`awk`, `grep`, `sed`, `date`).
- No abstractions until there is concrete, repeated pain.
- One file: `todo.sh`. No splitting unless the file becomes genuinely hard to navigate.

### Fail Fast, Fail Loud
- Invalid input aborts immediately with a clear error message and non-zero exit code.
- No silent failures, no partial state.
- Validate all user input at the entry point before any mutations.

### File-Backed Storage
- Tasks stored in a plain text file: `~/.todo/tasks.csv`.
- Format: `id,status,priority,title,created_at`
- File is human-readable and editable outside the tool.
- No locking; single-user tool, no concurrent access concern.

### Ship Fast and Iterate
- Get the core working first. Optimise based on real usage friction.
- Prefer readable code over clever code.

---

## Testing Expectations

- Manual testing is sufficient for this scope.
- If automated tests are added, they must run without installing external tools.
- Test by invoking `todo.sh` directly and asserting stdout/exit code.

---

## Artifact Layout

```
todo/
├── CONSTITUTION.md        # This file
├── todo.sh                # Single-file implementation
├── README.md              # Usage and examples
├── docs/
│   ├── domain.md          # DDD model (language, contexts, aggregates)
│   ├── architecture.md    # C4 diagrams
│   ├── DESIGN.md          # Emergent architecture from TDD
│   └── adr/               # Architecture Decision Records
└── .4dc/                  # Ephemeral increment context (gitignored)
    ├── increment.md
    ├── design.md
    └── learnings.md
```

---

## Delivery Practices

- Each increment is small enough to complete in one session.
- Before merging: run promote to update `docs/` with learnings.
- Delete `.4dc/` after every merge.
