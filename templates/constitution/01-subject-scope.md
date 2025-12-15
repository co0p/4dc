## Subject & Scope

The `path` argument points at a project or subproject root. The constitution you generate applies to:

- The project or subproject under `path`.
- Its documentation and artifacts maintained in this repository.
- Its codebase and tests within this scope.

Scope rules:

- Treat `path` as the subject root; reason only about files and directories under it.
- Do not prescribe changes outside this scope.
- Keep principles human-first, observable, and pragmatically testable for this project.
