````markdown
## Output Structure

The generated constitution MUST be written to a file named `CONSTITUTION.md` at the project root (`path`).

It MUST follow this structure:

```markdown
# Project Constitution

constitution-mode: <lite|medium|heavy>

## 1. Purpose and Scope

[Short description of what this project is and what this constitution is for.]

## 2. Implementation & Doc Layout

[Explicit description of where key artifacts live, for example:]

- **Increment artifacts**
  - Location: `increments/<slug>/`
  - Files:
    - `increment.md`
    - `design.md`
    - `implement.md`

- **Improve artifacts**
  - Location: `docs/improve/`
  - Filename pattern: `YYYY-MM-DD-improve.md`

- **ADR artifacts**
  - Location: `docs/adr/`
  - Filename pattern: `ADR-YYYY-MM-DD-<slug>.md`

- **Other docs (optional)**
  - Architecture notes: `docs/architecture/`
  - Runbooks / ops notes: `docs/ops/`
  - [Adjust to this project’s reality.]

## 3. Design & Delivery Principles

[Short, opinionated list of principles for this project, for example:]

- **Small, safe steps** (Kent Beck)
  - We prefer many small, reversible changes over large, risky ones.
  - Increments and implement plans should reflect this.

- **Refactoring as everyday work** (Fowler, Feathers)
  - We treat refactoring as part of normal work, not a separate phase.

- **Separation of concerns & stable boundaries** (Martin)
  - Domain logic, IO, and frameworks are kept separate where practical.

[And a few more, tailored to mode and project.]

## 4. Testing, CI/CD, and Observability

[If relevant, describe expectations at a high level, for example:]

- **Testing**
  - New changes should come with automated tests (unit/integration as appropriate).
- **CI/CD**
  - All changes should run through CI before merging.
- **Observability**
  - Important behavior should be visible through logs, metrics, or similar signals.

## 5. ADR and Improve Usage

[If relevant, describe how ADRs and Improve docs are used:]

- **ADRs**
  - Use ADRs in `docs/adr/` for significant architectural decisions.
- **Improve**
  - Use Improve docs in `docs/improve/` to reflect on system health and propose refactors.
```

### Acceptance (for `CONSTITUTION.md`)

The constitution is “good enough” when:

- **Clarity**
  - `constitution-mode` is clearly stated and matches the project’s reality.
  - Implementation & Doc Layout is explicit and correct for this repo.
  - Principles are few, concrete, and understandable.

- **Actionability**
  - It is obvious how Increment, Design, Implement, and Improve should behave under this constitution.
  - The document can be read end-to-end in a few minutes.

- **Focus**
  - The document avoids unnecessary theory and meta-commentary.
  - It contains no references to prompts, LLMs, or assistants.
````