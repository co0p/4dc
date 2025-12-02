# Project Constitution

## Vision
Enable friends to easily share personal items with trust and transparency.

## Mission
Deliver a mobile-first web app that lets people publish catalogs of items, request to borrow, and track borrow/return status with clear visibility and minimal friction.

## Core Values
- Simplicity over complexity
- Trust and transparency
- Reliability and correctness
- Privacy by default
- Small, incremental change

## Architectural Principles

### 1. Small, Frequent Releases _(Pillar: Delivery Velocity)_
**Statement:** Ship small changes weekly or faster with PRs under 300 lines, each including at least one verification (test or documented manual check).
**Rationale:** Faster feedback reduces risk and improves flow for a mobile-first product.
**In Practice:**
- Work on feature branches; merge to default via reviewed PRs.
- Target at least weekly releases; hotfix on demand.
- Each PR includes a unit/smoke test or a manual check note tied to acceptance criteria.

### 2. Protect Borrow/Return Flows _(Pillar: Test Strategy)_
**Statement:** Cover critical item catalog and borrow/return flows with automated tests; allow smoke/manual checks for UI glue.
**Rationale:** These flows are core to user value and must not regress.
**In Practice:**
- Domain functions for borrow/return have unit tests.
- One end-to-end smoke path: add item → borrow → return.
- Error paths log clearly and show actionable messages; manual checks documented when tests aren’t feasible.

### 3. Thin Routes, Isolated Domain _(Pillar: Design Integrity)_
**Statement:** Keep Express routes thin; business rules live in domain modules; data access behind adapters.
**Rationale:** Maintains testability and flexibility across UI/backends and simplifies change.
**In Practice:**
- Route handlers validate/dispatch; domain modules hold business logic.
- Persistence and external APIs use adapters to enable swap/upgrade.
- Domain modules avoid framework dependencies where practical.

### 4. Three-Strikes Before Abstraction _(Pillar: Simplicity First)_
**Statement:** Duplicate locally up to three occurrences before creating shared abstractions.
**Rationale:** Avoid premature abstractions that add overhead and slow iteration.
**In Practice:**
- Prefer clear duplication over early utility modules.
- Extract shared code only after the third repetition across modules.
- Perform refactors in small, focused PRs.

### 5. Wrap and Pin _(Pillar: Dependency Discipline)_
**Statement:** Minimize new dependencies; wrap third‑party libraries and pin versions for reproducible builds.
**Rationale:** Reduces supply‑chain and upgrade risk for Docker/Heroku deployments.
**In Practice:**
- Add dependencies only with rationale; prefer platform/standard libs.
- Provide adapter boundaries around third‑party libs/APIs.
- Pin versions; update via dedicated, reviewable PRs.

## Pillar Coverage
- ✓ Delivery Velocity (Principle #1)
- ✓ Test Strategy (Principle #2)
- ✓ Design Integrity (Principle #3)
- ✓ Simplicity First (Principle #4)
- ✓ Dependency Discipline (Principle #5)

## Update Process
- Propose changes via PR referencing this document and affected principles.
- For principle changes, include rationale, trade‑offs, and expected impact; request at least one reviewer.
- On approval, update this document and record date under Last Updated.
- For sweeping changes, prefer incremental PRs gated behind feature branches.

## Technical Decisions
### Languages
- JavaScript/TypeScript (Node.js) for server‑side logic, chosen for ecosystem, deployment simplicity, and team familiarity.
### Frameworks
- Express for HTTP routing; TailwindCSS for UI styling.
### Deployment
- Single Docker image; deploy to Heroku via container registry.
- Pinned dependencies; reproducible builds recommended with SBOM generation.

## Increments Location
- Base directory: `docs/increments/`
- Per‑increment folder structure: `docs/increments/<increment-folder>/increment.md`

## LLM Collaboration & Increment Workflow
- STOP gates: Clarification → Planned Files Summary confirmation → Drift Alert approval.
- JSON internal‑only: Structured outputs (questions/proposals/summaries) are for tooling/CI and not user‑facing.
- Non‑interactive phase: After file plan confirmation, the implementer may proceed autonomously until the next STOP gate or drift alert.
- Feature branch: Work on `feature/<increment-slug>`; no direct commits to default branch.
- Commit cadence: One commit per high‑level task; small and revertible.

## Scope Drift Management
- DRIFT ALERT: If work requires files/modules beyond the confirmed Planned Files Summary or increment scope, STOP and announce a drift.
- Scope update path: Propose a minimal scope change listing files and rationale, or split into a follow‑up increment.
- Design alignment: If acceptance criteria can’t be met as designed, request a design update before continuing.

## Testing & Verification Policy
- Cycle: Write Test → Implement → Validate → Commit; allow explicit manual checks when tests aren’t feasible.
- Critical‑path testing: Unit tests for domain logic; smoke/e2e for primary user flow.
- Verification mapping: Every task/subtask documents how it verifies acceptance criteria.

## Post‑Implementation Stabilization
- Documentation: Update README/usage and increment docs (design/implement notes, checkboxes).
- Hygiene: Add `.gitignore` entries for build outputs; untrack committed artifacts.
- Reproducibility: Ensure scripts build cleanly from a fresh checkout; document prerequisites.
- Packaging: Verify bundling/distribution if within acceptance criteria.
- Formatting/Linting: Run minimally on touched files.

## Merge & Release
- Merge strategy: Rebase/merge on feature branch; then merge to default and push.
- Tagging: If release‑worthy, create an annotated tag (e.g., `vX.Y.Z`) with brief notes.
- Cleanup: Delete local and remote feature branch after merge.

## Documentation & Traceability
- Required per‑increment docs: `increment.md`, `design.md`, `implement.md` with verification notes.
- Decision log: Capture key choices, trade‑offs, and drift alerts in `implement.md`.
- Traceability: Link tasks and verification to acceptance criteria.

## Roles & Decision Gates
- Sponsor: Confirms scope and Planned Files Summary; approves drift changes.
- Implementer (LLM/human): Executes plan, raises drift alerts, proposes scope adjustments.
- Reviewer (optional): Sanity‑checks stabilization steps and merge.

---
**Last Updated:** 2025-12-02
