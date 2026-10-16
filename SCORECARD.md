# 4DC Scorecard

Use this scorecard to evaluate the agent definitions in this repository. Run it when an agent feels off, after changing an agent file, or periodically to catch drift in the workflow design itself.

Each question inspects a specific property of the agent definitions. A no means the agent file needs attention.

## Idea agent

| # | Question |
|---|---|
| 1 | Does the agent explicitly refuse to name files, classes, libraries, or functions? |
| 2 | Does the agent require an actor and a goal before it can proceed? |
| 3 | Does the agent require an explicit out-of-scope list before handing off? |
| 4 | Does the agent reference the Ubiquitous Language and flag inconsistent terms? |
| 5 | Does the agent stop when the concept is clear, not when it is complete? |
| 6 | Does the hand-off document (`domain-concept.md`) capture all fields needed by `shape`? |

## Shape agent

| # | Question |
|---|---|
| 7 | Does the agent require a bounded context before writing any use case? |
| 8 | Does the agent write remaining use cases at summary level and expand only one at a time? |
| 9 | Does the agent refuse to name implementation files, classes, or libraries? |
| 10 | Does the agent name coupling risks when another context is involved? |
| 11 | Does the agent mark extensions as deferred rather than detailing them? |
| 12 | Does the hand-off document (`cycle-plan.md`) contain enough for `craft` to start without asking `shape` questions? |

## Craft agent

| # | Question |
|---|---|
| 13 | Does the agent require discussion before writing any code? |
| 14 | Does the agent require an agreed one-sentence step before implementation begins? |
| 15 | Does the agent write a failing test before writing production code? |
| 16 | Does the agent stop after each step and wait for human review? |
| 17 | Does the agent enforce separate `tidy:` and `feat:` commits? |
| 18 | Does the agent respond to `// @review:` comments before moving to the next step? |
| 19 | Does the agent append to `learnings.md` with a decision, reason, and rejected alternative? |

## Reflect agent

| # | Question |
|---|---|
| 20 | Does the agent read the project's existing doc structure before writing anything? |
| 21 | Does the agent write into existing project docs rather than creating 4DC-owned files? |
| 22 | Does the agent update `AGENTS.md` with new system patterns or metaphors? |
| 23 | Does the agent run a retrospective with all three questions? |
| 24 | Does the agent delete all three `.agent/` files at cycle close? |
| 25 | Does the agent refuse to commit `.agent/` files? |

## Language and consistency

| # | Question |
|---|---|
| 26 | Does every agent use `must` for requirements and avoid `should`, `could`, and `might`? |
| 27 | Does every agent use the same hand-off verb (`load`)? |
| 28 | Do file path references match the actual file locations in this repository? |
| 29 | Does every agent have a Language section with consistent rules? |
| 30 | Does every agent have an Influences section that matches the influences table in `README.md`? |

## Scoring

Count your no answers.

| No answers | Action |
|---|---|
| 0 | The agent definitions are clean. |
| 1 to 3 | Small gaps. Edit the specific agent files identified by the no answers. |
| 4 to 7 | One or more agents need a focused review. |
| 8 or more | The agent definitions have significant drift. Re-read the design principles in `README.md` before editing. |
