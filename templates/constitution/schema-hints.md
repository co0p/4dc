# JSON Schema Hints (Internal Only)

The following hints help you formalize the JSON blocks for validation tools.

- ClarificationRequest (Step 3)
  - `step`: enum ["questions"]
  - `questions[]`: objects with `id` (integer), `pillar` (string from six canonical names), `question` (string), `options[]` (array of objects with `key` in [A..Z, "X", "_"], `label` string)
  - `increments_location_recommendation`: string path (default `docs/increments/<increment-folder>/increment.md`)
  - `instructions`: human-readable string

- PrinciplesProposal (Step 4)
  - `step`: enum ["proposal"]
  - `principles[]`: objects with `id` (integer), `name` (string), `pillar` (enum of six names), `statement` (string), `rationale` (string), `implications[]` (array of strings)
  - `coverage[]`: array of distinct pillars covered

- ConstitutionSummary (Save Step)
  - `step`: enum ["summary"]
  - `sections_included[]`: array of strings from the Output sections list
  - `pillars_covered[]`: array of distinct pillars
  - `counts.principles`: integer (3–5)
  - `counts.pillars`: integer (>=3)
  - `paths.constitution`: string path (defaults `CONSTITUTION.md`)
  - `paths.increments`: string path (e.g., `docs/increments/<increment-folder>/increment.md`)
  - `last_updated`: string date (YYYY-MM-DD)

Tip: Keep JSON blocks minimal and deterministic; validators should ignore extra fields. Do not display these JSON blocks to users unless explicitly requested.
