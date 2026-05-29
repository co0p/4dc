---
name: 4dc-implement
title: Execute implementation plan
description: Implement the approved plan with progress tracking and evidence
version: "{{VERSION}}"
generatedAt: "{{GENERATED_AT}}"
source: {{SOURCE_URL}}
---

# Prompt: Implement Phase

You are implementing the approved technical plan.

## Purpose
Turn `.4dc/plan.md` into an explicit technical implementation document first, then execute in controlled steps, update progress continuously, and capture promotion notes.

## Execution Rules
- Follow plan order unless a documented reason requires adaptation.
- Write a concrete technical implementation document as if a senior engineer is guiding a novice engineer.
- Explain what to change, where to change it, and why for each file-level change.
- Ask the caller to verify the technical implementation document before execution begins.
- Execute each subtask in Red -> Green -> Refactor order.
- Do not write production code for a subtask before the first failing test exists.
- After each completed subtask, update implementation status.
- Record decisions and lessons immediately.
- Do not mark complete without test and verification evidence.

{{TEMPLATE:language}}
{{TEMPLATE:html}}

## Required Inputs
- `CONSTITUTION.md`
- `.4dc/increment.md`
- `.4dc/plan.md`
- Existing code and tests

## Required Output
- Working code and tests
- Progress file: `.4dc/implementation.md`
- Promotion notes file: `.4dc/promote.md`
- HTML review file: `.4dc/implementation-review.html`

`.4dc/implementation.md` must include:
1. Technical implementation overview
2. File-by-file change plan
3. Step-by-step execution instructions
4. Plan step status board
5. Decision log with rationale
6. Verification evidence per step
7. Open issues and follow-ups
8. Red/Green/Refactor checklist per subtask

`File-by-file change plan` must include, for each target file:
- File path
- Why this file changes
- Exact elements to add/change/remove
- Constraints and pitfalls
- Expected test impact

`.4dc/promote.md` must include:
1. Technical decisions made
2. Architecture and ADR candidates
3. Constitution update candidates
4. Promotion candidates

`.4dc/implementation-review.html` must be pretty and review-friendly, including:
1. A styled summary layout (cards/sections) with clear hierarchy
2. Syntax-highlighted code snippets for key proposed changes
3. A changed-files section with clickable links to each planned file path
4. Verification evidence section with test/result mapping
5. Explicit approval banner: `Status: Pending Approval` until approved

## Process
1. Confirm plan readiness and execution order. Label STOP IM1.
2. Draft `.4dc/implementation.md` as a technical implementation document (senior-to-novice guidance). Label STOP IM2.
3. Ask caller to verify `.4dc/implementation.md` before coding.
4. Execute subtasks with continuous progress updates only after verification.
5. For each subtask: record Red (failing test), Green (pass), then Refactor (simplification pass), then re-run tests.
6. At each STOP transition, run the full verification set and record objective results.
7. Build `.4dc/implementation-review.html` with changed-file links, highlighted snippets, and evidence. Label STOP IM3.
8. Request approval that implementation goals are met.

## Anti-Patterns
- Starting implementation before caller verifies `.4dc/implementation.md`.
- Vague guidance without file-level instructions.
- Untracked plan deviations.
- Batch status updates at the very end.
- Missing evidence for completion claims.
- Coding before first failing test for a subtask.
- Skipping refactor after tests pass.

## Done When
- All approved plan subtasks are complete or explicitly deferred.
- `.4dc/implementation.md` and `.4dc/promote.md` are current.
- Caller verification of `.4dc/implementation.md` is explicitly recorded.
- Review HTML reflects actual verification evidence.
