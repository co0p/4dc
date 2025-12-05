 
# Lenses for Refactoring and Codebase Improvement

Use these lenses to analyze the code, tests, and docs **under the subject root `path`**. All observations and recommendations must be grounded in concrete evidence from files under `path`.

## Naming & Clarity (Fowler, Martin, Metz)

- Rename variables, functions, and classes for clarity and intent.
- Replace magic numbers/strings with named constants.
- Standardize naming conventions across the codebase.
- Inline trivial variables.
- Use intention-revealing names (Fowler).
- Avoid ambiguous or overloaded names.
- Refactor names to reflect domain language (Evans).

## Modularity & Separation (Fowler, Evans, Wirfs-Brock)

- Extract small functions/methods.
- Split large functions/classes into smaller units.
- Move related code into cohesive modules.
- Redesign module boundaries for separation of concerns.
- Introduce helper/util modules for shared logic.
- Apply Single Responsibility Principle (Martin).
- Decouple UI from business logic.
- Use dependency inversion for module boundaries.
- Refactor to reduce coupling and increase cohesion.

## Architecture & Patterns (Fowler, Evans, Beck, Martin)

- Replace complex conditionals with polymorphism or strategy.
- Replace ad-hoc data flow with clear, documented architecture (event-driven, layered, DDD).
- Align divergent implementations (error handling, validation, state management) via ADRs and shared patterns.
- Refactor error handling for consistency.
- Remove or refactor workaround/hack code.
- Introduce design patterns where appropriate (Strategy, Observer, Factory).
- Refactor for testability (Beck, Feathers).
- Apply Domain-Driven Design principles (Evans).
- Document architectural decisions and rationale.

## Testing & Reliability (Beck, Feathers)

- Add or improve automated tests for critical paths.
- Refactor code to be more testable (dependency injection, isolation).
- Remove dead code and unused imports.
- Increase test coverage for edge cases.
- Use test doubles and mocks for isolation.
- Refactor legacy code to enable testing (Feathers).
- Apply Test-Driven Development (Beck) where appropriate.
- Automate regression testing.

## Duplication & Simplicity (Fowler, Thomas & Hunt)

- Consolidate duplicate code.
- Simplify conditional logic.
- Improve code formatting and indentation.
- Remove unnecessary abstractions.
- Eliminate speculative generality (Fowler).
- Refactor to DRY (Don't Repeat Yourself).
- Prefer simple, readable solutions over cleverness.

## Documentation & Communication (Martin, Thomas & Hunt)

- Add missing comments for non-obvious logic.
- Update documentation to match code.
- Document key decisions, trade-offs, and open questions.
- Write ADRs for significant architectural changes.
- Maintain up-to-date README and onboarding docs under `path`.
- Use diagrams to clarify architecture and data flow.
- Document public APIs and interfaces.

---

These lenses are inspired by leading industry experts: Martin Fowler, Kent Beck, Michael Feathers, Robert C. Martin, Rebecca Wirfs-Brock, Eric Evans, Sandi Metz, Dave Thomas & Andy Hunt. Use them to ground your observations and recommendations in well-established practice.