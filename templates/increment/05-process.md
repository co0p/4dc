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

   - Parse the user’s increment description as **product intent**, not a fixed technical solution.
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

6. Suggest Increment Outline and Folder Name → STOP 2

   - Propose a **high-level increment outline** that includes:
     - A working title for the increment.
     - A brief context summary.
     - A draft goal (outcome, scope, non-goals).
     - A rough list of tasks (WHAT-level items, not technical HOW).
     - Any notable risks, assumptions, and success/observability notes.
   - From the working title, derive a slug as follows:
     - Lowercase the title.
     - Replace any sequence of non-alphanumeric characters with a single hyphen (-).
     - Collapse repeated hyphens into one.
     - Trim leading and trailing hyphens.
   - Propose a folder name for this increment using only the slug:
     - `<slug>`
   - Propose the intended target path for this increment:
     - `<project-path>/increments/<slug>/increment.md`
     - For example: `examples/pomodoro/increments/demo-app-actions-and-quit-button/increment.md`.
   - Present together:
     - The outline.
     - The derived slug.
     - The folder name `<slug>`.
     - The full path to `increment.md`.
   - Label this as **STOP 2**.
   - Ask the user explicitly to:
     - Confirm or adjust the slug and folder name.
     - Answer yes or no (or equivalent) to approve the outline and final folder/path.

   Do not generate the full increment document until the user has approved the outline and the `<slug>` folder name and path at STOP 2.

7. Generate Increment Definition (After STOP 2 Approval)

   - Only after the user gives a clear affirmative at STOP 2 (for example: “yes”, “go ahead”, “looks good”):
     - Generate the full increment definition that:
       - Follows the structure described in the increment output structure template.
       - Implements the agreed outline, incorporating any user adjustments.
   - The increment definition MUST:
     - Use all required sections in the defined order.
     - Stay in WHAT-level language (no technical HOW, no file-level details).
     - Avoid any mention of prompts, LLMs, or this process.

8. Save or Present Increment in the Approved Folder

   - Use the user-approved folder name `<slug>` under an `increments` directory in the target project path.
   - The increment definition MUST be stored or intended at:
     - `<project-path>/increments/<slug>/increment.md`
   - If the environment supports file writes:
     - Create the `increments/<slug>/` directory if it does not exist.
     - Write the increment to `increment.md` in that folder.
   - If the environment is read-only:
     - Present the full path and content so the user can create the folder and file manually.
   - Confirm to the user the final path used or suggested.

9. Final Validation

   - Optionally recap:
     - The increment title.
     - The goal and main tasks.
     - The full folder and file path where the increment was (or should be) saved.