---
name: 4dc-plan
description: "Use after increment.md is approved. Converts increment intent into an ordered, verifiable technical execution plan with concrete subtasks."
---

# Plan Skill

## One Responsibility

Define HOW to deliver `.agent/increment.md` — an ordered sequence of actionable subtasks with explicit verification points.

---

## Expected Input

- `CONSTITUTION.md`
- `.agent/increment.md` (must be approved)
- Current codebase structure (read relevant files)

---

## Concrete Output

`.agent/plan.md` containing:
- **Goal**: copied from `increment.md` — one sentence
- **Approach**: 2–3 sentences on strategy, including architectural boundary and any performance-sensitive path; no code yet
- **Subtasks**: ordered list, each with:
  - Type: `[research]`, `[tidy]`, or `[behavior]`
  - Description (what, not how)
  - Verification step (how to confirm it’s done)
  - Dependencies on prior subtasks
  - Acceptance criteria covered
- **Risks**: known unknowns that could block execution

Required `.agent/plan.md` headings:
- `## Goal`
- `## Approach`
- `## Subtasks`
- `## Risks`

`[research]` is only for blocking unknowns.
`[tidy]` is structural only and must preserve observable behavior.
`[behavior]` changes observable behavior and must be verified by a failing test or failing executable check first.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT start implementation during this phase — no code, no file edits.
Do NOT write plan.md until the HTML review is approved.
Do NOT list subtasks without verification steps.
Every subtask must map to at least one acceptance criterion from increment.md.
Do NOT place a `[behavior]` subtask before the `[tidy]` subtasks it depends on.
</HARD-GATE>

---

## Process

1. **Read inputs** — `CONSTITUTION.md`, `.agent/increment.md`, relevant source files, and any current architecture or ADR docs touched by the change
2. **Identify risks** — unknown dependencies, test gaps, ambiguous requirements, architectural tension, and performance risks
3. **Draft subtasks** — ordered, each independently verifiable, sized for one focused work session, using `[research]` only when needed, `[tidy]` before `[behavior]`
4. **Generate `.agent/plan-review.html`** — present the plan with full traceability to acceptance criteria
5. **STOP** — wait for explicit approval or revision
6. **On approval** — write `.agent/plan.md`

{{TEMPLATE:html}}

---

## Checklist

- [ ] `increment.md` acceptance criteria read
- [ ] Relevant source files scanned
- [ ] Every subtask has a verification step
- [ ] Every acceptance criterion has a covering subtask
- [ ] Every `[behavior]` subtask names a failing test or failing executable check as its first verification step
- [ ] Any architectural or performance-sensitive change is reflected in the approach or risks
- [ ] Risks documented
- [ ] HTML review generated and shown
- [ ] User approval received
- [ ] `.agent/plan.md` written

---

## Handoff

Terminal artifact: `.agent/plan.md`
Next skill: `4dc-implement` — load `skills/implement/SKILL.md`
