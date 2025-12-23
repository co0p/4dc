## Output Structure

The generated increment definition MUST follow this structure and use these headings in this order.

1. Title

- First-level heading in the form:
   - `Increment: <Short, Product-Friendly Title>`
- The title should be:
   - Short and descriptive.
   - Understandable by product, engineering, and stakeholders.

2. User Story

- Capture the **primary user story** for this increment, in a form familiar to the project (for example: `As a <role>, I want <capability>, so that <benefit>.`).
- Ensure the story reflects the agreed scope of this increment, not a broader roadmap.
- If additional stories were identified but deferred, mention them only briefly as candidates for future increments (not as part of this increment).

3. Acceptance Criteria

- List the **agreed acceptance criteria** that must hold true for this user story to be considered satisfied.
- Each criterion should:
   - Describe observable behavior or evidence (WHAT happens, or what can be seen/measured).
   - Be phrased in clear, testable language (for example: “When X, then Y is visible/recorded/blocked.”).
- Where helpful, reuse structures and terminology from past increments in `docs/PRD.md` so the criteria stay consistent across the product.

4. Use Case

- Provide a **detailed use case** for the primary user story, structured along these lines:
   - **Actors:** Who is involved (users, systems, services).
   - **Preconditions:** What must already be true before the scenario starts.
   - **Main Flow:** A numbered sequence of steps that describes the typical successful path.
   - **Alternate / Exception Flows:** Short descriptions of only those alternate or error paths that are relevant to this increment.
- Keep the use case focused on behavior and interactions, not internal implementation details.

5. Context

- Explain the current situation or problem from a user or business perspective.
- Include:
   - Existing behavior or limitations that motivate this increment.
   - Any important background such as related work, earlier attempts, or relevant documents within the project.
   - Key constraints or assumptions (for example: time, scope, risk tolerance, regulatory limits).
- This section should give enough background that someone new to the increment understands **why it matters**.

6. Goal

- Describe the outcome or hypothesis:
   - What will be different for users, customers, or the business after this increment?
- Clarify the scope:
   - What this increment will do.
- Clarify the non-goals:
   - What this increment explicitly will not address.
- Briefly explain why this is a good, small increment:
   - It is small, coherent, and can be evaluated in a reasonable time.

7. Tasks

- Provide a **product-level** checklist of tasks. For each task, include:
   - `Task:` An outcome-level description of what should be true when the task is complete.
   - `User/Stakeholder Impact:` Who is affected and how their experience or workflow changes.
   - `Acceptance Clues:` Observable signs that the task is complete from a WHAT perspective (for example: a behavior is visible, a piece of information can be seen, or a simple check passes).
- Tasks MUST describe WHAT should be true, not the technical HOW.
- Avoid references to specific files, functions, branches, components, or implementation steps.

8. Risks and Assumptions

- List known risks, such as:
   - Potential user impact.
   - Product fit concerns.
   - Rollout or timing concerns.
- List key assumptions that, if wrong, could change the plan.
- Optionally mention high-level mitigations, still in outcome language (for example: “If adoption is low, we may follow up with user interviews”).

9. Success Criteria and Observability

- Describe how you will know this increment is successful:
   - Changes in metrics, events, or user behavior.
   - Evidence to look at after release.
- Describe what will be observed after release:
   - Which dashboards, logs, or reports will be checked.
   - Any simple checks in staging or production to confirm behavior.
- Keep this at the level of WHAT is observed, not HOW it is instrumented.

10. Process Notes

- Add high-level notes about how this increment should move through the workflow, without prescribing technical steps.
- Examples:
   - It should be implemented via small, safe changes over time.
   - It should go through the normal build, test, and release process used by the project.
   - It should be rolled out in a way that allows quick recovery if needed.
- Do not include concrete implementation instructions, code changes, or deployment commands.

11. Follow-up Increments (Optional)

- Briefly describe potential future increments that:
   - Extend this outcome.
   - Address related but out-of-scope work.
   - Further improve performance, reliability, or user experience.
- Each follow-up should be described as a possible future increment, not as part of the current one.

12. PRD Entry (for docs/PRD.md)

- Provide a concise **PRD entry** so this increment can be recorded in `docs/PRD.md` under the appropriate heading.
- Use a simple, stable structure such as:

   - `Increment ID:` The slug for this increment (for example: `short-and-long-break-actions`).
   - `Title:` The same short, product-friendly title used in the heading.
   - `Status:` An initial status such as `Proposed`, `In Progress`, or `Done` (the project may update this over time).
   - `Increment Folder:` The relative path to this increment’s folder from the project root (for example: `docs/increments/short-and-long-break-actions/`).
   - `User Story:` The primary user story (short form).
   - `Acceptance Criteria:` A short bullet list summarizing the agreed acceptance criteria.
   - `Use Case Summary:` A brief summary of the main flow from the detailed use case.

- This section is intended to be **copied into or synchronized with** `docs/PRD.md` and MUST:
   - Avoid references to prompts, LLMs, or assistants.
   - Use clear, product-oriented language.

---

The final increment definition MUST:

- Use the sections above in this order.
- Fill each section with project-specific content based on the scoped project and the increment description.
- Avoid references to prompts, LLMs, or assistants.
- Keep Tasks focused on WHAT, leaving the HOW to later phases and artifacts.
- If any section starts to describe internal components, data models, services, functions, classes, files, or specific modules, rewrite it to focus on observable behavior, outcomes, and evidence instead.
