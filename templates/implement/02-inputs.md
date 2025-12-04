## Inputs

The implementation plan MUST be grounded in:

1. The increment folder (current path)

   The `path` argument points to an **increment folder**, for example:

   - `<project-root>/increments/<slug>`

   Within this folder, the LLM MUST look for:

   - `increment.md` – the product-level WHAT (context, goal, tasks, risks, success criteria).
   - `design.md` – the agreed technical HOW (architecture, components, contracts, tests, rollout, observability).

   If `design.md` is missing:

   - The LLM MUST:
     - Treat this as a serious gap.
     - Ask the user whether to:
       - Proceed with a **lightweight plan** based on `increment.md` alone, or
       - Pause and create a design before continuing.
   - The implementation plan SHOULD be more conservative and higher-level if no design is provided.

2. The project codebase

   The increment folder sits inside a project root (for example: `examples/pomodoro`).

   The LLM MUST:

   - Treat the containing project as the **codebase** to be changed.
   - Inspect relevant code and tests under that project path in order to:
     - Understand where changes will likely go.
     - Identify existing patterns and conventions to follow.
     - See potential risks or dependencies.

3. The project constitution and practices

   If `CONSTITUTION.md` is present under the project root, the LLM MUST:

   - Respect:
     - Values and principles (for example: small changes, safety, observability).
     - Delivery, testing, and review expectations.
   - Ensure the implementation plan:
     - Proposes steps that can run through normal CI/CD paths.
     - Does not require fragile, one-off processes if they can be avoided.

4. Increment and design scope

   The plan MUST:

   - Stay within:
     - The increment’s **goal and tasks**.
     - The design’s **scope and boundaries**.
   - If new work is discovered that clearly goes beyond:
     - Call it out explicitly as:
       - A risk.
       - Follow-up work.
       - A candidate for a separate increment.

5. Team and workflow assumptions (from context where visible)

   Where visible from config, docs, or the constitution, the plan SHOULD respect:

   - How tests are usually run (commands, environments).
   - How feature flags, migrations, or config changes are normally introduced.
   - How code review and release typically work for this project.

   If these are not visible, the plan SHOULD:

   - Use common-sense defaults (for example: “run the existing test suite”, “use feature flags where appropriate”).
   - Flag assumptions so humans can adapt them.