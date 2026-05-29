---
name: 4dc-increment
argument-hint: short increment intent, for example "feature: export weekly summary"
title: Define next increment
description: Define a small, outcome-focused increment without technical implementation detail
version: "{{VERSION}}"
generatedAt: "{{GENERATED_AT}}"
source: {{SOURCE_URL}}
---

# Prompt: Increment Phase

You are defining the next delivery increment as a narrow, testable changeset.

## Purpose
Produce a clear WHAT and WHY for one increment, with measurable acceptance and explicit boundaries.

## Execution Rules
- Stay out of technical design and coding details.
- Keep scope to one increment that can be completed in one focused cycle.
- Use STOP gates and require explicit approval before final Markdown write.

{{TEMPLATE:language}}
{{TEMPLATE:html}}

## Required Inputs
- `CONSTITUTION.md`
- User intent and business context
- Current product behavior

## Required Output
- HTML review file: `.4dc/increment-review.html`
- Final file after approval: `.4dc/increment.md`

`.4dc/increment.md` must include:
1. Problem statement
2. Target outcome
3. In-scope and out-of-scope
4. Functional requirements (no technical details)
5. User stories
6. Acceptance criteria
7. Risks and assumptions
8. Completion signals
9. Planning game fields:
	- business value (1-5)
	- risk level (1-5)
	- slice size verdict (`small`, `too-large`)

## Process
1. Clarify goal, user impact, and boundaries. Label STOP I1.
2. Draft increment document with measurable acceptance criteria. Label STOP I2.
3. Run planning game sizing check.
	- If slice size verdict is `too-large`, split scope before continuing.
	- Keep only one shippable slice in this increment.
4. Produce `.4dc/increment-review.html` and request approval. Label STOP I3.
5. Write `.4dc/increment.md` after explicit approval.

## Anti-Patterns
- Hidden technical design choices.
- Multi-epic scope.
- Acceptance criteria that cannot be observed.
- Skipping value/risk/slice scoring before approval.

## Done When
- `.4dc/increment.md` is approved and written.
- Review HTML was shown before final write.
