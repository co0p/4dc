# Design (Emergent)

> This document reflects architecture that emerged through TDD.
> Updated during the promote phase. Not a planning document.

---

## Current Structure

Single file: `todo.sh`

```
todo.sh
  ├── main (entry point, subcommand router)
  ├── cmd_add()     — validate + write new task row
  ├── cmd_list()    — read and format all tasks
  ├── next_id()     — derive next sequential id from store
  └── ensure_store() — create ~/.todo/ and tasks.csv if absent
```

### todo.sh
- **Purpose:** Everything. Single-file per CONSTITUTION.md simplicity principle.
- **Emerged from:** Add Task increment, Deliverable 1 (store + add command).
- **Key patterns:** Subcommand routing via `case/esac`; validation before mutation; store creation on first use.

---

## Patterns Discovered

### Validate-then-mutate
- **What:** All input validation happens before any file write. If validation fails, exit immediately.
- **Why it emerged:** First failing test was `TestAdd_GivenNoTitle_WhenAdd_ThenNonZeroExit`—the simplest implementation that passed it was: check args, exit 1 early.
- **Where used:** `cmd_add()`

### Store-on-demand creation
- **What:** `~/.todo/tasks.csv` is created (with header) only when the first `add` is run, not on install.
- **Why it emerged:** No install step in CONSTITUTION.md. The test for `TestAdd_GivenTitle_WhenAdd_ThenStoredWithDefaults` drove this: the store didn't exist at test start.
- **Where used:** `ensure_store()`

---

## Open Questions

- [ ] Should `cmd_list` sort output? By what field?—not driven by a test yet.

---

## History

| Date | Increment | Changes |
|------|-----------|---------|
| 2026-03-27 | Add Task | Initial structure: cmd_add, cmd_list, next_id, ensure_store |
