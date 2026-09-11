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

- **Poppendieck — smallest shippable slice.** Pull the next slice from value, not from a backlog of ideas. The slice must be small enough to deliver in one cycle and verify end-to-end.
- **Beck — stories with binary acceptance.** A story is a promise of a conversation, but the acceptance criteria must be binary: it either passes or it does not. No "mostly works."
- **Poppendieck — decide as late as possible.** Do not fix the HOW here. The increment states the WHAT and WHY; the technical approach waits for the plan, where it can be informed by the codebase.

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
- **Acceptance criteria**: 2–5 binary, verifiable conditions derived from the use case; each criterion names observable proof rather than an implementation mechanism
- **Acceptance-test intent**: optional user-journey scenarios for larger increments; these clarify end-to-end evidence but are not a required artifact or completion gate unless `CONSTITUTION.md` says so
- **Out of scope**: explicit exclusions that prevent scope creep
- **Constitution constraints**: which guardrails apply to this increment
- **Roadmap entry**: the feature name and job story to add to `docs/roadmap.md` Partial section

Required `.agent/increment.md` headings:
- `## Use Case`
- `## Goal`
- `## Acceptance Criteria`
- `## Acceptance-Test Intent` (optional)
- `## Out Of Scope`
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
Acceptance tests are optional. Do not make an increment larger just to add them. If used, describe the user journey they would prove; do not prescribe test tooling here.
</HARD-GATE>

---

## Process

1. **Read context** — `CONSTITUTION.md`, `docs/roadmap.md`, and any prior `.agent/` files from the last cycle.
2. **Conversation: Elicit and propose the increment** — shape the job story, derive binary criteria, propose the scope, and iterate until the user confirms it.
3. **On approval** — write `.agent/increment.md` and move the feature to Partial in `docs/roadmap.md`.

---

## Checklist

- [ ] `CONSTITUTION.md` and `docs/roadmap.md` read
- [ ] Use case (job story) stated by customer
- [ ] Acceptance criteria derived from use case, not from technical assumptions
- [ ] Acceptance criteria are binary and verifiable
- [ ] Acceptance-test intent recorded when the increment is large enough to benefit from an end-to-end scenario (optional)
- [ ] Out-of-scope list is non-empty
- [ ] Roadmap entry (feature name + job story) identified
- [ ] User approval received
- [ ] `.agent/increment.md` written
- [ ] `docs/roadmap.md` updated: feature moved to Partial

---

## Handoff

Terminal artifact: `.agent/increment.md`
Next skill: `4dc-plan` — load `skills/plan/SKILL.md`
