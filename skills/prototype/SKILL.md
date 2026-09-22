---
name: 4dc-prototype
description: "Optional. Use after increment.md is approved when a blocking unknown needs a throwaway spike. Builds to learn, not to ship. Output is findings, not production code."
---

# Prototype Skill

## One Responsibility

Build a throwaway spike that resolves one named unknown, then record what was learned. The code is discarded; the knowledge feeds the plan.

---

## Foundations

- **Mary and Tom Poppendieck: cheapest useful experiment.** Resolve uncertainty with the smallest experiment that can produce credible evidence.
- **Mary and Tom Poppendieck: decide at the last responsible moment.** Delay reversible commitments until evidence is available, while making blocking decisions explicit when they become necessary.
- **Kent Beck: spike.** Use a time-boxed, throwaway experiment to answer one question rather than deliver production behavior.

---

## Expected Input

- `CONSTITUTION.md`
- `.agent/increment.md` (approved)
- One named unknown that blocks planning (e.g. "Can library X do Y within Zms?", "Is auth model A viable for our domain?")

---

## Concrete Output

`.agent/prototype.md` containing:
- **Question**: the one unknown this spike answers
- **Approach**: what was built and how (throwaway — not merged)
- **Finding**: the answer, with evidence (measurement, demo, reproduced behavior)
- **Recommendation**: what this means for the plan (proceed, change approach, split increment)
- **Disposal**: confirmation that prototype code was discarded or isolated

Required `.agent/prototype.md` headings:
- `## Question`
- `## Approach`
- `## Finding`
- `## Recommendation`
- `## Disposal`

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
Do NOT merge prototype code into the production branch.
Do NOT spend more than the agreed time-box without checking in.
Do NOT start the plan until the finding is recorded and the user confirms the recommendation.
Do NOT turn the prototype into the implementation — it is a learning artifact, not a head start.
</HARD-GATE>

---

## Process

1. **Name the unknown** — state the single question this spike answers and why it blocks planning.
2. **Agree a time-box** — propose a ceiling (e.g. 30 min, one session). If the ceiling is hit without an answer, stop and record what was learned.
3. **Build the throwaway** — write the minimum code to answer the question. It may live in a scratch directory, a branch, or a REPL — never in the production source tree.
4. **Capture the finding** — measure, demo, or reproduce. Record evidence in `.agent/prototype.md`.
5. **Recommend** — state what the plan should do differently given the finding.
6. **Dispose** — delete or isolate the prototype code. Confirm disposal in the finding.
7. **STOP** — present the finding and recommendation. Wait for the user to confirm before loading `4dc-plan`.

---

## Checklist

- [ ] One named unknown stated
- [ ] Time-box agreed
- [ ] Throwaway built (not in production source tree)
- [ ] Finding recorded with evidence
- [ ] Recommendation stated
- [ ] Prototype code disposed or isolated
- [ ] User confirmed the recommendation

---

## Handoff

Terminal artifact: `.agent/prototype.md` (transient — consumed by `4dc-plan`)
Next skill: `4dc-plan` — load `skills/plan/SKILL.md`

If the spike reveals the increment is infeasible, do not proceed to plan. Return to `4dc-increment` to rescope.