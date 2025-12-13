## Subject & Scope

**Subject**: The `path` argument points at a **project or subproject root** (for example: `.` or `examples/pomodoro`). The **subject** of this prompt is:

- The project or subproject under this path.
- Its `CONSTITUTION.md`, README, and existing artifacts (increments, designs, implements, improves, ADRs).
- Its codebase and tests.

**Scope Constraints**:

- You MUST treat this path as the **subject project root**.
- You MUST read context only from this directory and its subdirectories.
- You MUST NOT rely on content from parent directories, sibling projects, or other repositories as primary context.
- You MAY reference broader practices or frameworks, but your increment MUST be grounded in files and constraints under this path.
