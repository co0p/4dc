# Architecture

> Promoted from `.4dc/design.md` during the promote phase.
> Updated after each increment that changes the system structure.

---

## Context

```mermaid
C4Context
    title System Context – todo CLI
    Person(user, "User", "Runs commands in a terminal")
    System(todo, "todo CLI", "Manages a personal task list via bash commands")
    SystemDb_Ext(fs, "File System", "~/.todo/tasks.csv — plain text storage")
    Rel(user, todo, "Runs commands", "CLI")
    Rel(todo, fs, "Reads and writes", "POSIX file I/O")
```

---

## Container

```mermaid
C4Container
    title Container Diagram – todo CLI
    Person(user, "User")
    Container(script, "todo.sh", "Bash", "Single-file CLI. Parses subcommands, validates input, reads/writes task store.")
    ContainerDb(store, "tasks.csv", "Plain CSV", "~/.todo/tasks.csv. Rows: id,status,priority,title,created_at")
    Rel(user, script, "Invokes", "CLI args")
    Rel(script, store, "Reads/appends", "File I/O")
```

---

## Component

```mermaid
C4Component
    title Component Diagram – todo.sh
    Component(router, "Subcommand Router", "Bash case/esac", "Dispatches to the correct handler based on first argument")
    Component(validator, "Input Validator", "Bash functions", "Validates required arguments; exits 1 on invalid input")
    Component(store_rw, "Store Reader/Writer", "Bash + awk", "Creates store file if absent; appends rows; reads and formats rows")
    Component(id_gen, "Id Generator", "Bash + awk", "Reads max id from store; returns next sequential id")
    Rel(router, validator, "Calls before mutation")
    Rel(router, store_rw, "Calls after validation")
    Rel(store_rw, id_gen, "Calls on add")
```

---

## History

| Date | Increment | Changes |
|------|-----------|---------|
| 2026-03-27 | Add Task | Initial C4 diagrams: context, container, component |
