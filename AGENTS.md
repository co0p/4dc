# 4DC — working notes

This file is for contributors working on the agent definitions themselves, not for end users of 4dc.

## Repository layout

```
src/agents/       agent templates — edit here
src/parts/        shared sections and permission blocks
dist/agents/      rendered agents — what the installer ships to users
scripts/
  install.sh      installs dist/agents/ into a user's project
  release.sh      renders src/agents/ with src/parts/ into dist/agents/
```

Edit agent templates in `src/agents/` and shared text in `src/parts/`. A template line such as `<!-- include: language.md -->`, on a line of its own, is replaced during release by the contents of `src/parts/language.md`. Include paths resolve inside `src/parts/` only and must not contain `..`. A missing include or an unresolved marker stops the release.

Release also writes `version: "<YYYY-MM-DD>"` into each agent's frontmatter, using the release date. Do not put a version in the templates.

When a set of edits is ready to publish, run `scripts/release.sh`. That script renders every template into `dist/agents/`. Commit the result. The installer always copies from `dist/agents/`, so users only receive what has been deliberately released.

Do not edit `dist/agents/` directly. Changes made there will be overwritten the next time `release.sh` runs.

Each shared block lives in one file in `src/parts/`. Do not copy it back into the templates. If a block differs between agents, keep it in the template or add a separate part and include it where it belongs.

## How to use the agents

The agents work as a sequence, as a single phase, or as any combination. The human decides how far the work goes. There is no separate standalone mode.

Load whichever agent you need:

```
ideate  ->  shape  ->  craft  ->  reflect
```

Tell OpenCode which agent to load:

```text
Use the ideate agent.
```

The agent lists `.agent/` and reads whatever working docs fit. If upstream context is missing it asks for it directly rather than refusing to proceed. It writes its own doc to the same folder, names the natural next step once, and stops. You decide whether to take that step.

Working docs are named `<slug>-<phase>.md`, for example `.agent/onboard-the-user-shape.md`. The slug stays the same across the phases of one piece of work, so `craft` finds the approach `shape` wrote without being told.
