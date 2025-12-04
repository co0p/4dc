## Inputs

The increment MUST be grounded in:

1. **The project at the provided root path**

   The executing LLM MUST:

   - Treat the path argument as the **subject project root**.
   - Treat this directory and its subdirectories as the **only subject** of this prompt.
   - NOT rely on content from parent directories, sibling projects, or other repositories as primary context.
   - Within that scope, locate:
     - `CONSTITUTION.md` (if present) – this is the primary “WHY” and quality bar.
     - The main description artifact (e.g. `README.md`).
     - Any existing `design.md`, `implement.md`, `improve.md`, ADRs, or similar docs.
   - Infer, from within this scope:
     - What the product currently does and for whom.
     - Key constraints (technical, legal, operational).
     - Existing delivery practices (e.g. CI/CD, testing, release cadence).

2. **The increment description from the user**

   The user will provide a **short increment description** that may include:

   - A **problem statement** (what’s wrong or missing for users/business).
   - A **desired outcome** (what should be possible after this change).
   - Any **constraints** (timing, risk level, dependencies, etc.).

   The LLM MUST:

   - Treat this as **product intent**, not a fixed technical solution.
   - Resolve ambiguity by:
     - Asking targeted clarifying questions if critical information is missing.
     - Narrowing or splitting the idea into **one primary outcome** for this increment, plus optional follow-ups.

3. **Alignment with the Constitution**

   If `CONSTITUTION.md` is present, the LLM MUST:

   - Respect:
     - Values & principles (e.g. small changes, reliability, observability).
     - Delivery & testing expectations.
     - Any explicit “do not do” constraints.
   - Ensure the increment:
     - Does **not** contradict explicit values (e.g. no “big bang” if the constitution favors small steps).
     - Moves the project **toward** the constitution’s ideal behaviors.

4. **Context from recent work (optional but recommended)**

   If available within the scoped path, the LLM SHOULD consider:

   - Recent increments and `improve.md` documents.
   - ADRs relevant to the product area.
   - Open issues or TODOs in the repo.

   This helps avoid:

   - Duplicating existing work.
   - Proposing increments that conflict with recent decisions.
   - Ignoring known risks or pitfalls.

5. **Increment Qualities**

   The increment MUST be designed so that it is:

   - **Small** – something a team can realistically complete and ship soon.
   - **Testable** – success can be checked via tests, metrics, or clear behaviors.
   - **Releasable** – does not require special one-off processes or risky coordination.
   - **Observable** – we can see its impact (or lack of impact) after release.