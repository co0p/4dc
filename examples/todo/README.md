# todo – Example

This example walks through the 4dc workflow for a simple CLI todo tool written in bash.

It demonstrates the **design step** inserted between `increment` and `implement`, and shows how `promote` pushes learnings back to `docs/` on main.

---

## Workflow Walkthrough

### Step 1: Constitution

`CONSTITUTION.md` establishes:
- Pure bash, single file, no dependencies
- Fail-fast validation
- File-backed storage at `~/.todo/tasks.csv`
- `docs/` layout for permanent knowledge

### Step 2: Increment

`.4dc/increment.md` was produced for the first feature: **Add Task**.

Key outputs:
- 4 acceptance criteria (add with defaults, add with priority, error on missing title, list)
- 4 greppable test stubs (e.g. `TestAdd_GivenTitle_WhenAdd_ThenStoredWithDefaults`)
- 2 deliverables: storage+add, then list

### Step 3: Design

`.4dc/design.md` was produced before any code was written.

Key outputs:
- **Ubiquitous Language**: Task, Title, Status, Priority, Task Store, Id
- **Bounded Context**: Task Management (single context)
- **Domain Model**: Task aggregate (Mermaid class diagram), Status + Priority as value objects, TaskAdded event
- **C4 diagrams**: Context (user → todo.sh → filesystem), Container (todo.sh + tasks.csv), Component (router, validator, store-rw, id-gen)
- **Design decisions**: CSV storage rationale, sequential id, subcommand routing via case/esac
- **Open questions**: sort order for list, created_at format

### Step 4: Implement (TDD)

Implementation followed the design. Learnings captured in `.4dc/learnings.md`:
- Validate-then-mutate pattern → `docs/DESIGN.md`
- Store-on-demand creation pattern → `docs/DESIGN.md`
- CSV format decision → `docs/adr/`

### Step 5: Promote

Promote pushed learnings back to `docs/`:
- `docs/domain.md` ← domain model from `.4dc/design.md`
- `docs/architecture.md` ← C4 diagrams from `.4dc/design.md`
- `docs/DESIGN.md` ← emergent patterns from TDD

Then `.4dc/` was deleted.

---

## Files in This Example

```
examples/todo/
├── CONSTITUTION.md             # Project foundation
├── README.md                   # This file
├── docs/
│   ├── domain.md               # Promoted DDD model
│   ├── architecture.md         # Promoted C4 diagrams
│   └── DESIGN.md               # Promoted emergent architecture
└── .4dc/                       # Ephemeral (gitignored in real projects)
    ├── increment.md            # Shown here for illustration
    └── design.md               # Shown here for illustration
```

> In a real project, `.4dc/` is gitignored and deleted after merge. It is kept here only to show what the design session produced.

---

## Key Takeaway

Main branch holds the high-level picture (`CONSTITUTION.md`, `docs/`). `.4dc/` holds only the ephemeral working context for the active increment. Promote is the bridge that keeps the two in sync.