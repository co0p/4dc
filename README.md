<p align="center">
  <img src="assets/4dc-logo-traced.svg" alt="4DC logo" width="120" />
</p>

<p align="center"><strong>4dc — four discipline cycle</strong></p>
<p align="center"><em>A small-step workflow for building software with coding agents.</em></p>

> Do not trust AI output blindly. Review every proposed change before you accept it.

## What is 4DC?

4DC keeps a coding agent from jumping straight from a vague request to a large code change. It separates the work into four phases, each owned by a dedicated agent.

```
[ideate] ──► [shape] ──► [craft] ──► [reflect]
   ^                                     │
   └─────────────────────────────────────┘
```

Each agent hands off to the next. Nothing is built until the idea is clear. Nothing is planned in detail before building starts. Nothing is closed without reflection.

Every agent reports the result of each check it runs, even when the result is "none found". This way you always know what the agent looked at and what it did not find, rather than having to infer it from silence.

Each agent also works on its own. You do not need the full pipeline to ask `shape` for a pro/contra discussion, or `ideate` to think through an idea. If the artifact from a previous phase is missing, the agent asks you for the missing context directly, or accepts a pointer to something you already have, instead of refusing to proceed.

## The four agents

| Agent | Phase | What it does |
|---|---|---|
| `ideate` | Idea | Explores the problem with you at the domain level, then sketches the main success scenario. No code, no files, no implementation. |
| `shape` | Shape | Proposes candidate technical approaches and drives a pro/contra discussion before one is agreed. |
| `craft` | Build | Enriches the use case step by step through the craft loop. |
| `reflect` | Reflect | Disperses learnings into your project docs, runs a short retrospective, and closes the cycle. |

### The craft loop

The `craft` agent does not follow a plan. It finds the next step through conversation and repeats this loop until the use case goal is met.

```
discuss ──► agree ──► tdd ──► review ──► repeat
```

**Discuss.** The agent looks at the use case goal and the current code, then suggests the next meaningful step. You respond, redirect, or add context.

**Agree.** The agent states the step in one sentence. You confirm it before any code is written.

**TDD.** The agent writes a failing test, shows the failure, writes the minimum code to pass, then refactors if needed.

**Review.** The agent scans the changed files for `// @review: <instruction>` comments and states the result explicitly: each comment found and how it was addressed, or that none were found. You review the result. Approve to commit and repeat.

## How to install

Run this in your project root:

```bash
curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install.sh | bash
```

The script downloads the latest `main` branch, copies `dist/agents/` and `opencode.jsonc` into your project, and adds `.agent/` to `.gitignore`. It skips files that already exist. Pass `--force` to overwrite them.

```bash
curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install.sh | bash -s -- --force
```

## How to start

Tell OpenCode to load the `ideate` agent.

```text
Use the ideate agent.
```

The agent checks the current cycle state and tells you what to do next. If you are starting fresh, it opens the domain conversation.

## Cycle files

The agents write temporary files to `.agent/` during a cycle. All 4dc-owned files are prefixed with `4dc-`. Do not edit them by hand unless you know what you are doing.

| File | Written by | Consumed by |
|---|---|---|
| `.agent/4dc-ideate.md` | `ideate` | `shape` and `craft` — weaved into project docs by `reflect`, then deleted |
| `.agent/4dc-shape.md` | `shape` | `craft` and `reflect` — deleted at cycle close |
| `.agent/4dc-craft.md` | `craft` | `reflect` — weaved into project docs, then deleted |

The `reflect` agent deletes all three files when the cycle closes.

## Using an agent standalone

You can load any agent on its own, for a one-off discussion, without running the rest of the pipeline. For example:

```text
Use the shape agent: help me figure out the best way to onboard the user.
```

When an agent runs as part of the pipeline, it writes straight to its `.agent/4dc-*.md` file as described above, no questions asked. When it finds no pipeline file to write to or read from, it treats this as a standalone conversation instead of refusing to proceed. At the end of the conversation, it asks you where to save the outcome, and whether to save it at all. The default is a file under `./agents/` in your project root, named after the topic you discussed, for example `./agents/onboard-the-user.md`. You can point it to any other path, or tell it not to save anything, in which case it shows you the result in the chat instead.

## Scorecard

`SCORECARD.md` contains 40 yes/no questions grouped by agent and concern. It assesses the agent definitions themselves, not the output of a cycle. Run it after changing an agent file, or whenever an agent's behavior feels off, to check whether the agent files have drifted from the design. Each question targets a specific drift pattern. The scoring at the bottom tells you how to respond.

## Contributing

`AGENTS.md` in the repository root contains working notes for contributors. It describes the repository layout, how to edit and release agent files, how to run the scorecard, and how to use the pipeline to dog-food the agent definitions themselves.

## What 4DC does not do

4DC has no runtime dependency. It does not install Node.js, Python, a test runner, or an MCP server. Your project must already provide whatever testing and tooling it needs.

4DC does not create documentation upfront. Docs are created by the `reflect` agent when a cycle produces something worth recording, and only in locations consistent with your project's existing conventions.

## Influences

4DC borrows from people who shaped modern software development.

| Influence | `ideate` | `shape` | `craft` | `reflect` |
|---|---|---|---|---|
| **Cockburn** | Express the goal as an actor-goal pair; the domain rule must not depend on delivery mechanism | Summary level first; casual level when the build pulls it; name ports before naming adapters | Enrich toward the main success scenario; implement toward the port, keep the framework outside the domain | Check whether every use case reached its stated goal |
| **Kent Beck** | One idea at a time; stop when the concept is clear | The skeleton is a sketch, not a contract | Red-green-refactor within each step; simplest thing that works | Ask what changed about how we test next time |
| **Poppendieck** | Decide late; do not lock in a solution during ideation | Defer detail; keep the skeleton small | Pull from the use case; do not push from a plan | Identify waste in the cycle; propose removing it |
| **Fowler** | Name the bounded context before naming any behavior | Draw the rough boundary; flag coupling risks | Refactor after green; keep the structure honest | Update the architecture description to match what was built |
| **Evans** | Use and test the Ubiquitous Language | Every use case step uses domain language | Implement only what the domain model calls for | Record new or changed domain terms |
| **Crystal** | The conversation is the artifact at this stage | Enough design to start | Stop after each action; human feedback drives the next step | Short retrospective: what ambiguities survived, what quality gaps remain |
| **Dave Thomas** | Name the problem before naming the solution; a tracer bullet starts with a clear target | Design orthogonally; keep concerns separate at the boundary | Tracer bullet first: thin end-to-end before filling in detail; DRY within each step | Fix broken windows found during the cycle before closing it |
