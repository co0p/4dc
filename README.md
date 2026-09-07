<p align="center">
  <picture>
    <source srcset="assets/4dc-logo-traced-dark.svg" media="(prefers-color-scheme: dark)">
    <img src="assets/4dc-logo-traced.svg" alt="4DC logo" width="120" />
  </picture>
</p>

<p align="center"><strong>4dc – four discipline cycle</strong></p>
<p align="center"><em>A skill-based methodology for emergent, test-driven development with any coding agent.</em></p>

---

> [!CAUTION]
> DO NOT TRUST AI, and do not trust others' prompts. Running prompts from untrusted sources is a security risk. Watch https://media.ccc.de/v/39c3-agentic-probllms

## What is 4dc?

4dc is a **four-discipline cycle** grounded in Extreme Programming, Lean Software Development, and use-case thinking. It is delivered as a set of composable **skills** — one bounded responsibility and defined artifact set per phase.

Each skill has a single responsibility, a defined input, and a hard gate before it produces output. Skills are composable: use them individually via any prompt-capable agent, or use the included **orchestrator** (`AGENTS.md`) to automate phase detection and skill loading.

**Core bets:**
- Premature implementation is the root cause of most rework. Separate WHAT, HOW, and EXECUTION before touching code.
- Files are the source of truth between phases, not memory or conversation.
- No phase is complete without an observable artifact or test result.
- Evidence over claims — never "done" without proof.

---

## Installation

From the root of your project:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install-4dc.sh)"
```

This installs:

```
.agents/
  skills/
    constitution/SKILL.md
    increment/SKILL.md
    plan/SKILL.md
    implement/SKILL.md
    promote/SKILL.md
  AGENTS.md                 ← orchestrator (auto-detects phase, loads the right skill)
.agent/                     ← working directory (gitignored)
```

To regenerate the checked-in skill files from the template sources while maintaining this repository, run:

```bash
./scripts/generate-4dc.sh
```

---

## Two ways to use 4dc

### Option A — Orchestrator (recommended)

`.agents/AGENTS.md` is the installed entry point. Agent support for automatic instruction discovery varies; when it is not discovered automatically, explicitly attach or ask the agent to read `.agents/AGENTS.md` before starting.

The orchestrator inspects your workspace and determines the current phase:

| Condition | Phase | Loads |
|-----------|-------|-------|
| No `CONSTITUTION.md` | constitution | `.agents/skills/constitution/SKILL.md` |
| `CONSTITUTION.md` exists, no `.agent/increment.md` | increment | `.agents/skills/increment/SKILL.md` |
| `.agent/increment.md` exists, no `.agent/plan.md` | plan | `.agents/skills/plan/SKILL.md` |
| `.agent/plan.md` exists, implementation not complete | implement | `.agents/skills/implement/SKILL.md` |
| `.agent/implementation.md` status: complete | promote | `.agents/skills/promote/SKILL.md` |

You can also name a phase explicitly — the orchestrator loads it directly without checking conditions.

Stop gates are enforced: the orchestrator will not advance until the current phase has produced its artifact **and** you have explicitly approved it. Silence is not approval.

### Option B — Skills directly

Reference any skill file in your agent's context to run that phase in isolation:

```
# GitHub Copilot Chat
@workspace #.agents/skills/increment/SKILL.md "feature: add CSV export"

# Claude Code
/read .agents/skills/plan/SKILL.md

# Any prompt-capable agent
Attach .agents/skills/implement/SKILL.md to your context, then describe what to build
```

Useful when integrating 4dc into an existing workflow or running a single phase without the full cycle.

---

## The cycle

```
CONSTITUTION.md → .agent/increment.md → .agent/plan.md → .agent/implementation.md → docs/
```

Each arrow is a stop gate — explicit human approval before the next skill starts.

---

## Skills

### ⚖️ constitution — *Constitutional lawyer*

Asks probing questions, synthesises project-specific constraints into enforceable rules, and refuses to write a law that can't be tested. Never prescribes implementation — only guardrails.

**Responsibility:** Create or update `CONSTITUTION.md` and supporting documentation — the project's durable engineering guardrails.

**Input:** Existing `CONSTITUTION.md` (if any), `README.md`, project structure, existing docs, testing/deployment practices, ADRs.

**Output:**
- `CONSTITUTION.md` with engineering principles, architectural boundaries, testing strategy, documentation rules, and the `.agent/` lifecycle contract
- `docs/testing.md` — testing strategy, types, scope, gates, naming conventions
- `docs/deployment.md` — release triggers, versioning, deployment targets, rollback, monitoring
- `docs/adr/` — architecture decision records (if foundational decisions exist)
- `docs/roadmap.md` — feature tracking template (if not present)

**SDLC ideology:** XP — team agreements before code. Lean — eliminate ambiguity upstream. Every rule must be verifiable; no abstract slogans, no generic internet copy-paste.

**Hard gates:** No output until markdown review (`.agent/constitution-review.md`) is approved. No implementation detail. No rule that isn't grounded in this project's specific context.

---

### 🔬 increment — *Product analyst*

Relentlessly scopes down, insists on measurable acceptance criteria, refuses to let technical ideas contaminate the WHAT. If the slice is too large, it splits — no exceptions.

**Responsibility:** Define one small, outcome-focused increment: WHAT and WHY only, no technical detail.

**Input:** `CONSTITUTION.md`, `docs/roadmap.md`, customer job story (When / I want / So that), user intent.

**Output:** `.agent/increment.md` — use case, single-sentence goal, 2–5 binary acceptance criteria, explicit out-of-scope list, applicable constitution constraints, roadmap entry.

**SDLC ideology:** XP user stories with binary acceptance criteria. Lean smallest shippable slice. Use-case thinking: observable user outcome — no technical vocabulary in this artifact.

**Hard gates:** No file names, approaches, or code. No vague acceptance criteria. One increment per cycle. Acceptance criteria must derive from the use case, not be invented during this phase.

---

### 🏗️ plan — *Senior architect in planning mode*

Reads the codebase before saying anything, surfaces risks, writes subtasks precise enough that a junior can execute them without ambiguity. Stops before touching a file.

**Responsibility:** Define HOW — an ordered sequence of verifiable subtasks traceable to acceptance criteria.

**Input:** `CONSTITUTION.md`, `.agent/increment.md` (approved), current codebase structure and relevant source files.

**Output:** `.agent/plan.md` — goal, 2–3 sentence strategy, ordered subtasks (each marked `[research]`, `[tidy]`, or `[behavior]`) with verification step and dependency mapping, risks.

**SDLC ideology:** XP planning game — tasks sized for one focused session. Lean: pull from acceptance criteria, not push from ideas. Full traceability: requirement → subtask → test.

**Hard gates:** No code, no file edits. No subtask without a verification step. Every acceptance criterion must have a covering subtask. Markdown review (`.agent/plan-review.md`) must be approved before writing plan.md.

---

### 🔨 implement — *Disciplined craftsperson*

Follows the plan unless there is a documented reason not to, writes the failing test before touching production code, records every surprise immediately. Does not claim completion without proof.

**Responsibility:** Execute the plan in Red → Green → Refactor order with continuous progress tracking.

**Input:** `CONSTITUTION.md`, `.agent/increment.md`, `.agent/plan.md`, current test baseline.

**Output:**
- `.agent/implementation.md` — live progress log updated after each subtask; final `status: complete` or `status: blocked`
- `.agent/learnings.md` — decisions, deviations, surprises, and promote candidates

**SDLC ideology:** XP strict TDD — test first, always. Lean: minimal code to pass the test, then refactor. Evidence over claims. Continuous state — never batch-update progress.

**Hard gates:** No production code before a failing test. No subtask complete without objective evidence. No skipping subtasks without recording the reason.

Before marking implementation complete, the agent creates `.agent/implementation-review.md`, presents the final evidence and risks, and waits for explicit approval.

#### Tidy First within implement

When the plan contains `[tidy]` subtasks, implement runs them **before any `[behavior]` subtasks**, in two distinct loops:

**Loop 1 — Tidy (structural changes only)**

Each `[tidy]` subtask is a pure structural change: rename, extract, reorganise, inline — no new observable behavior. The rule is strict: a `[tidy]` commit must leave every test green, and must not mix with any behavior change. Commit each tidy subtask separately with a message that reads `tidy: <what changed>`.

This is Kent Beck's *Tidy First?* principle in practice: make the code easier to change, **then** change it. The `[tidy]` subtasks in the plan were already justified at the plan phase — if tidying doesn't directly serve the behavior change that follows, it doesn't belong in the plan.

**Loop 2 — Behavior (Red → Green → Refactor)**

After all `[tidy]` subtasks are committed, implement switches to `[behavior]` subtasks. Each one starts with a failing test. Constrain the AI's context to the minimum needed for that subtask — narrow context, narrow blast radius. Commit each passing subtask with a message that reads `feat: <what changed>` or `fix: <what changed>`.

**Why separate commits matter:** tidy commits can be reviewed, reverted, or cherry-picked independently of behavior changes. Mixed commits obscure intent and make code review harder. The separation is a communication discipline, not a formality.

---

### 📚 promote — *Librarian and retrospective facilitator*

Reads everything, judges what's worth keeping, proposes each promotion with a destination and rationale, only writes after explicit approval. Leaves the workspace clean.

**Responsibility:** Merge durable outcomes from `.agent/` into permanent project artifacts and close the cycle.

**Input:** `CONSTITUTION.md`, `.agent/increment.md`, `.agent/plan.md`, `.agent/implementation.md` (status: complete), `.agent/learnings.md`, existing `docs/`, ADR log.

**Output (per approval):** Updated `CONSTITUTION.md` (if guardrails need revision), new ADRs in `docs/adr/`, updated `docs/architecture.md` / `docs/domain.md` / `README.md` (for behavior or structure changes), updated `docs/roadmap.md` (feature moved from Partial to Done with test link), cleaned `.agent/`.

**SDLC ideology:** Lean — only promote what's verified. XP retrospective embedded in the delivery cycle. `.agent/` is scratchpad; `docs/` is truth.

**Promotion categories:**

| Type | Trigger | Destination |
|------|---------|-------------|
| Architecture decision | Non-obvious choice with lasting impact | `docs/adr/ADR-<date>-<slug>.md` |
| Guardrail update | Constitution rule violated or needs clarification | `CONSTITUTION.md` |
| Behaviour change | Public API, CLI, or user-facing behaviour changed | `README.md` |
| Test pattern | New approach worth standardising | `CONSTITUTION.md` testing section |
| Known issue | Found but not fixed this cycle | `docs/known-issues.md` |

**Hard gates:** No permanent doc until each candidate is individually approved. No promotion of unverified work. No `.agent/` cleanup until all promotions are confirmed written. Markdown review (`.agent/promotion-review.md`) must be approved before writing permanent docs.

---

## File contracts

```
CONSTITUTION.md              permanent · root · constitution writes it
.agent/increment.md          transient · per cycle · increment writes it
.agent/plan.md               transient · per cycle · plan writes it
.agent/implementation.md     transient · per cycle · implement writes it
.agent/learnings.md          transient · per cycle · implement appends to it
```

`.agent/` is gitignored. After promote completes, `.agent/` is cleared for the next cycle.

---

## Why This Matters

| Principle | What it prevents |
|-----------|------------------|
| **WHAT before HOW before CODE** | Building the wrong thing well. Scope is locked before implementation starts. |
| **Files are the source of truth** | Decisions living only in chat history or someone's memory. |
| **Evidence over claims** | Features marked "done" that don't actually pass their tests. |
| **Human approval at every gate** | Scope creep and silent, unreviewed changes. |
| **Durable docs, promoted deliberately** | Documentation rot — docs are only updated when something durable actually changed. |

---

## What Gets Created

### 📚 Permanent Documentation

These documents survive every cycle and form the project's permanent knowledge base:

| Document | Purpose |
|----------|---------|
| `CONSTITUTION.md` | Engineering guardrails — testing, architecture, release policy. Created once, revised rarely. |
| `docs/testing.md` | How we test, what needs coverage, the gate every feature must pass. |
| `docs/deployment.md` | How we release, roll back, and monitor. |
| `docs/architecture.md` | How the system is structured; updated when structure changes. |
| `docs/domain.md` | Shared vocabulary — the words the team and the product agree on. |
| `docs/adr/*.md` | A permanent record of significant decisions and why they were made. |
| `docs/roadmap.md` | Feature status with proof — the single source of truth for "what's shipped." |

### ⏳ Working Files (Temporary)

These files live in `.agent/` during the cycle and are cleared after promote:

| Document | Purpose |
|----------|---------|
| `.agent/increment.md` | This cycle's feature definition — goal and acceptance criteria. |
| `.agent/plan.md` | This cycle's technical execution steps. |
| `.agent/implementation.md` | Live progress log with evidence for every step. |
| `.agent/learnings.md` | Decisions and surprises captured for the promote phase. |

---

## Getting started

```bash
# 1. Install
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install-4dc.sh)"

# 2. Open your agent — AGENTS.md is detected automatically.
#    It will see that CONSTITUTION.md is missing and load the constitution skill.

# 3. Approve the constitution, then your agent moves to increment.
#    Or reference a skill directly:
#    @workspace #.agents/skills/increment/SKILL.md "feature: add CSV export"
```

---

## Real-World Example: GoPomodoro

[**gopomodoro**](https://github.com/co0p/gopomodoro) is a minimal Pomodoro timer built using 4dc from scratch. It demonstrates how 4dc delivers on its promise of emergent, test-driven development.