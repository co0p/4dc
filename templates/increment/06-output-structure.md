## Output Structure

The generated increment definition MUST follow this structure and use these headings in this order.

1. Title

- Format: `Increment: <Short, Product-Friendly Title>` as the first-level heading.
- The title should be:
  - Short and descriptive.
  - Understandable by product, engineering, and stakeholders.

2. Context

- Explain the current situation or problem from a user or business perspective.
- Include:
  - Existing behavior or limitations that motivate this increment.
  - Any important background such as related work, earlier attempts, or relevant documents within the project.
  - Key constraints or assumptions (for example: time, scope, risk tolerance, regulatory limits).
- This section should give enough background that someone new to the increment understands why it matters.

3. Goal

- Describe the outcome or hypothesis:
  - What will be different for users, customers, or the business after this increment?
- Clarify the scope:
  - What this increment will do.
- Clarify the non-goals:
  - What this increment explicitly will not address.
- Briefly explain why this is a good, small increment:
  - It is small, coherent, and can be evaluated in a reasonable time.

4. Tasks

- Provide a product-level checklist of tasks. For each task, include:
  - `Task:` An outcome-level description of what should be true when the task is complete.
  - `User/Stakeholder Impact:` Who is affected and how their experience or workflow changes.
  - `Acceptance Clues:` Observable signs that the task is complete from a WHAT perspective (for example: a behavior is visible, a piece of information can be seen, or a simple check passes).
- Tasks MUST describe WHAT should be true, not the technical HOW.
- Avoid references to specific files, functions, branches, components, or implementation steps.

5. Risks and Assumptions

- List known risks, such as:
  - Potential user impact.
  - Product fit concerns.
  - Rollout or timing concerns.
- List key assumptions that, if wrong, could change the plan.
- Optionally mention high-level mitigations, still in outcome language (for example: “If adoption is low, we may follow up with user interviews”).

6. Success Criteria and Observability

- Describe how you will know this increment is successful:
  - Changes in metrics, events, or user behavior.
  - Evidence to look at after release.
- Describe what will be observed after release:
  - Which dashboards, logs, or reports will be checked.
  - Any simple checks in staging or production to confirm behavior.
- Keep this at the level of WHAT is observed, not HOW it is instrumented.

7. Process Notes

- Add high-level notes about how this increment should move through the workflow, without prescribing technical steps.
- Examples:
  - It should be implemented via small, safe changes over time.
  - It should go through the normal build, test, and release process used by the project.
  - It should be rolled out in a way that allows quick recovery if needed.
- Do not include concrete implementation instructions, code changes, or deployment commands.

8. Follow-up Increments (Optional)

- Briefly describe potential future increments that:
  - Extend this outcome.
  - Address related but out-of-scope work.
  - Further improve performance, reliability, or user experience.
- Each follow-up should be described as a possible future increment, not as part of the current one.

---

The final increment definition MUST:

- Use the sections above in this order.
- Fill each section with project-specific content based on the scoped project and the increment description.
- Avoid references to prompts, LLMs, or assistants.
- Keep Tasks focused on WHAT, leaving the HOW to later phases and artifacts.
- If any section starts to describe internal components, data models, services, functions, classes, files, or specific modules, rewrite it to focus on observable behavior, outcomes, and evidence instead.