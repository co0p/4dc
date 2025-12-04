## Tasks

Describe the **work items at the “WHAT” level**, not the technical “HOW”.

Tasks should express **what needs to be true for this increment to be considered complete**, in terms that are clear to product, engineering, QA, and stakeholders.

The LLM MUST:

1. Express Tasks as Outcome-Oriented Work Items

   - Each task should describe **what must change or exist** from a product, user, or business point of view, for example:
     - “The login screen offers a ‘Forgot password?’ entry point.”
     - “Usage of the export feature is visible in analytics with basic volume metrics.”
     - “Support has a documented way to explain the new behavior to customers.”
   - Avoid:
     - Low-level implementation details (for example: database tables, functions, framework choices).
     - Step-by-step instructions or technical plans.

2. Use a Simple, Consistent Structure

   Each task SHOULD include:

   - Task: A short, action-oriented description of what will be true or available.
   - User/Stakeholder Impact: Whose experience or workflow is affected and how.
   - Acceptance Clues: How we might recognize that the task is “done” (from the WHAT perspective – for example: behavior visible, outcome demonstrable – without prescribing tests or code).

   Example format (shown as plain text to avoid ambiguity):

   - `Task:` Short, outcome-level description  
     - `User/Stakeholder Impact:` …  
     - `Acceptance Clues:` …

3. Keep Tasks Independent and Small (at the WHAT level)

   - Prefer more, smaller tasks rather than one huge catch-all.
   - Each task should:
     - Represent a **distinct observable aspect of the increment**.
     - Be small enough that it can be confidently delivered within this increment.

4. Cover Non-Feature Outcomes Where Relevant

   In addition to user-facing behavior, tasks MAY include:

   - Discoverability and communication:
     - What needs to be discoverable in UI, docs, or help.
     - What internal stakeholders (support, sales, operations) need to understand.
   - Evidence and learning:
     - What must be observable so that we can later tell if the increment helped (for example: “We can see basic usage numbers for X”).
     - Any product-level follow-up checks (for example: “Review metrics after N days to confirm adoption”).

   These tasks MUST still describe **WHAT** should be true, not HOW to instrument or implement it.

5. Defer Implementation Planning

   - Do not:
     - Describe code changes.
     - Specify test frameworks or CI details.
     - Propose pull request groupings, deployment commands, or rollout mechanisms.

If you find yourself describing **internal components, data models, services, functions, classes, files, or specific modules**, you are going too deep for an increment. Rephrase the task to focus on observable behavior, outcomes, and evidence instead.

The end result should be a **product-level checklist** of outcomes that, when all true, mean this increment’s goal has been achieved.