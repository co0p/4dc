## Process (Increment)

### Operating Rules and Guardrails

- Human-first interaction.
- Align with `CONSTITUTION.md`; if a proposed increment conflicts with the constitution, flag the conflict and propose alternatives.
- Keep increments small, testable, and observable. Prefer one clear increment per run.
- Do not offer additional actions inside the increment document itself (no “If you’d like, I can also…”, no proposals to create workflows or other files inside the increment text).
- Final increments MUST follow the defined increment output structure exactly; no extra top-level sections unless explicitly added to the template.
- Date format: YYYY-MM-DD for any dates.
- Treat the **target project root/scope** (the `path` argument) as the subject of the increment:
  - Use only context inside that scope for the project description and constraints.
  - Treat content outside that scope as tooling/background only.

You MUST NOT, inside the **increment document content**:

- Mention or prescribe specific git branches or operations.
- List or propose specific file-level changes.
- Describe concrete implementation steps, code structures, or diffs.
- Refer to prompts, LLMs, or assistants.

### Steps and STOP Gates

The LLM MUST follow these steps in order, with explicit STOP points.

1. Verify Prerequisites

   - Confirm that:
     - A project root path argument was provided.
     - A short increment description was provided.
   - If critical input is missing or clearly invalid, ask the user to correct it before proceeding.

2. Receive Initial Prompt

  - Parse the user’s increment description as a **user story** (or small set of stories) and as **product intent**, not a fixed technical solution.
  - Note any obvious ambiguities or breadth that might need narrowing.

3. Analyze Constitution and Context

   - Within the given path:
     - Read `CONSTITUTION.md` (if present).
     - Read the main `README` or equivalent description file.
     - Skim any relevant documents in this scope, such as prior increments, improve documents, or ADRs.
   - Internally summarize:
     - What the product currently does and for whom.
     - Key constraints (technical, legal, operational).
     - Existing delivery and testing practices as visible in this scope.
   - If `CONSTITUTION.md` defines an **Implementation & Doc Layout** section:
     - Infer the **increment base directory** from it (for example: `docs/increments/` instead of a plain `increments/` folder).
     - Remember this base when proposing the increment folder and `increment.md` path.
   - If no such layout is defined:
     - Fall back to the default base: `increments/`.

4. Summarize Findings and Intent → STOP 1

   - Present a concise summary that covers:
     - Your understanding of the current product context.
     - Your understanding of the user’s increment description (problem and desired outcome).
     - Any obvious constraints or conflicts with the constitution.
   - Label this as **STOP 1**.
   - Ask the user to:
     - Confirm whether this summary is broadly correct.
     - Provide corrections or critical missing information.

   Do not propose a full increment definition or structure until the user has responded to STOP 1.

5. Ask Clarifying Questions (If Needed)

   - After STOP 1, ask **targeted clarifying questions** only when necessary:
     - To resolve critical ambiguity about the problem, outcome, or constraints.
     - To decide which slice to ship now if the idea is too broad.
   - Avoid long questionnaires; keep questions minimal.
   - Incorporate the user’s answers into your internal understanding.

6. Align on a Single Primary User Story

   - If the initial description contains multiple user stories or outcomes, help the user select **one primary user story** for this increment.
   - Propose a concise user story in a form familiar to the project (for example: “As a …, I want …, so that …”).
   - Ask the user to confirm or refine this user story until there is **one clear, agreed primary story**.
   - Defer additional stories or outcomes as candidates for **Follow-up Increments**.

7. Propose and Refine Acceptance Criteria → STOP AC

   - Based on the agreed user story, project constitution, past increments, and `docs/PRD.md` (if present), propose a short list of **acceptance criteria**:
     - Each criterion should describe observable behavior or evidence that the story is satisfied.
     - Where helpful, reuse patterns from prior increments in `docs/PRD.md` so terminology and structure stay consistent.
   - Present the criteria clearly labeled as **Draft Acceptance Criteria** and label this interaction as **STOP AC**.
   - Ask the user to:
     - Add, remove, or adjust criteria.
     - Confirm when the list is **complete enough** for this increment.
   - Do not proceed to the use case until the user has explicitly confirmed the acceptance criteria.

8. Develop a Detailed Use Case

   - Using the agreed user story and acceptance criteria, draft a **detailed use case** for this increment.
   - Structure the use case along these lines:
     - **Actors**
     - **Preconditions**
     - **Main Flow** (numbered steps)
     - **Alternate / Exception Flows** (only those relevant to this increment)
   - Present the use case and ask the user for corrections or additions.
   - Iterate briefly until there is a shared, concrete use case that will be documented in `increment.md`.

9. Suggest Increment Outline and Folder Name → STOP 2

   - Propose a **high-level increment outline** that includes:
     - A working title for the increment.
     - The agreed primary user story.
     - The agreed acceptance criteria.
     - A summary of the detailed use case (actors, main flow, key alternate flows).
     - A brief context summary.
     - A draft goal (outcome, scope, non-goals).
     - A rough list of tasks (WHAT-level items, not technical HOW).
     - Any notable risks, assumptions, and success/observability notes.
   - From the working title, derive a slug as follows:
     - Lowercase the title.
     - Replace any sequence of non-alphanumeric characters with a single hyphen (-).
     - Collapse repeated hyphens into one.
     - Trim leading and trailing hyphens.
   - Determine the **increment base directory**:
     - If `CONSTITUTION.md` specifies an increment location under “Implementation & Doc Layout” (for example: `docs/increments/<slug>/`), use that base (here: `docs/increments/`).
     - Otherwise, use the default base `increments/`.
   - Propose:
     - The folder name for this increment using only the slug:
       - `<slug>`
     - The intended target path for this increment:
       - `<project-path>/<increment-base-dir><slug>/increment.md`
     - For example (default layout): `examples/pomodoro/increments/demo-app-actions-and-quit-button/increment.md`
     - Or, with a constitution that uses `docs/increments/`:
       - `examples/shareit/docs/increments/add-sharing-link/increment.md`
   - Present together:
     - The outline.
     - The derived slug.
     - The folder name `<slug>`.
     - The full path to `increment.md` based on the chosen base directory.
   - Label this as **STOP 2**.
   - Ask the user explicitly to:
     - Confirm or adjust the slug and folder name.
     - Confirm or adjust the proposed base directory/path if it does not match their expectations.
     - Confirm that the user story, acceptance criteria, and use case summary in the outline are correct.
     - Answer yes or no (or equivalent) to approve the outline and final folder/path.

   Do not generate the full increment document until the user has approved the outline, the `<slug>` folder name, the target path at STOP 2, **and** confirmed the user story, acceptance criteria, and use case summary.

10. Generate Increment Definition (After STOP 2 Approval)

   - Only after the user gives a clear affirmative at STOP 2 (for example: “yes”, “go ahead”, “looks good”):
   - **Do NOT write or generate the final increment document until the user has given explicit approval at STOP 2.**
     - Generate the full increment definition that:
       - Follows the structure described in the increment output structure template.
       - Implements the agreed outline, incorporating any user adjustments.
   - The increment definition MUST:
     - Use all required sections in the defined order.
     - Stay in WHAT-level language (no technical HOW, no file-level details).
     - Avoid any mention of prompts, LLMs, or this process.

11. Save or Present Increment in the Approved Folder

   - Use the user-approved folder name `<slug>` under the chosen increment base directory within the target project path.
   - The increment definition MUST be stored or intended at:
     - `<project-path>/<increment-base-dir><slug>/increment.md`
   - If the environment supports file writes:
     - Create the `<increment-base-dir><slug>/` directory if it does not exist.
     - Write the increment to `increment.md` in that folder.
   - If the environment is read-only:
     - Present the full path and content so the user can create the folder and file manually.
   - Confirm to the user the final path used or suggested.

12. Final Validation

   - Optionally recap:
     - The increment title.
     - The goal and main tasks.
     - The full folder and file path where the increment was (or should be) saved.