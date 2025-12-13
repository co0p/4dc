## Subject & Scope

**Subject**: The `path` argument points at an **increment folder** (for example: `.../increments/<slug>` or `.../docs/increments/<slug>`). This folder contains `increment.md`. The **subject** of this prompt is:

- The increment defined in that folder.
- The project codebase and constitution at the project root (the parent of the increment folder).

**Scope Constraints**:

- You MUST read:
  - `increment.md` in the increment folder.
  - `CONSTITUTION.md`, ADRs, code, and tests under the project root.
- You MUST treat the project root as the **primary context**.
- You MAY reference broader practices or frameworks, but your design MUST be grounded in files and architecture under the project root.
- You MUST NOT treat parent directories, sibling projects, or other repositories as your subject.
