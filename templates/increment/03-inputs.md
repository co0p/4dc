## Inputs

The increment MUST be grounded in:

1. The project at the provided root path

   The `path` argument points to a **project or subproject root** (for example: `.` or `examples/pomodoro`).

   The executing LLM MUST:

   - Treat this path as the **subject project root**.
   - Treat this directory and its subdirectories as the **only subject** of this prompt.
   - Not rely on content from parent directories, sibling projects, or other repositories as primary context.

    Within this scope, it should locate and use:

    - `CONSTITUTION.md` (if present) – the primary source of project values and guardrails.
    - The main description artifact (for example: `README.md` or similar).
    - `docs/PRD.md` (if present) – the primary product requirements document, including:
       - The list of past and current increments.
       - User stories and acceptance criteria associated with those increments.
       - Stable increment identifiers and links to increment folders.

2. The increment description from the user

   The user provides a **short increment description** that may include:

   - A problem statement (what is wrong or missing for users or the business).
   - A desired outcome (what should be possible after this change).
   - Any constraints (timing, risk level, dependencies, etc.).

   The LLM MUST:

   - Treat this as **product intent**, not a fixed technical solution.
   - Resolve ambiguity by:
     - Asking targeted clarifying questions if critical information is missing.
     - Narrowing or splitting the idea into **one primary outcome** for this increment, plus optional follow-ups.

3. Alignment with the Constitution

   If `CONSTITUTION.md` is present under the project root, the LLM MUST:

   - Respect:
     - Values and principles (for example: small changes, reliability, observability).
     - Delivery and testing expectations.
     - Any explicit “do not do” constraints.
   - Ensure the increment:
     - Does not contradict explicit values.
     - Moves the project **toward** the constitution’s ideal behaviors.
   - When the constitution defines an **Implementation & Doc Layout**:
     - Use that layout when proposing where increment artifacts should live (for example, `docs/increments/<slug>/increment.md` instead of a hard-coded default).

4. Context from recent work (optional but recommended)

   If available within the scoped path, the LLM SHOULD consider:

   - The "Increments" (or similar) section of `docs/PRD.md`, including past user stories and acceptance criteria.
   - ADRs relevant to the product area.
   - Open issues or TODOs in the repo.

   This helps avoid:

   - Duplicating existing work.
   - Proposing increments that conflict with recent decisions.
   - Ignoring known risks or pitfalls.

5. Increment Qualities

   The increment MUST be designed so that it is:

   - **Small** – something a team can realistically complete and ship soon.
   - **Testable** – success can be checked via tests, metrics, or clear behaviors.
   - **Releasable** – does not require special one-off processes or risky coordination.
   - **Observable** – we can see its impact (or lack of impact) after release.