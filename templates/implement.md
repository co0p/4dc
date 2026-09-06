---
name: 4dc-implement
description: "Use after plan.md is approved. Executes the plan in Red→Green→Refactor order, tracking progress and capturing learnings after each subtask."
---

# Implement Skill

## One Responsibility

Execute `.agent/plan.md` in controlled steps, maintain continuous progress state, and record every decision and deviation as it happens.

---

## Expected Input

- `CONSTITUTION.md`
- `.agent/increment.md` (approved)
- `.agent/plan.md` (approved)
- Current test baseline (run tests before touching anything)

---

## Concrete Output

- **`.agent/implementation.md`** — live progress log, updated after each subtask; final status either `status: complete` or `status: blocked`
- **`.agent/learnings.md`** — decisions made, deviations from plan, and lessons for `promote`

Required `implementation.md` headings:
- `## Baseline`
- `## Subtasks`
- `## Final Verification`

Required `learnings.md` headings:
- `## Decisions`
- `## Deviations`
- `## Surprises`
- `## Promote Candidates`

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT write production code for a subtask before a failing test exists for it (Red→Green→Refactor).
Do NOT mark a subtask complete without objective evidence (test output, command result, or observable behavior).
Do NOT skip or reorder subtasks without documenting the reason in learnings.md.
Do NOT proceed to the next subtask if the current one fails verification.
Do NOT mix structural tidying and behavior change in the same subtask.
</HARD-GATE>

---

## Process

1. **Establish baseline** — run existing tests; record pass/fail state in `implementation.md`
2. **Execute `[research]` subtasks first when present** — resolve only the blocking unknown, record the evidence, and return to the approved plan
3. **Execute `[tidy]` subtasks before any `[behavior]` subtask they support**
   a. Make the structural change
   b. Run the relevant tests — confirm they stay green
   c. Update `implementation.md` with evidence
   d. Record any design or architecture implications in `learnings.md`
4. **For each `[behavior]` subtask in plan.md:**
   a. Write the failing test (Red)
   b. Run it — confirm it fails for the right reason
   c. Write minimal production code (Green)
   d. Run the narrowest relevant tests, then the broader suite needed by the constitution
   e. Refactor if needed — run tests again
   f. Update `implementation.md`: mark subtask complete with evidence
   g. Append any decisions, deviations, surprises, or promote candidates to `learnings.md`
5. **Final verification** — run the full test suite or constitution-defined release gate; confirm all acceptance criteria from `increment.md` are met
6. **Mark complete** — set `status: complete` in `implementation.md`

## Tidy First Rule

When the approved plan includes `[tidy]` subtasks, use Kent Beck's Tidy First rule:

- Finish the structural preparation before behavior work that depends on it.
- Keep every `[tidy]` subtask behavior-preserving and test-green.
- If a tidy step would change observable behavior, it was mislabeled and must move to `[behavior]`.

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
type: tidy | behavior | research
status: complete
evidence: `npm test` — 12 passing, 0 failing (added test: <name>)

### 2. <subtask name>
status: in-progress
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

{{TEMPLATE:html}}

---

## Checklist

- [ ] Baseline test run recorded
- [ ] `[tidy]` subtasks, if any, stayed behavior-preserving and test-green
- [ ] Each subtask follows Red→Green→Refactor
- [ ] Each subtask has objective completion evidence
- [ ] All acceptance criteria met
- [ ] `implementation.md` status set to `complete`
- [ ] `learnings.md` has promote candidates listed

---

## Handoff

Terminal artifacts: `.agent/implementation.md` (status: complete) + `.agent/learnings.md`
Next skill: `4dc-promote` — load `skills/promote/SKILL.md`
