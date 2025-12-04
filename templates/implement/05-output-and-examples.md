## Output Structure and Examples

The generated **implementation plan** MUST be written to a file named `implement.md` in the current increment folder (for example: `.../increments/<slug>/implement.md`).

The implementation plan document MUST follow this structure:

1. Implementation Plan Title

- First-level heading in the form:
  - `Implementation Plan: <Short, Descriptive Title>`
- The title should usually align with or closely follow the increment and design titles.

2. Context and Inputs

- Briefly restate:
  - The increment’s goal and main tasks (WHAT).
  - The design’s high-level approach (HOW).
- List the key inputs this plan assumes:
  - `increment.md`.
  - `design.md`.
  - Any critical ADRs or documents.
- Mention any important assumptions (for example: availability of staging, feature flag system).

3. High-Level Approach

- Summarize:
  - The main workstreams (for example: backend, frontend, data migration, observability).
  - How these workstreams relate to the increment’s goal and design.
- Keep this section short; it is a narrative overview, not detailed steps.

4. Workstreams and Steps

For each workstream, provide:

- A short description of the workstream’s purpose.
- A numbered list of **concrete steps**, in a sensible order.  
  Each step should clearly state:
  - What is being changed (at a high but concrete level — components, areas of code, tests).
  - Why the step is needed (how it supports the design and increment goal).
  - Any immediate validation to perform (for example: run tests, check a specific behavior locally).

Example pattern (conceptual, not literal text):

- Workstream: Backend API changes  
  1. Introduce new handler and routes for the feature behind a flag.  
  2. Update validation and error handling for the new endpoint.  
  3. Add unit tests for the handler and edge cases.  
  4. Add integration tests for the full request/response flow.

Steps SHOULD:

- Be small and reviewable.
- Avoid coupling multiple unrelated changes into a single step.
- Reference parts of the codebase in a way that is easy for engineers to locate (for example: component names, module names), without requiring exact line numbers.

Steps are ordered:

- Within each workstream, steps are listed in the **recommended execution order**.
- Unless the user specifies otherwise, an assistant or engineer starting work from this plan SHOULD:
  - Begin with the **first workstream** listed, and
  - Take the **first step** in that workstream that is not yet marked as done.

5. Testing Plan

- Summarize how testing fits into the steps:
  - Which test suites to run and when.
  - New tests to add or extend (unit, integration, end-to-end, regression).
- Note:
  - Any special test data, fixtures, or environment configuration.
  - How to avoid or manage flakiness.

6. CI/CD and Rollout Plan

- Describe how the steps will move through CI/CD:
  - Any expected changes to pipelines, if applicable.
  - Expected commands or jobs to run.
- For rollout:
  - How to enable the new behavior:
    - Feature flags or configuration switches.
    - Staged or canary rollout, if appropriate.
  - When to remove temporary guards or flags (if part of this increment or later).
- For rollback:
  - What to do if a particular step fails in production.
  - Which changes are easy to revert and how.
  - How to quickly disable the behavior without full rollback (if possible).

7. Observability and Post-Deployment Checks

- Logging and metrics:
  - Which log lines or metrics should be in place before rollout.
  - Any alerts or dashboards that should be updated.
- Validation steps:
  - What to verify in:
    - Local/dev environments.
    - Staging or pre-production.
    - Production after deployment (for example: metrics, logs, traces, user behavior).
- Tie these checks back to:
  - The increment’s success criteria.
  - The design’s observability guidance.

8. Risks, Dependencies, and Coordination

- List:
  - Dependencies on other teams, systems, or changes.
  - Risks that could make implementation difficult or slow.
- For each major risk:
  - Briefly state:
    - Why it is a risk.
    - How to mitigate it (for example: order of steps, extra tests, flags).
- Note any coordination tasks:
  - Communication with support, ops, or other teams.
  - Timing constraints (for example: maintenance windows).

9. Follow-up and Cleanup

- Document:
  - Any cleanup steps that can be done after the main rollout (for example: removing deprecated paths).
  - Any tech debt that this increment will create and how to address it later.
- Suggest where a follow-up increment or design might be appropriate.

---

## How to Use This Plan

This implementation plan is meant to be used as a **backlog for this increment**.

After `implement.md` exists:

- When starting from this plan:
  - By default, start with the **first workstream** listed and the **first step** in that workstream that is not yet completed.
  - If the team has indicated a different starting point (for example, via comments, checkmarks, or additional notes), follow that instead.
- For each chosen step:
  - Read the step, its workstream, and any linked testing/rollout/observability notes.
  - Open your editor and terminal in the project root.
  - Make the corresponding code, test, and configuration changes.
  - Run the tests or checks mentioned for that step.
  - Update your normal tracking surface (for example: PR checklist, issue checklist, or notes) to reflect that the step is completed.
- Keep the plan and reality in sync:
  - If you discover that a step needs to be split, merged, or adjusted, update `implement.md` or note the change in your tracking tool.
  - If new work beyond the increment’s scope appears, treat it as:
    - A risk to call out, and/or
    - A candidate for a **new increment**.

The 4dc loop is:

- **Increment** – define the product-level WHAT in `increment.md`.
- **Design** – define the technical HOW in `design.md`.
- **Implement** – define the ordered steps in `implement.md`.
- **Improve** – periodically analyze and improve the overall system.

Actual code changes are done by humans (and their usual coding tools) by following the steps in `implement.md`, one by one.

---

### Examples (Conceptual)

Good implementation plans using this structure typically:

- Are directly traceable to a **single increment** and its **design**.
- Break work into:
  - A handful of clear workstreams.
  - Concrete, ordered steps within each workstream.
- Include:
  - Specific test additions and when to run them.
  - A practical rollout and rollback approach.
  - Concrete post-deployment checks tied to success criteria.

They are **practical documents** that engineers can follow day-to-day, while still leaving room for human judgment and adaptation.