## Process

Follow this process to produce an `implement.md` that is:

- Strictly grounded in `increment.md` and `design.md`.
- Organized as small, testable work items.
- Usable as a backlog for the increment.

The `path` argument for this prompt points at an **increment folder** (for example: `.../increments/<slug>`), containing `increment.md` and `design.md`. The **project codebase** and other documents live under the project root.

### Phase 1 – Gather and Summarize (STOP 1)

1. Gather Context

   - Read:
     - `CONSTITUTION.md` — values, principles, guardrails.
     - `increment.md` — context, goal, non-goals, risks, acceptance criteria.
     - `design.md` — technical approach, components, contracts, test/CI/observability plans.
   - Optionally review:
     - Relevant ADRs.
     - Recent `improve.md` documents affecting this area.
   - Inspect existing **code and tests** that `design.md` references, under the project root.

2. Restate Increment and Design (Briefly)

   - In a few sentences, restate:
     - The increment’s main goal and non-goals.
     - The high-level design approach (components, flows, key contracts).
   - Call out any **critical constraints** from constitution or design that affect implementation (e.g., must not break API X, must preserve behavior Y).

3. Summarize Implementation Perspective → STOP 1

   - Present a concise summary that covers:
     - What must be true for implementation to be considered done (mapped to increment/design).
     - The main parts of the system that will see code changes.
     - Any obvious implementation risks or tricky areas.
   - Clearly label this as **STOP 1**.
   - Ask the user to:
     - Confirm whether this understanding is correct.
     - Add any constraints or preferences that affect how work should be split (e.g., “prefer backend-first”, “front-end work after API”).

   **Do not** start listing detailed work items until the user has responded to STOP 1.

### Phase 2 – Identify Workstreams and Steps (STOP 2)

4. Identify Workstreams

   - Group the work logically into 2–5 **workstreams**, for example:
     - “Domain/Backend changes”
     - “UI/Tray updates”
     - “Tests & fixtures”
     - “Observability/metrics”
   - For each workstream:
     - Name the relevant components and files.
     - Note which sections of `design.md` they relate to.

5. Propose Work Items (High-Level)

   - For each workstream, propose a list of **small, testable work items**:
     - Each item should have:
       - A short, actionable title.
       - A reference to the relevant `design.md` section/decision.
       - Target files/modules.
       - A brief description of the intended change.
       - A testing/verification angle.
   - Ensure each work item is:
     - Scoped to be done in a focused session where practical.
     - Independently valuable or at least leaves the system in a coherent state.

6. Order the Work Items

   - Propose an **execution order** across workstreams, considering:
     - Risk reduction (tests and safe abstractions early).
     - Dependencies (infrastructure before UI, etc.).
     - Opportunities for partial value (early visible wins).
   - This ordering will become the backbone of the final `implement.md`.

7. Present Workstreams and Draft Steps → STOP 2

   - Present:
     - The list of workstreams.
     - The proposed work items within each workstream.
     - The suggested overall order (a simple numbered list or phased grouping).
   - Clearly label this as **STOP 2**.
   - Ask the user explicitly:
     - Whether the grouping and order match their expectations.
     - Whether any workstreams or items should be added, removed, merged, or re-prioritized.

   **Do not** generate the final `implement.md` until the user has approved this outline.

### Phase 3 – Write the Implementation Plan After YES

8. Produce the Final `implement.md` (After STOP 2 Approval)

   - Only after the user gives a clear affirmative response at STOP 2:
     - Generate `implement.md` that follows the output structure (see output template).
     - Implement the agreed outline, with any adjustments from user feedback.

   - While writing:
     - Do **not** introduce new architectural concepts or redesign decisions.
     - Do **not** restate the full design; refer to it in a focused way (per-step references).
     - Do **not** mention prompts, LLMs, or this process.

### Phase 4 – Final Check

9. Validate the Plan

   - Ensure:
     - Each step references:
       - A design decision/section.
       - Target files/modules.
       - Tests to add/update/run.
     - The steps are small, concrete, and can be executed independently.
     - The plan can be reasonably executed with XP practices (TDD, pairing, CI).

   - If anything is missing or unclear:
     - Add or adjust steps.
     - Optionally ask the user focused clarifying questions.

Once validated, present the final `implement.md` content. The host environment or user will save it in the increment folder.