---
name: 4dc-subtask-plan
description: "Use before each pending implementation subtask. Investigates local unknowns, agrees the immediate implementation slice, and starts execution after one explicit approval."
---

# Subtask Plan Skill

## One Responsibility

Turn the next approved-plan subtask into a small, immediate implementation agreement between developers. Resolve local unknowns, discuss the exact sequence and evidence with the user, record the approved mini-plan, then run the bounded Tidy or Red-Green-Refactor sequence for this one subtask without further routine confirmation.

## Foundations

{{FOUNDATION:cockburn-communication}}
{{FOUNDATION:jeffries-card-conversation-confirmation}}
{{FOUNDATION:poppendieck-pull-small-batches}}
{{FOUNDATION:poppendieck-decide-late}}
{{FOUNDATION:fowler-two-hats}}

## Expected Input

- `.agent/plan.md` (approved)
- `.agent/implementation.md` with the next subtask in `state: pending`
- The current subtask's files, references, tests, dependencies, and acceptance criteria
- Current code and test state in only that narrow scope

## Concrete Output

Update the current entry in `.agent/implementation.md`:

```markdown
state: approved
mini_plan:
  intent: <one sentence>
  files:
    - <path and exact symbol or region>
  steps:
    - <ordered investigation, structural, red, green, and refactor actions that apply>
  verification:
    - <commands or reproducible checks and expected evidence>
  observability:
    - <signal change and verification, or not applicable with reason>
  risks:
    - <local risk, assumption, or none>
  non_goals:
    - <explicitly excluded adjacent work>
approved: <ISO date or conversation reference>
```

The mini-plan may refine HOW within the approved subtask. It may not expand acceptance criteria, add a new outcome, or silently change the cycle plan.

Research is folded into this skill:
- Before approval, inspect code, docs, history, and run non-mutating diagnostics needed to make the mini-plan credible.
- If a safe experiment must change files, include the disposable experiment and cleanup in the mini-plan. Execute it only after approval, record the result, remove scratch artifacts, and continue when the finding stays within scope.
- If the finding changes scope, acceptance criteria, architecture, or safety, stop and return to the relevant approval gate.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT edit production code, tests, or project documentation during this skill.
Do NOT set `state: approved` until the user explicitly approves the presented mini-plan.
Do NOT plan more than the current subtask.
Do NOT implement directly from `state: pending`; every tidy and behavior subtask requires this conversation.
If current code invalidates the approved cycle plan, record the discrepancy and return to plan approval rather than improvising.
After approval, do NOT ask for routine confirmation during Red, Green, Refactor, verification, commits, or progress tracking within this subtask. Return to conversation before the next subtask's mini-plan or when a stop condition is reached.
</HARD-GATE>

---

## Process

1. **Select the current subtask** — the first `state: pending` entry whose dependencies are complete. Mark its todo item `in_progress`; it remains in progress through implementation.
2. **Investigate narrow context** — read only its files, references, governing docs, history, and current tests. Run non-mutating diagnostics needed to answer local implementation questions. Record findings in the draft mini-plan and confirm the cycle plan still matches the code.
3. **Draft the mini-plan** — state the intent, findings, exact files and symbols, ordered moves, test evidence, observability impact, local risks, and explicit non-goals. For behavior work, name the first failing acceptance-facing or focused test and the minimal expected production path; for tidy work, state the behavior-preservation proof. Include any approved post-approval experiment as an explicit first step with cleanup and decision rules.
4. **Developer conversation** — present the mini-plan concisely. Ask one focused question where a choice remains. Iterate until the user explicitly approves it.
5. **Persist approval** — write the approved `mini_plan`, `approved`, and `state: approved` fields into `.agent/implementation.md`.
6. **Bounded execution** — `[tidy]` goes to `4dc-tidy` and `[behavior]` to `4dc-tdd-red`. Continue through the implementation-skill transitions for this one subtask without asking the user again. Each skill updates `implementation.md`. Return to `subtask-plan` for the next pending subtask, or to conversation if a stop condition triggers.

## Checklist

- [ ] Current pending subtask and dependencies confirmed
- [ ] Todo item marked `in_progress`
- [ ] Local unknowns investigated and findings included without broad re-planning
- [ ] Mini-plan names intent, exact scope, ordered moves, verification, observability, risks, and non-goals
- [ ] Plan drift or new scope returned to the appropriate approval gate
- [ ] User explicitly approved the mini-plan
- [ ] `implementation.md` records the mini-plan and `state: approved`
- [ ] No implementation change made

## Handoff

Updated artifact: `.agent/implementation.md` (current subtask `state: approved` with an approved mini-plan)
Next skill: detected from the approved subtask type and executed within the bounded subtask sequence.
