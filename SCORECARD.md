# 4DC Scorecard

This scorecard assesses the consistency of the agent definitions in this repository. It does not assess the output of a cycle. Run it after changing an agent file, when an agent's behavior feels off, or periodically to catch drift in the workflow design itself.

Each question inspects a specific property of the agent definitions. A no means the agent file needs attention.

## Ideate agent

| # | Question |
|---|---|
| 1 | Does the agent explicitly refuse to name files, classes, libraries, or functions? |
| 2 | Does the agent require an actor and a goal before it can proceed? |
| 3 | Does the agent require an explicit out-of-scope list before handing off? |
| 4 | Does the agent reference the Ubiquitous Language and flag inconsistent terms? |
| 5 | Does the agent stop when the concept is clear, not when it is complete? |
| 6 | Does the hand-off document (`.agent/4dc-ideate.md`) capture all fields needed by `shape`, including the main success scenario? |

## Shape agent

| # | Question |
|---|---|
| 7 | Does the agent require a bounded context before proposing an approach? |
| 8 | Does the agent propose multiple candidate approaches and discuss their trade-offs before committing to one? |
| 9 | Does the agent refuse to name implementation files, classes, or libraries? |
| 10 | Does the agent name coupling risks when another context is involved? |
| 11 | Does the agent state the reason each rejected alternative was set aside? |
| 12 | Does the hand-off document (`.agent/4dc-shape.md`) contain enough for `craft` to start without asking `shape` questions? |

## Craft agent

| # | Question |
|---|---|
| 13 | Does the agent require discussion before writing any code? |
| 14 | Does the agent require an agreed one-sentence step before implementation begins? |
| 15 | Does the agent write a failing test before writing production code? |
| 16 | Does the agent stop after each step and wait for human review? |
| 17 | Does the agent enforce separate `tidy:` and `feat:` commits? |
| 18 | Does the agent respond to `// @review:` comments before moving to the next step? |
| 19 | Does the agent append to `.agent/4dc-craft.md` with a decision, reason, and rejected alternative? |

## Reflect agent

| # | Question |
|---|---|
| 20 | Does the agent read the project's existing doc structure before writing anything? |
| 21 | Does the agent write into existing project docs rather than creating 4DC-owned files? |
| 22 | Does the agent weave new system patterns or metaphors into the appropriate project doc? |
| 23 | Does the agent run a retrospective with all three questions (ambiguities, quality gaps, better output)? |
| 24 | Does `shape` leave `4dc-ideate.md` in place for `reflect` to weave and delete? |
| 25 | Does `reflect` delete `4dc-ideate.md`, `4dc-shape.md`, and `4dc-craft.md` at cycle close? |

## Standalone use

| # | Question |
|---|---|
| 32 | Does every agent accept missing upstream pipeline files by asking for the missing context or a pointer, instead of refusing to proceed? |
| 33 | Does every agent have an Output section that distinguishes pipeline use (write straight to `.agent/4dc-*.md`, no question asked) from standalone use (ask where and whether to save)? |
| 34 | Does the standalone default path follow the same convention (`./agents/<topic-slug>.md`) across every agent? |
| 35 | Does every agent's Status updates section report whether it ran as part of the pipeline or standalone, and where the result was saved? |

## Language and consistency

| # | Question |
|---|---|
| 26 | Does every agent use `must` for requirements and avoid `should`, `could`, and `might`? |
| 27 | Does every agent use the same hand-off verb (`load`)? |
| 28 | Do file path references match the actual file locations in this repository? |
| 29 | Does every agent have a Language section with consistent rules? |
| 30 | Does every agent have an Influences section that matches the influences table in `README.md`? |
| 31 | Does every agent have a Status updates section that lists checks to report explicitly, including negative results? |

## Implicit method borrowing

The agents borrow from established thinkers and methods but do not name those methods. This keeps interactions conversational and avoids turning the agent into a method coach. The checks below verify that the discipline is maintained.

| # | Question |
|---|---|
| 36 | Does every Influences entry state a behavioral principle rather than a method name, framework name, or step label (for example: no `red-green-refactor`, `double diamond`, `Jobs to Be Done`, `ports and adapters`, `tracer bullet` as a label)? |
| 37 | Does the agent body (outside the Influences section) avoid naming the methods or frameworks those thinkers are associated with? |
| 38 | Does each Influences entry describe what the agent must do, not who the person is or what they published? |
| 39 | When a new thinker is added to any Influences section, is the same thinker absent from the body text of the agent? |
| 40 | Are the behavioral principles in the Influences section consistent with the procedural instructions in the same agent? A principle that contradicts the steps it is meant to inform is a drift signal. |

## Repository layout

| # | Question |
|---|---|
| 41 | Are all working agent files in `src/agents/` and not edited directly in `dist/agents/` or `.opencode/agents/`? |
| 42 | Does `dist/agents/` contain the last output of `scripts/release.sh` and nothing else? |
| 43 | Are `dist/agents/` and `.opencode/agents/` identical? |
| 44 | Does `scripts/install.sh` copy from `dist/agents/` and not from `.opencode/agents/` or `src/agents/`? |

## Scoring

Count your no answers.

| No answers | Action |
|---|---|
| 0 | The agent definitions are clean. |
| 1 to 3 | Small gaps. Edit the specific agent files identified by the no answers. |
| 4 to 8 | One or more agents need a focused review. |
| 9 or more | The agent definitions have significant drift. Re-read the design principles in `README.md` before editing. |
