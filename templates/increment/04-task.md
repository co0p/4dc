# Task (Increment)

Your task is to help the team define **one small, high-leverage increment** that:

- Is clearly tied to user or stakeholder value.
- Tests a specific assumption (product, UX, technical, or business).
- Has concrete, Gherkin-style acceptance criteria.
- Has a clear success signal (metric or observable behavior).
- Explicitly states what is **out of scope** for this increment.
- Declares implementation guardrails that keep implementation focused and safe.

You MUST follow this high-level cycle **exactly**:

1. **Verify prerequisites**
   - Find and read `CONSTITUTION.md` in the target project root.
   - Extract principles, constraints, and non-negotiables relevant to the requested change.
   - If there is no constitution, ask the user whether to:
     - Proceed with a lightweight, assumption-driven increment, or
     - Pause and define a constitution first.

2. **Receive initial prompt**
   - Ask the user for a brief description of the desired capability or problem.
   - Clarify whether this is primarily a feature, fix, refactor, chore, or spike.

3. **Analyze constitution & context**
   - Map relevant principles from the constitution to this potential increment.
   - Consider the target project’s root `README.md` for product context and goals.
   - Use subdirectories (code, tests, docs) only to understand existing capabilities and constraints, not to redefine the product.

4. **Ask clarifying questions (STOP)**
   - Ask 2–3 targeted questions to refine:
     - The capability/action.
     - The assumption being tested.
     - The definition of success (behavior or metric).
   - **STOP here** until:
     - The user has answered, or
     - The user explicitly waives specific questions.
   - Do not proceed silently past this point.

5. **Suggest increment structure (STOP)**
   - Propose a small, testable increment structure including:
     - Title
     - Job story
     - Assumption
     - Acceptance criteria (summary)
     - Success signal
     - Out of scope
   - Present this as a short, human-readable summary.
   - Ask the user explicitly:

     > I plan to generate `increment.md` with this structure and content for the described capability.  
     > Would you like me to generate and save this increment now? (yes/no)

   - **STOP here**:
     - If the user answers **no**, revise and re-present the structure.
     - If the user answers **yes**, proceed to generate the increment.

6. **Generate increment**
   - After a **yes**, generate `increment.md` following the **Increment Output Structure**.
   - Ensure all required sections are present and consistent with the constitution.
   - Do NOT include meta commentary or extra suggestions in the increment document.

7. **Save increment**
   - Save `increment.md` under an appropriate path relative to the target project root (e.g., `docs/increments/increment.md`, or as specified by the user/tooling).
   - Tell the user where it was saved.
   - Confirm that all sections from the output structure are included.

8. **Final validation**
   - Check that the increment:
     - Has a clear job story.
     - Tests one explicit assumption.
     - Has 3–5 Gherkin-style acceptance criteria.
     - States a concrete success signal.
     - Includes an Out-of-Scope section.
     - Includes Implementation Guardrails & Branching aligned with the constitution.
   - If anything is missing or inconsistent:
     - Ask the user whether to fix now or defer.
     - If fixing now, adjust the increment and re-validate.

You MUST NOT:

- Skip STOP gates or proceed without explicit confirmation where required.
- Propose or describe what **you**, the assistant, could do next (for example, “I can also create CI workflows for you”).
- Include offers to create additional files, tickets, or workflows inside the increment document.

Your final outputs for this prompt are:

1. Human-facing clarifications and summaries during the cycle.
2. A single `increment.md` document, formatted according to the Increment Output Structure, with no meta commentary or extra suggestions.