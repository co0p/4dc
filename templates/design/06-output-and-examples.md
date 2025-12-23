## Output Structure and Examples

The generated **technical design** MUST be written to a file named `design.md` in the **current increment folder** (the folder pointed to by `path`, which already contains `increment.md` and, per the project’s constitution, will also hold `implement.md` for this increment).

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
- The level of detail should align with the project’s constitution and `constitution-mode` (for example, a lighter test plan for a `lite` demo app, a more formal plan for critical services).

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
- For `lite` mode constitutions, it is acceptable for this section to focus mainly on clear, structured logs and simple checks; for heavier modes, it may be more elaborate if required.

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

13. Machine-Readable Artifacts (when applicable)

- When the increment introduces or changes **APIs, events, or schemas**, the design MUST include clearly delimited, machine-readable artifacts such as:
  - OpenAPI or JSON/YAML snippets for HTTP or RPC APIs.
  - JSON Schema or equivalent definitions for key payloads.
  - Event or message schemas for pub/sub or queue-based integrations.
- These artifacts SHOULD be presented as fenced code blocks, labeled with their intended purpose and (optionally) target file path, for example:

  - `# API Contract: POST /widgets (intended file: api/contracts/widgets-openapi.yaml)`
    ```yaml
    # OpenAPI fragment
    openapi: 3.0.0
    ...
    ```

- When the increment affects **system structure or architecture**, the design SHOULD also include updated **Mermaid-based C4 diagrams** that can be applied to `ARCHITECTURE.md` or the project’s architecture docs:
  - At least one **C4 Container (Level 2)** diagram for the relevant part of the system.
  - One or more **C4 Component (Level 3)** diagrams for key containers impacted by this increment.
- These diagrams SHOULD be presented as fenced `mermaid` code blocks with a short caption indicating where they belong, for example:

  - `# Architecture Diagram (C4 Container L2, intended for ARCHITECTURE.md)`
    ```mermaid
    %% Mermaid C4 container view here
    ```

  - `# Architecture Diagram (C4 Component L3 for <Container>, intended for ARCHITECTURE.md)`
    ```mermaid
    %% Mermaid C4 component view here
    ```

- Machine-readable artifacts MUST remain consistent with the narrative in:
  - Proposed Solution.
  - Architecture and Boundaries.
  - Contracts and Data.
  - Testing and Safety Net.

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
- Understand risks, trade-offs, and follow-up options, all in the context of the existing codebase and the project’s constitution.