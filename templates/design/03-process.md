## Process

Follow this process to produce a `design.md` that is aligned with the constitution and the current increment, and that keeps a human in the loop.

### Phase 1 – Gather & Summarize (STOP 1)

1. **Gather Context**

   - Read and internalize:
     - `CONSTITUTION.md` — values, principles, guardrails, delivery expectations.
     - The current `increment.md` — context, goal, tasks (WHAT), risks, success criteria.
   - Optionally review:
     - Relevant ADRs.
     - Existing `design.md` documents for related areas.
     - Recent `improve.md` documents that mention this part of the system.

2. **Restate Problem and Scope (Briefly)**

   - In a few sentences, restate:
     - What problem this design is solving.
     - The outcome the increment targets.
     - The scope and non-goals from `increment.md`.

3. **Summarize Findings → STOP 1**

   - Present a concise summary that covers:
     - Your understanding of the problem and scope.
     - Which parts of the system are likely involved.
     - Any key constraints or assumptions visible from `CONSTITUTION.md`, `increment.md`, and existing docs.
   - Clearly label this as **STOP 1**.
   - Ask the user to:
     - Confirm whether this summary is broadly correct.
     - Provide corrections or add missing, critical context.

   Do **not** proceed to proposing a full design until the user has responded to STOP 1.

4. **Ask Targeted Clarifying Questions (If Needed)**

   - After presenting the findings, ask **brief, targeted questions** only if:
     - Critical information is missing or ambiguous (e.g. performance constraints, data sensitivity).
     - There is a conflict between `CONSTITUTION.md` and `increment.md` that must be resolved.
   - Avoid long questionnaires; keep questions minimal and specific.
   - Incorporate the user’s answers into your internal understanding before proceeding.

### Phase 2 – Propose Design & Outline (STOP 2)

5. **Identify Involved Components and Boundaries**

   - Determine which:
     - Modules, packages, services, or layers are impacted.
     - External systems (datastores, queues, APIs) are involved.
   - Note any existing boundaries that must be respected (from the constitution).

6. **Propose a Technical Approach**

   - Describe:
     - How responsibilities will be distributed across components.
     - Any new components or changes to existing ones.
     - The main data flows (inputs, outputs, key transformations).
   - Keep the approach:
     - As simple as possible.
     - Constrained to the increment’s scope.

7. **Define Contracts & Interfaces**

   - Specify:
     - New or changed APIs, function signatures, events, or schemas.
   - Clarify:
     - What remains stable.
     - How backward compatibility will be preserved where necessary.

8. **Plan the Safety Net (Testing)**

   - Enumerate:
     - Which **unit tests** are needed (per component).
     - Which **integration / end-to-end tests** are needed (per flow or contract).
   - Include:
     - Any regression tests required for known bugs.
     - Any special test data/fixtures.

9. **Consider CI/CD and Rollout**

   - Note:
     - Whether existing pipelines are sufficient or need updates.
     - Any required configuration or environment changes.
   - Describe:
     - How this change can be rolled out safely:
       - Feature flags?
       - Gradual rollout?
       - Internal dogfooding first?
     - How it can be rolled back:
       - Reverting code.
       - Toggling configuration.

10. **Specify Observability**

    - Define:
      - Logs needed (what to log and with what context).
      - Metrics (counters, histograms, gauges) that reflect:
        - Usage.
        - Performance.
        - Errors and unusual conditions.
    - Mention:
      - Any alerts or dashboards that should be created or updated.

11. **Summarize Proposed Design Outline → STOP 2**

    - Before writing the full `design.md`, present a **section-by-section outline** summarizing:
      - The high-level solution and which components are involved.
      - Key contracts/data changes.
      - Testing strategy.
      - CI/CD and rollout considerations.
      - Observability and operations aspects.
      - Major risks, trade-offs, and follow-up ideas.
    - Map this outline clearly onto the sections defined in `05-output-and-examples.md`.
    - Clearly label this as **STOP 2**.
    - Ask the user explicitly to:
      - Answer yes/no (or equivalent) to confirm the outline.
      - Suggest adjustments (add/remove/strengthen/weaken points) if needed.

    Do **not** generate the full `design.md` until the user has approved this outline.

### Phase 3 – Write the Design After YES

12. **Produce the Final `design.md` (After STOP 2 Approval)**

    - Only after the user gives a clear affirmative response at STOP 2 (e.g. “yes”, “go ahead”, “looks good”):
      - Generate `design.md` that:
        - Follows the structure defined in `05-output-and-examples.md`.
        - Implements the agreed outline, including any adjustments from user feedback.
    - While writing:
      - Do **not** introduce new, major decisions that were not in the approved outline.
      - Do not mention prompts, LLMs, or this process.
      - Keep the document clear, concise, and directly traceable to `CONSTITUTION.md` and `increment.md`.

If the user does **not** approve the outline at STOP 2:

- Update the outline based on their feedback.
- Re-present it and wait for approval before generating the final design.