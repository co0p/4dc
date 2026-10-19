# 4DC — working notes

This file is for contributors working on the agent definitions themselves, not for end users of 4dc.

## Repository layout

```
src/agents/       working copies of the agent files — edit here
dist/agents/      published copies — what the installer ships to users
.opencode/agents/ local copies loaded by OpenCode — kept in sync with dist/ by release.sh
scripts/
  install.sh      installs dist/agents/ into a user's project
  release.sh      promotes src/agents/ into dist/ and .opencode/agents/
```

Edit agent files in `src/agents/`. When a set of edits is ready to publish, run `scripts/release.sh`. That script copies `src/agents/` into both `dist/agents/` and `.opencode/agents/`. Commit the result. The installer always copies from `dist/agents/`, so users only receive what has been deliberately released.

Do not edit `dist/agents/` or `.opencode/agents/` directly. Changes made there will be overwritten the next time `release.sh` runs.

## How to use the agents

The agents work as a pipeline or standalone.

### Pipeline use

Load each agent in order. Each one reads the output of the previous phase from `.agent/` and writes its own output there.

```
ideate  ->  shape  ->  craft  ->  reflect
```

Tell OpenCode which agent to load:

```text
Use the ideate agent.
```

The agent checks the current cycle state and tells you what to do next. When a phase is done, it tells you to load the next agent.

### Standalone use

Any agent can be loaded on its own without running the full pipeline. If the upstream artifact is missing, the agent asks for the missing context directly rather than refusing to proceed.

```text
Use the shape agent: help me think through the best approach for onboarding a new user.
```

At the end of a standalone conversation, the agent asks where to save the result and whether to save it at all. The default is `./agents/<topic-slug>.md` in your project root.

### Dog-fooding: using the agents to improve the agents

The pipeline can be pointed at the agent definitions themselves. This is the recommended way to make deliberate changes to `src/agents/`.

The craft agent's red-green loop applies to prose as well as code. Before an edit, name the behavior you want the agent file to have and why the current text falls short. That gap is your failing test. Make the edit. Re-read the file against the intended behavior. If it now holds, the test passes.

The reflect agent's rule — do not write into files owned by 4dc — must be suspended explicitly when the pipeline is used this way. State this at the start of the reflect session so the agent knows it is operating in meta-cycle mode.

Keep a human in the loop. The pipeline assumes a person who can redirect and approve each step. Running it without a person removes the check that catches agent overconfidence.
