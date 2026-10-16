# 4DC Agents

This file is the orchestrator for the Four Discipline Cycle. It also holds the System Metaphor for the project, which grows over time as the `reflect` agent closes cycles.

## How to use

Read the state table below. Load the agent that matches your current state.

| State | Signal | Load |
|---|---|---|
| Starting a new idea | no files in `.agent/` | `idea` |
| Idea is captured | `.agent/domain-concept.md` exists | `shape` |
| Skeleton is ready | `.agent/cycle-plan.md` exists | `craft` |
| Build is accepted | branch merged or user signals done | `reflect` |
| Cycle is closed | `.agent/` is empty | `idea` |

To load an agent, tell OpenCode: `Use the idea agent` (or `shape`, `craft`, `reflect`).

## Cycle files

These files live in `.agent/` and are never committed.

| File | Written by | Consumed by |
|---|---|---|
| `domain-concept.md` | `idea` | `shape` |
| `cycle-plan.md` | `shape` | `craft` |
| `learnings.md` | `craft` | `reflect` |

## System Metaphor

This section is maintained by the `reflect` agent. It describes how this project thinks about itself. It starts empty and grows with each completed cycle.

<!-- reflect agent writes here -->
