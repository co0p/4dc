# Architecture — [Project Name]

C4 Level 2: Container diagram. Updated when structural boundaries change.

> This is a durable orientation guide, not a component inventory. It explains the system boundary, runtime containers, important communication paths, and constraints so a reader can reason about change safely.

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

## Reading and Update Guidance

Explain the architectural reasoning that matters to contributors: why the containers are separated, which boundaries must remain stable, and what kinds of changes require an ADR or an update to this document. Do not duplicate class lists, endpoint lists, or deployment instructions.

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

The diagram is a current model, not a historical record. Remove stale paths and obsolete containers rather than preserving them for context.

**Last updated:** [YYYY-MM-DD] — [brief reason]
