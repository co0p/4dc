<p align="center">
  <img src="assets/4dc-logo-traced.svg" alt="4DC logo" width="120" />
</p>

<p align="center"><strong>4dc — four discipline cycle</strong></p>
<p align="center"><em>A small-step workflow for building software with coding agents.</em></p>

> Do not trust AI output blindly. Review every proposed change before you accept it.

## What is 4DC?

4DC keeps a coding agent from jumping straight from a vague request to a large code change. It separates the work into four phases, each owned by a dedicated agent.

```
[idea] ──► [shape] ──► [craft] ──► [reflect]
  ^                                    │
  └────────────────────────────────────┘
```

Each agent hands off to the next. Nothing is built until the idea is clear. Nothing is planned in detail before building starts. Nothing is closed without reflection.

## Inside craft

The `craft` agent does not follow a plan. It finds the next step through conversation and repeats this cycle until the use case goal is met.

```
discuss ──► agree ──► tdd ──► review ──► repeat
```

**Discuss.** The agent looks at the use case goal and the current code, then suggests the next meaningful step. You respond, redirect, or add context.

**Agree.** The agent states the step in one sentence. You confirm it before any code is written.

**TDD.** The agent writes a failing test, shows the failure, writes the minimum code to pass, then refactors if needed.

**Review.** You review the result. Leave `// @review: <instruction>` comments on anything that needs changing. Approve to commit and repeat.

## The four agents

| Agent | Phase | What it does |
|---|---|---|
| `idea` | Idea | Explores the problem with you at the domain level. No code, no files, no implementation. |
| `shape` | Shape | Turns the domain concept into a use-case skeleton using Cockburn use cases. |
| `craft` | Build | Enriches the skeleton step by step. This is done in the craft cycle: discuss, agree, tdd, review, repeat. |
| `reflect` | Reflect | Disperses learnings into your project docs, runs a short retrospective, and closes the cycle. |

## How to install

Copy the `.opencode/` directory and `AGENTS.md` into your project root.

```bash
cp -r /path/to/4dc/.opencode /your/project/
cp /path/to/4dc/AGENTS.md /your/project/
```

Add `.agent/` to your project's `.gitignore`. That directory holds temporary cycle files and must never be committed.

```bash
echo ".agent/" >> .gitignore
```

If you use the `opencode.jsonc` from this repository, copy it too. It registers the four agents with OpenCode and sets their display colors.

```bash
cp /path/to/4dc/opencode.jsonc /your/project/
```

## How to start

Tell OpenCode to read `AGENTS.md` and load the `idea` agent.

```text
Read AGENTS.md and use the idea agent.
```

The orchestrator in `AGENTS.md` detects your current cycle state and tells you which agent to load next. If you are starting fresh, it points you to `idea`.

## Cycle files

The agents write temporary files to `.agent/` during a cycle. These files are the memory between agents. Do not edit them by hand unless you know what you are doing.

| File | Written by | Consumed by |
|---|---|---|
| `.agent/domain-concept.md` | `idea` | `shape` |
| `.agent/cycle-plan.md` | `shape` | `craft` |
| `.agent/learnings.md` | `craft` | `reflect` |

The `reflect` agent deletes all three files when the cycle closes.

## Scorecard

`SCORECARD.md` contains 27 yes/no questions grouped by phase. Use it at the end of a cycle or whenever the work feels off. Each question targets a specific drift pattern. The scoring at the bottom tells you how to respond.

## What 4DC does not do

4DC has no runtime dependency. It does not install Node.js, Python, a test runner, or an MCP server. Your project must already provide whatever testing and tooling it needs.

4DC does not create documentation upfront. Docs are created by the `reflect` agent when a cycle produces something worth recording, and only in locations consistent with your project's existing conventions.

## Influences

4DC borrows from people who shaped modern software development.

| Influence | `idea` | `shape` | `craft` | `reflect` |
|---|---|---|---|---|
| **Cockburn** | Express the goal as an actor-goal pair; the domain rule must not depend on delivery mechanism | Summary level first; casual level when the build pulls it; name ports before naming adapters | Enrich toward the main success scenario; implement toward the port, keep the framework outside the domain | Check whether every use case reached its stated goal |
| **Kent Beck** | One idea at a time; stop when the concept is clear | The skeleton is a sketch, not a contract | Red-green-refactor within each step; simplest thing that works | Ask what changed about how we test next time |
| **Poppendieck** | Decide late; do not lock in a solution during ideation | Defer detail; keep the skeleton small | Pull from the use case; do not push from a plan | Identify waste in the cycle; propose removing it |
| **Fowler** | Name the bounded context before naming any behavior | Draw the rough boundary; flag coupling risks | Refactor after green; keep the structure honest | Update the architecture description to match what was built |
| **Evans** | Use and test the Ubiquitous Language | Every use case step uses domain language | Implement only what the domain model calls for | Record new or changed domain terms |
| **Crystal** | The conversation is the artifact at this stage | Enough design to start | Stop after each action; human feedback drives the next step | Short retrospective: what slowed us, what helped |
| **Dave Thomas** | Name the problem before naming the solution; a tracer bullet starts with a clear target | Design orthogonally; keep concerns separate at the boundary | Tracer bullet first: thin end-to-end before filling in detail; DRY within each step | Fix broken windows found during the cycle before closing it |
