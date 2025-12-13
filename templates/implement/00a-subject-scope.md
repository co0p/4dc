## Subject & Scope

**Subject**: The `path` argument points at an **increment folder** (for example: `.../increments/<slug>` or `.../docs/increments/<slug>`). This folder contains `increment.md` and `design.md`. The **subject** of this prompt is:

- The increment defined in that folder.
- The approved design in `design.md`.
- The project codebase and constitution at the project root (the parent of the increment folder).

**Scope Constraints**:

- You MUST read:
  - `increment.md` and `design.md` in the increment folder.
  - `CONSTITUTION.md`, ADRs, code, and tests under the project root.
- You MUST treat `design.md` as **authoritative** for this increment's technical approach.
- You MUST NOT redesign components, contracts, or data flows.
- You MUST treat the project root as the **primary context**.
- You MUST NOT treat parent directories, sibling projects, or other repositories as your subject.
