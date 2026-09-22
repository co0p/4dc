# Architecture — [Project Name]

Durable orientation for how this system is shaped. Two diagrams are mandatory: C4 Level 1 (System Context) and C4 Level 2 (Container). A third section (Container Internals, C4 Level 3) is populated on demand when a Refactor increment promotes internal restructuring worth preserving.

> This is a durable orientation guide, not a component inventory. It explains the system boundary, its context, runtime containers, important communication paths, and constraints so a reader can reason about change safely.

---

## System Context (C4 Level 1)

**System:** [Project Name]
**Purpose:** [One sentence: what problem does this system solve?]

### Users

| Actor | Type | Interaction |
|-------|------|-------------|
| [Actor A] | [user, operator, administrator, developer] | [what they do with the system — one line] |
| [Actor B] | [type] | [interaction] |

### External Systems

| System | Direction | Purpose |
|--------|-----------|---------|
| [External System X] | inbound / outbound / bidirectional | [what it provides or consumes] |
| [External System Y] | direction | [purpose] |

### Context Diagram

```
                  [Actor A]                      [Actor B]
                     │                              │
                     ▼                              ▼
              ┌──────────────────────────────────────────┐
              │           [Project Name]                 │
              │                                          │
              │        [Purpose in one line]             │
              └──────────────────────────────────────────┘
                     │                              │
                     ▼                              ▼
             [External System X]           [External System Y]
```

---

## Containers (C4 Level 2)

A container is any separately runnable or deployable unit: a process, a script, a service, a database.

| Container | Technology | Responsibility |
|-----------|-----------|----------------|
| [Container A] | [e.g. Go binary, bash script, Node.js process] | [What it does — one line] |
| [Container B] | [technology] | [responsibility] |
| [Container C] | [technology] | [responsibility] |

### Container Diagram

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

### Communication

| From | To | Protocol / Mechanism | Notes |
|------|----|----------------------|-------|
| [Container A] | [Container B] | [e.g. stdout pipe, HTTP, file read] | [any constraint] |
| [User] | [Container A] | [e.g. CLI args, browser] | |

### Data Stores

| Store | Type | Owned by | Schema / Format |
|-------|------|----------|-----------------|
| [Store A] | [e.g. SQLite file, CSV, in-memory] | [Container A] | [brief description] |

If no persistent data store exists, state that explicitly:
> This system is stateless. No persistent data store.

---

## Container Internals (C4 Level 3)

Populated on demand. Each subsection describes the internal component structure of one container. Added or updated only when a Refactor increment promotes internal restructuring worth preserving durably.

> No entries yet. Populated when a Refactor increment produces a component view worth keeping.

<!--
Template for each container that gets an entry:

### [Container Name] — Components

| Component | Responsibility |
|-----------|----------------|
| [Component A] | [one line] |
| [Component B] | [one line] |

Interactions between components:
- [Component A] calls [Component B] via [mechanism]
- [Component C] observes [Component A] events

Notes:
- What this decomposition preserves (boundary, dependency direction, testability seam).
- Why it is worth recording here rather than only in code.
-->

---

## Key Constraints

Constraints that affect all containers and must not be violated:

- [e.g. "No network calls — runs entirely offline"]
- [e.g. "Single binary, no install step"]
- [e.g. "All state is held in browser memory; nothing is written to a server"]

---

## Reading and Update Guidance

Explain the architectural reasoning that matters to contributors: why the containers are separated, which boundaries must remain stable, and what kinds of changes require an ADR or an update to this document. Do not duplicate class lists, endpoint lists, or deployment instructions.

---

## Out of Scope

Explicitly name what this document does NOT cover:

- C4 Level 4 (Code) — the codebase itself
- Deployment topology (see `docs/deployment.md`)
- CI/CD pipeline
- Operational signals (see `docs/observability.md`)

---

## Update Policy

Update this file when:
- A user, external system, or interaction at Level 1 changes.
- A container is added, removed, or its technology changes.
- A communication path between containers changes.
- A Refactor increment promotes a component view into the Container Internals section.

Do NOT update for internal refactors that stay within a container and are not promoted, new features within an existing container, or test changes.

The diagram is a current model, not a historical record. Remove stale paths and obsolete containers rather than preserving them for context.

**Last updated:** [YYYY-MM-DD] — [brief reason]
