---
name: 4dc-refactor
description: "Improve the design of the code that just passed its test, without changing behavior. Tests must stay green throughout. Commits as refactor: <what>. Advances to the next subtask."
---

# Refactor Skill

## One Responsibility

Take the code that just made the current test case green and improve its design — without changing behavior. Tests stay green throughout. Then record the case as complete and either activate the next case in the same subtask or advance.

---

## Foundations

- **Martin Fowler: behavior-preserving refactoring.** Improve internal structure under a green test suite without changing observable behavior.
- **Martin Fowler: two hats.** Keep behavior changes and structural improvements separate so each remains understandable and verifiable.
- **Kent Beck: the refactor pass.** After reaching green, deliberately inspect the design and improve it only when the improvement is useful.
- **Mary and Tom Poppendieck: eliminate waste.** Build, preserve, and document only work that contributes verified value or necessary learning.

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

## Language and Interaction Rules

- Use plain, direct language. Keep output scannable.
- Prefer short sentences and bullets.
- State decisions, actions, blockers, and evidence. Omit motivational, decorative, and generic advice.
- Do not repeat the user's request, loaded instructions, or handoff contents.
- Explain a choice only when it affects scope, risk, verification, or the next handoff.
- Ask one focused question at a time. Do not use broad questionnaires.
- State assumptions explicitly. If required evidence is missing or contradictory, ask rather than inventing an answer.
- Use the project's domain language in product artifacts. Keep internal workflow and agent terminology out of permanent product documentation.
- Refer to files, symbols, commands, states, and evidence precisely. Avoid vague terms such as "works", "correct", or "should be fine".
- End a phase response with the decision needed, the blocker, or the next handoff. During approved autonomous implementation, continue without routine confirmation.

## Execution Contract

- Never copy internal workflow names, skill names, phase names, orchestrator terms, `.agent/` paths, or `.agents/` paths into permanent product artifacts.
- Before writing a permanent artifact, scan it for internal workflow references and remove them.
- Produce only the artifact for this phase. Do not leak work from a later phase into this one.
- Treat tests, architecture notes, ADRs, and user-facing docs as first-class communication artifacts.
- Gather only enough context to identify the governing constraints, the target artifact, and the cheapest validation step. Then act.
- Resolve conflicts in this order: explicit user approval, approved prior-phase artifacts, `CONSTITUTION.md`, this skill.
- Low-risk actions: reads, searches, diffs, and local validation commands.
- Medium-risk actions: local reversible edits to phase artifacts.
- High-risk actions: destructive file operations, external side effects, or skipping a stop gate. Require explicit approval first.
- If a required input is missing or contradictory, ask one focused question or stop and wait for explicit approval. Do not invent missing facts.
- Before finishing, run the phase checklist and confirm every required section is present.

---

<HARD-GATE>
Do NOT change observable behavior. If any test goes red, you changed behavior — revert and try again.
Do NOT add new tests or new production features. This is the refactor hat, not the behavior hat.
Do NOT skip the question "can the design be improved?" — even if the answer is no, the question must be asked and recorded.
Do NOT set the active test case to `state: complete` before either the refactor commit hash or an explicit "no refactor needed" note is recorded in `implementation.md`.
Do NOT mark the subtask complete until every planned test case has completed Red → Green → Refactor and all tests are green.
Keep the refactor strictly proportional to the Green change just made. Cross-file reshaping, dependency inversion, or boundary changes belong in a Refactor Increment, not here. When such reshaping is warranted, record it in `learnings.md` as a candidate Refactor increment instead of doing it now.
Commit as `refactor: <what changed>` — never `feat:` or `fix:`. If no refactoring was needed, skip the commit and record `refactor: none — <reason>` instead.
</HARD-GATE>

---

## Process

1. **Read the current subtask** from `.agent/implementation.md` — the `[behavior]` subtask in `state: in-progress` with an active test case in `state: green`.
2. **Run the tests.** Confirm they are green before you start.
3. **Ask the refactor question:** can the design be improved without changing behavior? Look for: duplication, unclear naming, long methods, deep nesting, missing abstraction, poor separation of concerns.
4. **If no improvement is needed:** skip the commit. Record `refactor: none — <one-line reason>` for the active test case as evidence. Advance to step 8.
5. **If improvement is needed:** make the change. One move at a time — extract method, rename, inline, move. Run the tests after each move. If they go red, revert — you changed behavior.
6. **Run the full test suite** required by the constitution.
7. **Commit** as `refactor: <what changed>` and record `refactor:` the commit hash and message in `implementation.md`.
8. **Set the active test case to `state: complete`** with test output evidence. This transition happens only after step 4 or step 7 has recorded either the "no refactor needed" note or the commit hash.
9. **Append learnings** if design decisions or promote candidates emerged.
10. **Advance within the subtask:** if a pending test case remains, set it as `active_test` with `state: pending`; the approved subtask mini-plan remains valid and `tdd-red` starts its cycle without another user gate. If all cases are complete, set the behavior subtask `state: complete`, **mark the todo item `completed`**, and route the next pending subtask through `subtask-plan`.

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
- [ ] If refactored: tests stayed green after each move, committed as `refactor:`, hash recorded in `implementation.md`
- [ ] If no refactoring needed: `refactor: none — <reason>` recorded in `implementation.md`
- [ ] Active test case marked `state: complete` (only after commit hash or "no refactor" note is present)
- [ ] Next pending test case activated, or behavior subtask marked complete and todo item marked `completed` when all cases are complete
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