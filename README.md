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

4dc is a **four-discipline cycle** grounded in Extreme Programming, Lean Software Development, and use-case thinking. It is delivered as a set of composable **skills** — one skill per phase, one concrete output artifact per skill.

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
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install-4dc-skills.sh)"
```

This installs:

```
skills/
  constitution/SKILL.md
  increment/SKILL.md
  plan/SKILL.md
  implement/SKILL.md
  promote/SKILL.md
AGENTS.md                 ← orchestrator (auto-detects phase, loads the right skill)
.agent/                   ← working directory (gitignored)
```

---

## Two ways to use 4dc

### Option A — Orchestrator (recommended)

`AGENTS.md` is the entry point. Any compatible agent (GitHub Copilot, Claude Code, Cursor, etc.) reads it automatically on startup.

The orchestrator inspects your workspace and determines the current phase:

| Condition | Phase | Loads |
|-----------|-------|-------|
| No `CONSTITUTION.md` | constitution | `skills/constitution/SKILL.md` |
| `CONSTITUTION.md` exists, no `.agent/increment.md` | increment | `skills/increment/SKILL.md` |
| `.agent/increment.md` exists, no `.agent/plan.md` | plan | `skills/plan/SKILL.md` |
| `.agent/plan.md` exists, implementation not complete | implement | `skills/implement/SKILL.md` |
| `.agent/implementation.md` status: complete | promote | `skills/promote/SKILL.md` |

You can also name a phase explicitly — the orchestrator loads it directly without checking conditions.

Stop gates are enforced: the orchestrator will not advance until the current phase has produced its artifact **and** you have explicitly approved it. Silence is not approval.

### Option B — Skills directly

Reference any skill file in your agent's context to run that phase in isolation:

```
# GitHub Copilot Chat
@workspace #skills/increment/SKILL.md "feature: add CSV export"

# Claude Code
/read skills/plan/SKILL.md

# Any prompt-capable agent
Attach skills/implement/SKILL.md to your context, then describe what to build
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

**Responsibility:** Create or update `CONSTITUTION.md` — the project's durable engineering guardrails.

**Input:** Existing `CONSTITUTION.md` (if any), `README.md`, project structure.

**Output:** `CONSTITUTION.md` with engineering principles, architectural boundaries, testing strategy, documentation rules, and the `.agent/` lifecycle contract.

**SDLC ideology:** XP — team agreements before code. Lean — eliminate ambiguity upstream. Every rule must be verifiable; no abstract slogans, no generic internet copy-paste.

**Hard gates:** No output until HTML review is approved. No implementation detail. No rule that isn't grounded in this project's specific context.

---

### 🔬 increment — *Product analyst*

Relentlessly scopes down, insists on measurable acceptance criteria, refuses to let technical ideas contaminate the WHAT. If the slice is too large, it splits — no exceptions.

**Responsibility:** Define one small, outcome-focused increment: WHAT and WHY only, no technical detail.

**Input:** `CONSTITUTION.md`, user intent (one sentence).

**Output:** `.agent/increment.md` — single-sentence goal, 2–5 binary acceptance criteria, explicit out-of-scope list, applicable constitution constraints.

**SDLC ideology:** XP user stories with binary acceptance criteria. Lean smallest shippable slice. Use-case thinking: observable user outcome — no technical vocabulary in this artifact.

**Hard gates:** No file names, approaches, or code. No vague acceptance criteria. One increment per cycle.

---

### 🏗️ plan — *Senior architect in planning mode*

Reads the codebase before saying anything, surfaces risks, writes subtasks precise enough that a junior can execute them without ambiguity. Stops before touching a file.

**Responsibility:** Define HOW — an ordered sequence of verifiable subtasks traceable to acceptance criteria.

**Input:** `CONSTITUTION.md`, `.agent/increment.md`, current codebase structure.

**Output:** `.agent/plan.md` — goal, 2–3 sentence strategy, ordered subtasks each with a verification step and dependency mapping, risks.

**SDLC ideology:** XP planning game — tasks sized for one focused session. Lean: pull from acceptance criteria, not push from ideas. Full traceability: requirement → subtask → test.

**Hard gates:** No code, no file edits. No subtask without a verification step. Every acceptance criterion must have a covering subtask.

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

**Input:** All `.agent/` artifacts, existing `docs/`, ADR log.

**Output (per approval):** Updated `CONSTITUTION.md`, new ADRs in `docs/adr/`, updated `README.md`, cleaned `.agent/`.

**SDLC ideology:** Lean — only promote what's verified. XP retrospective embedded in the delivery cycle. `.agent/` is scratchpad; `docs/` is truth.

**Promotion categories:**

| Type | Trigger | Destination |
|------|---------|-------------|
| Architecture decision | Non-obvious choice with lasting impact | `docs/adr/ADR-<date>-<slug>.md` |
| Guardrail update | Constitution rule violated or needs clarification | `CONSTITUTION.md` |
| Behaviour change | Public API, CLI, or user-facing behaviour changed | `README.md` |
| Test pattern | New approach worth standardising | `CONSTITUTION.md` testing section |
| Known issue | Found but not fixed this cycle | `docs/known-issues.md` |

**Hard gates:** No permanent doc until that candidate is individually approved. No promotion of unverified work. No `.agent/` cleanup until all promotions are confirmed written.

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

## Getting started

```bash
# 1. Install
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install-4dc-skills.sh)"

# 2. Open your agent — AGENTS.md is detected automatically.
#    It will see that CONSTITUTION.md is missing and load the constitution skill.

# 3. Approve the constitution, then your agent moves to increment.
#    Or reference a skill directly:
#    @workspace #skills/increment/SKILL.md "feature: add CSV export"
```

---

## Real-World Example: GoPomodoro

[**gopomodoro**](https://github.com/co0p/gopomodoro) is a minimal Pomodoro timer built using 4dc from scratch. It demonstrates how 4dc delivers on its promise of emergent, test-driven development.