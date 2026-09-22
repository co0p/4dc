---
name: 4dc-refactor
description: "Improve the design of the code that just passed its test, without changing behavior. Tests must stay green throughout. Commits as refactor: <what>. Advances to the next subtask."
---

# Refactor Skill

## One Responsibility

Take the code that just made the current test case green and improve its design — without changing behavior. Tests stay green throughout. Then record the case as complete and either activate the next case in the same subtask or advance.

---

## Foundations

{{FOUNDATION:fowler-behavior-preserving-refactoring}}
{{FOUNDATION:fowler-two-hats}}
{{FOUNDATION:beck-refactor-pass}}
{{FOUNDATION:poppendieck-eliminate-waste}}

---

## Expected Input

- `.agent/plan.md` (approved)
- `.agent/implementation.md` with the current `[behavior]` subtask in `state: in-progress` and its `active_test` in `state: green` (test passes, minimal code written, not yet refactored)
- `CONSTITUTION.md` testing strategy

**Narrow context:** load only the files named in the current subtask's `files:` and `references:` fields in `plan.md`. Do not re-scan the codebase — the plan already did that work.

---

## Concrete Output

Updates `.agent/implementation.md` for the current subtask:
- The active test case: `state: complete`
- `evidence:` test output confirming green after refactoring (or a note that no refactoring was needed)
- `refactor:` the commit hash and `refactor: <what changed>` message (omitted if no changes were made)
- If unfinished test cases remain: set `active_test` to the next pending case and leave the subtask `state: in-progress`
- If all test cases are complete: set the behavior subtask `state: complete`

Appends to `.agent/learnings.md` if design decisions or promote candidates emerged.

---

## Scope Boundary

This skill does **one thing**: improve the design of code that already passes its test.

- It does NOT add new behavior.
- It does NOT write new tests.
- It does NOT change what the code does — only how it is organized.
- It does NOT handle `[tidy]` subtasks.

The distinction: `tidy` prepares structure before behavior; `refactor` improves design after behavior. Both are behavior-preserving, but they serve different moments in the cycle.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT change observable behavior. If any test goes red, you changed behavior — revert and try again.
Do NOT add new tests or new production features. This is the refactor hat, not the behavior hat.
Do NOT skip the question "can the design be improved?" — even if the answer is no, the question must be asked and recorded.
Do NOT mark the subtask complete until every planned test case has completed Red → Green → Refactor and all tests are green.
Commit as `refactor: <what changed>` — never `feat:` or `fix:`. If no refactoring was needed, skip the commit and record that decision.
</HARD-GATE>

---

## Process

1. **Read the current subtask** from `.agent/implementation.md` — the `[behavior]` subtask in `state: in-progress` with an active test case in `state: green`.
2. **Run the tests.** Confirm they are green before you start.
3. **Ask the refactor question:** can the design be improved without changing behavior? Look for: duplication, unclear naming, long methods, deep nesting, missing abstraction, poor separation of concerns.
4. **If no improvement is needed:** record `evidence: no refactoring needed — design is sufficient` for the active test case and continue.
5. **If improvement is needed:** make the change. One move at a time — extract method, rename, inline, move. Run the tests after each move. If they go red, revert — you changed behavior.
6. **Run the full test suite** required by the constitution.
7. **Record evidence** in `implementation.md`: set the active test case `state: complete`, with test output confirming green and what was refactored.
8. **Commit** as `refactor: <what changed>`.
9. **Append learnings** if design decisions or promote candidates emerged.
10. **Advance within the subtask:** if a pending test case remains, set it as `active_test` with `state: pending`; the approved subtask mini-plan remains valid and `tdd-red` starts its cycle automatically without another user gate. If all cases are complete, set the behavior subtask `state: complete`, **mark the todo item `completed`**, and route the next pending subtask through `subtask-plan`.

### When all subtasks are complete

- Run **final verification** — the full test suite or constitution-defined release gate. Confirm all acceptance criteria from `increment.md` are met.
- Run every required acceptance test from `.agent/plan.md`. Record each result as `passed`, `failed`, `skipped`, or `not-applicable`; anything other than `passed` blocks completion unless the user explicitly approves and records an exception.
- Present the final evidence and remaining risks. Wait for explicit approval.
- On approval, set `implementation.md` top-level `status: complete`. The next skill is `4dc-promote`.

---

## Checklist

### Per subtask
- [ ] Tests green before starting
- [ ] Refactor question asked (duplication, naming, structure, separation)
- [ ] If refactored: tests stayed green after each move
- [ ] If no refactoring needed: decision recorded
- [ ] Active test case marked `state: complete` with evidence
- [ ] Next pending test case activated, or behavior subtask marked complete and todo item marked `completed` when all cases are complete
- [ ] Committed as `refactor: <what changed>` (or commit skipped with recorded reason)
- [ ] Learnings appended if decisions or candidates emerged

### When all subtasks complete
- [ ] Final verification run (full suite or release gate)
- [ ] All acceptance criteria from `increment.md` met
- [ ] Required acceptance tests recorded and passing, or an explicit user-approved exception recorded
- [ ] User approval received
- [ ] `implementation.md` top-level `status` set to `complete`
- [ ] `learnings.md` has promote candidates listed

---

## Handoff

Updated artifacts: `.agent/implementation.md` (current subtask `state: complete`) + `.agent/learnings.md`
Next skill (if subtasks remain): detected by the orchestrator from the next subtask's type and state
Next skill (if all complete): `4dc-promote` — load `skills/promote/SKILL.md`
