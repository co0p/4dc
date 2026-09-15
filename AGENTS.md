# 4dc Agent Orchestrator

You operate under the **4dc methodology** — a four-discipline cycle:

```
constitution → increment → [prototype?] → plan → implement → [tidy] → tdd-red → tdd-green → refactor → promote
                   |                                              ↑___________________________________|
                   |                                              [adr?]              [research → tdd-green]
                   ↓
            branch: increment/<slug>
            (all subtask commits land here)
                   ↓
               promote:
            1. docs promotion
            2. final tidy pass
            3. squash-merge → main, OR push branch → PR
```

- `prototype` is optional — load it only when a blocking unknown needs a throwaway spike before planning.
- `adr` is on-demand — load it whenever a structural, hard-to-reverse decision emerges during any phase.
- `implement` runs exactly once per cycle — it scaffolds `.agent/implementation.md` and populates the todo list.
- `tidy`, `tdd-red`, `tdd-green`, and `refactor` loop per subtask until all are `state: complete`.

Read this file completely before doing any work, then load the skill for the current phase.

---

## Core Principles

These apply across all phases. Skills do not repeat them.

The methodology draws on three traditions: **Kent Beck** (XP — test-first, team agreements, Tidy First, the planning game), **Mary Poppendieck** (Lean Software Development — eliminate waste, decide as late as possible, pull from value, small batches), and **Martin Fowler** (evolutionary architecture, refactoring as behavior-preserving design improvement, two-hats discipline).

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
- **Use the internal todo list throughout the implementation loop.** The `implement` skill populates it; every subsequent skill keeps it current. Mark `in_progress` when a subtask starts, `completed` only when `implementation.md` records `state: complete` with a commit hash. One item `in_progress` at a time. The todo list is the user's live progress view — accuracy is not optional.

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
| No `CONSTITUTION.md` | **constitution** | `skills/constitution/SKILL.md` |
| `CONSTITUTION.md` exists, no `.agent/increment.md` | **increment** | `skills/increment/SKILL.md` |
| `.agent/increment.md` exists, user requests a spike | **prototype** *(optional)* | `skills/prototype/SKILL.md` |
| `.agent/increment.md` exists, no `.agent/plan.md` | **plan** | `skills/plan/SKILL.md` |
| User names a structural decision to capture | **adr** *(on-demand)* | `skills/adr/SKILL.md` |
| `.agent/plan.md` exists, no `.agent/implementation.md` | **implement** | `skills/implement/SKILL.md` |
| `.agent/plan.md` exists, current `[tidy]` subtask `state: pending` | **tidy** | `skills/tidy/SKILL.md` |
| `.agent/plan.md` exists, current `[behavior]` subtask `state: pending` | **tdd-red** | `skills/tdd-red/SKILL.md` |
| `.agent/plan.md` exists, current `active_test` `state: red` or `[research]` `state: pending` | **tdd-green** | `skills/tdd-green/SKILL.md` |
| `.agent/plan.md` exists, current `active_test` `state: green` | **refactor** | `skills/refactor/SKILL.md` |
| `.agent/implementation.md` marked `status: complete` | **promote** | `skills/promote/SKILL.md` |

**If the user explicitly names a phase, load that skill directly without checking conditions.**

**Implementation loop:** during the implementation phase, read `.agent/implementation.md` to find the current subtask by type and state:
- `[tidy]` + `state: pending` → load `tidy`
- `[behavior]` with `active_test: pending` → load `tdd-red`
- `[behavior]` with `active_test: red` → load `tdd-green` (writes minimal code to pass, sets active test `state: green`)
- `[behavior]` with `active_test: green` → load `refactor` (improves design, sets active test `state: complete`, activates the next test or completes the subtask)
- `[research]` + `state: pending` → load `tdd-green` (investigates, sets `state: complete`)
- All subtasks `state: complete` → `refactor` runs final verification and sets `status: complete`

At every subtask transition: mark the current todo item `in_progress` before starting, `completed` only after `implementation.md` records `state: complete` with a commit hash.

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
| `.agent/prototype.md` | transient, per cycle (optional) | `prototype` skill writes it |
| `.agent/plan.md` | transient, per cycle | `plan` skill writes it |
| `.agent/implementation.md` | transient, per cycle | `implement` skill creates it; `tidy`, `tdd-red`, `tdd-green`, and `refactor` update it |
| `.agent/learnings.md` | transient, per cycle | `tidy`, `tdd-red`, `tdd-green`, `refactor`, and `adr` skills append to it |
| `docs/adr/ADR-*.md` | permanent | `adr` skill writes it |

All `.agent/` files are lowercase. The `.agent/` directory is gitignored by default.

---

## Approval

Before writing a phase's final artifact, propose the outcome in the conversation and pause for explicit approval. This applies in every phase.

**Approval semantics:** An explicit user statement in the conversation, such as "looks good" or "proceed," is approval. Silence is not approval. When in doubt, ask.

---

## Skill Loading

After determining the current phase, read the skill file fully before beginning:

```
Read skills/<phase>/SKILL.md now.
```

The skill file contains the detailed process. This file handles orchestration only — phase detection, stop gates, shared contracts.
