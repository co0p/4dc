# Engineering Constitution for Pomodoro (Tray App)

## Purpose

This CONSTITUTION exists to make explicit the engineering principles and trade-offs that guide development of the Pomodoro tray app. It should be used as the first reference when making design, testing, dependency, and release decisions so day-to-day work remains aligned with the product's intent: a minimal, reliable macOS menu-bar Pomodoro implemented as a single Go binary.

## Context

- Product / domain:
  - A lightweight Pomodoro timer that lives in the macOS menu bar. The product prioritizes minimal configuration, predictable behavior, and a small footprint.
- Team:
  - Small, maintenance-focused contributors or a solo maintainer. Preference for short feedback loops and incremental changes over heavy upfront design.
- Non-negotiables:
  - Platform: macOS-only (menu bar / tray experience).
  - Distribution: Single binary artifacts are the primary release format.
  - Privacy: Offline by default; no telemetry or automatic reporting.
  - Implementation: Go (golang) is the primary implementation language.

## Our Principles and Trade-offs

We prioritize delivering a focused, dependable experience with minimal overhead. That means favoring straightforward implementations and tiny, well-tested core logic while allowing the UI glue to remain pragmatic and platform-specific.

- Speed vs safety: prefer fast, small, reversible changes that keep the app useful; invest in tests for core timer semantics and config handling to avoid regressions.
- Short-term delivery vs long-term maintainability: prefer simple, clean code for core components and accept limited, recorded shortcuts in non-critical UI glue when necessary. Pay down those shortcuts intentionally.
- Experimentation vs stability: permit short-lived experiments behind feature flags or in branches, but only release experiments when they meet the minimal quality expectations defined in the Test Strategy.

### Default Trade-off Rules

- When in doubt between **shipping faster** and **polishing the design**, we usually:
  - Ship a small, well-scoped change with automated tests for core logic and a clear release note. Defer cosmetic polishing to a follow-up if it can be done without user-visible regressions.
- When in doubt between **adding a dependency** and **building it ourselves**, we usually:
  - Prefer a small, actively maintained Go library that reduces risk or code surface area. Avoid large frameworks or dependencies that pull significant transitive noise.
- When in doubt between **adding tests now** and **moving on**, we usually:
  - Add at least unit tests for core logic and integration tests for non-trivial state transitions. For purely UI glue changes, rely on manual verification but add unit tests for any moved or extracted logic.

---

## The 6 Pillars of Our Engineering

### 1. Delivery Velocity

We aim for rapid, low-risk iterations that keep the app useful and small.

- Desired iteration speed: short cycles (days to a couple of weeks) for small features and fixes; emergency fixes are allowed as patches.
- Typical size of changes: single-concern PRs that implement one feature or fix and include tests for core behaviors.
- Release cadence and acceptable risk per release: frequent single-binary releases are acceptable; each release must preserve core timer semantics and basic config compatibility unless a breaking change is explicitly documented.

**We optimize for:**
- Small, test-backed changes that deliver user-facing value quickly.

**We accept the following risks:**
- Minor, documented UI regressions that do not break timer semantics or user data.

**We avoid:**
- Large, multi-feature releases that increase regression risk without adequate testing.

### 2. Test Strategy

Testing focuses on ensuring core timer correctness and preventing regressions in state and configuration handling.

**Minimum expectations:**
- Unit tests covering timer logic (tick behavior, start/stop/reset semantics), configuration parsing, and any non-trivial state transitions.
- A small integration test or smoke check that validates the main run loop and persistence (if present) in CI where practical.
- Manual UI checks for menu-bar interactions, notifications, and platform behaviors before release.

**When moving fast, we are allowed to:**
- Rely on manual verification for superficial UI tweaks, provided core logic is covered by automated tests.

**We never skip tests for:**
- Changes that touch timer semantics, persistence, or configuration formats that affect user data or expected behavior.

### 3. Design Integrity

Structure code to keep the domain (timer semantics and scheduling) independent from platform UI code.

**We strive for:**
- A small, testable core package that implements all timer rules and state transitions.
- A thin platform adapter that implements menu-bar UI, notifications, and OS integration.

**We are okay with:**
- Some messiness in platform glue (menu-bar integration and macOS-specific bits) as long as boundaries between core logic and UI remain clear and covered by tests where feasible.

**Red flags that trigger redesign or refactoring:**
- Core logic depending on platform APIs or UI code.
- Repeated, duplicated timer logic spread across UI adapters.

### 4. Simplicity First

The app exists to provide a single, clear job: reliable Pomodoro timing. Complexity should only be introduced when it provides clear value.

**We prefer:**
- The simplest thing that could possibly work, then iterate.

**We add abstraction only when:**
- Multiple features or code paths duplicate the same logic, or when an abstraction significantly reduces risk or test surface.

**We treat complexity as acceptable when:**
- It is necessary for correctness, performance within the single-binary constraint, or significant developer productivity gains.

### 5. Technical Debt Boundaries

Shortcuts are allowed when they are small, documented, and scheduled for repayment.

**Allowed short-term shortcuts:**
- Quick UI hacks for prototype or experimental behavior that are flagged in the issue tracker and bounded by scope.

**Debt must be recorded when:**
- A change sacrifices clarity, testability, or introduces duplicated logic to speed delivery.

**We commit to paying down debt when:**
- The debt interferes with adding features, causes repeated bugs, or when it accumulates past a small, agreed threshold (for example, more than two related debt items blocking core improvements). Payment should be scheduled into the next minor release or an explicit refactor sprint.

### 6. Dependency Discipline

Dependencies must be evaluated for size, transitive impact, maintenance, and security.

**We add a new dependency only when:**
- It materially reduces risk or work, has minimal transitive surface, and is actively maintained in the Go ecosystem.

**We isolate dependencies by:**
- Encapsulating third-party usage behind small adapters/interfaces so the core domain remains dependency-free and easy to test.

**We avoid:**
- Adding large frameworks or cross-platform UI stacks that bloat the single-binary philosophy or leak framework concepts into the domain model.

---

## How We Use This Constitution

- How work is chosen and sliced: prefer small, user-visible increments that can be landed, tested, and released independently. Each change should map to a single user benefit and include a note if it introduces technical debt.
- How designs are evaluated: any design should be measured by clarity, testability, and minimal impact on binary size and runtime behavior. Designs that obscure core semantics get rejected.
- How implementation and testing decisions are made: core logic always requires automated tests; UI-only adjustments can be landed with manual verification but should be followed by refactors or tests if repeated.
- When to refactor, pay down debt, or revisit architecture: refactor when duplicate logic or bugs indicate boundary erosion, or when new features force repeated changes across modules. Schedule debt repayment into the next minor release when practical.

## Amendments and Evolution

- This CONSTITUTION should be revisited when there is a major product shift (for example, expanding beyond macOS), sustained team growth, or repeated friction in day-to-day work.
- Amendments should be documented with a date, author, and short rationale. Use incremental versioning (e.g., `v1.0 — 2025-12-03`) and keep the document under source control at the project root.

## References and Inspirations

- The Pomodoro Technique (timeboxing and focused work principles).
- Unix philosophy and KISS (Keep It Simple, Stupid) for small, composable tools.
- Design sensibilities of small macOS menu-bar utilities: minimal UI, quick discoverability, and low configuration.

## Open Questions

- Confirm pillar priorities: current default is `Simplicity First > Delivery Velocity ≈ Design Integrity > Test Strategy > Dependency Discipline > Technical Debt Boundaries`.
- Packaging signatures and distribution channels: should we require notarization or Homebrew integration in the future?
- Support expectations: what level of user support and incident response is expected for releases?
