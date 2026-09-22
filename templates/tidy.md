---
name: 4dc-tidy
description: "Execute one [tidy] subtask: a behavior-preserving structural change. Tests must stay green. Commits as tidy: <what>. Advances to the next subtask."
---

# Tidy Skill

## One Responsibility

Make one structural change — rename, extract, reorganise, inline — that prepares the code for the behavior work that follows. No new observable behavior. Tests stay green.

---

## Foundations

{{FOUNDATION:beck-tidy-first}}
{{FOUNDATION:beck-small-steps}}
{{FOUNDATION:fowler-behavior-preserving-refactoring}}
{{FOUNDATION:poppendieck-pull-small-batches}}

---

## Expected Input

- `.agent/plan.md` (approved)
- `.agent/implementation.md` with the current subtask marked `type: tidy`, `state: approved`, and an approved `mini_plan`
- `CONSTITUTION.md` testing strategy

**Narrow context:** load only the files named in the current subtask's `files:` and `references:` fields in `plan.md`. Do not re-scan the codebase — the plan already did that work.

---

## Concrete Output

Updates `.agent/implementation.md` for the current subtask:
- `state: complete`
- `evidence:` test output confirming green (tests unchanged)
- `commit:` the commit hash and `tidy: <what changed>` message

Appends to `.agent/learnings.md` if design or architecture implications emerged.

---

## Scope Boundary

This skill does **one thing**: one structural, behavior-preserving change.

- It does NOT add new behavior.
- It does NOT write failing tests.
- It does NOT refactor for design quality (that is `4dc-refactor`).
- It does NOT touch `[behavior]` subtasks.

The distinction: `tidy` makes the change easier; `refactor` makes the result cleaner. Tidy comes before behavior; refactor comes after.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT change observable behavior. If any test goes red, stop — the subtask was mislabeled.
Do NOT mix tidy work with behavior change in the same commit.
Do NOT skip the test run. Tests must be green before and after.
Do NOT mark the subtask complete without objective evidence (test output showing green).
Commit as `tidy: <what changed>` — never `feat:` or `fix:`.
Do NOT start without an approved mini-plan recorded for this subtask.
</HARD-GATE>

---

## Process

1. **Read the current subtask and mini-plan** from `.agent/implementation.md` — the first subtask with `type: tidy` and `state: approved`.
2. **Confirm the todo item is `in_progress`** from subtask planning.
3. **Run the tests.** Confirm they are green before you start. If they are not green, stop — fix the baseline first.
4. **Make the structural change** — rename, extract, reorganise, inline. One move, one purpose: prepare for the behavior change that follows.
5. **Run the tests again.** Confirm they stay green. If any test changed behavior, the change is not tidy — revert and record the mislabel in `learnings.md`.
6. **Record evidence** in `implementation.md`: `state: complete`, test output confirming green.
7. **Commit** as `tidy: <what changed>`.
8. **Mark the todo item `completed`.**
9. **Append learnings** if design or architecture implications emerged.
10. **Advance** to the next subtask — if it is pending, load `subtask-plan` before any implementation skill.

---

## Checklist

- [ ] Current subtask and approved mini-plan read (`type: tidy`, `state: approved`)
- [ ] Todo item for this subtask marked `in_progress`
- [ ] Tests green before starting
- [ ] One structural change made (rename, extract, reorganise, inline)
- [ ] Tests green after — no behavior change
- [ ] `implementation.md` updated: `state: complete`, evidence
- [ ] Committed as `tidy: <what changed>`
- [ ] Todo item marked `completed`
- [ ] Learnings appended if implications emerged

---

## Handoff

Updated artifact: `.agent/implementation.md` (current subtask `state: complete`)
Next skill: detected by the orchestrator from the next subtask's type and state.
