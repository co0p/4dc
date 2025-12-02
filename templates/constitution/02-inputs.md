## Inputs

You have access to:

- This repository’s contents (including README, existing CONSTITUTION.md, ADRs, docs, code, and configuration files).
- Any answers the user provides during this interaction.

From these, you MUST build and maintain the following **internal notes**. You may show them to the user for confirmation and refinement:

1. **Team / product context** (`team_and_product_context`)
   - What this project is about.
   - Who it serves.
   - High-level domain and problem space.

2. **Team values, preferences, and constraints** (`team_values_and_constraints`)
   - How the team appears to balance speed vs safety.
   - Any explicit or implicit quality bars.
   - Constraints inferred from README, docs, and code (e.g., production-critical, experimental).

3. **Existing engineering practices / examples** (`existing_practices_and_examples`)
   - How the team currently reviews, tests, deploys, refactors, and documents.
   - Inferred from scripts, workflows, folder structure, and docs.

4. **Inspirations / reference materials** (`inspirations_and_references`)
   - Any frameworks, books, or methodologies explicitly referenced (e.g. Kent Beck, DORA, Clean Architecture).
   - Any implicit influences you can reasonably infer.

5. **Known non-negotiables** (`non_negotiables`)
   - Compliance, security, regulatory or uptime constraints, if discoverable.
   - Otherwise, questions you must ask the user to clarify.