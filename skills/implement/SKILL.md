---
name: 4dc-implement
description: "Run once after plan.md is approved. Scaffolds implementation.md, populates the todo list, and hands off to subtask-plan for the first developer conversation."
---

# Implement Skill

## One Responsibility

Bootstrap the implementation loop: read the approved plan, create `.agent/implementation.md` with every subtask in its initial state, populate the **internal todo list**, then hand off to `subtask-plan`. This skill runs exactly once per cycle; each subtask is mini-planned and approved later, immediately before implementation.

---

## Foundations

- **Alistair Cockburn: information radiators.** Keep work state and evidence visible so coordination does not depend on private context.
- **Mary and Tom Poppendieck: pull small batches from value.** Work on one small, valuable unit at a time so feedback is fast and correction is cheap.
- **Kent Beck: small steps.** Make each change small, independently verifiable, and cheap to reverse.

---

## Expected Input

- `.agent/plan.md` (must be approved)
- `CONSTITUTION.md` (for testing strategy reference)

---

## Concrete Output

### `.agent/implementation.md`

Scaffolded from `plan.md`. Every subtask from `## Subtasks` becomes an entry in the initial state:

```markdown
# Implementation: <goal from plan.md>

status: in-progress
branch: <branch from plan.md>
started: <ISO date>

## Baseline
Tests before: establish by running the full suite before any changes

## Subtasks

### 1. <subtask name from plan.md>
type: <tidy | behavior>
state: pending
tests:                         # behavior subtasks only
  - id: <id>
    name: <name>
    file: <file>
    state: pending
active_test: <first test id>   # behavior subtasks only

### 2. <subtask name>
type: <type>
state: pending
...
```

Rules:
- Copy subtask names, types, and test lists verbatim from `plan.md`. Do not paraphrase.
- All subtasks start `state: pending`.
- For `[behavior]` subtasks, copy the full `tests` list and set `active_test` to the first test id.
- Do not add fields not present in the template above — the implement skills own those fields.
- `subtask-plan` adds `mini_plan`, `approved`, and changes `state` to `approved` after the user confirms the immediate approach.

### Internal todo list

Populate the todo list immediately after writing `implementation.md`. One item per subtask, in plan order. Use the subtask name and type as the label:

```
[tidy]     1. <subtask name>          pending
[behavior] 2. <subtask name>          pending
[behavior] 3. <subtask name>          pending
```

**Todo list discipline — applies for the entire implementation loop, not just this skill:**

- Mark a todo `in_progress` the moment the skill for that subtask starts. Only one item is `in_progress` at a time.
- Mark a todo `completed` only after `implementation.md` records `state: complete` for that subtask and the commit hash is present.
- Never mark a subtask complete based on intent — only on evidence in `implementation.md`.
- If a subtask is split or re-ordered during the loop, update the todo list to match before continuing.
- The todo list is the user's real-time view of cycle progress. Keep it accurate.

---

## Scope Boundary

This skill does **one thing**: initialise the tracking artifacts and hand off.

- It does NOT write production code.
- It does NOT write tests.
- It does NOT make any structural change to the codebase.
- It does NOT run the full test suite (that baseline is established as the first act of the first implement skill).

---

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
Do NOT create implementation.md until plan.md is confirmed approved.
Do NOT paraphrase subtask names or test ids — copy them verbatim from plan.md.
Do NOT mark any todo in_progress until the corresponding skill has actually started work.
Do NOT skip populating the todo list — it is a required output of this skill, not optional.
Do NOT hand a pending subtask directly to tidy, tdd-red, or tdd-green. It must pass through subtask-plan.
</HARD-GATE>

---

## Process

1. **Read `.agent/plan.md`** — extract goal, branch, and the full ordered subtask list including test case ids for behavior subtasks.
2. **Scaffold `.agent/implementation.md`** — create the file using the template above. Subtask names and test lists copied verbatim.
3. **Populate the todo list** — one item per subtask, all `pending`, in plan order, labeled with type and name.
4. **Hand off** — load `skills/subtask-plan/SKILL.md` for the first pending subtask. Its approved mini-plan determines the later Tidy, Red, Green, and Refactor work.

---

## Checklist

- [ ] `.agent/plan.md` read and confirmed approved
- [ ] Goal and branch copied from `plan.md` into `implementation.md`
- [ ] Every subtask from `plan.md` present in `implementation.md` with correct type and `state: pending`
- [ ] Every `[behavior]` subtask has its `tests` list and `active_test` set
- [ ] Todo list populated with one item per subtask, all `pending`
- [ ] No code written, no tests written, no structural changes made

---

## Handoff

Terminal artifact: `.agent/implementation.md` (scaffolded, all subtasks `state: pending`)
Todo list: populated, all items `pending`
Next skill: `4dc-subtask-plan` for the first pending subtask.