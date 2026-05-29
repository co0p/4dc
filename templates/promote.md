---
name: 4dc-promote
title: Promote learnings to trunk artifacts
description: Merge durable outcomes into permanent documentation and architecture records
version: "{{VERSION}}"
generatedAt: "{{GENERATED_AT}}"
source: {{SOURCE_URL}}
---

# Prompt: Promote Phase

You are promoting durable outcomes from the 4dc working set into permanent project artifacts.

## Purpose
Convert transient implementation learnings into stable documentation and governance updates before merge.

## Execution Rules
- Review all `.4dc` artifacts before proposing promotions.
- Present each promotion candidate with destination and rationale.
- Require explicit approval before writing permanent docs.

{{TEMPLATE:language}}
{{TEMPLATE:html}}

## Required Inputs
- `CONSTITUTION.md`
- `.4dc/increment.md`
- `.4dc/plan.md`
- `.4dc/implementation.md`
- `.4dc/promote.md`
- Existing permanent docs and ADRs

## Required Output
- HTML report: `.4dc/promotion-report.html`
- Approved updates to permanent artifacts (for example CONSTITUTION, ADRs, design docs, OpenAPI, observability docs)
- Cleanup summary for `.4dc` transient files
- Retrospective delta in `.4dc/promote.md` including:
	- what improved delivery speed
	- what caused defects or friction
	- one concrete process adjustment for the next cycle

Promotion destinations to evaluate:
1. `CONSTITUTION.md`
2. ADR collection
3. OpenAPI specification
4. Visual design guide
5. Personas
6. Deployment strategy
7. Testing decisions
8. Observability documentation
9. C4 diagrams (system, container, component)
10. README and other canonical docs

## Process
1. Extract promotion candidates and classify durable vs transient. Label STOP PR1.
2. Draft exact target updates and traceability mapping. Label STOP PR2.
3. Produce `.4dc/promotion-report.html` for deep review. Label STOP PR3.
4. Apply only approved updates.
5. Suggest that the user empties `.4dc` after each promote and confirm the cleanup decision.
6. Record one next-cycle process adjustment in `.4dc/promote.md`.

## Anti-Patterns
- Promoting without traceability to evidence.
- Dropping unresolved decisions.
- Deleting transient artifacts before approval.

## Done When
- Approved promotions are written.
- `.4dc/promotion-report.html` is available and reviewed.
- A post-promote suggestion to empty `.4dc` was given.
- Retrospective delta is captured in `.4dc/promote.md`.
- Cleanup actions are explicitly confirmed.
