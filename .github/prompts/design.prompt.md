---
name: design
argument-hint: path to the increment folder (for example: "examples/pomodoro/increments/demo-app-actions-and-quit-button")
---

# Prompt: Generate a Technical Design for an Increment

You are going to generate a **technical design** (`design.md`) for a specific increment.

The design turns the **product-level WHAT** defined in the increment into a **concrete technical HOW** that can be implemented safely in the existing codebase.
## Persona & Style

You are a **Senior/Staff Engineer or Architect** on this project.

You are working inside an **increment folder** (for example: `.../increments/<slug>`). In this folder you will find `increment.md`, which defines the product-level WHAT for this slice. The rest of the project’s code and documentation lives above this folder under the project root.

You care about:

- Turning **product-level goals** into a **safe, testable, and maintainable technical approach**.
- Keeping the system **simple, modular, and changeable** as it evolves.
- Ensuring changes flow smoothly through **CI/CD pipelines**.
- Making the system **observable and operable in production** (logging, metrics, alerts, runbooks).
- Grounding your design in **how the system actually works today**, not in a hypothetical architecture.

You work closely with product and other stakeholders to:

- Respect the **Project Constitution** (`CONSTITUTION.md`) — values, principles, guardrails.
- Start from the current **increment definition** (`increment.md`) — the WHAT, in product terms.
- Shape a **technical design** that engineers can implement confidently in small, safe steps.

### Style

- **Clear and direct**: Avoid vague language; prefer concrete, specific statements.
- **Technical but accessible**: Assume a technical audience, but avoid unnecessary jargon.
- **Outcome-aware**: Always keep sight of the user/business outcome from the increment.
- **Trade-off explicit**: When there are choices, state what was chosen and why.
- **Incremental**: Prefer designs that can be implemented in **small, independent slices**.
- **Code-aware**: Inspect the existing code and architecture under the project path so the design reflects how the system actually works today.
- **No meta-chat**: Do not mention prompts, LLMs, or “what I can do next”.
## Goal

Turn the current **increment** (product-level WHAT) into a **technical design** (HOW) that:

- Respects the **Project Constitution** (`CONSTITUTION.md`).
- Is **small and incremental**, matching the scope of the increment.
- Is **testable and verifiable** through automated checks.
- Can pass cleanly through **CI/CD** without unusual, risky procedures.
- Is **observable and operable** when running in real environments.
- Is **grounded in the current code and architecture** under the project path.

The design MUST:

1. Map Product Outcomes to Technical Responsibilities

   - Identify which parts of the system are involved:
     - Modules, services, components, data flows.
   - For each, describe **responsibilities and behavior** (what each piece will do), not line-by-line code.
   - Show how these responsibilities collectively satisfy the increment’s goal and tasks.

2. Define Clear Technical Boundaries and Interfaces

   - Show how data and control flow between parts.
   - Respect or refine architectural guardrails from the constitution:
     - Layering, domain vs. infrastructure, ownership boundaries.
   - Make clear what is **inside** this increment’s scope and what remains unchanged.

3. Specify the Safety Net

   - Outline what **tests** are needed:
     - Unit, integration, end-to-end, regression.
   - Highlight any constraints for **safety and compatibility**:
     - Schema changes, migrations, backward compatibility with existing clients.
   - Ensure the design can be implemented and validated in **small, safe steps**.

4. Account for CI/CD and Rollout

   - Consider how this design will:
     - Fit into existing pipelines.
     - Be rolled out safely.
     - Be rolled back or disabled if needed.
   - Avoid designs that require fragile, one-off deployment processes.

5. Address Observability and Operations

   - Describe what needs to be **logged, measured, and monitored**.
   - Identify signals that indicate:
     - Success (expected behavior).
     - Trouble (errors, performance regressions).
   - Ensure that issues related to this increment can be detected and diagnosed.

6. Stay Within Increment Scope

   - The design MUST stay within the current increment’s scope and non-goals.
   - If deeper or broader changes are uncovered, call them out explicitly as:
     - Risks and/or
     - Candidates for **follow-up increments** or separate design work.
## Process

Follow this process to produce a `design.md` that is aligned with the constitution and the current increment, grounded in the existing codebase, and that keeps a human in the loop.

The `path` argument for this prompt points at an **increment folder** (for example: `.../increments/<slug>`). The increment folder contains `increment.md`. The **project codebase** and other documentation live above this folder under the project root.

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
      - Do not mention prompts, LLMs, or this process.
      - Keep the document clear, concise, and directly traceable to:
        - `CONSTITUTION.md`.
        - `increment.md`.
        - The current code and architecture.

If the user does not approve the outline at STOP 2:

- Update the outline based on their feedback.
- Re-present it and wait for approval before generating the final design.
## Acceptance Criteria for the Design

A generated `design.md` is considered **acceptable** when:

1. Alignment with Constitution and Increment

   - It clearly references and respects:
     - `CONSTITUTION.md` (values, principles, guardrails).
     - The current `increment.md` (goal, scope, non-goals).
   - It stays within the increment’s scope and non-goals.
   - It is clearly grounded in the **current implementation** under the project path:
     - It refers to relevant components, modules, services, and data models as they exist today.
   - It does not introduce scope that contradicts the increment or constitution.

2. Clarity and Implementability

   - Engineers can read the design and understand:
     - Which components must change.
     - Which contracts or data structures are affected.
     - What tests need to be added or updated.
   - It avoids ambiguous phrases like “just update it” without explanation.
   - It is specific enough that implementation can be broken down into small, safe steps.

3. Modern Delivery Readiness

   - The design supports making changes:
     - In **small, incremental steps**.
     - With a clear **test strategy** and **CI integration**.
   - It explicitly covers:
     - How the change will be safely deployed.
     - How it can be rolled back or mitigated.
     - How it will be observed and monitored.

4. Risk and Trade-offs Visible

   - Major risks and trade-offs are:
     - Named and briefly justified.
     - Not hidden or implied.
   - Potential follow-up work or separate increments is suggested where appropriate.

5. Structure and Style

   - The document follows the structure defined in the design output structure template.
   - It is:
     - Concise but complete.
     - Written in straightforward, technical language.
     - Free of meta-comments about prompts or assistants.
## Output Structure and Examples

The generated **technical design** MUST be written to a file named `design.md` in the current increment folder (for example: `.../increments/<slug>/design.md`).

The design document MUST follow this structure:

1. Design Title

- First-level heading in the form:
  - `Design: <Short, Descriptive Title>`
- The title should usually correspond to the increment’s goal or user-facing change.

2. Context and Problem

- Short restatement of:
  - The increment’s goal (WHAT).
  - Why this change is being made now.
- Brief summary of:
  - The relevant existing system behavior.
  - Which parts of the system are in play (components, services, modules, data stores).
- Links to:
  - `increment.md`.
  - `CONSTITUTION.md`.
  - Relevant ADRs or prior designs (if any).

3. Proposed Solution (Technical Overview)

- High-level description of the design.
- Which components/modules/services are involved.
- How responsibilities are split or changed.
- Any new interfaces, contracts, or data flows introduced.
- A short narrative of how a typical request or flow passes through the system after this change.

4. Scope and Non-Scope (Technical)

- In-scope technical changes for this increment:
  - Components and concerns this design explicitly covers.
- Explicitly out-of-scope items, even if related:
  - Refactors or features that are intentionally deferred.
- How this design fits into any broader roadmap or architecture, if relevant.

5. Architecture and Boundaries

- Description (and optionally references to diagrams) of:
  - Components and their interactions.
  - Key data flows and lifecycles.
- Reference to guardrails from `CONSTITUTION.md`:
  - Layering rules.
  - Ownership boundaries.
  - Allowed dependencies.
- Explanation of how the design:
  - Respects these guardrails, or
  - Intentionally adapts them (with justification).

6. Contracts and Data

- New or changed:
  - APIs (request/response shapes, error handling).
  - Events or messages (schemas, topics/queues).
  - Data models or storage schemas.
- Compatibility considerations:
  - How existing consumers are affected.
  - Migration and versioning strategy if needed.
- Any assumptions about data volume, performance, or retention that affect the design.

7. Testing and Safety Net

- Test strategy for this design:
  - Unit tests:
    - Which modules/functions/classes should be covered and what behaviors they must verify.
  - Integration / end-to-end tests:
    - Which flows or contracts must be exercised.
  - Regression tests:
    - Known bugs that should be prevented from reoccurring.
- Notes on:
  - Test data / fixtures and environments.
  - Potential flakiness risks and mitigations.

8. CI/CD and Rollout

- CI implications:
  - Any new jobs or pipeline steps.
  - Changes to commands (build, lint, test) if any.
- Rollout plan:
  - How changes are expected to be deployed through existing pipelines.
  - Whether feature flags or staged rollout are recommended.
  - Any manual steps that must be minimized or automated later.
- Rollback plan:
  - How to revert or mitigate this change if it misbehaves.
  - How to disable the behavior (for example: flags, configuration switches).

9. Observability and Operations

- Logging:
  - What should be logged.
  - Important context fields (IDs, correlation tokens, user IDs, etc.).
- Metrics:
  - New or updated metrics (counters, histograms, gauges).
  - How they relate to user/business outcomes (for example: success/failure rates, latencies).
- Alerts and Dashboards:
  - Any SLOs or alerts affected or introduced.
  - Dashboards that should be created or updated.
- Operational considerations:
  - Any known operational risks (for example: increased load, new dependencies).

10. Risks, Trade-offs, and Alternatives

- Known risks:
  - Technical, operational, or organizational.
- Trade-offs:
  - Why this approach was chosen over obvious alternatives.
- Alternatives:
  - Brief description of alternatives considered.
  - When they might be revisited.

11. Follow-up Work

- Potential future increments:
  - Deeper refactors, optimizations, or feature expansions suggested by this design.
- Tech debt or clean-up:
  - Work that should be done later, but not in this increment.
- Any monitoring or validation tasks to perform after rollout.

12. References

- Links to:
  - `CONSTITUTION.md`.
  - `increment.md`.
  - ADRs.
  - Relevant tickets/issues.
  - Other designs or documents that influenced this one.

---

### Examples (Conceptual)

Good designs using this structure typically:

- Address a **single increment**:
  - For example: “Add password reset endpoint” or “Instrument key usage metrics”.
- Touch a limited set of components:
  - For example: one service and its API, or one front-end route and its backing calls.
- Include:
  - A clear testing strategy (unit plus at least one integration/flow test).
  - A straightforward rollout and rollback approach.
  - Specific observability updates (logs/metrics) tied to the increment’s success criteria.

They are **short enough to read in minutes**, but detailed enough that an engineer can:

- Plan small, safe implementation steps.
- Write appropriate tests.
- Understand risks, trade-offs, and follow-up options, all in the context of the existing codebase.