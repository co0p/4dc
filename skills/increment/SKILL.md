---
name: 4dc-increment
argument-hint: "short increment intent, e.g. 'feature: export weekly summary'"
description: "Use after CONSTITUTION.md exists and before any implementation. Defines the next narrow, testable behavior change — WHAT and WHY only, no technical detail."
---

# Increment Skill

## One Responsibility

Define one small, outcome-focused increment with measurable acceptance criteria and explicit out-of-scope boundaries.

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
- **Acceptance criteria**: 2–5 binary, verifiable conditions derived from the use case
- **Out of scope**: explicit exclusions that prevent scope creep
- **Constitution constraints**: which guardrails apply to this increment
- **Roadmap entry**: the feature name and job story to add to `docs/roadmap.md` Partial section

---

<HARD-GATE>
Do NOT write acceptance criteria before the use case is stated — criteria must derive from the job story.
Do NOT include technical design, file names, implementation approaches, or coding detail in increment.md.
Do NOT start a plan or any implementation work during this phase.
Do NOT approve an increment with vague acceptance criteria ("works correctly", "feels right").
One increment per cycle — if scope expands, split into separate increments.
</HARD-GATE>

---

## Process

1. **Read context** — `CONSTITUTION.md`, `docs/roadmap.md`, any prior `.agent/` files from the last cycle
2. **Elicit the use case** — ask the customer for a job story (_When / I want / So that_); if they provide only a vague intent, help them shape it into a job story before proceeding
3. **Derive acceptance criteria** from the job story — ask 1–2 clarifying questions if criteria are not yet binary
4. **Generate `.agent/increment-review.html`** — show use case, goal, acceptance criteria, and roadmap entry
5. **STOP** — wait for explicit approval or revision requests
6. **On approval** — write `.agent/increment.md` and move the feature to Partial in `docs/roadmap.md`

---

## Checklist

- [ ] `CONSTITUTION.md` and `docs/roadmap.md` read
- [ ] Use case (job story) stated by customer
- [ ] Acceptance criteria derived from use case, not from technical assumptions
- [ ] Acceptance criteria are binary and verifiable
- [ ] Out-of-scope list is non-empty
- [ ] Roadmap entry (feature name + job story) identified
- [ ] HTML review generated and shown
- [ ] User approval received
- [ ] `.agent/increment.md` written
- [ ] `docs/roadmap.md` updated: feature moved to Partial

---

## Handoff

Terminal artifact: `.agent/increment.md`
Next skill: `4dc-plan` — load `skills/plan/SKILL.md`
