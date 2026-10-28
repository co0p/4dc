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

Each agent hands off to the next. Nothing is built until the use case is clear. Nothing is planned in detail before building starts. Nothing is closed without reflection.

Every agent reports the result of each check it runs, even when the result is "none found". This way you always know what the agent looked at and what it did not find, rather than having to infer it from silence.

Each agent also works on its own. You do not need the full pipeline to ask `shape` for a pro/contra discussion, or `ideate` to turn a need into a use case. There is no separate standalone mode. Every agent reads what already exists in `.agent/`, asks you for anything missing, or accepts a pointer to something you already have, instead of refusing to proceed. You decide when a phase is done and which phase comes next.

## The four agents

| Agent | Phase | What it does |
|---|---|---|
| `ideate` | Use case | Turns a backlog item, story, or need into one sprint-ready use case with actors, a main success scenario, extensions, and acceptance criteria. No code, no implementation. |
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

The agent lists `.agent/` and tells you what to do next. If no doc fits your idea, it opens the domain conversation.

## Working files

Every phase reads and writes its working docs in `.agent/` at the project root. The folder is the only place the agents look for state, and the installer adds it to `.gitignore`.

Each doc is named `<slug>-<phase>.md`. The slug is a few lowercase words taken from the idea, goal, or decision, and it stays the same across the phases of one piece of work. For example:

| File | Written by | Read by |
|---|---|---|
| `.agent/onboard-the-user-ideate.md` | `ideate` | `shape`, `craft`, `reflect` |
| `.agent/onboard-the-user-shape.md` | `shape` | `craft`, `reflect` |
| `.agent/onboard-the-user-craft.md` | `craft` | `reflect` |
| `.agent/onboard-the-user-reflect.md` | `reflect` | you, when no project doc fits the retrospective |

`reflect` deletes the ideate, shape, and craft docs when the work closes, after you confirm the plan. It leaves the reflect doc and every unrelated file alone.

Several efforts can live in the folder at the same time. An agent asks which slug you mean when more than one fits. When a shape doc exists, `craft` reads it instead of asking you to restate the approach.

If you used an earlier version of 4dc, move any `.agent/4dc-*.md` file you want to keep to `.agent/<slug>-<phase>.md`, or remove it. Those old names match the new pattern and would otherwise be read as current state.

## Working one phase at a time

You can load any agent on its own, for a one-off discussion, without running the other phases. For example:

```text
Use the shape agent: help me figure out the best way to onboard the user.
```

There is no separate standalone mode and no separate output location. The agent lists `.agent/`, reads any doc that fits the question, asks you for what is missing, and saves its own doc to the same folder. At the end of the conversation it suggests a slug and a path; you can rename the slug, or decline a file, in which case the agent shows the result in the chat instead.

## Contributing

`AGENTS.md` in the repository root contains working notes for contributors. It describes the repository layout and how to edit and release agent files.

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
