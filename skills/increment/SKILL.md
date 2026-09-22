---
name: 4dc-increment
argument-hint: "short increment intent, e.g. 'feature: export weekly summary'"
description: "Use after CONSTITUTION.md exists and before any implementation. Defines the next narrow, testable changeset — WHAT and WHY only, no technical detail."
---

# Increment Skill

## One Responsibility

Define one small, outcome-focused increment with measurable acceptance criteria and explicit out-of-scope boundaries.

---

## Foundations

- **Ron Jeffries: card, conversation, confirmation.** Keep the written request concise, develop shared understanding through conversation, and prove the outcome with concrete acceptance evidence.
- **Bill Wake: INVEST.** Shape work so it is independent, negotiable, valuable, estimable, small, and testable.
- **Mary and Tom Poppendieck: pull small batches from value.** Work on one small, valuable unit at a time so feedback is fast and correction is cheap.
- **Mary and Tom Poppendieck: decide at the last responsible moment.** Delay reversible commitments until evidence is available, while making blocking decisions explicit when they become necessary.
- **Alistair Cockburn: communication as coordination.** Use focused conversation to expose assumptions and reach shared understanding before commitment.

---

## Expected Input

- `CONSTITUTION.md`
- `docs/roadmap.md`
- User intent (one sentence or short phrase describing the desired outcome)
- **Customer-authored use case** — a job story written before acceptance criteria are defined

---

## Concrete Output

`.agent/increment.md` containing:
- **Use case**: job story in the form _"When [situation], I want to [action], so that [outcome]."_ Written or confirmed by the customer before criteria are defined.
- **Goal**: one sentence — the observable outcome (distilled from the use case)
- **Mode**: one of `Behavior`, `Refactor`, or `Chore`. Determines what acceptance evidence looks like (see below).
- **Branch**: the git branch name for this increment, in the form `increment/<slug>` (e.g. `increment/token-refresh`)
- **Acceptance criteria**: 2–5 binary, verifiable conditions derived from the use case; each criterion names observable proof rather than an implementation mechanism
- **Acceptance-test intent**: required. Non-technical description of what an observer sees when each criterion holds. Content depends on mode:
  - `Behavior`: describe the new user-observable outcome — user action, precondition, observable result
  - `Refactor`: name the existing acceptance tests that must remain green as the regression anchor (behavior preserved)
  - `Chore`: name the actor (developer, operator, CI) and the observable outcome at their boundary (script exit code, generated artifact, log line, config check)
- **Out of scope**: explicit exclusions that prevent scope creep
- **Smallest-slice rationale**: why removing any remaining criterion would make the outcome unusable, unverifiable, or not independently releasable
- **Constitution constraints**: which guardrails apply to this increment
- **Roadmap entry**: the feature name and job story to add to `docs/roadmap.md` In Progress section

Required `.agent/increment.md` headings:
- `## Use Case`
- `## Goal`
- `## Mode`
- `## Branch`
- `## Acceptance Criteria`
- `## Acceptance-Test Intent`
- `## Out Of Scope`
- `## Smallest-Slice Rationale`
- `## Constitution Constraints`
- `## Roadmap Entry`

### Mode Reference

| Mode | Delivers | Acceptance evidence at Promote |
|------|----------|-------------------------------|
| `Behavior` | New or changed user-observable behavior | New acceptance test + full existing suite green |
| `Refactor` | Internal structure change, behavior preserved | Named existing acceptance tests remain green, no new AT |
| `Chore` | Developer- or operator-facing change (build, CI, tooling, docs) | Acceptance test at the named actor's boundary + full existing suite green |

Exploratory work with no production output is not an Increment. Use the `prototype` phase.

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
Do NOT write acceptance criteria before the use case is stated — criteria must derive from the job story.
Do NOT include technical design, file names, implementation approaches, or coding detail in increment.md.
Do NOT start a plan or any implementation work during this phase.
Do NOT approve an increment with vague acceptance criteria ("works correctly", "feels right").
One increment per cycle — if scope expands, split into separate increments.
Do NOT propose the final increment after only restating the initial request. Ask one focused question at a time until the user, trigger, outcome, failure behavior, and scope boundary are explicit.
Do NOT keep two independently valuable outcomes in one increment. Present the smaller first slice and move the remainder to out of scope or the roadmap.
Do NOT call an increment smallest without performing the subtraction test: remove each criterion in turn and split it out unless its removal makes the remaining outcome unusable, unverifiable, or not independently releasable.
Exactly one Mode must be declared (`Behavior`, `Refactor`, or `Chore`). Exploratory work belongs in the `prototype` phase, not an increment.
Acceptance-Test Intent is required for every mode:
- `Behavior`: describe the new observable outcome in non-technical language. No test files, no tooling, no code.
- `Refactor`: name the existing acceptance tests that anchor the regression proof.
- `Chore`: name the actor and the observable outcome at their boundary.
</HARD-GATE>

---

## Process

1. **Read context** — `CONSTITUTION.md`, `docs/roadmap.md`, and any prior `.agent/` files from the last cycle.
2. **Conversation: Discover the outcome** — ask one focused question at a time. Establish who experiences the problem, the triggering situation, the smallest valuable outcome, the expected failure or boundary behavior, and what must not be included. Do not infer a missing answer from conventions.
3. **Determine the Mode** — from the discussion, classify the increment as `Behavior`, `Refactor`, or `Chore`. If the work has no observable outcome anywhere (user, operator, developer, or CI), it is a `prototype`, not an increment.
4. **Slice by subtraction** — draft the criteria, remove each one in turn, and split any independently releasable value into a later increment. State the smallest-slice rationale and deferred outcomes.
5. **Draft Acceptance-Test Intent** — matched to the Mode:
   - `Behavior`: what a user sees when each criterion holds
   - `Refactor`: which existing acceptance tests must keep passing to prove behavior is preserved
   - `Chore`: which actor observes what outcome at which boundary
6. **Conversation: Propose the increment** — present the job story, mode, binary criteria, acceptance-test intent, out-of-scope list, and smallest-slice rationale. Iterate until the user explicitly confirms it.
7. **On approval** — write `.agent/increment.md`, move the feature to In Progress in `docs/roadmap.md`, and create the branch: `git checkout -b increment/<slug>`.

---

## Checklist

- [ ] `CONSTITUTION.md` and `docs/roadmap.md` read
- [ ] Use case (job story) stated by customer
- [ ] Mode declared (`Behavior`, `Refactor`, or `Chore`); exploratory work routed to `prototype` instead
- [ ] Acceptance criteria derived from use case, not from technical assumptions
- [ ] Acceptance criteria are binary and verifiable
- [ ] Acceptance-Test Intent recorded and matches the declared Mode:
  - Behavior: describes the new observable outcome without technical detail
  - Refactor: names the existing acceptance tests that anchor the regression proof
  - Chore: names the actor and the observable outcome at that boundary
- [ ] Out-of-scope list is non-empty
- [ ] User, trigger, outcome, failure boundary, and exclusions were explicitly discussed
- [ ] Subtraction test performed for every acceptance criterion
- [ ] Smallest-slice rationale explains why the remaining criteria cannot be split further
- [ ] Branch name derived (`increment/<slug>`)
- [ ] Roadmap entry (feature name + job story) identified
- [ ] User approval received
- [ ] `.agent/increment.md` written
- [ ] `docs/roadmap.md` updated: feature moved to In Progress
- [ ] Branch created: `git checkout -b increment/<slug>`

---

## Handoff

Terminal artifact: `.agent/increment.md`
Next skill: `4dc-plan` — load `skills/plan/SKILL.md`