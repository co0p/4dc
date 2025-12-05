# LLM–Human Interaction: Improve Step Findings Style Reference

The improve step is **findings-driven**, not questionnaire-driven. You analyze the codebase under `path`, then present factual observations organized by lenses.

## Interaction Flow

1. **Initial Acknowledgment**

   - Briefly restate the subject: the project rooted at `path`.
   - Clarify that you will analyze files under `path` using the defined lenses.

2. **STOP 1 – Context and Assessment Summary**

   After your initial review of project docs and key code areas under `path`:

   - Present:
     - A short summary of the project’s purpose and main capabilities.
     - A concise initial assessment (constitution/design alignment, quality, key risks).
   - Then STOP and ask:
     - Whether anything is misunderstood about the project context.
     - Whether there are areas under `path` that the user wants you to **prioritize** (without skipping your own analysis).

   **Wait for user feedback before proceeding.**

3. **Findings Presentation**

   Present each finding as a factual observation with evidence from the codebase:

   ### Finding 1: [Title]
   - **Lens:** [Naming/Modularity/Architecture/Testing/Duplication/Documentation]
   - **Observation:** [Factual description of what was found]
   - **Evidence:** [Specific file/line references from the codebase under `path`]
   - **Recommendation:** [Suggested action based on industry best practices]

   ### Finding 2: [Title]
   - **Lens:** [...]
   - **Observation:** [...]
   - **Evidence:** [...]
   - **Recommendation:** [...]

   Repeat as needed for all important findings.

4. **STOP 2 – Plan Outline**

   - Summarize:
     - The main groups of findings by lens.
     - The outline for the **Improvements** section of the improve document (which Improvement proposals you plan to include, and at what priority).
   - Ask the user to confirm:
     - Whether this outline reflects their priorities.
     - Whether you should proceed to generate the final improve document.

   **Do not generate the final improve document until the user explicitly says “yes” (or equivalent approval).**

5. **After STOP 2**

   - If the user asks for changes, update the outline and ask for confirmation again.
   - Only after an explicit “yes” generate the final `improve.md` artifact content.

---

Always present findings with facts and evidence. Do not ask the user what to look for or what to improve; let the lenses and your analysis of files under `path` guide recommendations.