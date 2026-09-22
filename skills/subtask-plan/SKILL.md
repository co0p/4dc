---
name: 4dc-subtask-plan
description: "Use before each pending implementation subtask. Investigates local unknowns, agrees the immediate implementation slice, and starts autonomous execution after one explicit approval."
---

# Subtask Plan Skill

## One Responsibility

Turn the next approved-plan subtask into a small, immediate implementation agreement between developers. Resolve local unknowns, discuss the exact sequence and evidence with the user, record the approved mini-plan, and then execute the whole subtask autonomously through Tidy or Red-Green-Refactor.

## Foundations

- **Alistair Cockburn: communication as coordination.** Use focused conversation to expose assumptions and reach shared understanding before commitment.
- **Ron Jeffries: card, conversation, confirmation.** Keep the written request concise, develop shared understanding through conversation, and prove the outcome with concrete acceptance evidence.
- **Mary and Tom Poppendieck: pull small batches from value.** Work on one small, valuable unit at a time so feedback is fast and correction is cheap.
- **Mary and Tom Poppendieck: decide at the last responsible moment.** Delay reversible commitments until evidence is available, while making blocking decisions explicit when they become necessary.
- **Martin Fowler: two hats.** Keep behavior changes and structural improvements separate so each remains understandable and verifiable.

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
- If a safe experiment must change files, include the disposable experiment and cleanup in the mini-plan. Execute it only after approval, record the result, remove scratch artifacts, and continue autonomously when the finding stays within scope.
- If the finding changes scope, acceptance criteria, architecture, or safety, stop and return to the relevant approval gate.

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
Do NOT edit production code, tests, or project documentation during this skill.
Do NOT set `state: approved` until the user explicitly approves the presented mini-plan.
Do NOT plan more than the current subtask.
Do NOT implement directly from `state: pending`; every tidy and behavior subtask requires this conversation.
If current code invalidates the approved cycle plan, record the discrepancy and return to plan approval rather than improvising.
After approval, do NOT ask for routine confirmation during Red, Green, Refactor, verification, commits, or progress tracking. Continue until the subtask is complete or a stop condition is reached.
</HARD-GATE>

---

## Process

1. **Select the current subtask** — the first `state: pending` entry whose dependencies are complete. Mark its todo item `in_progress`; it remains in progress through implementation.
2. **Investigate narrow context** — read only its files, references, governing docs, history, and current tests. Run non-mutating diagnostics needed to answer local implementation questions. Record findings in the draft mini-plan and confirm the cycle plan still matches the code.
3. **Draft the mini-plan** — state the intent, findings, exact files and symbols, ordered moves, test evidence, observability impact, local risks, and explicit non-goals. For behavior work, name the first failing acceptance-facing or focused test and the minimal expected production path; for tidy work, state the behavior-preservation proof. Include any approved post-approval experiment as an explicit first step with cleanup and decision rules.
4. **Developer conversation** — present the mini-plan concisely. Ask one focused question where a choice remains. Iterate until the user explicitly approves it.
5. **Persist approval** — write the approved `mini_plan`, `approved`, and `state: approved` fields into `.agent/implementation.md`.
6. **Execute autonomously** — `[tidy]` goes to `4dc-tidy` and `[behavior]` to `4dc-tdd-red`. Continue through all implementation-skill transitions without asking the user again. Each skill updates `implementation.md`; return to `subtask-plan` only after the current subtask is complete and the next one is pending.

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
Next skill: detected from the approved subtask type and executed autonomously until the subtask is complete.