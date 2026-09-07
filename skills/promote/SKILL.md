---
name: 4dc-promote
description: "Use after implementation.md is marked complete. Reviews all .agent/ artifacts, proposes promotions to permanent docs, and closes the cycle."
---

# Promote Skill

## One Responsibility

Merge durable outcomes from the `.agent/` working set into permanent project artifacts before the branch merges.

---

## Expected Input

- `CONSTITUTION.md`
- `.agent/increment.md`
- `.agent/plan.md`
- `.agent/implementation.md` (status: complete)
- `.agent/learnings.md`
- Existing `docs/`, ADR log, and any project documentation

---

## Concrete Output

One or more of the following, per approval:
- Updated `CONSTITUTION.md` (if guardrails need revision)
- New ADR in `docs/adr/` (for significant architectural decisions)
- Updated `docs/architecture.md` (if runtime structure, dependencies, or performance-critical paths changed)
- Updated `docs/domain.md` (if domain language changed or new concepts appeared)
- Updated `README.md` or other docs (for changed behavior or usage)
- Updated `docs/roadmap.md` — feature moved from Partial to Done, acceptance test link added
- Deleted or archived `.agent/` files after promotion (keeping `.agent/` clean for next cycle)

Required review outputs:
- Promotion candidates are listed individually with destination path, rationale, and approval status.
- The promotion review explicitly states whether architecture, domain language, testing guidance, and performance documentation changed or stayed unchanged.

## Execution Contract

- Produce only the artifact for this phase. Do not leak work from a later phase into this one.
- Treat tests, architecture notes, ADRs, and user-facing docs as first-class communication artifacts.
- Gather only enough context to identify the governing constraints, the target artifact, and the cheapest validation step. Then act.
- Resolve conflicts in this order: explicit user approval, approved prior-phase artifacts, `CONSTITUTION.md`, this skill.
- Low-risk actions: reads, searches, diffs, and local validation commands.
- Medium-risk actions: local reversible edits to phase artifacts.
- High-risk actions: destructive file operations, external side effects, or skipping a stop gate. Require explicit approval first.
- If a required input is missing or contradictory, ask one focused question or stop at the review gate. Do not invent missing facts.
- Before finishing, run the phase checklist and confirm every required section is present.

---

<HARD-GATE>
Do NOT write permanent docs until each promotion candidate has been individually approved.
Do NOT promote guesses or plans — only promote what was actually built and verified.
Do NOT delete .agent/ files until all promotions are written and confirmed.
Present each candidate separately with destination path and rationale.
Do NOT leave permanent docs stale when the implementation changed architecture, domain language, or performance-critical behavior.
</HARD-GATE>

---

## Process

1. **Read all `.agent/` artifacts** — full review of increment, plan, implementation, and learnings.
2. **Conversation: Propose promotions** — identify candidates and state what each is, its destination, and why it is durable. Iterate until the user says to proceed.
3. **Generate `.agent/promotion-review.md`** — include the required Markdown review sections and an individual approval checkbox for every candidate.
4. **STOP** — present the review and wait for explicit per-candidate approval.
5. **On approval** — write each approved permanent artifact.
6. **Clean up** — archive or delete `.agent/` files for this cycle.

## Markdown Review Contract

Use `.agent/promotion-review.md`. Include **Objective**, **Inputs Reviewed**, **Proposed Output Summary**, **Promotion Candidates** (type, destination, rationale, approval), **Risks and Trade-offs**, **Open Questions**, and **Approval Decision**. Record each explicit conversational approval in the review.

---

## Promotion Categories

| Type | Trigger | Destination |
|------|---------|-------------|
| Architecture decision | Non-obvious choice with lasting impact | `docs/adr/ADR-<date>-<slug>.md` |
| Guardrail update | Constitution rule violated, needs clarification | `CONSTITUTION.md` |
| Architecture sync | Runtime containers, dependency direction, or performance-critical paths changed | `docs/architecture.md` |
| Behavior change | Public API, CLI, or user-facing behavior changed | `README.md` |
| Feature shipped | Acceptance tests pass; feature complete | `docs/roadmap.md` — move to Done, add acceptance test link |
| Test pattern | New testing approach worth standardizing | `CONSTITUTION.md` testing section |
| Performance contract | A latency, throughput, cost, or scaling expectation changed | `CONSTITUTION.md` or `docs/architecture.md` |
| Known issue | Found but not fixed this cycle | `docs/known-issues.md` |
| New domain concept | A concept, event, or rule used in code/tests that has no shared definition | `docs/domain.md` (create using the template in the Appendix if absent) |
| Structural change | A container added, removed, or re-wired | `docs/architecture.md` (create using the template in the Appendix if absent) |

---

## Checklist

- [ ] All `.agent/` artifacts read
- [ ] Promotion candidates identified and categorized
- [ ] Markdown review generated covering all candidates
- [ ] User approval received per candidate
- [ ] Each approved artifact written to permanent location
- [ ] Architecture, domain language, testing guidance, and performance documentation either updated or explicitly marked unchanged
- [ ] `.agent/` files cleaned up

---

## Handoff

Terminal artifacts: permanent docs updated, `.agent/` clean
Cycle complete. Next action: `4dc-increment` for the next cycle — load `skills/increment/SKILL.md`

---

## Appendix: Document Templates

Use these verbatim as the starting content when creating a new document for the first time.

### Template: docs/domain.md

```markdown
# Domain Vocabulary

Shared language for this project. Updated during promote phases.

---

## Concepts

### [ConceptName]

**Definition:** [One sentence in domain terms].

**State:**
- `[field]` — [what it represents and valid values]

**Rules:**
- [Invariant or constraint in domain language]

**Related:**
- [Other concepts this connects to]
- [Events it raises]

---

## Domain Events

### [EventName]

**When:** [What triggers this event in domain terms]
**Payload:** [Key fields involved]
**Consumers:** [Who or what reacts in the domain]

---

## Rules and Constraints

[System-wide invariants, state transitions, cross-concept rules described in domain language, not implementation details]

---

## Example

### User

**Definition:** A person who has authenticated and is interacting with the system.

**State:**
- `email` — contact address; must be unique
- `status` — `active`, `suspended`, or `deleted`

**Rules:**
- Email uniqueness enforced across the system
- Status can transition: `active` ↔ `suspended` → `deleted` (one-way)
- Cannot be deleted if currently has active sessions

**Related:**
- Session (one User has many Sessions)
- UserCreated, UserSuspended, UserDeleted (events raised)


### Session

**Definition:** An authenticated connection between a User and the system.

**State:**
- `user_id` — reference to owning User
- `status` — `active`, `expired`, or `revoked`
- `expires_at` — when the session becomes invalid

**Rules:**
- Expiry time is set at creation and cannot be changed
- Session cannot outlive the User who owns it
- Multiple Sessions can exist for one User

**Related:**
- User (many Sessions per User)
- SessionCreated, SessionExpired, SessionRevoked (events raised)

### UserCreated

**When:** A new User completes authentication for the first time.
**Payload:** `user_id`, `email`, timestamp
**Consumers:** Email service, audit log, user preference initialization

### SessionExpired

**When:** A Session's expiry time elapses or is explicitly revoked.
**Payload:** `session_id`, `user_id`, reason (timeout, logout, or user deletion)
**Consumers:** Cache invalidation, session cleanup, session analytics

---

## System Rules

**User Deletion:** When a User is deleted, all their Sessions are revoked in the same operation.

**Session Lifetime:** A Session cannot outlast the User who created it. If a User is deleted while Sessions exist, those Sessions are revoked.

**Email Uniqueness:** No two active Users can share an email address.
```

### Template: docs/architecture.md

```markdown
# Architecture — [Project Name]

C4 Level 2: Container diagram. Updated when structural boundaries change.

> This is not a design doc. It answers one question: what are the runtime containers, what do they do, and how do they communicate?

---

## Context (C4 Level 1 summary)

**System:** [Project Name]
**Users:** [Who uses it — one line each]
**Purpose:** [One sentence: what problem does this system solve?]

---

## Containers

A container is any separately runnable or deployable unit: a process, a script, a service, a database.

| Container | Technology | Responsibility |
|-----------|-----------|----------------|
| [Container A] | [e.g. Go binary, bash script, Node.js process] | [What it does — one line] |
| [Container B] | [technology] | [responsibility] |
| [Container C] | [technology] | [responsibility] |

---

## Container Diagram

```
┌─────────────────────────────────────────────────────────┐
│  [System Name]                                          │
│                                                         │
│  ┌──────────────────┐         ┌──────────────────────┐  │
│  │  [Container A]   │──────▶  │  [Container B]       │  │
│  │                  │  [how]  │                      │  │
│  │  [technology]    │         │  [technology]        │  │
│  └──────────────────┘         └──────────────────────┘  │
│           │                             │                │
│           ▼                             ▼                │
│  ┌──────────────────┐         ┌──────────────────────┐  │
│  │  [Container C]   │         │  [External System]   │  │
│  │  [technology]    │         │  (out of scope)      │  │
│  └──────────────────┘         └──────────────────────┘  │
└─────────────────────────────────────────────────────────┘

         [User / Actor]
              │
              ▼ [interaction description]
         [Container A]
```

---

## Communication

| From | To | Protocol / Mechanism | Notes |
|------|----|----------------------|-------|
| [Container A] | [Container B] | [e.g. stdout pipe, HTTP, file read] | [any constraint] |
| [User] | [Container A] | [e.g. CLI args, browser] | |

---

## Data Stores

| Store | Type | Owned by | Schema / Format |
|-------|------|----------|-----------------|
| [Store A] | [e.g. SQLite file, CSV, in-memory] | [Container A] | [brief description] |

If no persistent data store exists, state that explicitly:
> This system is stateless. No persistent data store.

---

## Key Constraints

Constraints that affect all containers and must not be violated:

- [e.g. "No network calls — runs entirely offline"]
- [e.g. "Single binary, no install step"]
- [e.g. "All state is held in browser memory; nothing is written to a server"]

---

## Out of Scope

Explicitly name what this diagram does NOT cover:

- Internal component structure of each container (C4 Level 3 — not written unless needed)
- Deployment topology
- CI/CD pipeline

---

## Update Policy

Update this file when:
- A container is added, removed, or its technology changes
- A communication path between containers changes
- A new external system dependency is added

Do NOT update for internal refactors, new features within an existing container, or test changes.

**Last updated:** [YYYY-MM-DD] — [brief reason]
```