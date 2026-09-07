---
name: 4dc-increment
argument-hint: "short increment intent, e.g. 'feature: export weekly summary'"
description: "Use after CONSTITUTION.md exists and before any implementation. Defines the next narrow, testable changeset — WHAT and WHY only, no technical detail."
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
- **Acceptance criteria**: 2–5 binary, verifiable conditions derived from the use case; each criterion names observable proof rather than an implementation mechanism
- **Out of scope**: explicit exclusions that prevent scope creep
- **Constitution constraints**: which guardrails apply to this increment
- **Roadmap entry**: the feature name and job story to add to `docs/roadmap.md` Partial section

Required `.agent/increment.md` headings:
- `## Use Case`
- `## Goal`
- `## Acceptance Criteria`
- `## Out Of Scope`
- `## Constitution Constraints`
- `## Roadmap Entry`

## Execution Contract

- Produce only the artifact for this phase. Do not leak work from a later phase into this one.
- Treat tests, architecture notes, ADRs, and user-facing docs as first-class communication artifacts.
- Gather only enough context to identify the governing constraints, the target artifact, and the cheapest validation step. Then act.
- Resolve conflicts in this order: explicit user approval, approved prior-phase artifacts, `CONSTITUTION.md`, this skill.
- Low-risk actions: reads, searches, diffs, and local validation commands.
- Medium-risk actions: local reversible edits to phase artifacts.
- High-risk actions: destructive file operations, external side effects, or skipping a stop gate. Require explicit approval first.
- If a required input is missing or contradictory, ask one focused question or stop at the review gate. Do not invent missing facts.
- Before finishing, run the phase checklist and confirm every required section is present.

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

1. **Read context** — `CONSTITUTION.md`, `docs/roadmap.md`, and any prior `.agent/` files from the last cycle.
2. **Conversation: Elicit and propose the increment** — shape the job story, derive binary criteria, propose the scope, and iterate until the user confirms it.
3. **Generate `.agent/increment-review.md`** — include the required Markdown review sections, use case, criteria, out-of-scope list, and roadmap entry.
4. **STOP** — present the review and wait for explicit approval.
5. **On approval** — write `.agent/increment.md` and move the feature to Partial in `docs/roadmap.md`.

## Markdown Review Contract

Use `.agent/increment-review.md`. Include **Objective**, **Inputs Reviewed**, **Proposed Output Summary**, **Risks and Trade-offs**, **Open Questions**, and **Approval Decision**. An explicit conversational approval is sufficient; record it in the Approval Decision section.

---

## Checklist

- [ ] `CONSTITUTION.md` and `docs/roadmap.md` read
- [ ] Use case (job story) stated by customer
- [ ] Acceptance criteria derived from use case, not from technical assumptions
- [ ] Acceptance criteria are binary and verifiable
- [ ] Out-of-scope list is non-empty
- [ ] Roadmap entry (feature name + job story) identified
- [ ] Markdown review generated and shown
- [ ] User approval received
- [ ] `.agent/increment.md` written
- [ ] `docs/roadmap.md` updated: feature moved to Partial

---

## Handoff

Terminal artifact: `.agent/increment.md`
Next skill: `4dc-plan` — load `skills/plan/SKILL.md`