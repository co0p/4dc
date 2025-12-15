````markdown
## Process

Follow this process to produce a `CONSTITUTION.md` that is scaled to the project and grounded in the existing codebase and context.

The `path` argument points at the **project root**.

### Phase 1 – Inspect and Propose Mode (STOP 1)

1. Inspect the Project

   - Read:
     - Any existing `README.md` under `path`.
     - Any visible `docs/`, `increments/`, `adr/`, or `improve`-like directories.
     - CI configuration files where present (for example: `.github/workflows`, `ci/`).
   - Skim the code layout:
     - Primary language(s).
     - Size/complexity signals (for example: number of packages/modules, presence of services).
   - Note:
     - Whether this looks like:
       - A small script or demo.
       - A single-service application.
       - A larger, multi-module or multi-team system.

2. Propose a Constitution Mode

   - Based on the above, propose one of:
     - **lite** — for:
       - Small scripts, tools, demos.
       - Low criticality, mostly internal use.
       - Simple CI/testing.
     - **medium** — for:
       - Typical product/service.
       - Some real users or customers.
       - CI, tests, and observability matter.
     - **heavy** — for:
       - Long-lived, critical, multi-team or regulated systems.
       - Strong uptime / compliance expectations.
       - More formal ADR and Improve usage.

   - Explain briefly:
     - Why you think this mode fits.
     - What this mode implies in practice (1–2 bullets).

3. Summarize Findings and Mode Suggestion → STOP 1

   - Present a short summary:
     - What this project appears to be (size, type).
     - The **proposed Constitution Mode** and why.
   - Clearly label this as **STOP 1**.
   - Ask the user to confirm or change the mode:
     - “Does `lite / medium / heavy` feel right? If not, which mode would you choose and why?”

   **Wait for user input** before continuing.

### Phase 2 – Draft Principles and Layout (STOP 2)

4. Confirm Mode and Capture It

   - Use the user’s choice as the final `constitution-mode`.
   - This value will be written near the top of `CONSTITUTION.md` and guide the rest of the document.

5. Propose Implementation & Doc Layout

   - Propose a default layout, adapted to what you saw under `path`. For example:

     ```markdown
     ## Implementation & Doc Layout

     - **Increment artifacts**
       - Location: `increments/<slug>/`
       - Files:
         - `increment.md`
         - `design.md`
         - `implement.md`

     - **Improve artifacts**
       - Location: `docs/improve/`
       - Filename pattern: `YYYY-MM-DD-improve.md`

     - **ADR artifacts**
       - Location: `docs/adr/`
       - Filename pattern: `ADR-YYYY-MM-DD-<slug>.md`
     ```

   - If the repo already has a layout that’s close to this (for example: existing `docs/adr`, `increments/`), adapt your proposal to match.

6. Select and Tailor Principles

   - From the palette below, choose a **small subset** appropriate to the chosen mode and project type.
     - For `lite`: focus on 2–3 core principles.
     - For `medium`: 4–6 principles.
     - For `heavy`: more principles, especially around DORA/observability/ADRs.

   - Palette (you will tailor and rephrase for this project):

     - **Small, safe steps** (Kent Beck)
       - Prefer many small, reversible changes over large, risky ones.
     - **Refactoring as everyday work** (Fowler, Feathers)
       - It is normal and expected to refactor code to keep it clean and simple.
     - **Separation of concerns & stable boundaries** (Martin)
       - Domain logic, IO, and frameworks are kept separate where practical.
     - **Lean flow & limited WIP** (Poppendiecks)
       - Avoid huge, multi-week increments; keep work flowing in small slices.
     - **DORA-aware delivery** (Forsgren, Humble, Kim)
       - Favor changes that reduce lead time and change failure rate; keep MTTR low.
     - **Responsibility-driven design & DDD** (Wirfs-Brock, Evans, Metz)
       - Components have clear responsibilities and use domain language consistently.
     - **Pragmatic DRY & simplicity** (Thomas & Hunt, Fowler)
       - Remove real duplication; avoid speculative abstraction.
     - **Behavioral acceptance** (BDD / Gherkin)
       - Express important behaviors in Given/When/Then style where helpful.

   - For each chosen principle:
     - Write a short **name and description**.
     - Optionally add 1–2 examples that make sense for this project.

7. Draft Constitution Outline → STOP 2

   - Draft an outline for `CONSTITUTION.md`, something like:

     ```markdown
     # Project Constitution

     constitution-mode: <lite|medium|heavy>

     ## 1. Purpose and Scope
     ## 2. Implementation & Doc Layout
     ## 3. Design & Delivery Principles
     ## 4. Testing, CI/CD, and Observability (if relevant)
     ## 5. ADR and Improve Usage (if relevant)
     ```

   - Briefly summarize each section’s intended content, including:
     - Chosen mode.
     - Proposed layout.
     - Selected principles.

   - Label this as **STOP 2** and ask the user:
     - Whether they want to adjust:
       - The layout.
       - The set of principles.
       - The level of emphasis on DORA/observability.

   **Wait for explicit approval** before writing the final `CONSTITUTION.md`.

### Phase 3 – Write `CONSTITUTION.md` After YES

8. Produce the Final `CONSTITUTION.md` (After STOP 2 Approval)

   - Only after the user clearly approves the outline:
     - Generate `CONSTITUTION.md` that:
       - Follows the agreed outline and layout.
       - Uses the chosen mode and principles.
       - Does **not** mention prompts, LLMs, or this process.

   - Keep it:
     - Short and readable.
     - Concrete and opinionated.
     - Directly usable by Increment, Design, Implement, and Improve prompts.

9. Final Sanity Check

   - Ensure that:
     - `constitution-mode` is clearly stated.
     - Implementation & Doc Layout is explicit and matches (or sensibly adjusts) the current repo.
     - Principles are:
       - Few in number.
       - Clearly described.
       - Mapped to concrete behavior where possible.

   - If anything feels overly heavy for the chosen mode, simplify.

Return the full `CONSTITUTION.md` content as the final output.
````