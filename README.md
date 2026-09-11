<p align="center">
  <img src="assets/4dc-logo-traced.svg" alt="4DC logo" width="120" />
</p>

<p align="center"><strong>4dc - four discipline cycle</strong></p>
<p align="center"><em>A small-step workflow for building software with coding agents.</em></p>

> Do not trust AI output blindly. Review every proposed change, command, and test result.

## Why 4dc?

4dc keeps a coding agent from jumping straight from a vague request to a large code change. It separates:

```text
what the user needs -> how to build it -> one small change -> proof -> durable knowledge
```

This gives a junior developer a clear path:

- The request becomes one small, testable increment.
- The plan names the files, symbols, tests, and risks before code changes.
- The agent works in small steps instead of guessing across the repository.
- Tests and handover files show what happened and why.
- Important decisions are preserved instead of being lost in chat.

4dc borrows practical ideas from people who shaped modern software development:

- **Kent Beck:** test-first development, small steps, Tidy First, and the planning game.
- **Mary Poppendieck:** remove waste, pull work from user value, decide late when possible, and keep batches small.
- **Martin Fowler:** evolutionary architecture, explicit architecture decisions, and refactoring as behavior-preserving design improvement.

The skills are not a replacement for judgment. They are guardrails that make good judgment easier to apply repeatedly.

## Installation

Run this from the root of the project that will use 4dc:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install-4dc.sh)"
```

The installer requires:

- Bash 3.2 or newer
- `curl`
- `tar`
- `mktemp`
- `grep`
- A coding agent that can read Markdown instructions and run the project's commands

4dc has no runtime package. It does not install Node.js, Python, a test runner, a language toolchain, an MCP server, or an agent application. The project must already provide whatever its own testing and deployment documentation requires.

The installer needs network access once to download the repository. After installation, the skills work from local files.

It installs:

```text
.agents/
  AGENTS.md                 orchestrator
  skills/
    constitution/SKILL.md
    increment/SKILL.md
    prototype/SKILL.md
    plan/SKILL.md
    adr/SKILL.md
    tidy/SKILL.md
    tdd-red/SKILL.md
    tdd-green/SKILL.md
    refactor/SKILL.md
    promote/SKILL.md
.agent/                     temporary handover directory
```

`.agent/` is added to `.gitignore`. Existing `.agents/AGENTS.md` is preserved by default. To update it explicitly:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install-4dc.sh)" -- --force
```

The skill files are updated on every installation. The installer fails if a required skill is missing instead of installing an incomplete workflow. It does not delete unrelated files in `.agents/`.

## Handover Documents

The files are the memory between skills. Read the current handover before acting. Do not rely on chat history alone.

### Permanent documents

| Document | Purpose |
|----------|---------|
| `CONSTITUTION.md` | Project rules: architecture, testing, performance, documentation, release. |
| `docs/testing.md` | Test strategy, commands, evidence, risks, and confidence gaps. |
| `docs/deployment.md` | Release, deployment, rollback, configuration, and health checks. |
| `docs/architecture.md` | Current system boundaries, containers, communication, and data stores. |
| `docs/domain.md` | Shared business terms, events, and rules. |
| `docs/ui.md` | Permanent UI, interaction, visual, accessibility, and content decisions. Required for projects with a UI. |
| `docs/adr/ADR-*.md` | Significant decisions, alternatives, rationale, and consequences. |
| `docs/roadmap.md` | User outcomes and delivery status backed by evidence. |

### Temporary cycle documents

| Document | Written by | Purpose |
|----------|------------|---------|
| `.agent/increment.md` | `increment` | The user outcome, acceptance criteria, and scope boundary. |
| `.agent/prototype.md` | `prototype` | Finding from an optional throwaway experiment. |
| `.agent/plan.md` | `plan` | Detailed file-level implementation map. |
| `.agent/implementation.md` | implementation skills | Current subtask, active test, state, and evidence. |
| `.agent/learnings.md` | implementation skills | Decisions, deviations, surprises, and promotion candidates. |

The temporary files are cleared or archived after `promote` completes. The plan is deliberately detailed so implementation skills can load only the files and references needed for the current subtask.

## The Cycle

```text
constitution -> increment -> [prototype?] -> plan -> [tidy?]
                                      -> tdd-red -> tdd-green -> refactor -> promote
```

`adr` is on-demand whenever a structural, hard-to-reverse, or non-obvious decision emerges. Each phase has an explicit approval gate. Silence is not approval.

The orchestrator reads the workspace and loads the next skill. If automatic discovery is unavailable, ask the agent to read `.agents/AGENTS.md` first.

## Skills By Phase

### 1. Constitution

**Skill:** `constitution`

Use when `CONSTITUTION.md` is missing or needs revision.

**Does:** Establishes project-specific engineering guardrails. It also creates the permanent documentation baseline when bootstrapping a project.

**Does not:** Choose the implementation for a feature.

**Handoff:** `CONSTITUTION.md` and supporting `docs/` files.

### 2. Increment

**Skill:** `increment`

Use after the constitution exists.

**Does:** Turns a request into one small user outcome with a job story, binary acceptance criteria, and explicit out-of-scope items.

**Does not:** Name files, classes, libraries, or implementation approaches.

Acceptance-test scenarios are optional. For a large feature, describe useful user journeys, but treat them as advisory unless the constitution or an explicit decision makes them required.

**Handoff:** `.agent/increment.md`.

### 3. Prototype (Optional)

**Skill:** `prototype`

Use only when a blocking unknown needs investigation before planning.

**Does:** Builds a time-boxed throwaway spike and records the finding.

**Does not:** Produce production code or become a shortcut around planning.

**Handoff:** `.agent/prototype.md`.

### 4. Plan

**Skill:** `plan`

Use after the increment, and after the prototype if one was needed.

**Does:** Reads the codebase broadly and creates the implementation map:

- Approach and architecture boundary
- Complete file list: new, modify, touch, delete
- Ordered `[research]`, `[tidy]`, and `[behavior]` subtasks
- Symbols, line ranges, and document references
- Test file and test cases for behavior work
- One `active_test` at a time for multi-case behavior subtasks
- Dependencies, risks, and acceptance-criteria coverage
- Optional, non-blocking acceptance scenarios for larger features

**Does not:** Edit production code.

**Handoff:** `.agent/plan.md`.

### 5. ADR (On Demand)

**Skill:** `adr`

Use when a decision changes a lasting boundary, dependency direction, technology choice, data model, or ownership of behavior.

**Does:** Records one decision, its context, alternatives, rationale, and consequences.

**Does not:** Document small implementation details or duplicate existing rules.

**Handoff:** `docs/adr/ADR-YYYYMMDD-<slug>.md`.

### 6. Implementation Loop

The implementation skills work one subtask at a time. They read only the files named by the plan.

#### Tidy

**Skill:** `tidy`

Executes one behavior-preserving structural change before behavior work. Tests stay green. Commit prefix: `tidy:`.

#### TDD Red

**Skill:** `tdd-red`

Writes one failing test for the current `active_test`. It confirms a real behavior failure and writes no production code.

#### TDD Green

**Skill:** `tdd-green`

Writes the minimum production code needed to pass the current test. It does not refactor. Behavior commits use `feat:` or `fix:`.

#### Refactor

**Skill:** `refactor`

Improves the design after green without changing behavior. Tests stay green. Commit prefix: `refactor:`. It activates the next test case or advances to the next subtask.

For multiple cohesive test cases:

```text
active test pending -> tdd-red -> red
red -> tdd-green -> green
green -> refactor -> complete
repeat for the next test case
```

### 7. Promote

**Skill:** `promote`

Use after all implementation subtasks pass final verification and `.agent/implementation.md` is approved as complete.

**Does:** Moves durable knowledge into permanent documentation, updates the roadmap, records ADRs or architecture changes, and cleans the temporary cycle files.

**Does not:** Promote guesses or unverified plans. Optional acceptance scenarios provide supplementary evidence and do not block by default.

## Direct Use

The orchestrator is recommended. You can also load a skill directly:

```text
Read `.agents/skills/increment/SKILL.md` and define the next increment.
Read `.agents/skills/plan/SKILL.md` and create the approved implementation plan.
Read `.agents/skills/tdd-red/SKILL.md` and handle the current active test.
```

Keep responses and artifacts direct: state the action, evidence, blocker, or handoff. Do not add generic motivation or repeat the handover contents.

## Maintainer Notes

This repository contains `templates/`, generated `skills/`, installer scripts, and `VALIDATION.md`. To regenerate skills after changing templates:

```bash
./scripts/generate-4dc.sh
```

`VALIDATION.md` is for maintaining this repository. It is not installed into consuming projects.
