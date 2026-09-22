---
name: 4dc-tdd-green
description: "Make the current failing test pass with minimal code. No refactoring. Sets state: green and continues autonomously to refactor."
---

# TDD Green Skill

## One Responsibility

Write the minimal production code that makes the current failing `active_test` pass. Nothing more. A behavior subtask may contain several cohesive test cases; complete one case, then hand off to `4dc-refactor` before activating the next case.

- For a `[behavior]` subtask with an active test in `state: red`: write minimal code to pass that test, set the test to `state: green`, and hand off to `4dc-refactor`.

---

## Foundations

{{FOUNDATION:beck-green}}
{{FOUNDATION:fowler-two-hats}}
{{FOUNDATION:poppendieck-eliminate-waste}}

---

## Expected Input

- `.agent/plan.md` (approved)
- `.agent/implementation.md` with one behavior subtask in `state: red` and an approved `mini_plan`
- `CONSTITUTION.md` testing strategy

**Narrow context:** load only the files named in the current subtask's `files:` and `references:` fields in `plan.md`. Do not re-scan the codebase — the plan already did that work.

---

## Concrete Output

Updates `.agent/implementation.md` for the current subtask:
- For `[behavior]`: active test `state: green`, `evidence:` test output showing pass, `commit:` hash with `feat:` or `fix:` prefix

Appends to `.agent/learnings.md` if decisions, deviations, or surprises emerged.

---

## Scope Boundary

This skill does **one thing**: make the test pass with minimal code.

- It does NOT refactor (that is `4dc-refactor`).
- It does NOT write failing tests (that is `4dc-tdd-red`).
- It does NOT handle `[tidy]` subtasks (that is `4dc-tidy`).
- It does NOT write permanent docs (that is `4dc-promote`).
- It does NOT run final verification (that is `4dc-refactor`, when all subtasks are complete).

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT write more code than needed to pass the test. No speculative generality, no "while I'm here" features.
Do NOT refactor in this skill. Improving the design is a separate hat — load `4dc-refactor` next.
Do NOT mark a `[behavior]` subtask `state: complete` — set the active test to `state: green` and hand off to refactor. The subtask becomes complete only after every planned test case has completed its own Red → Green → Refactor loop.
Do NOT mix behavior change and refactoring in the same commit.
Commit behavior work as `feat: <what changed>` or `fix: <what changed>`. Never `refactor:` or `tidy:`.
Behavior work inherits the mini-plan approved before Red. Continue without routine user confirmation.
</HARD-GATE>

---

## Process

### For a `[behavior]` subtask in `state: red`

1. **Read the failing `active_test`** from the test file recorded in `implementation.md`.
2. **Write minimal production code** to make the test pass. Add only what the test requires — no extra methods, no speculative abstraction, no "I'll need this later."
3. **Run the test.** Confirm it passes.
4. **Run the narrowest relevant tests**, then the broader suite required by the constitution.
5. **Record evidence** in `implementation.md`: set the active test to `state: green`, with test output showing pass. Leave the other test cases unchanged.
6. **Commit** as `feat: <what changed>` or `fix: <what changed>`.
7. **Append learnings** — decisions, deviations, surprises, promote candidates.
8. **Continue autonomously.** Load `4dc-refactor` to complete this active test case. Do not request user confirmation at this transition.

---

## implementation.md Structure

```markdown
# Implementation: <increment goal>

status: in-progress  <!-- or: complete | blocked -->
started: <ISO date>

## Baseline
Tests before: X passing, Y failing

## Subtasks

### 1. <subtask name>
type: behavior
state: in-progress
tests:
  - id: provider-success
    name: <test name>
    file: <test file>
    state: complete
    evidence: `npm test` — 12 passing, 0 failing
    commit: <hash> — feat: <what changed>
    refactor: <hash> — refactor: <what improved>
  - id: provider-not-found
    name: <test name>
    file: <test file>
    state: red
    evidence: fails — <assertion reason>
active_test: provider-not-found

### 2. <subtask name>
type: tidy
state: complete
evidence: `npm test` — 12 passing, 0 failing (tests unchanged)
commit: <hash> — tidy: <what changed>

### 3. <subtask name>
type: behavior
state: in-progress
tests:
  - id: <case-id>
    name: <test name>
    file: <test file>
    state: green
    evidence: `npm test` — 12 passing, 0 failing
    commit: <hash> — feat: <what changed>
active_test: <case-id>

### 4. <subtask name>
type: behavior
state: in-progress
tests:
  - id: <case-id>
    name: <test name>
    file: <test file>
    state: pending
active_test: <case-id>
```

---

## learnings.md Structure

```markdown
# Learnings: <increment goal>

## Decisions
- <decision>: <rationale>

## Deviations
- Subtask N: <what changed and why>

## Surprises
- <unexpected finding>

## Promote Candidates
- <ADR, architecture update, domain-language update, test pattern, or performance contract worth keeping>
```

---

## Checklist

### Per `[behavior]` subtask
- [ ] Minimal code written to pass the test (no speculative generality)
- [ ] Narrowest relevant tests run, then broader suite
- [ ] Active test case marked `state: green` with evidence
- [ ] Committed as `feat:` or `fix:`
- [ ] Learnings appended (decisions, deviations, surprises, promote candidates)
- [ ] After refactor, activate the next pending test case or mark the behavior subtask complete when all cases are complete

---

## Handoff

Updated artifacts: `.agent/implementation.md` (current `[behavior]` subtask `state: green`) + `.agent/learnings.md`
Next skill (after `[behavior]` green): `4dc-refactor` — load `skills/refactor/SKILL.md`
