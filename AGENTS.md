# 4dc Agent Orchestrator

You operate under the **4dc methodology** — a four-discipline cycle:

```
constitution → increment → plan → implement → promote
```

Read this file completely before doing any work, then load the skill for the current phase.

---

## Core Principles

These apply across all phases. Skills do not repeat them.

- Use plain, direct language. Keep output scannable.
- Ask focused questions; never a broad questionnaire.
- One clarifying question at a time. If evidence is missing, ask once.
- Source code and committed docs are the source of truth.
- Communication is the primary value stream: preserve intent, decisions, and evidence in files, tests, and permanent docs.
- Design before code: state the desired behavior, architectural boundary, and any performance-critical constraint before implementation starts.
- Tests before code for behavior changes. Structural tidying may come first, but only when it preserves observable behavior and keeps tests green.
- Documentation is part of the application. Keep architecture, domain language, testing guidance, and ADRs aligned with the implemented system.
- Project artifacts are product-specific. Do not copy this repository's internal workflow name, phase sequence, or orchestrator terminology into an application's `CONSTITUTION.md`, README, ADRs, or other project documentation.
- Never claim work is complete without objective evidence.
- Forward-only change: do not preserve backward compatibility unless explicitly requested.
- For work with more than three meaningful tasks or unknown dependencies: publish a short task plan, execute in verified steps, update progress after each step.

## Instruction Resolution

When instructions pull in different directions, resolve them in this order:

1. Explicit user approval and current request
2. Current phase stop gates and approved phase artifacts
3. `CONSTITUTION.md`
4. This orchestrator's defaults

If a conflict still cannot be resolved, ask one focused question.

## Action Risk Ladder

- Low risk: reads, searches, diffs, local validation commands
- Medium risk: local reversible edits to the current phase artifact
- High risk: destructive operations, external side effects, or skipping a review gate

High-risk actions require explicit approval.

---

## Phase Detection

Inspect the workspace and determine the current phase:

| Condition | Phase | Load skill |
|-----------|-------|------------|
| No `CONSTITUTION.md` | **constitution** | `.agents/skills/constitution/SKILL.md` |
| `CONSTITUTION.md` exists, no `.agent/increment.md` | **increment** | `.agents/skills/increment/SKILL.md` |
| `.agent/increment.md` exists, no `.agent/plan.md` | **plan** | `.agents/skills/plan/SKILL.md` |
| `.agent/plan.md` exists, implementation not complete | **implement** | `.agents/skills/implement/SKILL.md` |
| `.agent/implementation.md` marked `status: complete` | **promote** | `.agents/skills/promote/SKILL.md` |

**If the user explicitly names a phase, load that skill directly without checking conditions.**

---

## Stop Gates

You MUST NOT advance to the next phase until:
1. The current phase has produced its output artifact, AND
2. The user has explicitly approved it.

Silence is not approval. "Looks good" is approval. When in doubt, ask.

---

## Handoff Contracts

Files used as handoff contracts between phases:

| File | Scope | Owner |
|------|-------|-------|
| `CONSTITUTION.md` | permanent, root | `constitution` skill writes it |
| `.agent/increment.md` | transient, per cycle | `increment` skill writes it |
| `.agent/plan.md` | transient, per cycle | `plan` skill writes it |
| `.agent/implementation.md` | transient, per cycle | `implement` skill writes it |
| `.agent/learnings.md` | transient, per cycle | `implement` skill appends to it |

All `.agent/` files are lowercase. The `.agent/` directory is gitignored by default.

---

## Markdown Review Contract

Before writing a phase's final artifacts, create `.agent/<phase>-review.md`, show its contents, and pause for explicit approval. This applies in every phase.

**Workflow (MANDATORY):**
1. Discuss and refine the proposed outcome with the user.
2. Generate `.agent/<phase>-review.md`.
3. Show the review and STOP for explicit approval.
4. Record the approval in the review's Approval Decision section.
5. Only then write the phase's final artifacts.

**Required review sections:**
1. Objective
2. Inputs Reviewed
3. Proposed Output Summary
4. Risks and Trade-offs
5. Open Questions
6. Approval Decision

**Approval semantics:** An explicit user statement in the conversation, such as “looks good” or “proceed,” is approval. Record that decision in the review file; the user does not need to edit a checkbox themselves.

---

## Skill Loading

After determining the current phase, read the skill file fully before beginning:

```
Read .agents/skills/<phase>/SKILL.md now.
```

The skill file contains the detailed process. This file handles orchestration only — phase detection, stop gates, shared contracts.
