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

{{FOUNDATION:jeffries-card-conversation-confirmation}}
{{FOUNDATION:wake-invest}}
{{FOUNDATION:poppendieck-pull-small-batches}}
{{FOUNDATION:poppendieck-decide-late}}
{{FOUNDATION:cockburn-communication}}

---

## Expected Input

- `CONSTITUTION.md`
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

### Mode Reference

| Mode | Delivers | Acceptance evidence at Promote |
|------|----------|-------------------------------|
| `Behavior` | New or changed user-observable behavior | New acceptance test + full existing suite green |
| `Refactor` | Internal structure change, behavior preserved | Named existing acceptance tests remain green, no new AT |
| `Chore` | Developer- or operator-facing change (build, CI, tooling, docs) | Acceptance test at the named actor's boundary + full existing suite green |

Exploratory work with no production output is not an Increment. Use the `prototype` phase.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT write acceptance criteria before the use case is stated — criteria must derive from the job story.
Do NOT include technical design, file names, implementation approaches, or coding detail in increment.md.
Do NOT start a plan or any implementation work during this phase.
Do NOT approve an increment with vague acceptance criteria ("works correctly", "feels right").
One increment per cycle — if scope expands, split into separate increments.
Do NOT propose the final increment after only restating the initial request. Ask one focused question at a time until the user, trigger, outcome, failure behavior, and scope boundary are explicit.
Do NOT keep two independently valuable outcomes in one increment. Present the smaller first slice and defer the remainder to a future increment.
Do NOT call an increment smallest without performing the subtraction test: remove each criterion in turn and split it out unless its removal makes the remaining outcome unusable, unverifiable, or not independently releasable.
Exactly one Mode must be declared (`Behavior`, `Refactor`, or `Chore`). Exploratory work belongs in the `prototype` phase, not an increment.
Acceptance-Test Intent is required for every mode:
- `Behavior`: describe the new observable outcome in non-technical language. No test files, no tooling, no code.
- `Refactor`: name the existing acceptance tests that anchor the regression proof.
- `Chore`: name the actor and the observable outcome at their boundary.
</HARD-GATE>

---

## Process

1. **Read context** — `CONSTITUTION.md` and any prior `.agent/` files from the last cycle.
2. **Conversation: Discover the outcome** — ask one focused question at a time. Establish who experiences the problem, the triggering situation, the smallest valuable outcome, the expected failure or boundary behavior, and what must not be included. Do not infer a missing answer from conventions.
3. **Determine the Mode** — from the discussion, classify the increment as `Behavior`, `Refactor`, or `Chore`. If the work has no observable outcome anywhere (user, operator, developer, or CI), it is a `prototype`, not an increment.
4. **Slice by subtraction** — draft the criteria, remove each one in turn, and split any independently releasable value into a later increment. State the smallest-slice rationale and deferred outcomes.
5. **Draft Acceptance-Test Intent** — matched to the Mode:
   - `Behavior`: what a user sees when each criterion holds
   - `Refactor`: which existing acceptance tests must keep passing to prove behavior is preserved
   - `Chore`: which actor observes what outcome at which boundary
6. **Conversation: Propose the increment** — present the job story, mode, binary criteria, acceptance-test intent, out-of-scope list, and smallest-slice rationale. Iterate until the user explicitly confirms it.
7. **On approval** — write `.agent/increment.md` and create the branch: `git checkout -b increment/<slug>`.

---

## Checklist

- [ ] `CONSTITUTION.md` read
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
- [ ] User approval received
- [ ] `.agent/increment.md` written
- [ ] Branch created: `git checkout -b increment/<slug>`

---

## Handoff

Terminal artifact: `.agent/increment.md`
Next skill: `4dc-plan` — load `skills/plan/SKILL.md`
