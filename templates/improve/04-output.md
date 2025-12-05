# Improve Output Format

The improve output is a human-readable artifact, not a conversation transcript. It must be concise and parsable by the `/increment` prompt.

- The subject is the project rooted at `path`.
- The artifact must **not** mention prompts, LLMs, or assistants.
- The artifact must follow the structure below exactly (section headings and order).
- Each “Improvement” is a **proposal for future work** that the team may or may not pick up as a new increment.

## Output Schema: docs/improve/YYYY-MM-DD-improve.md

The improve artifact is stored by date under:

- Directory: `docs/improve/`
- File name: `<YYYY-MM-DD>-improve.md` (the date when the analysis was performed).

The content of each improve file must follow this structure:

```markdown
# Improve: [Short Title For This Improvement Cycle]

## 1. Assessment
- **Constitution Alignment:** [Brief evaluation]
- **Design Alignment:** [Brief evaluation]
- **Quality:** [Brief evaluation]
- **Risks:** [List]

## 2. Lessons
- **Worked Well:** [List]
- **To Improve:** [List]
- **Emerging Patterns:** [List]

## 3. Improvements

#### Improvement 1: [Title]
- **Lens:** [Naming/Modularity/Architecture/Testing/Duplication/Documentation]
- **Priority:** [H/M/L]
- **Effort:** [X min]
- **Files:** `path/to/file.ext`, `another/path/file.ext`
- **Change:** [Specific change description]
- **Increment Hint (optional):** [Suggested increment title or capability this could become]

#### Improvement 2: [Title]
- **Lens:** [...]
- **Priority:** [...]
- **Effort:** [...]
- **Files:** `path/to/file.ext`
- **Change:** [...]
- **Increment Hint (optional):** [...]
```

Notes:

- Each improvement is a separate `#### Improvement N: ...` section with explicit file references under `path`.
- The **Files** list must use concrete, existing paths under the subject root.
- The **Change** description must be specific enough to implement without re-interpreting the intent.
- “Increment Hint” is optional and provides a convenient starting point for future increments.
- ADRs are created as separate, independent artifacts using the ADR Output Template. They are **not** part of the improve file content.

### Acceptance (for the improve artifact)

The improve document is “good enough” when:

- **Scope**
  - All observations and file paths refer to content under `path`.
  - No changes are proposed outside the subject project root.

- **Alignment**
  - Assessment clearly references project Constitutions/Designs where they exist under `path`.
  - Lessons and improvements are grounded in concrete evidence from the codebase.

- **Clarity**
  - All sections (Assessment, Lessons, Improvements) are present and non-empty.
  - Each Improvement includes lens, priority, effort, files, and a precise change description.
  - Each Improvement reads as a **proposal** that could become an increment, not as an instruction that has already been executed.
  - The document contains no references to prompts, LLMs, or assistants.