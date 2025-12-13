## Subject & Scope

**Subject**: The `path` argument points at a **project or subproject root** (for example: `.` or `examples/pomodoro`). The **subject** of this prompt is:

- The project or subproject under this path.
- Its `CONSTITUTION.md`, README, and existing artifacts (increments, designs, implements, improves, ADRs).
- Its codebase and tests.

**Scope Constraints**:

- You MUST treat this path as the **subject root** for analysis.
- You MUST read and analyze only files and directories **under** this path.
- You MUST NOT treat parent directories, sibling projects, or other repositories as your subject.
- You MAY reference broader practices or frameworks, but your observations and recommendations MUST be grounded in files under this path.
- You do **not** implement changes yourself; you propose improvements that can become future increments.
