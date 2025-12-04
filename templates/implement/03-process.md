## Process

Follow this process to produce an `implement.md` that is:

- Directly traceable to `increment.md` and `design.md`.
- Grounded in the actual codebase.
- Broken into small, safe, verifiable steps.
- Easy for humans to review and adapt.

The `path` argument for this prompt points at an **increment folder** (for example: `.../increments/<slug>`). The increment folder contains `increment.md` and, usually, `design.md`. The **project codebase** and other documentation live above this folder under the project root.

### Phase 1 – Gather and Summarize (STOP 1)

1. Gather Context

   - Read and internalize:
     - `increment.md` in this folder — context, goal, tasks (WHAT), risks, success criteria.
     - `design.md` in this folder — technical HOW, architecture, components, contracts, tests, rollout, observability.
   - Under the project root (containing this increment folder), optionally review:
     - `CONSTITUTION.md` — values, principles, guardrails, delivery expectations.
     - Relevant ADRs or prior designs.
     - Recent `improve.md` documents that mention this part of the system.
   - Inspect relevant **code and tests** under the project path:
     - Focus on components, modules, services, and data paths that:
       - The design says will change, or
       - Are clearly adjacent to those components.
     - Note:
       - Entry points (APIs, commands, UI routes).
       - Data models and persistence layers involved.
       - Existing tests and how they are organized.

2. Restate Problem, Scope, and Design (Briefly)

   - In a few sentences, restate:
     - The problem and outcome from `increment.md`.
     - The main technical approach from `design.md`.
     - The primary areas of the codebase that will be touched.

3. Summarize Findings and Constraints → STOP 1

   - Present a concise summary that covers:
     - Your understanding of:
       - The increment’s goal, tasks, and non-goals.
       - The design’s main ideas and boundaries.
     - Which parts of the system appear most relevant to implementation.
     - Any constraints or risks visible from:
       - The constitution.
       - Existing tests and code structure.
   - Clearly label this as **STOP 1**.
   - Ask the user to:
     - Confirm whether this summary is broadly correct.
     - Provide corrections or highlight any constraints you may have missed.

   Do not proceed to proposing a detailed step-by-step plan until the user has responded to STOP 1.

4. Ask Targeted Clarifying Questions (If Needed)

   - After presenting the findings, ask **brief, targeted questions** only if:
     - There are unclear priorities between tasks.
     - There are ambiguous sequencing or dependency issues (for example: data migrations vs. feature flags).
     - There are notable uncertainties about environments (staging, production, CI).
   - Avoid long questionnaires; keep questions minimal and specific.
   - Incorporate the user’s answers into your internal understanding before proceeding.

### Phase 2 – Propose Implementation Outline (STOP 2)

5. Identify Workstreams and Dependencies

   - Group work into **workstreams** (for example: “backend API changes”, “frontend UI changes”, “data migration”, “observability updates”).
   - For each workstream, identify:
     - The main components / areas of code affected.
     - Any obvious dependencies on other workstreams.
   - Ensure that:
     - Workstreams respect the increment’s scope and non-goals.
     - Workstreams reflect the design’s architecture and boundaries.

6. Propose Step-Level Outline

   - For each workstream, propose a **high-level sequence of steps**, such as:
     - Introduce new code paths behind flags.
     - Add tests for new behavior.
     - Migrate data or configuration.
     - Remove deprecated paths once traffic has moved.
   - Keep steps:
     - Small enough to be mapped to one or a few pull requests.
     - Ordered to reduce risk (for example: additive changes before destructive ones).

7. Integrate Testing, CI, and Rollout into the Outline

   - For relevant steps, note:
     - When to add or update tests (unit, integration, end-to-end).
     - When to run specific test suites or commands.
     - How the change will flow through CI.
   - Integrate rollout considerations:
     - Use of feature flags or configuration switches.
     - Staged or canary rollouts, if appropriate.
     - Points where it is safe to pause, validate, or roll back.

8. Integrate Observability and Post-Release Checks

   - For relevant steps, note:
     - Where logging and metrics changes will be implemented.
     - What should be verified in staging or pre-production.
     - What should be checked in production after release (for example: metrics, dashboards, logs).

9. Summarize Proposed Implementation Outline → STOP 2

   - Before writing the full `implement.md`, present a **section-by-section outline** covering:
     - Workstreams and their purpose.
     - Key steps in each workstream, in rough order.
     - Where tests, CI, rollout, and observability fit into the sequence.
     - Major risks or decision points.
   - Map this outline clearly onto the implementation output structure (sections for overview, steps, risks, validation, etc.).
   - Clearly label this as **STOP 2**.
   - Ask the user explicitly to:
     - Answer yes/no (or equivalent) to confirm the outline.
     - Suggest adjustments (add/remove/merge/split/reorder steps) if needed.

   Do not generate the full `implement.md` until the user has approved this outline.

### Phase 3 – Write the Implementation Plan After YES

10. Produce the Final `implement.md` (After STOP 2 Approval)

    - Only after the user gives a clear affirmative response at STOP 2 (for example: “yes”, “go ahead”, “looks good”):
      - Generate `implement.md` that:
        - Follows the structure defined in the implementation output structure template.
        - Implements the agreed outline, including any adjustments from user feedback.
    - While writing:
      - Do not introduce new, major decisions that were not in the approved outline, unless clearly flagged as newly discovered risks or options.
      - Do not mention prompts, LLMs, or this process.
      - Keep the document clear, concise, and directly traceable to:
        - `increment.md`.
        - `design.md`.
        - The current code and architecture.

If the user does not approve the outline at STOP 2:

- Update the outline based on their feedback.
- Re-present it and wait for approval before generating the final implementation plan.