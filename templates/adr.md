# Architecture Decision Record (ADR) Template

ADRs document significant architectural decisions and their rationale. They serve as a permanent record of _why_ the system is structured a particular way, not just _what_ was built.

---

## File Naming Convention

Save ADRs in `docs/adr/` with a timestamp and slug:

```
docs/adr/ADR-YYYYMMDD-slug.md
```

Example:
- `docs/adr/ADR-20260907-split-testing-from-constitution.md`
- `docs/adr/ADR-20260905-cache-invalidation-strategy.md`

---

## Template: Use This Structure

```markdown
# ADR-YYYYMMDD — [Decision Title]

**Decision:** [One-sentence statement of what you decided]

**Date:** YYYY-MM-DD

**Status:** Accepted | Rejected | Superseded | Deferred

**Context**

[Why does this decision matter? What problem are we solving? What constraints exist?]

Example:
- A new feature requires storing user preferences; we need to decide where and how.
- Current approach (in-memory cache) does not survive application restarts.
- Performance requirement: lookups must be < 10ms on average.
- Team skill: no DBA expertise; prefer managed services.

**Decision**

[What did we decide to do? Be specific.]

Example:
- Store user preferences in a transactional SQL database (PostgreSQL on AWS RDS)
- Use an ORM (SQLAlchemy / Prisma / Entity Framework) to manage schema and migrations
- Cache in-process with TTL of 5 minutes to meet latency requirement
- Implement read replicas for analytics queries to avoid impacting transactional workload

**Consequences (Positive)**

[What gets better?]

- Preferences survive application restarts
- Easy to query and report on user behavior
- Horizontal scaling: multiple app servers share same database
- Standard tooling and team knowledge

**Consequences (Negative)**

[What gets harder or more complex?]

- Added operational burden: database backups, monitoring, security patching
- Network latency: every preference lookup requires a round trip (mitigated by cache)
- Cost: managed database service is not free
- Schema migrations: changes require careful coordination with application deployments

**Alternatives Considered**

1. **Keep in-memory cache** — Simple but data loss on restart; ruled out by requirement.
2. **Use Redis** — Fast and battle-tested, but adds another infrastructure service to operate.
3. **Use file-based storage (SQLite)** — No operational overhead but poor multi-process scaling.
4. **Use cloud document store (Firestore / DynamoDB)** — Pay-per-request pricing could be costly at scale.

**Rationale for Decision Over Alternatives**

We chose PostgreSQL because:
- Team has existing RDS infrastructure and operational patterns
- SQL queries are predictable for future analytics
- Cost is predictable and lower than document databases at our anticipated scale
- ORM options provide familiar abstractions

**Related Decisions**

- [ADR-YYYYMMDD-caching-strategy.md](#) — explains the in-process cache layer
- [ADR-YYYYMMDD-database-migrations.md](#) — explains schema versioning approach

**References**

- [Relevant architecture doc](#)
- [Relevant RFC or issue link](#)
- [External standard or reference](#)

**Questions for Review**

- Is this the right data model for future use cases? (Reviewers: [name])
- Should we add read replicas from day one? (Ops: [name])
- Migration strategy: how do we backfill existing users? (Tech lead: [name])
```

---

## When to Write an ADR

Write an ADR when:

- You make a **structural choice** that affects multiple parts of the system (new layer, new service, new tech)
- The choice involves **trade-offs** (performance vs. simplicity, flexibility vs. cost)
- The choice is **difficult to reverse** (schema design, language choice, external dependency)
- The choice is **not obvious** to someone new to the codebase

Do **not** write an ADR for:
- Small implementation details (choice of variable name, ordering of function arguments)
- Bug fixes or maintenance patches
- Decisions that are already captured in CONSTITUTION.md or public docs

---

## How ADRs Connect to 4dc Phases

**Constitution phase:** High-level architectural principles and boundaries are documented in `CONSTITUTION.md`. ADRs extend this with decisions about _specific_ components or patterns.

**Promote phase:** When a significant decision emerges during implementation, it is captured in `.agent/learnings.md` and promoted as a new ADR file in `docs/adr/` during the promote phase.

**Reference:** `CONSTITUTION.md` points to `docs/adr/` as the permanent log of architectural decisions; newly added ADRs are linked from there.

---

## ADR Review and Approval

Before an ADR is merged:
1. It must be reviewed by at least one person familiar with the affected system
2. All "Questions for Review" must be answered or marked as deferred
3. Status must be set to "Accepted" (or "Deferred" if not yet implemented)

After it is merged:
- Link it from `docs/ARCHITECTURE.md` or an index
- Update any related ADRs with "Supersedes" or "Related Decisions" links
- If a related decision is reversed, update the old ADR's status to "Superseded"
