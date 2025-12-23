# Lenses for Refactoring and Codebase Improvement

Use these lenses to analyze the code, tests, and docs **under the subject root `path`**. All observations and recommendations must be grounded in concrete evidence from files under `path`.

Favor **small, safe refactorings** that can be implemented incrementally, in line with XP practices: change the code in small steps, keep tests passing, and keep the system working after each change.

When `CONSTITUTION.md` defines modes or different “levels of rigor”, scale how deep you go with each lens accordingly:

- In **lighter improvement passes** (for example, demo apps, prototypes, or low‑risk examples), prefer a small number of high‑value, quick wins.
- In **deeper improvement passes** (for example, production services or critical systems), allow more thorough, systematic improvements.

## Naming & Clarity (Fowler, Martin, Metz)

- Rename variables, functions, and classes for clarity and intent.
- Replace magic numbers/strings with named constants.
- Standardize naming conventions across the codebase.
- Inline trivial variables.
- Use intention-revealing names (Fowler).
- Avoid ambiguous or overloaded names.
- Refactor names to reflect domain language (Evans).

**Mode guidance:**

- **Lighter passes:** Focus on the most confusing or frequently touched areas; propose a handful of renames or clarifications with clear benefit.
- **Deeper passes:** Systematically align key domain terms and naming conventions across modules, especially on public APIs and core domain types.

## Modularity & Separation (Fowler, Evans, Wirfs-Brock)

- Extract small functions/methods.
- Split large functions/classes into smaller units.
- Move related code into cohesive modules.
- Redesign module boundaries for separation of concerns.
- Introduce helper/util modules for shared logic (only when duplication is clear and repeated).
- Apply Single Responsibility Principle (Martin).
- Decouple UI from business logic.
- Use dependency inversion for module boundaries.
- Refactor to reduce coupling and increase cohesion.

**Mode guidance:**

- **Lighter passes:** Identify one or two high-impact extractions (for example, a very long function or god object) that can be split safely.
- **Deeper passes:** Tackle broader boundary issues (for example, splitting a monolithic module into clearer submodules) while still keeping refactorings incremental.

## Architecture & Patterns (Fowler, Evans, Beck, Martin)

- Replace complex conditionals with polymorphism or strategy where it clearly simplifies behavior.
- Replace ad-hoc data flow with clear, documented architecture (event-driven, layered, DDD, etc.).
- Align divergent implementations (error handling, validation, state management) via ADRs and shared patterns.
- Refactor error handling for consistency.
- Remove or refactor workaround/hack code.
- Introduce design patterns where appropriate (Strategy, Observer, Factory), avoiding speculative over-engineering.
- Refactor for testability (Beck, Feathers).
- Apply Domain-Driven Design principles (Evans).
- Document architectural decisions and rationale where they affect many changes.
- Keep architectural **documentation** (for example, `ARCHITECTURE.md`) accurate and useful, including high-level and mid-level diagrams.

When an `ARCHITECTURE.md` (or similarly named architecture document) exists or is required by the constitution:

- Verify that it includes at least:
  - A **C4 Container (Level 2)** view describing key containers and their relationships.
  - One or more **C4 Component (Level 3)** views for important containers.
- Check that Mermaid diagrams in this document accurately reflect:
  - The current services, modules, and data stores under `path`.
  - The actual communication paths and dependencies observed in the code and configuration.
- Identify and propose small, incremental updates to keep diagrams and narrative synchronized with the implementation.

**Mode guidance:**

- **Lighter passes:** Highlight a small number of architectural inconsistencies or hacks, plus any obviously outdated diagrams, and propose targeted, low-risk cleanups (including concrete diagram tweaks).
- **Deeper passes:** Propose more systematic alignment of patterns (for example, a consistent error-handling approach) and more comprehensive updates to architecture documentation and C4 diagrams, potentially leading to ADRs and multiple future increments.

## Testing & Reliability (Beck, Feathers)

- Add or improve automated tests for critical paths.
- Refactor code to be more testable (dependency injection, isolation).
- Remove dead code and unused imports.
- Increase test coverage for edge cases that have caused or could cause bugs.
- Use test doubles and mocks for isolation where appropriate.
- Refactor legacy code to enable testing (Feathers).
- Apply Test-Driven Development (Beck) where appropriate for new or changed behavior.
- Automate regression testing for previously broken scenarios.
- Ensure test suites run fast enough to support frequent commits.

**Mode guidance:**

- **Lighter passes:** Focus on a few critical flows or recent bug areas; propose small additions or fixes to tests that clearly increase confidence.
- **Deeper passes:** Address systemic testing gaps (for example, missing tests on a whole domain module) and propose structural changes that improve testability and speed.

## Duplication & Simplicity (Fowler, Thomas & Hunt)

- Consolidate duplicate code where it is clearly beneficial.
- Simplify conditional logic.
- Improve code formatting and indentation.
- Remove unnecessary abstractions.
- Eliminate speculative generality (Fowler).
- Refactor to DRY (Don't Repeat Yourself) with restraint, favoring clarity and local duplication until patterns are proven.
- Prefer simple, readable solutions over cleverness.

**Mode guidance:**

- **Lighter passes:** Target the most obvious duplication or complexity in frequently touched areas; avoid wide refactors.
- **Deeper passes:** Tackle broader simplifications (for example, unifying similar modules or flows) where the benefit outweighs temporary churn.

## Documentation & Communication (Martin, Thomas & Hunt)

- Add missing comments for non-obvious logic.
- Update documentation to match code.
- Document key decisions, trade-offs, and open questions.
- Write ADRs for significant architectural changes.
- Maintain up-to-date README and onboarding docs under `path`.
- Use diagrams to clarify architecture and data flow where they add value.
- Document public APIs and interfaces (inputs, outputs, error behaviors).

**Mode guidance:**

- **Lighter passes:** Propose a few high-leverage documentation updates (for example, fixing critical README gaps or adding a short runbook section).
- **Deeper passes:** Suggest more complete documentation improvements (for example, better module-level docs, diagrams, or ADRs for important patterns).

## Delivery & Flow (Beck, Humble & Farley)

Focus on how easily and safely changes move through the system:

- **PR / change size and scope**
  - Are changes generally small and focused, or large and tangled?
  - Are there frequent “kitchen sink” PRs that mix unrelated concerns?
- **Branching and release habits**
  - Are branches short-lived and merged frequently?
  - Is there unnecessary long-lived branching that slows integration?
- **CI speed and reliability**
  - Are CI pipelines fast enough for frequent commits?
  - Are there flaky tests or noisy checks that slow teams down or erode trust?
- **Automation vs. manual steps**
  - Are there repetitive manual steps for build, test, or release that should be scripted?
  - Is it easy to run the same checks locally as in CI?
- **Bottlenecks and friction**
  - Are there areas where a small change requires editing many files or touching many services?
  - Are there “fear zones” in the codebase that people avoid changing?

Improvements from this lens should:

- Make it easier to ship **small, safe changes**.
- Reduce friction in CI and review (for example, faster tests, clearer scripts).
- Align with the project’s principles around small, frequent releases in `CONSTITUTION.md`.

**Mode guidance:**

- **Lighter passes:** Identify 1–2 concrete friction points (for example, a slow test suite or a common manual step) and propose small automation or test-scope tweaks.
- **Deeper passes:** Look at overall branching, CI structure, and release practices, and propose more structural improvements (for example, reorganizing test suites or simplifying pipelines).

## Dependencies & Operability (Pragmatic, Nygard)

Focus on external dependencies, configuration, and how the system runs in real environments:

- **Library and dependency hygiene**
  - Are there heavy or outdated dependencies that conflict with principles like “wrap and pin” or similar?
  - Are dependency versions pinned and upgrades manageable?
  - Are third-party libraries wrapped behind adapters, or used directly throughout the codebase?
- **Configuration & environments**
  - Is configuration explicit, documented, and easy to override per environment (dev/staging/prod)?
  - Are environment-specific behaviors clear and testable?
- **Build & deployment paths**
  - Are Dockerfiles / build scripts simple and reproducible?
  - Are there fragile manual steps in deploying or running the system?
  - Is local development setup straightforward and consistent with production?
- **Observability vs. constitution**
  - Are logs structured and useful for debugging?
  - Are critical paths at least minimally instrumented (errors, key requests)?
  - For more demanding systems, are basic metrics and/or alerts present where the constitution expects them?

Improvements from this lens should:

- Reduce operational risk (fewer surprises in production or test environments).
- Make dependencies easier to reason about, update, and replace.
- Align concrete practices with dependency and observability principles defined in `CONSTITUTION.md`.

**Mode guidance:**

- **Lighter passes:** Focus on obvious wins (for example, pinning a few key dependencies, wrapping one critical third-party client, or fixing a glaring logging gap).
- **Deeper passes:** Propose more systematic improvements (for example, cleaning up Docker/build scripts, standardizing configuration handling, or aligning logging/metrics across major components).