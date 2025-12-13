## Goal

Turn the current **increment** (product-level WHAT) into a **technical design** (HOW) that:

- Respects the **Project Constitution** (`CONSTITUTION.md`), including its **Implementation & Doc Layout** and any `constitution-mode` (for example: `lite`, `medium`, `heavy`).
- Is **small and incremental**, matching the scope of the increment.
- Is **testable and verifiable** through automated checks (keeping the constitution’s expectations in mind — lighter for `lite` mode, richer for `medium`/`heavy`).
- Can pass cleanly through **CI/CD** without unusual, risky procedures.
- Is **observable and operable** when running in real environments, at the level the constitution expects.
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
   - When `constitution-mode` is `lite`, keep the safety net focused and pragmatic (for example: essential unit tests and a small integration test where it matters most); for `medium`/`heavy`, follow stronger testing expectations.

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
   - For `lite` constitutions, this may be as simple as clear, structured request and error logging; for `medium`/`heavy`, it may include metrics and alerts.

6. Stay Within Increment Scope

   - The design MUST stay within the current increment’s scope and non-goals.
   - If deeper or broader changes are uncovered, call them out explicitly as:
     - Risks and/or
     - Candidates for **follow-up increments** or separate design work.

7. Stay at the Design Level, Not Implementation Tasks

   - The design MUST NOT be an implementation task list.
   - Do **not** describe step-by-step edit sequences, git operations, or a chronological plan.
   - Do **not** specify:
     - Which files to edit in what order.
     - Per-file actions or changes.
     - Step sequences or PR groupings.
     - Deployment commands or rollout scripts.
   - Focus on:
     - Components and responsibilities.
     - Interfaces and data flows.
     - Test strategy (behaviors and coverage expectations, not test file names or sequences).
     - CI/CD and observability as constraints and targets (what should be true), not as implementation steps.
   - Leave **concrete work steps** to the Implement phase.