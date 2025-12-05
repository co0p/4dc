# Process

This section defines **how** to run the improve analysis and generate a dated improvement document for the project rooted at `path`.

## 0. Subject and Scope

- Treat the given `path` as the **subject root**.
- You may read:
  - `CONSTITUTION.md`, README, and other docs under `path`.
  - Existing increment, design, implement, and improve documents under `path`.
  - Code and tests under `path`.
- Do **not** treat parent directories, sibling projects, or other repositories as your subject.
- Your goal is to analyze the health of the system under `path` and propose improvements; you do **not** implement changes yourself.
- When `CONSTITUTION.md` defines a `constitution-mode`, use it to scale:
  - How many improvements you propose.
  - How ambitious they are (still favoring small, safe refactorings across all modes).

---

## 1. Receive Initial Improvement Request

1. Restate the request in your own words:
   - Confirm that the user wants a **holistic improvement and learning pass** on the project at `path`.
   - Mention that the output will be a **dated improve document** under `docs/improve/`.

2. Optionally ask a small number of critical clarifying questions if needed:
   - For example, which environment is primary (web, CLI, service), or which area is most business-critical.
   - Keep questions minimal and focused.

---

## 2. Analyze Project Context and Assess Implementation

Review key project context under `path`:

- `CONSTITUTION.md` and similar guiding documents.
- Readme / high-level docs and onboarding material under `path`.
- Recent increments/designs/implements/improves under `path`.
- Representative code and tests (especially around recent changes).

### Assessment Tasks

- **Evaluate vs. Constitution:** Assess how well the implementation adheres to the project's core principles and constraints. Prepare to summarize this as per-principle **1–5 star ratings** with short rationales.
- **Evaluate vs. Design Goals:** Assess whether the implementation meets the intended design approach, component boundaries, and data flow.
- **Quality Evaluation:** Assess code quality, readability, maintainability, and testability.
- **Identify Risks:** List technical debt, potential bugs, performance concerns, or security issues.
- **Identify Architectural Opportunities:** Note opportunities for improved structure, patterns, or abstractions.

### STOP 1 – Context and Assessment Summary

After context analysis and assessment:

1. Provide a brief summary that includes:
   - The project’s purpose and main capabilities as seen under `path`.
   - Tech stack and notable architectural choices.
   - A concise initial assessment (constitution/design alignment, quality, key risks).
   - The detected `constitution-mode` (if present), and what that implies about how heavy or light the improvements should be.

2. Ask the user to:
   - Confirm or correct your understanding of the context.
   - Optionally highlight specific areas under `path` to prioritize (files, modules, domains).

3. **STOP and wait** for the user’s response before proceeding with deeper analysis.

---

## 3. Analyze Codebase Through Lenses

After STOP 1 and any clarifications:

- Analyze the codebase under `path` using the context-based lenses described in the **Lenses** section:
  - Naming & Clarity
  - Modularity & Separation
  - Architecture & Patterns
  - Testing & Reliability
  - Duplication & Simplicity
  - Documentation & Communication

- Identify and list actionable improvement suggestions, grouped by these lens contexts.
- For each suggestion:
  - Reference the relevant lens group.
  - Provide a clear rationale inspired by industry best practices.
  - Ground your rationale in specific observations from files under `path`.
  - Prefer proposals that can be implemented as **small, safe refactorings** rather than risky “big bang” changes.

Do **not** ask the user what to look for; use your analysis and the lenses to recommend improvements.

### Lessons and ADR Candidates

While you analyze:

- Capture **Lessons**:
  - What worked well.
  - What could be improved.
  - Emerging patterns that appear repeatedly.
- Surface **ADR candidates** when:
  - Divergent implementations or patterns appear (e.g., different error handling strategies).
  - A shared architectural direction would reduce confusion and duplication.

For each ADR candidate:

- Describe the observation and why it matters.
- Recommend whether an ADR should be created.
- When the user agrees that an ADR is needed, generate its content using the **ADR template described in this prompt** (the “ADR Output Template” section), as a separate document.

ADRs are independent artifacts and are **not** embedded into the improve document.

---

## 4. Draft the Improvement Plan Outline

Based on findings and lessons:

1. Draft an outline for the dated improve artifact following the Improve Output Format:

   - **Assessment:** Key bullet points:
     - Constitution and design alignment, including which constitution principles are most relevant to rate.
     - Quality.
     - Risks.
   - **Lessons:**
     - Worked Well.
     - To Improve.
     - Emerging Patterns.
   - **Improvements:**
     - A list of improvement **proposals**, each with:
       - A working title.
       - Lens.
       - Rough priority (H/M/L).
       - Likely file(s) to touch.
       - A sense of effort aligned to the project’s mode (e.g. quick wins in `lite`, some deeper refactors in `heavy`).

2. Ensure each proposal is:
   - Concrete enough to be turned into an increment.
   - Grounded in specific findings.
   - Appropriately sized (small to medium chunks of work), favoring incremental refactorings.

### STOP 2 – Outline Approval

- Present the outline to the user.
- Ask explicitly:
  - Whether the Assessment and Lessons reflect what they care about now.
  - Whether the proposed Improvements and their rough priorities make sense.
  - Whether you should proceed to generate the final dated improve document.

If the user requests changes:

- Update the outline (e.g., re-order, drop, refine proposals, or adjust ambition to better fit `constitution-mode`).
- Reconfirm before moving on.

Do **not** generate or overwrite the final improve content until the user explicitly says “yes” (or equivalent).

---

## 5. Generate the Dated Improve Document (After STOP 2 Approval)

Once the user explicitly approves the outline:

1. Generate the full content for a new dated improve file strictly following the Improve Output Format:

   - **Assessment:**
     - Per-constitution **1–5 star ratings** with a short rationale for each relevant principle.
     - A brief narrative plus bullet points for design alignment, quality, and risks.
   - **Lessons:** Three lists:
     - Worked Well.
     - To Improve.
     - Emerging Patterns.
   - **Improvements:** Each proposal as a separate `#### Improvement N: ...` section with:
     - Lens.
     - Priority.
     - Effort estimate (e.g., “30 min”, “2–3 h”).
     - Explicit file paths under `path`.
     - A specific, actionable change description.
     - Optional Increment Hint, phrased so it could be used directly as the short description for a future increment.

2. Ensure the text:

   - Does **not** mention prompts, LLMs, or assistants.
   - Reads as if written directly by the team.
   - Refers only to files and concepts under `path` as the subject.
   - Makes it easy to pick any Improvement and turn it into an increment.

---

## 6. Final Validation and Storage

Before presenting the final improve document content:

**Verification Checklist:**

- Assessment section includes:
  - Constitution and design alignment where applicable.
  - Per-constitution star ratings with concise rationales.
  - Quality and risk summary.
- Lessons section documents:
  - What worked well.
  - What to improve.
  - Emerging patterns.
- Each improvement proposal is a separate section with:
  - Lens, priority, effort.
  - Explicit file paths under `path`.
  - A specific, actionable change description.
  - Optional Increment Hint, if helpful.
- ADR candidates, if any, are proposed separately and not embedded in the improve file.
- The document contains no references to prompts, LLMs, or assistants.
- The **number and ambition** of improvements feel appropriate for the project’s `constitution-mode` (for example, a handful of quick, high-value refactors in `lite`).

If any of these items are missing or unclear, revise the plan or ask the user focused clarifying questions before treating the artifact as complete.

Finally, present the improve document as the complete content for a new file at:

- `docs/improve/<YYYY-MM-DD>-improve.md`

where `<YYYY-MM-DD>` is the date of this analysis (ISO 8601). Do **not** assume you are physically writing the file; generate the complete document content so that the host environment or user can create or update the file.