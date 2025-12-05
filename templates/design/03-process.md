## Process

Follow this process to produce a `design.md` that is aligned with the constitution and the current increment, grounded in the existing codebase, and that keeps a human in the loop.

The `path` argument for this prompt points at an **increment folder** (for example: `.../increments/<slug>`). The increment folder contains `increment.md`. The **project codebase** and other documents (such as `CONSTITUTION.md`, ADRs, and prior designs) live under the project root.

### Phase 1 – Gather and Summarize (STOP 1)

1. Gather Context

   - Read and internalize:
     - `CONSTITUTION.md` — values, principles, guardrails, delivery expectations.
     - The current `increment.md` in this folder — context, goal, tasks (WHAT), risks, success criteria.
   - Optionally review, under the project root:
     - Relevant ADRs.
     - Existing `design.md` documents for related areas.
     - Recent `improve.md` documents that mention this part of the system.
   - Inspect relevant **code and tests** under the project path:
     - Focus on components, modules, services, and data flows that:
       - Are directly involved in fulfilling the increment’s goal and tasks, or
       - Are upstream or downstream dependencies of those parts.
     - Note obvious constraints and patterns:
       - Frameworks, data models, layering rules, ownership boundaries.
       - Existing conventions that the design should follow or intentionally refactor.

2. Restate Problem and Scope (Briefly)

   - In a few sentences, restate:
     - What problem this design is solving.
     - The outcome the increment targets.
     - The scope and non-goals from `increment.md`.
     - Which parts of the existing system appear to be involved.

3. Summarize Findings → STOP 1

   - Present a concise summary that covers:
     - Your understanding of the problem and scope from the increment.
     - Which parts of the system (components, modules, services, data stores) are likely involved.
     - Any key constraints or assumptions visible from `CONSTITUTION.md`, `increment.md`, existing docs, and the current code.
   - Clearly label this as **STOP 1**.
   - Ask the user to:
     - Confirm whether this summary is broadly correct.
     - Provide corrections or add missing, critical context.

   Do not proceed to proposing a full design until the user has responded to STOP 1.

4. Ask Targeted Clarifying Questions (If Needed)

   - After presenting the findings, ask **brief, targeted questions** only if:
     - Critical information is missing or ambiguous (for example: performance constraints, data sensitivity, external dependencies).
     - There is a conflict between `CONSTITUTION.md` and `increment.md` that must be resolved.
   - Avoid long questionnaires; keep questions minimal and specific.
   - Incorporate the user’s answers into your internal understanding before proceeding.

### Phase 2 – Propose Design and Outline (STOP 2)

5. Identify Involved Components and Boundaries

   - Determine which:
     - Modules, packages, services, or layers are impacted.
     - External systems (datastores, queues, APIs) are involved.
   - Note any existing boundaries that must be respected (from the constitution and current architecture).
   - Call out any component-level refactors that seem necessary to support the increment.

6. Propose a Technical Approach

   - Describe:
     - How responsibilities will be distributed across components.
     - Any new components or changes to existing ones.
     - The main data flows (inputs, outputs, key transformations).
   - Keep the approach:
     - As simple as possible.
     - Constrained to the increment’s scope and non-goals.
     - Implementable in **small, safe steps**.
   - **Do not** turn this into a chronological list of code edits or tasks.  
     Focus on *structure and behavior*, not on “Step 1/Step 2” sequences.

7. Define Contracts and Interfaces

   - Specify:
     - New or changed APIs, function or method signatures, events, or schemas.
   - Clarify:
     - What remains stable.
     - How backward compatibility will be preserved where necessary.
   - Highlight any versioning, migration, or compatibility strategies if needed.

8. Plan the Safety Net (Testing)

   - Enumerate:
     - Which **unit tests** are needed (per component or function).
     - Which **integration / end-to-end tests** are needed (per flow or contract).
   - Include:
     - Any regression tests required for known bugs.
     - Any special test data/fixtures or environments.
   - Note any risks around test flakiness and how to mitigate them.

9. Consider CI/CD and Rollout

   - Note:
     - Whether existing pipelines are sufficient or need updates.
     - Any required configuration or environment changes.
   - Describe:
     - How this change can be rolled out safely:
       - Feature flags.
       - Gradual rollout.
       - Internal dogfooding first.
     - How it can be rolled back or disabled:
       - Reverting code.
       - Toggling configuration or flags.
       - Falling back to previous behavior.

10. Specify Observability

    - Define:
      - Logs needed (what to log and with what context, such as IDs, correlation tokens).
      - Metrics that reflect:
        - Usage.
        - Performance.
        - Errors and unusual conditions.
    - Mention:
      - Any alerts or dashboards that should be created or updated.

11. Summarize Proposed Design Outline → STOP 2

    - Before writing the full `design.md`, present a **section-by-section outline** summarizing:
      - The high-level solution and which components are involved.
      - Key contracts and data changes.
      - Testing strategy.
      - CI/CD and rollout considerations.
      - Observability and operations aspects.
      - Major risks, trade-offs, and follow-up ideas.
    - Map this outline clearly onto the sections defined in the design output structure.
    - Clearly label this as **STOP 2**.
    - Ask the user explicitly to:
      - Answer yes/no (or equivalent) to confirm the outline.
      - Suggest adjustments (add/remove/strengthen/weaken points) if needed.

    Do not generate the full `design.md` until the user has approved this outline.

### Phase 3 – Write the Design After YES

12. Produce the Final `design.md` (After STOP 2 Approval)

    - Only after the user gives a clear affirmative response at STOP 2 (for example: “yes”, “go ahead”, “looks good”):
      - Generate `design.md` that:
        - Follows the structure defined in the design output structure template.
        - Implements the agreed outline, including any adjustments from user feedback.
    - While writing:
      - Do not introduce new, major decisions that were not in the approved outline.
      - Do not introduce step-by-step implementation instructions or task lists.
      - Do not mention prompts, LLMs, or this process.
      - Keep the document clear, concise, and directly traceable to:
        - `CONSTITUTION.md`.
        - `increment.md`.
        - The current code and architecture.

If the user does not approve the outline at STOP 2:

- Update the outline based on their feedback.
- Re-present it and wait for approval before generating the final design.