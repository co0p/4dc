## Goal

Turn the approved **design** for this increment into an **ordered implementation plan** that:

- Respects the **Project Constitution** and the agreed `design.md`.
- Breaks work into **small, XP-style tasks** that can be executed via TDD and pairing.
- Identifies **files/modules** to touch for each task.
- Specifies **tests** and checks for each task.
- Can be executed incrementally and safely, while keeping the system in a working state.

The implementation plan MUST:

1. Stay Within Design and Increment Scope

   - Treat `design.md` as the **authoritative technical plan** for this increment.
   - Treat `increment.md` as the **scope guardrail** for product outcomes.
   - Do **not** redesign components, contracts, or data flows here.
   - If the design appears problematic, call it out as a **risk** or **follow-up increment**, not as a change to make in this plan.

2. Produce Small, Testable Work Items

   - Each step should:
     - Be small enough to execute in a single focused session.
     - Identify clear **code locations** (files/modules).
     - Include a **testing/verification** angle (tests to add/update/run).
   - Steps should be written so that:
     - A developer can pick one step and complete it independently.
     - The system stays in a releasable or quickly-fixable state after each step where practical.

3. Preserve Traceability to the Design

   - Each work item should clearly reference:
     - The relevant section(s) or decision(s) in `design.md` that it implements.
   - The plan should make it obvious:
     - How the design will be realized incrementally.
     - Which parts of the design are covered by which tasks.

4. Support CI/CD and Review

   - Steps should naturally map to:
     - Small PRs or commits.
     - Clear CI expectations (which checks should pass).
   - Where helpful, note:
     - Groupings of tasks that can form a single PR.
     - Suggested points for partial rollouts or feature-flag flips.