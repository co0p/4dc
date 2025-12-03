# Process (Increment)

## Operating Rules and Guardrails

- Human-first interaction.
- Align with `CONSTITUTION.md`; if a proposed increment violates the constitution, flag the conflict and propose alternatives.
- Keep increments small, testable, and observable. Prefer one clear increment per run.
- Follow the Task section’s cycle **exactly**.
- Respect STOP gates:
  - At the clarifying-questions step, do NOT proceed until questions are answered or the user explicitly waives them.
  - At the “Suggest Increment Structure” step, present a concise plan and obtain an explicit **yes/no** before generating and saving the increment document.
- Do NOT offer additional actions in the increment document itself (no “If you’d like, I can also…”, no proposals to create workflows or other files inside the increment text).
- Final increments MUST follow the Increment Output Structure exactly; no extra top-level sections unless explicitly added to the template.
- Date format: `YYYY-MM-DD` for any dates.
- Treat the **target project root/scope** as the subject of the increment:
  - Use only context inside that scope for the project description and constraints.
  - Treat content outside that scope as tooling/background only.

The detailed steps to follow are:

1. Verify Prerequisites
2. Receive Initial Prompt
3. Analyze Constitution & Context
4. Ask Clarifying Questions (STOP)
5. Suggest Increment Structure (STOP)
6. Generate Increment
7. Save Increment
8. Final Validation

For each step, follow the detailed instructions from the Task section, ensuring you do not skip or reorder steps, and that STOP gates are respected.