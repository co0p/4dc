## Output Structure and Examples

The generated implementation plan MUST be written to a file named `implement.md` in the current increment folder (for example: `.../increments/<slug>/implement.md`).

The plan MUST follow this structure:

1. Implement Title

- First-level heading in the form:
  - `Implement: <Short, Descriptive Title>`
- Usually corresponds to the increment’s goal or the design title.

2. Context (Very Short)

- 2–4 bullet points summarizing:
  - The increment’s goal and key non-goals.
  - The main design approach (one or two sentences).
- Links to:
  - `increment.md`
  - `design.md`
  - `CONSTITUTION.md` (by filename only; no absolute paths).

3. Workstreams

- A short list of named workstreams, for example:

  ```markdown
  ## 1. Workstreams

  - Workstream A – Domain/Backend changes
  - Workstream B – Tray/UI updates
  - Workstream C – Tests and fixtures
  - Workstream D – Observability and metrics
  ```

4. Steps (XP-style Tasks)

Each step is a small, concrete work item:

```markdown
## 2. Steps

### Step 1: [Short actionable task title]
- **Workstream:** [A/B/C/D]
- **Based on Design:** [Reference to design section/decision, e.g. "Design §5: Architecture and Boundaries – Tray label updates"]
- **Files:** `path/to/file.go`, `another/path/file_test.go`
- **Actions:**
  - [Concrete code-level action 1]
  - [Concrete code-level action 2]
- **Tests:**
  - [Tests to add/update]
  - [CI commands to run, e.g. `go test ./...`]

### Step 2: [Short actionable task title]
- **Workstream:** [...]
- **Based on Design:** [...]
- **Files:** [...]
- **Actions:**
  - [...]
- **Tests:**
  - [...]
```

5. Notes on Rollout and Validation

- Optional, but recommended:

  ```markdown
  ## 3. Rollout & Validation Notes

  - Suggested grouping into PRs:
    - PR 1: Steps 1–3 (backend domain changes + tests)
    - PR 2: Steps 4–5 (tray UI + tests)
  - Suggested validation checkpoints:
    - After Step 3: [what to manually or automatically verify]
    - After Step 5: [what to manually or automatically verify]
  ```

### Acceptance (for the implement.md artifact)

The implementation plan is “good enough” when:

- **Traceability**
  - Every step references at least one part of `design.md`.
  - It is clear how executing all steps will realize the design.

- **Granularity**
  - Steps are small and concrete.
  - A developer can pick any step and understand:
    - Which files to touch.
    - What changes to make at a high level.
    - What tests to write or adjust.

- **XP-friendly**
  - Steps naturally support TDD and pairing.
  - The plan can be executed incrementally with CI, leaving the system in a working or quickly recoverable state.

- **Clarity**
  - The document follows the structure above.
  - It contains no references to prompts, LLMs, or assistants.