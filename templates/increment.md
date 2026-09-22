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
- `docs/roadmap.md`
- User intent (one sentence or short phrase describing the desired outcome)
- **Customer-authored use case** — a job story written before acceptance criteria are defined

---

## Concrete Output

`.agent/increment.md` containing:
- **Use case**: job story in the form _"When [situation], I want to [action], so that [outcome]."_ Written or confirmed by the customer before criteria are defined.
- **Goal**: one sentence — the user-observable outcome (distilled from the use case)
- **Branch**: the git branch name for this increment, in the form `increment/<slug>` (e.g. `increment/token-refresh`)
- **Acceptance criteria**: 2–5 binary, verifiable conditions derived from the use case; each criterion names observable proof rather than an implementation mechanism
- **Acceptance-test intent**: optional user-journey scenarios for larger increments; these clarify end-to-end evidence but are not a required artifact or completion gate unless `CONSTITUTION.md` says so
- **Out of scope**: explicit exclusions that prevent scope creep
- **Smallest-slice rationale**: why removing any remaining criterion would make the outcome unusable, unverifiable, or not independently releasable
- **Constitution constraints**: which guardrails apply to this increment
- **Roadmap entry**: the feature name and job story to add to `docs/roadmap.md` Partial section

Required `.agent/increment.md` headings:
- `## Use Case`
- `## Goal`
- `## Branch`
- `## Acceptance Criteria`
- `## Acceptance-Test Intent` (optional)
- `## Out Of Scope`
- `## Smallest-Slice Rationale`
- `## Constitution Constraints`
- `## Roadmap Entry`

{{SHARED:execution-contract}}

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
Acceptance tests are optional. Do not make an increment larger just to add them. If used, describe the user journey they would prove; do not prescribe test tooling here.
</HARD-GATE>

---

## Process

1. **Read context** — `CONSTITUTION.md`, `docs/roadmap.md`, and any prior `.agent/` files from the last cycle.
2. **Conversation: Discover the outcome** — ask one focused question at a time. Establish who experiences the problem, the triggering situation, the smallest valuable outcome, the expected failure or boundary behavior, and what must not be included. Do not infer a missing answer from conventions.
3. **Slice by subtraction** — draft the criteria, remove each one in turn, and split any independently releasable value into a later increment. State the smallest-slice rationale and deferred outcomes.
4. **Conversation: Propose the increment** — present the job story, binary criteria, out-of-scope list, and smallest-slice rationale. Iterate until the user explicitly confirms it.
5. **On approval** — write `.agent/increment.md`, move the feature to Partial in `docs/roadmap.md`, and create the branch: `git checkout -b increment/<slug>`.

---

## Checklist

- [ ] `CONSTITUTION.md` and `docs/roadmap.md` read
- [ ] Use case (job story) stated by customer
- [ ] Acceptance criteria derived from use case, not from technical assumptions
- [ ] Acceptance criteria are binary and verifiable
- [ ] Acceptance-test intent recorded when the increment is large enough to benefit from an end-to-end scenario (optional)
- [ ] Out-of-scope list is non-empty
- [ ] User, trigger, outcome, failure boundary, and exclusions were explicitly discussed
- [ ] Subtraction test performed for every acceptance criterion
- [ ] Smallest-slice rationale explains why the remaining criteria cannot be split further
- [ ] Branch name derived (`increment/<slug>`)
- [ ] Roadmap entry (feature name + job story) identified
- [ ] User approval received
- [ ] `.agent/increment.md` written
- [ ] `docs/roadmap.md` updated: feature moved to Partial
- [ ] Branch created: `git checkout -b increment/<slug>`

---

## Handoff

Terminal artifact: `.agent/increment.md`
Next skill: `4dc-plan` — load `skills/plan/SKILL.md`
