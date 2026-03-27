# Increment: Add Task

**Type:** feature
**Status:** In progress

---

## User Story

As a user, I want to add a task with a title and optional priority, so that I can track what I need to do.

---

## Context

This is the first increment. There is no existing functionality. The goal is to get the most basic happy path working: add a task, see it stored.

---

## Acceptance Criteria

- [ ] Given a title, when I run `todo.sh add "Buy milk"`, then a task is stored with status `open` and priority `normal`
  → `TestAdd_GivenTitle_WhenAdd_ThenStoredWithDefaults`
- [ ] Given a title and priority flag, when I run `todo.sh add "Fix bug" --priority high`, then the task is stored with priority `high`
  → `TestAdd_GivenTitleAndPriority_WhenAdd_ThenStoredWithPriority`
- [ ] Given no title, when I run `todo.sh add`, then the command exits with a non-zero code and prints an error
  → `TestAdd_GivenNoTitle_WhenAdd_ThenNonZeroExit`
- [ ] Given a stored task, when I run `todo.sh list`, then the task appears in output with its id, status, priority, and title
  → `TestList_GivenStoredTask_WhenList_ThenTaskAppears`

---


## Use Case

**Actor:** User (CLI)
**Preconditions:** `todo.sh` is executable. `~/.todo/` may or may not exist.

**Main Flow:**
1. User runs `todo.sh add "Buy milk"`
2. Tool creates `~/.todo/tasks.csv` if it doesn't exist
3. Tool assigns next available id
4. Tool writes row: `id,open,normal,Buy milk,<timestamp>`
5. Tool prints: `Added task #1: Buy milk`

**Alternate Flows:**
- If `--priority high|normal|low` is provided, use that value instead of `normal`
- If no title argument, print `Error: title is required` and exit 1

**Postconditions:** Task row exists in `~/.todo/tasks.csv`

---

## Deliverables

### Deliverable 1: Storage and add command *(Walking Skeleton)*
- **Provides:** Thin end-to-end path: `add` writes a row, `list` returns it—proving CSV storage and command routing work together
- **Criteria:** AC 1, AC 2, AC 3
- **Shippable:** Can add tasks; errors on bad input
- **Status:** Not started
- **Success tests:**
  - `TestAdd_GivenTitle_WhenAdd_ThenStoredWithDefaults`
  - `TestAdd_GivenTitleAndPriority_WhenAdd_ThenStoredWithPriority`
  - `TestAdd_GivenNoTitle_WhenAdd_ThenNonZeroExit`

### Deliverable 2: List command
- **Provides:** Readable output for stored tasks
- **Criteria:** AC 4
- **Shippable:** Can add and list tasks end-to-end
- **Status:** Not started
- **Success tests:**
  - `TestList_GivenStoredTask_WhenList_ThenTaskAppears`

---

## Out of Scope

- Edit, complete, delete subcommands (separate increments)
- Filtering or sorting (separate increment)
- Categories or tags (separate increment)

---

## Promotion Checklist

- [ ] Storage format worth documenting in CONSTITUTION.md?
- [ ] Any design that emerged worth capturing in docs/DESIGN.md?
- [ ] ADR needed for CSV format choice?
