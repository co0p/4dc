# Output Structure (Increment)

You MUST:

- Output only the increment specification document in Markdown, using the structure defined in this file.
- NOT include any meta commentary about what you (the assistant) could do next (for example, "If you'd like, I can also add...", "Next, I can create...", "I can generate a workflow").
- NOT include suggestions for additional files, CI workflows, or other automation tasks inside the increment. Those may be implied by principles, but not offered as actions by you.

Return the result as **Markdown** with the following structure:

```markdown
# [Increment Title]

## Job Story
**When** [situation]  
**I want to** [action]  
**So I can** [outcome]

**Assumption Being Tested:** [Specific hypothesis for this increment]

## Acceptance Criteria
- **Given** [precondition]  
  **When** [action]  
  **Then** [observable outcome]
- **Given** [error condition]  
  **When** [action]  
  **Then** [error handling outcome]
<!-- Additional scenarios as needed, keeping total criteria typically between 3–5. -->

## Success Signal
[How we know this increment works – a metric or concrete observation]

## Out of Scope
- [What this increment does NOT include to keep focus tight]

## Implementation Guardrails & Branching
- Feature branch: `feature/<increment-slug>`; no direct commits to default branch.
- Planned Files Summary: confirm the planned file changes before coding (STOP gate).
- DRIFT ALERT: STOP on out-of-scope changes; propose a minimal update or a follow-up increment.
- Verification: map tasks to acceptance criteria with tests or explicit manual checks.
- Stabilization: docs and hygiene (e.g., `.gitignore`, reproducible builds) completed on the feature branch before merge.