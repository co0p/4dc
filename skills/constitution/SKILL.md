---
name: 4dc-constitution
description: "Use when CONSTITUTION.md is missing or needs updating. Reads project context, asks focused questions, and produces project guardrails and SDLC standards."
---

# Constitution Skill

## One Responsibility

Create or update `CONSTITUTION.md` — the project's durable engineering guardrails.

The generated constitution must describe only the application. Do not include this repository's internal workflow name, phase sequence, agent names, or transient artifact paths.

---

## Expected Input

- Existing `CONSTITUTION.md` (if present)
- `README.md`
- Current project structure and any existing docs

---

## Concrete Output

### Primary Artifact: `CONSTITUTION.md`

`CONSTITUTION.md` containing:
1. Engineering principles grounded in XP, lean software development, and use-case thinking
2. Architectural boundaries, dependency direction, and performance-critical paths
3. **Testing strategy reference** — points to `docs/testing.md` (test types in scope, what must have tests, the gate that must be green before promote, naming conventions)
4. **Performance envelope** — stated latency, throughput, cost, or scale expectations, or an explicit `N/A` with rationale
5. **Release and deployment reference** — points to `docs/deployment.md` (how a release is triggered, versioning scheme, deployment target(s), rollback procedure)
6. Documentation rules, architecture sync rules, and ADR policy
7. Documentation and delivery expectations appropriate to the project

Required `CONSTITUTION.md` headings:
- `## Engineering Principles`
- `## Architecture Boundaries`
- `## Testing Strategy`
- `## Performance Envelope`
- `## Documentation And ADR Policy`
- `## Release And Deployment`
- `## Delivery and Documentation`

### Supporting Documents

**`docs/testing.md`** — testing practices for this project:
- The reasoning behind the testing strategy and the risks it is intended to control
- Guidance for choosing test depth at architectural and user-facing boundaries
- Actual commands for local, CI, acceptance, and release checks, with setup and interpretation guidance
- Evidence required before a change is considered complete
- Known confidence gaps, maintenance practices, and reasons for environment-specific checks
- No inventory of individual tests, test-case lists, or coverage targets unless a project-specific decision genuinely requires one

**`docs/deployment.md`** — deployment and release procedures for this project:
- The deployment model, its rationale, and the operational assumptions it relies on
- Release triggers, versioning decisions, ownership, and required evidence
- Actual deployment and rollback runbooks with verification and recovery guidance
- Configuration and secret-handling principles without secret values
- Health signals, alert actions, and meaningful operational risks
- No historical release log or generic checklist detached from this project's procedure

**`docs/adr/`** — Architecture Decision Records:
- Decisions with rationale and consequences
- Indexed from `CONSTITUTION.md`
- Created using the ADR template when foundational decisions exist
- Updated when architectural decisions emerge
- Each ADR explains the context, decision, alternatives, trade-offs, and consequences; it is not an implementation diary

**`docs/architecture.md`** — C4 architecture view:
- Required for every project, even when the system is small
- Must contain a current C4 Level 2 container view (or an equivalent explicitly labeled diagram)
- Describes runtime containers, responsibilities, communication paths, and data stores

**`docs/domain.md`** — Domain glossary:
- Required for every project, even when the vocabulary is initially small
- Defines shared concepts, domain events, and system rules in business language
- Must not be replaced by an ADR, README, or implementation-specific notes

### Secondary Artifact

`docs/roadmap.md` (created from the template in the Appendix if it does not exist yet)

### Documentation Baseline

Before creating or updating the constitution, audit the repository for the complete permanent documentation baseline:

- `CONSTITUTION.md`
- `docs/testing.md`
- `docs/deployment.md`
- `docs/architecture.md` with a C4 Level 2 container view
- `docs/domain.md` with the project's glossary
- `docs/adr/`
- `docs/roadmap.md`

Missing baseline documents are constitution outputs; they are not optional follow-up work. Existing documents must be checked for the required content rather than accepted solely because the path exists.

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
Do NOT write `CONSTITUTION.md` until the Markdown review in `.agent/constitution-review.md` has been explicitly approved.
Do NOT ask more than 5 questions per round.
Do NOT include implementation details — CONSTITUTION.md contains guardrails, not recipes.
Do NOT copy generic principles from the internet. Every rule must be justified by this project's specific context.
</HARD-GATE>

---

## Process

1. **Read project context** — scan `README.md`, existing `CONSTITUTION.md`, directory structure, any ADRs or docs, and existing deployment or testing practices
2. **Conversation: Propose the guardrails** — summarize the proposed testing and deployment strategies, identify foundational ADR candidates, and ask whether the direction feels right. Iterate until the user says “looks good” or “proceed.”
3. **Generate `.agent/constitution-review.md`** — include the required Markdown review sections and proposed outputs.
4. **STOP** — present the review and wait for explicit approval.
5. **On approval:**
   - Write `CONSTITUTION.md` with references to supporting documents
   - Create `docs/testing.md` with project-specific testing practices
   - Create `docs/deployment.md` with project-specific deployment procedures
   - Create or update `docs/architecture.md` with the current C4 Level 2 container view
   - Create `docs/domain.md` with the project's initial glossary, even if only a few concepts are known
   - Create initial `docs/adr/` structure if foundational decisions exist
   - Create `docs/roadmap.md` if not present

## Markdown Review Contract

Use `.agent/constitution-review.md`. Include **Objective**, **Inputs Reviewed**, **Proposed Output Summary**, **Risks and Trade-offs**, **Open Questions**, and **Approval Decision**. An explicit conversational approval is sufficient; record it in the Approval Decision section.

---

## Checklist

- [ ] Existing docs read (including any deployment or testing practices)
- [ ] Foundational ADRs identified (if any exist)
- [ ] Proposed testing strategy, deployment strategy, and ADR candidates discussed
- [ ] Markdown review generated and shown
- [ ] User approval received
- [ ] `CONSTITUTION.md` written with Testing, Performance, and Release sections populated (with references to supporting docs)
- [ ] `docs/testing.md` created with project-specific practices
- [ ] `docs/deployment.md` created with project-specific procedures
- [ ] `docs/adr/` directory created with index link from `CONSTITUTION.md` (populate with foundational decisions if identified)
- [ ] `docs/roadmap.md` created if not present

---

## Handoff

Terminal artifacts:
- `CONSTITUTION.md` — guardrails and governance
- `docs/testing.md` — testing procedures and practices
- `docs/deployment.md` — deployment and release procedures
- `docs/adr/` structure — architectural decisions
- `docs/roadmap.md` — product roadmap

Future work should reference the testing, deployment, and architecture documents defined by the project constitution. New ADRs should be added when architectural decisions emerge.

After this skill, continue with the next approved work item using the applicable project workflow.

---

## Appendix: Document Templates

Use these verbatim as the starting content when creating a new document for the first time.

### Template: docs/testing.md

```markdown
# Testing

Guide to making reliable testing decisions for this project. Explain why the test strategy is shaped this way, how a developer should choose the cheapest test that gives sufficient confidence, and how to run the relevant checks. Update when the architecture, risk profile, or test workflow changes.

---

## Testing Approach and Rationale

Explain the risks the test strategy is designed to control and the boundaries where each kind of test provides confidence. Prefer principles and decision guidance over inventories. For example, explain why domain rules are tested without infrastructure, why persistence boundaries need integration checks, or why an acceptance test exercises a complete user job story.

---

## Choosing Test Depth

<!--
Describe how to decide whether a change needs a focused check, an integration check, an acceptance scenario, a performance measurement, or no new test. Tie the decision to user risk, architectural boundaries, determinism, and failure cost.
Do not maintain a catalog of individual tests or report a coverage percentage here.
-->

---

## Test Design Conventions

<!--
Describe conventions that make tests communicate behavior: naming, fixture ownership, isolation, determinism, test data, and how user-facing assertions should avoid implementation details. Keep examples small and illustrative rather than listing the suite.
-->

---

## Running the Checks

<!--
Document the actual commands for fast local feedback, the complete pre-merge gate, and any setup required for acceptance or environment-dependent checks. Explain when to use each command and how to interpret failures. Commands must be maintained as executable guidance, not illustrative placeholders.
-->

---

## Evidence Required Before Promotion

<!--
State the evidence required before a change is considered complete. Focus on behavior, risk, and reproducibility. Do not use line coverage or a list of passing tests as a substitute for explaining why the evidence is sufficient.
-->

---

## Automation and Feedback Loops

<!--
Explain where checks run (local, CI, release), what feedback each loop provides, and how failures are triaged. Record the rationale for any intentionally manual or environment-specific check.
-->

---

## Known Risks and Gaps

<!--
Document meaningful confidence gaps, why they exist, and what signal would justify changing the approach. Do not turn this section into a test inventory or coverage report.
-->

---

## Maintenance Guidance

<!--
Explain how tests are kept deterministic, how flakiness is handled, when fixtures or helpers should be changed, and how this guide itself is updated when the testing rationale changes.
-->
```

### Template: docs/deployment.md

```markdown
# Deployment

Guide to releasing and operating this project. Explain the deployment model, why it is appropriate, how to execute it safely, and how to recover. Update whenever the release or operational model changes.

Describe the release unit, the target, ownership, and the operational assumptions. Explain why this deployment shape is used.
- This is a Node.js backend service deployed to AWS Lambda
- This is a React frontend deployed to Vercel
- This is a Go CLI tool distributed via Homebrew and GitHub Releases
-->

Explain the release trigger and versioning decision, including who can release and what evidence is required first. Avoid a changelog or release-history list here.
- Manual: Tag a commit with `v<major>.<minor>.<patch>` and push to origin; CI builds and publishes
- Automatic: Merge to `main` triggers a release with semantic versioning based on conventional commits
- Versioning scheme: Semantic Versioning (SemVer) for libraries, CalVer for applications
-->
Describe the environments and their purpose, including the differences that matter for safe verification. Do not use this section as an environment inventory without explaining the deployment model.
Or for CLI:
- macOS: Homebrew tap `example/tap/tool`
- Linux: GitHub Releases, apt repository
- Windows: GitHub Releases, Scoop bucket
-->

---

## Deployment Procedure

Document the actual release procedure as a short runbook, with prerequisites, commands or links, verification signals, and ownership. Explain why the ordering protects users and data. Keep the checklist operational; put rationale in surrounding prose.
Describe rollback triggers, the recovery action, data implications, and who decides. Explain any forward-only or irreversible operation and the recovery alternative.

For CLI releases:
- Yanked versions: Tag with `v<version>` and mark as yanked in release notes
- Users on old version: Keep supporting previous major version for [N] months
-->

---

## Deployment Checklist

Keep only the small set of release decisions and checks that are specific to this project. Do not turn this into a repeated list of every test, deployment, or release ever performed.

Explain configuration ownership, secret handling, safe defaults, and the reason for separating deploy-time configuration from source code. Never record secret values.
- Environment variables are injected at deployment time from [source]
- Configuration file: `.env.production` (not checked in), managed by [process]
- Database connection string: Retrieved from [secrets manager] at startup
Explain how operators know a release is healthy, which signals matter, and what action an alert should trigger. Record thresholds only when they are real, justified, and maintained.
- Error rates are monitored in Datadog; alert if error rate > 1% for 5 min
- Latency p99 is tracked; alert if > [threshold]
- Database connection pool is monitored; alert if exhausted
- Disk space is monitored; alert if < 10% free
- Deployment notifications sent to [Slack channel]
-->

Document constraints that materially affect release safety and the mitigation or follow-up needed. Do not preserve obsolete procedures as historical reference.
- Deployment is not atomic: old and new code may run simultaneously for [duration]
- Secrets rotation requires [manual step]
- Large deployments > 100MB take [N] minutes; monitor for timeout
-->
```

### Template: docs/adr.md

When creating a new Architecture Decision Record, use this template:

```markdown
# Architecture Decision Record (ADR) Template

ADRs document significant architectural decisions and their rationale. They serve as a concise guide to _why_ the system is structured a particular way, not a changelog, implementation diary, or list of completed work.

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

[What did we decide to do? Be specific about the boundary or principle, but avoid duplicating implementation details that belong in code or runbooks.]

Example:
- Store user preferences in a transactional SQL database (PostgreSQL on AWS RDS)
- Use an ORM (SQLAlchemy / Prisma / Entity Framework) to manage schema and migrations
- Cache in-process with TTL of 5 minutes to meet latency requirement
- Implement read replicas for analytics queries to avoid impacting transactional workload

**Consequences (Positive)**

[What gets better, and why is that useful to the project?]

- Preferences survive application restarts
- Easy to query and report on user behavior
- Horizontal scaling: multiple app servers share same database
- Standard tooling and team knowledge

**Consequences (Negative)**

[What gets harder, riskier, or more constrained?]

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

## How ADRs Connect to Project Work

**Project initialization:** High-level architectural principles and boundaries are documented in `CONSTITUTION.md`. ADRs extend this with decisions about _specific_ components or patterns.

**Later development:** When a significant decision emerges during implementation, it is captured in project decision notes and recorded as a new ADR file in `docs/adr/`.

**Reference:** `CONSTITUTION.md` points to `docs/adr/` as the permanent log of architectural decisions; newly added ADRs are linked from there.

---

## ADR Review and Approval

Before an ADR is merged:
1. It must be reviewed by at least one person familiar with the affected system
2. All "Questions for Review" must be answered or marked as deferred
3. Status must be set to "Accepted" (or "Deferred" if not yet implemented)

After it is merged:
- Link it from the project's architecture or ADR index
- Update any related ADRs with "Supersedes" or "Related Decisions" links
- If a related decision is reversed, update the old ADR's status to "Superseded"
```

### Template: docs/roadmap.md

```markdown
# Roadmap

Product direction and sequencing guide. Each entry explains the user outcome, current confidence, and ordering rationale. Keep it concise and decision-oriented; detailed implementation status belongs in phase artifacts and code evidence.

> A feature moves to **Done** only when its user outcome is verified and the evidence is linked here.
> Source of truth: if a feature is not in Done with a passing test link, it is not considered shipped.

---

## Done

<!--
Each entry follows this pattern:

### [Feature name — short, user-visible]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Evidence:** [link to the smallest durable verification record](path/to/evidence)
- **Use case:** [docs/usecases/use-case-slug.md](docs/usecases/use-case-slug.md) *(if promoted)*
- **Delivered:** [increment slug or YYYY-MM-DD]
-->

---

## Partial

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Evidence:** pending — define the verification approach in the approved plan
- **Increment:** [increment slug]
-->

---

## Planned

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Why now / ordering:** [dependencies, user value, open questions, or sequencing rationale]
-->

---

## Rules

- Features move left to right: Planned → Partial → Done. Never skip Partial.
- A feature enters Partial when its increment is approved.
- A feature enters Done only when its evidence link resolves to a current verification record.
- Do not add implementation detail here — link to the use case or ADR for that.
- If a planned feature is no longer needed, remove it and note the removal in `learnings.md` for that cycle.
```