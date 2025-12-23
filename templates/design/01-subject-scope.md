## Subject & Scope

**Subject**: The `path` argument points at an **increment folder** (for example: `.../increments/<slug>` or `.../docs/increments/<slug>`). This folder contains `increment.md`. The **subject** of this prompt is:

- The increment defined in that folder.
- The project codebase and constitution at the project root (the parent of the increment folder).
- The project’s key documentation under the project root, in particular:
  - Product requirements such as `docs/PRD.md` (if present).
  - Architecture documentation such as `ARCHITECTURE.md` (if present), including any diagrams.

**Scope Constraints**:

- You MUST read:
  - `increment.md` in the increment folder.
  - `CONSTITUTION.md`, ADRs, code, and tests under the project root.
  - `docs/PRD.md` and `ARCHITECTURE.md` (or their project-specific equivalents) when they exist under the project root.
- You MUST treat the project root as the **primary context**.
- You MAY reference broader practices or frameworks, but your design MUST be grounded in files, architecture, and requirements under the project root.
- You MUST NOT treat parent directories, sibling projects, or other repositories as your subject.
