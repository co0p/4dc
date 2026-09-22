---
name: 4dc-tdd-red
description: "Write exactly one failing test for the current [behavior] subtask. Confirm it fails for the right reason. Do not write production code. Do not refactor."
---

# TDD Red Skill

## One Responsibility

Write one failing test for the current `active_test` in the `[behavior]` subtask and confirm it fails for the right reason. Then stop. A subtask may contain several cohesive test cases; this skill advances one case at a time.

---

## Foundations

{{FOUNDATION:beck-red}}
{{FOUNDATION:jeffries-card-conversation-confirmation}}
{{FOUNDATION:freeman-pryce-tests-guide-design}}

---

## Expected Input

- `.agent/plan.md` (approved)
- `.agent/implementation.md` with the current subtask marked `state: approved`, `type: behavior`, and an approved `mini_plan`
- `CONSTITUTION.md` testing strategy (test depth, naming, isolation conventions)

**Narrow context:** load only the files named in the current subtask's `files:` and `references:` fields in `plan.md`. Do not re-scan the codebase — the plan already did that work.

---

## Concrete Output

Updates `.agent/implementation.md` for the current subtask:
- The subtask: `state: in-progress` (transitions from `approved`)
- The active test case: `state: red`
- `test:` the failing test name and location
- `evidence:` the test runner output showing the failure and the assertion reason
- The subtask remains `in-progress` until every planned test case is complete

Appends to `.agent/learnings.md` only if a decision or surprise emerged while writing the test (e.g. the subtask needs splitting, or the interface is unclear).

---

## Scope Boundary

This skill does **one thing**: write the failing test.

- It does NOT write production code.
- It does NOT refactor.
- It does NOT make the test pass.
- It does NOT touch other subtasks.
- It does NOT handle `[tidy]` subtasks.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT write production code in this skill — not even a stub that makes the test pass.
Do NOT skip the failure-confirmation step. The test must run and fail for the right reason.
Do NOT write tests for more than one active test case per invocation.
Do NOT proceed to `4dc-tdd-green` if the test fails for the wrong reason (syntax error, missing import, wrong assertion). Fix the test first.
Do NOT start without an approved mini-plan recorded for this subtask.
</HARD-GATE>

---

## Process

1. **Read the current subtask and mini-plan** from `.agent/implementation.md` — the approved `[behavior]` subtask with an unfinished `tests` list. Select its `active_test`, or the first test case with `state: pending`.
2. **Confirm the todo item is `in_progress`** from subtask planning.
3. **Transition the subtask** to `state: in-progress` if it is still `state: approved` from mini-planning.
4. **Read the testing strategy** in `CONSTITUTION.md` and `docs/testing.md` — choose the cheapest test depth that gives sufficient confidence at this boundary.
5. **Write one test case** for `active_test` that specifies one example of the behavior. Name it in domain language so the test reads as a specification. Do not implement or activate the other cases yet.
6. **Run the test.** Confirm it fails. Read the failure message — it must fail because the behavior does not exist, not because of a setup or import error.
7. **Record evidence** in `.agent/implementation.md`:
   - Confirm subtask is `state: in-progress`
   - Set the active test case `state: red`
   - Record `test:` the test name and file
   - Record `evidence:` the failure output (test name + assertion reason)
8. **Continue.** Load `4dc-tdd-green` for this active test case. Do not request user confirmation at this transition.

If the test reveals a local implementation surprise that stays within the approved mini-plan's scope, record it in `implementation.md` and `learnings.md`, adapt the remaining steps, and continue. Stop only if it changes acceptance criteria or scope, requires an unapproved structural decision, or makes the approved approach unsafe.

---

## Checklist

- [ ] Current `[behavior]` subtask and `active_test` read from `implementation.md`
- [ ] Todo item for this subtask marked `in_progress`
- [ ] Testing strategy consulted (depth, naming, isolation)
- [ ] One `active_test` written that specifies one example of the behavior
- [ ] Test run and confirmed failing for the right reason
- [ ] Active test case updated: `state: red`, test name, evidence
- [ ] `active_test` remains the only active test case
- [ ] No production code written
- [ ] Any interface surprise recorded in `learnings.md`

---

## Handoff

Updated artifact: `.agent/implementation.md` (current subtask `state: in-progress`, active test `state: red`)
Next skill: `4dc-tdd-green` — load `skills/tdd-green/SKILL.md`
