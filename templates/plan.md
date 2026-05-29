---
name: 4dc-plan
title: Plan technical execution
description: Convert increment intent into an implementation strategy with concrete subtasks
version: "{{VERSION}}"
generatedAt: "{{GENERATED_AT}}"
source: {{SOURCE_URL}}
---

# Prompt: Plan Phase

You are producing the technical plan for the approved increment.

## Purpose
Define HOW to deliver `.4dc/increment.md` safely using explicit subtasks, sequencing, and verification points.

## Execution Rules
- Switch to plan mode immediately for this phase.
- Keep every task actionable, ordered, and verifiable.
- Do not start implementation in this phase.
- Use STOP gates and explicit approval before writing final Markdown.

{{TEMPLATE:language}}
{{TEMPLATE:html}}

## Required Inputs
- `CONSTITUTION.md`
- `.4dc/increment.md`
- Existing code, tests, and architecture docs

## Required Output
- HTML review file: `.4dc/plan-review.html`
- Final file after approval: `.4dc/plan.md`

`.4dc/plan.md` must include:
1. Technical approach summary
2. Constraints and dependencies
3. Task breakdown (ordered subtasks)
4. Verification strategy per subtask
5. Risks and mitigations
6. Traceability from requirements to subtasks
7. Test-first mapping:
	- first failing test for each subtask
	- acceptance-criteria-to-test mapping
8. Spike plan for uncertainty hotspots (if any):
	- time-box
	- question to answer
	- expected output and next safe step

## Process
1. Analyze impacted areas and constraints. Label STOP P1.
2. Draft full technical plan and subtask map. Label STOP P2.
3. Add first failing test for each subtask before implementation approval.
4. For unknowns, define a time-boxed spike and capture explicit exit criteria.
5. Produce `.4dc/plan-review.html` and request approval. Label STOP P3.
6. Write `.4dc/plan.md` only after explicit approval.

## Anti-Patterns
- Starting coding in the plan phase.
- Tasks without verification.
- Missing dependency or risk coverage.
- Subtasks without first failing tests.
- Unknowns without spike definition.

## Done When
- `.4dc/plan.md` is approved and written.
- Every functional requirement maps to at least one plan task.
- Review HTML was shown before final write.
