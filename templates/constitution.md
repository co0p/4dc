---
name: 4dc-constitution
title: Create or adjust project constitution
description: Capture project guardrails and SDLC standards through focused discovery
version: "{{VERSION}}"
generatedAt: "{{GENERATED_AT}}"
source: {{SOURCE_URL}}
---

# Prompt: Constitution Phase

You are defining or updating `CONSTITUTION.md` for this project.

## Purpose
Create practical guardrails that guide day-to-day engineering decisions and documentation quality.

## Execution Rules
- Read existing repository docs and structure before asking questions.
- Ask 3-5 high-value questions per round.
- Use STOP gates. Do not write final Markdown without explicit approval.
- Keep decisions concrete, testable, and project-specific.

{{TEMPLATE:language}}
{{TEMPLATE:html}}

## Required Inputs
- Existing `CONSTITUTION.md` (if present)
- `README.md`
- Current project structure and delivery constraints

## Required Output
- HTML review file: `.4dc/constitution-review.html`
- Final file after approval: `CONSTITUTION.md`

`CONSTITUTION.md` must include:
1. Engineering principles grounded in XP, lean software development, and use-case thinking.
2. Architectural boundaries and dependency direction.
3. Testing strategy and quality gates.
4. Documentation rules and ADR policy.
5. SDLC artifact expectations:
   - OpenAPI specification
   - ADRs
   - Visual design guide
   - Personas
   - Deployment strategy
   - Testing decisions
   - Observability
   - C4 diagrams (system, container, component)

## Process
1. Context discovery and baseline summary. Label STOP C1.
2. Draft constitution outline and policy table. Label STOP C2.
3. Produce `.4dc/constitution-review.html` and request approval. Label STOP C3.
4. Write `CONSTITUTION.md` only after explicit approval.

## Anti-Patterns
- Abstract slogans with no enforcement rule.
- Architecture policies that cannot be verified.
- Skipping SDLC artifact expectations.

## Done When
- `CONSTITUTION.md` exists and reflects approved content.
- Review HTML was shown before final write.
- No unresolved contradictions remain.
