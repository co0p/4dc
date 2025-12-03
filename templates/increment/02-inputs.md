# Inputs and Scope (Increment)

You have access to:

- The repository contents as exposed by the tools that call you.
- The project’s `CONSTITUTION.md` and other root-level docs in the **target project root**.
- Any answers the user provides during this interaction.
- A **project root path argument** provided by the calling tool (for example `"."` or `"examples/pomodoro"`).

You MUST apply the same scoping rules as the constitution prompt:

1. **Target project root**

   - Treat the project root path argument as the **project root directory for the TARGET project**.
   - Files that live **directly** in this directory (such as `README.md`, `CONSTITUTION.md`, `LICENSE`, and other top-level markdown or configuration files) define the project context.
   - Use these root-level files for:
     - Product / domain understanding.
     - High-level goals and constraints.
     - Engineering principles and non-negotiables.

2. **Subdirectories under the target root**

   - Subdirectories (e.g., `src/`, `docs/`, `.github/`, `examples/`, `templates/`, `tests/`, etc. under the project root) MUST NOT override the primary product description.
   - Use them only to understand:
     - Existing capabilities and code structure relevant to the increment.
     - Engineering practices (tests, CI, workflows).
     - Implementation details that may constrain or inform the increment.

3. **Files outside the target root**

   - Treat files outside the project root path as:
     - Tooling, frameworks, or background material.
   - You MUST NOT:
     - Treat them as the subject of the increment.
     - Copy product descriptions from them into this project’s increment.
     - Mention the host/framework repo (e.g., `4dc`) in the increment, unless it is an explicit runtime dependency or architectural element of the project.

When inferring context, you MUST:

- Use the target project’s `CONSTITUTION.md` as the primary source of:
  - Principles and trade-offs.
  - Pillar interpretations.
  - Non-negotiables and constraints.
- Use the target project’s root `README.md` (if present) for:
  - Product and user context.
  - High-level goals the increment should support.
- Use code, tests, and docs under subdirectories only to:
  - Understand where and how the increment might land.
  - Avoid proposing increments that obviously conflict with existing structure.

If critical context is missing or ambiguous (e.g., no constitution, unclear user), you MUST ask targeted clarifying questions before finalizing an increment proposal.