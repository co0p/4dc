---
name: 4dc-plan
description: "Use after increment.md is approved. Converts increment intent into an ordered, verifiable technical execution plan with file-level detail so implement skills need minimal context loading."
---

# Plan Skill

## One Responsibility

Define HOW to deliver `.agent/increment.md` — an ordered sequence of actionable subtasks with explicit verification points, file-level scope, and context references. The plan is the map; the implement skills follow it without re-discovering the codebase.

---

## Foundations

{{FOUNDATION:beck-planning-game}}
{{FOUNDATION:poppendieck-pull-small-batches}}
{{FOUNDATION:poppendieck-decide-late}}
{{FOUNDATION:freeman-pryce-tests-guide-design}}
{{FOUNDATION:beck-tidy-first}}

---

## Expected Input

- `CONSTITUTION.md`
- `.agent/increment.md` (must be approved)
- The codebase — read the files that the increment touches, not just the directory listing. The plan must cite specific files, symbols, and line ranges so the implement skills can load narrow context.

---

## Concrete Output

`.agent/plan.md` containing:

- **Goal**: copied from `increment.md` — one sentence
- **Branch**: copied from `increment.md` — the `increment/<slug>` branch for this cycle
- **Approach**: 2–3 sentences on strategy, including architectural boundary and any performance-sensitive path; no code yet
- **Design** (required when data shapes, call flow, or architecture change; omit only for purely structural tidying):
  - **Data Models**: new or modified types, structs, schemas, or domain objects — named fields with types and invariants. Include before/after when modifying existing shapes.
  - **Call / Data Flow**: the sequence of calls or data transformations this increment introduces, from entry point to persistence or output. Prose or a numbered sequence. Names must match actual symbols in the codebase.
  - **Component View (C4 Level 3)**: required when the increment adds, changes, or removes components inside a container. Names the affected container (from `docs/architecture.md`), lists the components involved with one-line responsibilities, and states the interactions between them. New, modified, or removed components must be flagged. Symbol names must match actual codebase symbols. Skip only when no internal structure changes; state the omission explicitly.
  - **Error / Edge-case Inventory**: every error condition and boundary case the increment must handle, derived by walking the call flow and asking "what can fail here?" For each: the condition, the expected response or recovery, and which behavior subtask covers it. Cases not covered by any subtask are gaps — resolve them before approving the plan.
  - **Observability Intent**: the events, metrics, or log lines this increment should emit, stated as a short list. For each: event name, severity/level, and the data it carries. Omit only when the project's constitution explicitly exempts observability for this type of change.
  - **Architecture Delta**: how the container view in `docs/architecture.md` changes — new containers, removed containers, new communication paths, changed dependency direction, or "no change" if none. When the Component View shows restructuring worth preserving durably, flag it as a promotion candidate for `docs/architecture.md`'s Container Internals section.
- **Files**: the complete set of files this increment touches, each labeled by role:
  - `new` — will be created
  - `modify` — existing file that will change
  - `touch` — read for context, unlikely to change (reference only)
  - `delete` — will be removed
- **Subtasks**: ordered list, each with:
  - Type: `[tidy]` or `[behavior]`
  - Description (what, not how)
  - Files: which files this subtask touches (subset of the Files section)
  - References: specific symbols, line ranges, or sections to look at — `path/to/file.ts:42` or `docs/architecture.md#containers`
  - Tests: for `[behavior]` subtasks, a cohesive list of test cases; each case has an id, test file, and test name (Red handles one active case at a time)
  - Verification step (how to confirm it's done)
  - Dependencies on prior subtasks
  - For `[behavior]` subtasks: `acceptance criteria:` covered — one or more criterion ids from increment.md
  - For `[tidy]` subtasks: `supports:` reference — either a `[behavior]` subtask in this plan, or a specific rule in a durable source (see Tidy Rules)
- **Acceptance Tests**: executable or manually reproducible feature-level checks that realize the Increment's Acceptance-Test Intent. Content depends on the Increment's Mode:
  - `Behavior`: one or more new acceptance tests covering every criterion. Each has an id, preconditions, action, expected observable outcome, evidence location or procedure, criteria covered, and gate status.
  - `Refactor`: no new acceptance test. Names the existing acceptance tests from the Increment that anchor the regression proof, with their location and how they will be run.
  - `Chore`: one or more acceptance tests at the actor's boundary declared in the Increment (script exit code, generated artifact hash, log line, config check).
- **Context Map**: pointers to the governing docs for this increment — constitution sections, ADRs, architecture sections, domain entries that the implement skills must respect
- **Risks**: known unknowns that could block execution
- **Planning Decisions**: choices made during planning that a future reader would not infer from the plan alone. For each: the option chosen, the alternatives that were live, and the reason. Captures the "why this approach and not another" before it is lost.

Required `.agent/plan.md` headings:
- `## Goal`
- `## Branch`
- `## Approach`
- `## Design` (omit only when no data shape, call flow, or architecture boundary changes)
- `## Files`
- `## Subtasks`
- `## Context Map`
- `## Acceptance Tests`
- `## Risks`
- `## Planning Decisions`

`[tidy]` is structural only and must preserve observable behavior.
`[behavior]` changes observable behavior and must be verified by a failing test or failing executable check first.

Research is not an implementation subtask type. Resolve plan-blocking unknowns before approval, using `prototype` when a disposable experiment is required. Local implementation uncertainty is investigated while preparing the affected subtask's mini-plan and recorded there.

### Tidy Rules

A `[tidy]` subtask is a small, single-purpose structural change with immediate payoff in this cycle. Timing options (Beck, *Tidy First?*):
- **Tidy first**: a `[tidy]` subtask placed before the behavior it enables.
- **Tidy after**: not a subtask — handled by the Refactor skill after Green, proportional to the behavior just added.
- **Tidy later**: not part of this plan — deferred as a candidate Refactor or Chore increment.

Every `[tidy]` subtask has a `supports:` reference. Two allowed forms:

1. A `[behavior]` subtask in this plan: `supports: subtask 3 (Refresh token before expiry)`
2. A specific rule in a durable source that the code currently violates. Durable sources:
   - `CONSTITUTION.md` — engineering guardrails and boundaries
   - `docs/testing.md` — testing rules
   - `docs/architecture.md` — dependency direction, container boundaries, communication rules
   - `docs/domain.md` — domain vocabulary
   - `docs/deployment.md` — deploy-time constraints reflected in code
   - `docs/observability.md` — signal shape, naming, level, privacy
   - `docs/ui.md` — shared UI, interaction, accessibility, content
   - `docs/adr/ADR-*.md` — recorded architectural decisions
   - any other `docs/*.md` that carries a durable rule

A `supports:` reference to a durable source must cite the section anchor, not the whole file. Examples:
- `supports: CONSTITUTION.md#Architecture-Boundaries`
- `supports: docs/architecture.md#Auth-Container`
- `supports: docs/domain.md#Session`
- `supports: docs/adr/ADR-20260105-idp-choice.md`

Invalid `supports:` targets:
- Bare filenames without a section anchor (e.g. `supports: CONSTITUTION.md`).
- Any `.agent/` path (those are transient, not durable).
- Vague reasoning without a cited rule (e.g. "improves readability").

Proportionality rule: a `[tidy]` subtask supporting a `[behavior]` subtask must be strictly smaller than that behavior. If the reshaping would be disproportionate, it is not a tidy — it belongs in its own **Refactor Increment**.

Cross-file reshaping, dependency inversion, and boundary changes are not `[tidy]` subtasks. They belong in a Refactor Increment.

### Why file-level detail belongs in the plan

The plan is the one phase that reads the whole codebase to find the approach. The implement skills (tdd-red, tdd-green, tidy, refactor) should load only the files named in their subtask — not re-scan the tree. This keeps implement context narrow and fast, and makes the plan's intent verifiable: if a subtask lists no files, it is not actionable.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT start implementation during this phase — no code, no file edits.
Do NOT write `plan.md` until the user explicitly approves the proposed plan.
Do NOT list subtasks without verification steps.
Do NOT list subtasks without a Files field — every subtask must name the files it touches.
Every `[behavior]` subtask maps to at least one acceptance criterion from increment.md.
Every `[tidy]` subtask names a `supports:` reference. It must be one of:
  - a `[behavior]` subtask in this plan that the tidy enables, or
  - a specific rule in a durable source (CONSTITUTION.md, a docs/*.md file, or an ADR under docs/adr/), cited with a section anchor.
Bare filenames, `.agent/` paths, and vague reasoning without a cited rule are not valid `supports:` targets.
A `[tidy]` subtask supporting a `[behavior]` subtask must be strictly smaller than that behavior. Cross-file reshaping, dependency inversion, or boundary changes belong in a Refactor Increment, not in a `[tidy]` subtask.
Every `[behavior]` subtask must name its test file and test name.
Every `[behavior]` subtask may contain multiple test cases, but they must describe one cohesive behavior. Track each case separately with `id`, `name`, `file`, and `state`; set exactly one `active_test` at a time.
Acceptance tests are separate from focused unit or integration test cases. Their shape depends on the Increment's Mode:
- `Behavior`: every acceptance criterion must be covered by at least one new acceptance test. Each is a promotion gate unless the user explicitly approves a documented exception.
- `Refactor`: no new acceptance test. The existing acceptance tests named in the Increment are the promotion anchor and must remain green.
- `Chore`: every acceptance criterion is covered by an acceptance test at the actor's boundary declared in the Increment. Each is a promotion gate.
Do NOT place a `[behavior]` subtask before the `[tidy]` subtasks it depends on.
File paths must be specific (`src/auth/login.ts`), not globs or directory names alone.
References must point to specific symbols, line ranges, or doc sections — not "see the auth module."
Do NOT omit `## Design` when the increment adds or changes data shapes, call flow, or architecture boundaries. Symbol names in the call flow must match actual codebase symbols, not invented names.
Do NOT omit the Component View (C4 Level 3) when the increment adds, changes, or removes components inside a container. Skip only when no internal structure changes and state the omission explicitly.
Do NOT write "no change" in Architecture Delta without checking `docs/architecture.md` first.
Do NOT leave error conditions undiscovered — walk every step in the Call/Data Flow and ask "what can fail here?" before the plan is approved. Any uncovered case is a gap and must appear in Risks or be added to a subtask.
Do NOT omit Observability Intent unless the constitution explicitly exempts this type of change.
Do NOT leave `## Planning Decisions` empty — if the approach was obvious with no alternatives considered, state that explicitly rather than omitting the section.
Do NOT use vague acceptance-test outcomes such as "works", "succeeds", or "is correct". Name the externally visible output, state transition, response, or side effect and the evidence that proves it.
Do NOT create `[research]` subtasks. Resolve blocking uncertainty before plan approval or identify it as investigation inside the affected tidy or behavior subtask's future mini-plan.
</HARD-GATE>

---

## Process

1. **Read inputs** — `CONSTITUTION.md`, `.agent/increment.md`, and any current architecture, design, ADR, or domain docs touched by the change.
2. **Scan the codebase** — find the files the increment touches. Read them well enough to name specific symbols, line ranges, and the boundaries between modules. This is the one phase that reads broadly; the implement skills will read narrowly.
3. **Identify risks** — surface unknowns, performance-sensitive paths, and cross-cutting concerns before drafting.
4. **Draft `## Design`** — for any increment that changes data shapes, call flow, or architecture: name new or modified types with fields and invariants, trace the call sequence from entry point to output using actual codebase symbols, walk every step in the flow and record every failure condition and edge case with its expected response and covering subtask (gaps go to Risks), list the observability events the change should emit, and state the architecture delta against `docs/architecture.md`. Omit only for pure structural tidying with no behavior change.
5. **Draft the Files section** — list every file the increment will create, modify, touch for reference, or delete. Label each by role. Include `docs/ui.md` when shared UI decisions change.
6. **Draft the Context Map** — link the constitution sections, ADRs, architecture sections, and domain entries that govern this change. The implement skills load these by reference, not by re-reading the whole doc.
7. **Draft ordered subtasks** — each with type, description, files, references, test (for behavior), verification, dependencies, and acceptance criteria coverage.
8. **Realize the Increment's Acceptance-Test Intent** — turn the approved intent into concrete acceptance tests, matched to the Increment's Mode:
   - `Behavior`: create new acceptance tests covering every criterion. Specify setup, action, exact observable result, evidence location or manual procedure, and gate status.
   - `Refactor`: name the existing acceptance tests declared in the Increment as the regression anchor. Record their location and how they will be run. Do not add new acceptance tests.
   - `Chore`: create acceptance tests at the actor's boundary named in the Increment (script exit code, generated artifact, log line, config check).
   Walk each criterion to ensure it has proof rather than only lower-level test coverage.
9. **Record `## Planning Decisions`** — before proposing the plan, capture every non-obvious choice made during planning: approach selection, subtask sequencing decisions, trade-offs accepted. If the approach was unambiguous with no real alternatives, state that.
10. **Conversation: Propose the plan** — present the plan and iterate until the user confirms it covers the increment and the file scope is correct.
11. **On approval** — write `.agent/plan.md`. Then load `skills/implement/SKILL.md` to scaffold `.agent/implementation.md` and populate the todo list.

---

## plan.md Structure

```markdown
# Plan: <increment goal>

## Goal
<one sentence, copied from increment.md>

## Branch
`increment/<slug>` — copied from increment.md

## Approach
<2–3 sentences: strategy, architectural boundary, performance-sensitive path>

## Design

### Data Models

<!-- New or modified types. Include before/after for modifications. -->

**New: `TokenClaims`**
```
TokenClaims {
  subject:   string      // user identifier
  expiresAt: timestamp   // Unix seconds; must be > issuedAt
  issuedAt:  timestamp
  scopes:    string[]    // non-empty
}
```

**Modified: `Session`** (adds `claims` field)
```
Before: Session { id, userId, createdAt }
After:  Session { id, userId, createdAt, claims: TokenClaims }
```

### Call / Data Flow

<!-- Numbered sequence from entry point to output. Symbol names must match the codebase. -->

1. `AuthMiddleware.handle(request)` — extracts bearer token from `Authorization` header
2. `TokenValidator.validate(token) → TokenClaims` — decodes and verifies signature; throws `TokenExpiredError` if past `expiresAt`
3. `SessionStore.load(claims.subject) → Session` — looks up or creates session
4. If `claims.expiresAt - now < 60s`: `TokenRefresher.refresh(session) → TokenClaims` — calls IdP silent-refresh endpoint, updates `session.claims`
5. `request.context.session = session` — attaches session to request context for downstream handlers

### Error / Edge-case Inventory

<!-- Walk the call flow step by step and state every failure condition. Each entry must name the covering subtask or flag as a gap. -->

| Condition | Expected response | Covered by |
|-----------|-------------------|------------|
| Token signature invalid | Reject with 401, do not create session | Subtask 2 |
| Token expired and IdP refresh fails | Reject with 401, log warning with `subject` | Subtask 2 |
| Token near-expiry but IdP is unreachable | Allow request with current session; schedule background retry | Subtask 2 — **gap: retry not yet planned; add to Risks** |
| `expiresAt` before `issuedAt` in claims | Reject with 401; `TokenClaims` invariant violation | Subtask 2 |
| `SessionStore.load` returns null (new user) | Create new session; continue | Subtask 2 |

### Observability Intent

<!-- One line per event. Omit only if the constitution explicitly exempts observability for this change type. -->

| Event | Level | Data |
|-------|-------|------|
| `auth.token.validated` | debug | `subject`, `expiresAt`, latency_ms |
| `auth.token.refresh_attempted` | info | `subject`, `idp`, success: bool |
| `auth.token.refresh_failed` | warn | `subject`, `idp`, error_code |
| `auth.token.invalid` | warn | reason, token_prefix (first 8 chars) |

### Architecture Delta

<!-- Describe container-level changes. State "no change" explicitly if none. -->

No new containers. `TokenRefresher` is a new module inside the existing `auth` container — not a separate deployable. The auth container gains a new outbound call to the IdP refresh endpoint (already present in the container diagram; no diagram update needed). If the IdP endpoint proves unreachable in tests, an ADR is required before promote.

## Files

| File | Role | Notes |
|------|------|-------|
| `src/auth/login.ts` | modify | add token refresh logic |
| `src/auth/session.ts` | modify | extend session type |
| `src/auth/__tests__/login.test.ts` | new | test file for behavior subtasks |
| `docs/architecture.md` | touch | reference — container view for auth |
| `src/legacy/token.ts` | delete | replaced by new refresh logic |

## Subtasks

### 1. Extract token validation to its own module
type: [tidy]
files:
  - `src/auth/login.ts` (extract from lines 42–68)
  - `src/auth/token.ts` (new — extracted code moves here)
references:
  - `src/auth/login.ts:42` — `validateToken` function
  - `docs/architecture.md#containers` — auth container boundary
verification: `npm test` — all existing tests pass (behavior preserved)
depends on: —
supports: subtask 2 (Refresh token before expiry) — extraction makes the refresh path testable in isolation

### 2. Refresh token before expiry
type: [behavior]
files:
  - `src/auth/token.ts` (add `refreshIfNeeded`)
  - `src/auth/__tests__/token.test.ts` (new test)
tests:
  - id: refresh-near-expiry
    file: `src/auth/__tests__/token.test.ts`
    name: `refreshes token when within 60s of expiry`
    state: pending
  - id: preserve-valid-token
    file: `src/auth/__tests__/token.test.ts`
    name: `preserves token when more than 60s from expiry`
    state: pending
active_test: refresh-near-expiry
references:
  - `src/auth/token.ts:12` — current token type
  - `CONSTITUTION.md#performance-envelope` — latency budget for auth
verification: `npm test -- token` — fails first, then passes after implementation
depends on: 1
acceptance criteria: AC-2 (token refreshes before expiry)

## Context Map

- `CONSTITUTION.md#testing-strategy` — test depth and naming conventions for auth
- `CONSTITUTION.md#performance-envelope` — latency budget
- `docs/architecture.md#containers` — auth container boundary and dependencies
- `docs/domain.md#session` — domain definition of a session
- `docs/adr/ADR-20260105-idp-choice.md` — why this IdP was chosen

## Acceptance Tests

These feature-level tests prove the approved increment. They block promotion unless an explicit approved exception is recorded.

### AT-1: Successful account lookup

- criterion: AC-1
- user action: request account information
- precondition: provider returns a valid account
- expected outcome: account information is shown
- evidence: `tests/acceptance/account-lookup.test.ts` — assertion on rendered account identifier and balance
- gate: required
- state: planned

### AT-2: Provider unavailable

- criterion: AC-3
- user action: request account information
- precondition: provider times out
- expected outcome: a temporary service error is shown
- evidence: `tests/acceptance/account-lookup.test.ts` — assertion on the temporary-error response and absence of account data
- gate: required
- state: planned

## Risks
- Token refresh may race with in-flight requests (concurrency)
- IdP rate limit on refresh calls (investigate while mini-planning subtask 2; return to plan if it changes scope)
- Near-expiry background retry not yet planned (flagged in Error Inventory)

## Planning Decisions

<!-- Choices made during planning that would not be obvious from the plan alone. Captured here before they are lost. -->

| Decision | Chosen | Rejected | Reason |
|----------|--------|----------|--------|
| Where to put refresh logic | New `TokenRefresher` module in auth container | Inside `AuthMiddleware` | Keeps middleware thin; `TokenRefresher` is independently testable |
| Refresh trigger threshold | 60 seconds before expiry | 30s, 120s | Matches IdP's recommended pre-refresh window from ADR-20260105 |
| Error on IdP unreachable | Allow request, warn | Hard reject | Availability over strict freshness; consistent with constitution performance envelope |
```

---

## Checklist

- [ ] `increment.md` acceptance criteria read
- [ ] Relevant source files scanned (not just directory listing)
- [ ] `## Design` written when data shapes, call flow, or architecture boundaries change:
  - [ ] Data Models: new/modified types named with fields and invariants; before/after shown for modifications
  - [ ] Call / Data Flow: numbered sequence from entry point to output; symbol names match codebase
  - [ ] Component View (C4 Level 3): container named, components and their responsibilities listed, interactions stated; new/modified/removed components flagged. Omission explicitly justified when no internal structure changes.
  - [ ] Error / Edge-case Inventory: every failure condition walked from the call flow; each entry names the covering subtask or is flagged as a gap in Risks
  - [ ] Observability Intent: events, levels, and data listed; omitted only with constitutional justification
  - [ ] Architecture Delta: container-level changes stated, or "no change" confirmed against `docs/architecture.md`; component restructuring worth preserving flagged as a promotion candidate
- [ ] Files section lists every file the increment touches, labeled by role
- [ ] Every subtask has a Files field naming the files it touches
- [ ] Every subtask has a References field with specific symbols, line ranges, or doc sections
- [ ] Every `[behavior]` subtask has a cohesive `tests` list with `id`, `file`, `name`, and `state` for each case
- [ ] Every `[behavior]` subtask has exactly one `active_test`
- [ ] Every `[behavior]` subtask lists at least one covered acceptance criterion
- [ ] Every `[tidy]` subtask has a `supports:` reference: either a `[behavior]` subtask in this plan or a durable rule (CONSTITUTION.md, docs/*.md, or docs/adr/) cited with a section anchor
- [ ] No `[tidy]` subtask uses a bare filename, `.agent/` path, or vague reasoning as its `supports:` value
- [ ] Every `[tidy]` supporting a `[behavior]` subtask is strictly smaller than that behavior; larger reshaping is deferred to a Refactor Increment
- [ ] Related test cases are grouped together; unrelated behavior is split into another subtask
- [ ] Every subtask has a verification step
- [ ] Every acceptance criterion has a covering subtask
- [ ] Acceptance tests cover every acceptance criterion through an observable entry point
- [ ] Every acceptance test states preconditions, action, exact observable outcome, evidence, and gate status
- [ ] Any non-blocking acceptance test has an explicit user-approved exception and rationale
- [ ] Context Map links the governing constitution sections, ADRs, architecture, and domain docs
- [ ] Any architectural or performance-sensitive change is reflected in the approach or risks
- [ ] Risks documented
- [ ] `## Planning Decisions` written: each non-obvious choice records the option chosen, alternatives considered, and reason
- [ ] User approval received
- [ ] `.agent/plan.md` written

---

## Handoff

Terminal artifact: `.agent/plan.md`

Next phase: **implement**

Load `skills/implement/SKILL.md`. The implement skill scaffolds `.agent/implementation.md` from this plan, populates the todo list, and hands off to the first implement skill (`tidy`, `tdd-red`, or `tdd-green`). Only after `implementation.md` exists can the orchestrator route to the appropriate implementation skill.
