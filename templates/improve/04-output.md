# Improve Output Format

The improve output is a human-readable artifact, not a conversation transcript. It must be concise and parsable by the `/increment` prompt.

- The subject is the project rooted at `path`.
- The artifact must **not** mention prompts, LLMs, or assistants.
- The artifact must follow the structure below exactly (section headings and order).
- Each “Improvement” is a **proposal for future work** that the team may or may not pick up as a new increment.
- When `CONSTITUTION.md` defines a `constitution-mode`, ensure the number and size of Improvements are appropriate for that mode (fewer, smaller proposals in `lite`; more comprehensive coverage allowed in `heavy`).

## Output Schema: docs/improve/YYYY-MM-DD-improve.md

The improve artifact is stored by date under:

- Directory: `docs/improve/`
- File name: `<YYYY-MM-DD>-improve.md` (the date when the analysis was performed).

The content of each improve file must follow this structure.

Top-level structure:

- Heading: `# Improve: [Short Title For This Improvement Cycle]`

Sections in order:

1. `## 1. Assessment`
2. `## 2. Lessons`
3. `## 3. Improvements`

### 1. Assessment

Content:

- `**Constitution Alignment:**`  
  For each relevant principle from `CONSTITUTION.md` under `path`, give a 1–5 star rating and a brief rationale, for example:  
  - `Principle: Browser-native UX – ★★★☆☆ (mostly followed, but some direct DOM manipulation bypasses shared helper)`  
  - `Principle: Observable workflows – ★★☆☆☆ (limited logging, no tracing on critical paths)`
- `**Design Alignment:**` brief evaluation of how the current implementation aligns with key design goals or architecture.
- `**Quality:**` brief evaluation of code/test quality, readability, maintainability.
- `**Risks:**` bullet list of key risks (technical debt, potential bugs, performance, security, operability).

### 2. Lessons

Content:

- `**Worked Well:**` list of things that went well (patterns, practices, areas of code).
- `**To Improve:**` list of pain points or problems to address.
- `**Emerging Patterns:**` list of patterns that are appearing across the codebase (good or bad).

### 3. Improvements

Each improvement is a separate section with the following structure (shown here as plain text, not a fenced code block):

- Heading: `#### Improvement N: [Title]`  
  - `**Lens:**` one of `Naming`, `Modularity`, `Architecture`, `Testing`, `Duplication`, `Documentation` (or a closely related variant).  
  - `**Priority:**` `H` / `M` / `L`.  
  - `**Effort:**` rough time estimate, e.g. `30 min`, `2–3 h`, `1–2 days`.  
  - `**Files:**` comma-separated list of concrete paths under `path`, e.g. ``path/to/file.ext``, ``another/path/file.ext``.  
  - `**Change:**` specific change description, clear enough that someone can implement it without re-interpreting intent, and small enough to be done as one or a few safe refactorings.  
  - `**Increment Hint (optional):**` one-sentence suggested increment title or description that could be used directly as the short description for a future increment.

Example (conceptual, not for copy-paste):

- `#### Improvement 1: Consolidate duplicate catalog serializers`  
  - `**Lens:** Duplication`  
  - `**Priority:** M`  
  - `**Effort:** 1–2 h`  
  - `**Files:** examples/shareit/backend/routes/catalog.js, examples/shareit/backend/routes/admin-catalog.js`  
  - `**Change:** Extract a shared serializer function for catalog item JSON shape and use it in both routes to avoid divergence in field names and formats.`  
  - `**Increment Hint (optional):** Refactor catalog serializers into a shared helper to remove duplication across routes.`

Notes:

- Each improvement is a distinct `#### Improvement N: ...` section.
- The **Files** list must use concrete, existing paths under the subject root.
- The **Change** description must be specific enough to implement without re-interpreting the intent, and should lend itself to implementation in small, safe steps rather than a giant rewrite.
- “Increment Hint” is optional and provides a convenient starting point for future increments:
  - Phrase it so it can be dropped directly into the Increment prompt as the short increment description.
- ADRs are created as separate, independent artifacts using the ADR Output Template. They are **not** part of the improve file content.

---

## Rating Scale (for Constitution Alignment)

When assigning stars to constitution principles:

- `★☆☆☆☆` – Poor alignment; the principle is frequently violated or effectively absent.
- `★★☆☆☆` – Weak alignment; some places follow the principle, many do not.
- `★★★☆☆` – Mixed alignment; the principle is visible but inconsistently applied.
- `★★★★☆` – Good alignment; the principle is followed in most relevant areas.
- `★★★★★` – Strong alignment; the principle is consistently and clearly followed.

For each rated principle:

- Include **one short rationale** (1–2 sentences).
- Reference concrete evidence under `path` where helpful (e.g. key modules, patterns, or omissions).

---

## Acceptance (for the improve artifact)

The improve document is “good enough” when:

- **Scope**
  - All observations and file paths refer to content under `path`.
  - No changes are proposed outside the subject project root.

- **Alignment**
  - Assessment clearly references project Constitutions/Designs where they exist under `path`.
  - Constitution Alignment includes per-principle star ratings with concise rationales.
  - Lessons and improvements are grounded in concrete evidence from the codebase.

- **Clarity**
  - All sections (Assessment, Lessons, Improvements) are present and non-empty.
  - Each Improvement includes lens, priority, effort, files, and a precise change description.
  - Each Improvement reads as a **proposal** that could become an increment, not as an instruction that has already been executed.
  - The document contains no references to prompts, LLMs, or assistants.

- **Mode-awareness**
  - The **number and ambition** of improvements feel appropriate for the project’s `constitution-mode` (for example, a handful of quick, high-value refactors in `lite`).