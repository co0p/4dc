## Acceptance Criteria for the Implementation Plan

A generated `implement.md` is considered **acceptable** when:

1. Alignment with Increment and Design

   - It clearly references and respects:
     - `increment.md` (goal, scope, tasks, non-goals, success criteria).
     - `design.md` (architecture, components, contracts, tests, rollout, observability).
   - It stays within the increment’s scope and non-goals.
   - It does not contradict major design decisions without clearly calling them out as risks or issues.

2. Clarity and Actionability

   - Engineers can read the plan and understand:
     - What concrete steps need to be taken.
     - Roughly in what order.
     - How those steps map to parts of the codebase.
   - Steps are:
     - Small enough to be implementable and reviewable.
     - Written in straightforward, unambiguous language.
   - The plan avoids vague instructions like “just refactor X” without more detail.

3. Safety and Delivery Readiness

   - The plan supports:
     - Small, incremental changes.
     - A clear testing flow (which tests to add or run when).
     - Smooth integration into existing CI/CD pipelines.
   - It explicitly covers:
     - How to safely roll out the change.
     - How to roll back or mitigate issues.
     - How to verify success during and after rollout.

4. Observability and Validation

   - The plan includes:
     - Steps to update logging and metrics as needed.
     - Checks to perform in:
       - Local/dev environments.
       - Staging or pre-production.
       - Production after deployment.
   - It ties validation steps back to:
     - The increment’s success criteria.
     - The design’s observability plan.

5. Structure and Style

   - The document follows the structure defined in the implementation output structure template.
   - It is:
     - Concise but complete.
     - Written for a technical audience.
     - Free of meta-comments about prompts, LLMs, or this process.