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
- Test types and scope (unit, integration, acceptance, performance, etc.)
- What requires tests (behavior changes, bug fixes, refactoring rules)
- Test location and naming conventions
- How to run tests locally and in CI
- Test gate before promote (constitution-level requirement)
- Known gaps and limitations

**`docs/deployment.md`** — deployment and release procedures for this project:
- Deployment overview (what and where)
- Release triggers and versioning scheme
- Deployment targets (staging, production, etc.)
- Step-by-step deployment procedure with checklist
- Rollback procedure and recovery strategy
- Environment configuration (secrets, variables, config files)
- Monitoring and alerts
- Known limitations and risks

**`docs/adr/`** — Architecture Decision Records:
- Decisions with rationale and consequences
- Indexed from `CONSTITUTION.md`
- Created using the ADR template when foundational decisions exist
- Updated when architectural decisions emerge

### Secondary Artifact

`docs/roadmap.md` (created from the template in the Appendix if it does not exist yet)

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

Document on testing practices for this project. Updated when new test patterns emerge or when constitution-level testing rules change.

---

## Test Types and Scope

<!--
Define the types of tests used in this project and what is required for each change.
Example categories:
- **Unit tests** — test single functions or methods in isolation
- **Integration tests** — test components working together
- **Acceptance tests** — test user-visible behavior against job stories
- **Performance tests** — test latency, throughput, or resource constraints
- **Contract tests** — test API or data contracts
-->

---

## What Requires Tests

<!--
State clearly what _must_ have test coverage:
- All behavior changes must have a failing test first (Red → Green → Refactor)
- Bug fixes must have a regression test before the fix
- Refactoring must keep existing tests green
- Structural tidying (`[tidy]` subtasks) must not change observable behavior
-->

---

## Test Location and Naming Conventions

<!--
Example:
- Test files live next to production code or in a `tests/` directory
- Test file naming: `<module>_test.<ext>` or `test_<module>.<ext>`
- Test function naming: `test_<unit>_<scenario>` or `describe('<unit>', () => { it('...')
- Helper functions: `setup_<fixture>()`, `assert_<condition>()`
-->

---

## Running Tests

<!--
How do developers run tests locally?
Example:
```bash
npm test              # All tests
npm test -- --watch  # Watch mode
npm test -- <pattern> # Filter by pattern
```

Or for other languages:
```bash
go test ./...
pytest
python -m unittest discover
```
-->

---

## Test Gate Before Promote

<!--
What must be green before code is promoted to permanent docs?
This is the _constitution-level gate_.
Example:
- All unit and integration tests must pass
- Code coverage must be ≥ 80% for changed files
- No flaky tests
- Performance tests must not regress
-->

---

## Continuous Integration / Testing Automation

<!--
Is there a CI/CD pipeline? What does it test?
Example:
- GitHub Actions / GitLab CI / Jenkins runs tests on every PR
- Coverage reports generated and must meet threshold
- Performance benchmarks tracked
- Linting and static analysis gates the merge
-->

---

## Known Test Gaps or Limitations

<!--
Document any areas where testing is incomplete or infeasible.
Example:
- Real-time features are not tested because [reason]
- External service integrations use mocks, not real endpoints
- UI tests are manual because [reason]
- Performance testing only covers [scenario], not [scenario]
-->

---

## Test Maintenance and Flakiness

<!--
How do you keep tests reliable?
Example:
- Flaky tests are tracked in [location] and fixed before merge
- Tests with external dependencies use retries with [strategy]
- Test data is seeded from [source], reset after each test
- Slow tests are isolated and run separately
-->
```

### Template: docs/deployment.md

```markdown
# Deployment

Document on deployment practices and runbooks for this project. Updated whenever deployment procedure, target environment, or rollback strategy changes.

---

## Deployment Overview

<!--
One-sentence summary of what gets deployed and where.
Example:
- This is a Node.js backend service deployed to AWS Lambda
- This is a React frontend deployed to Vercel
- This is a Go CLI tool distributed via Homebrew and GitHub Releases
-->

---

## Release Triggers and Versioning

<!--
How is a release triggered? How is it versioned?
Example:
- Manual: Tag a commit with `v<major>.<minor>.<patch>` and push to origin; CI builds and publishes
- Automatic: Merge to `main` triggers a release with semantic versioning based on conventional commits
- Versioning scheme: Semantic Versioning (SemVer) for libraries, CalVer for applications
-->

---

## Deployment Target(s)

<!--
Where does code run in production?
Example:
- Production: `https://api.example.com` (AWS ECS, us-east-1)
- Staging: `https://staging-api.example.com` (AWS ECS, us-east-1)
- Development: Local development environment only

Or for CLI:
- macOS: Homebrew tap `example/tap/tool`
- Linux: GitHub Releases, apt repository
- Windows: GitHub Releases, Scoop bucket
-->

---

## Deployment Procedure

<!--
Step-by-step runbook for deploying a release.
Example:

### Prerequisites
- [ ] All tests pass locally and in CI
- [ ] Code reviewed and merged to `main`
- [ ] Tag pushed with `git tag -a v<version> -m "Release <version>" && git push origin v<version>`

### Deploy Steps
1. CI pipeline is triggered by the tag push
2. Build artifacts are created: `dist/` for frontend, Docker image for backend
3. Artifact is published to target registry (npm, Docker Hub, GitHub Releases)
4. If deployment to production is automatic:
   - ECS task definition is updated with new image tag
   - Service is updated and old tasks are drained gracefully (30s drain timeout)
   - Health checks pass before considering deployment complete
5. If manual approval is needed:
   - Ops team receives notification in [Slack/email/deployment dashboard]
   - Approval triggers the above steps

### Verification
- [ ] Health checks pass on target environment
- [ ] Smoke tests pass (e.g., `curl https://api.example.com/health`)
- [ ] Logs show no errors
- [ ] Key metrics (latency, error rate) are nominal
-->

---

## Rollback Procedure

<!--
How do you undo a bad deployment?
Example:
- Automatic rollback: If health checks fail, ECS automatically reverts to previous task definition
- Manual rollback: `git revert <commit>`, tag with `v<version>-hotfix.1`, push; CI redeploys
- Database migrations: Forward-only; data rollback requires [process]
- Feature flags: Bad behavior can be disabled without redeployment via [system]

For CLI releases:
- Yanked versions: Tag with `v<version>` and mark as yanked in release notes
- Users on old version: Keep supporting previous major version for [N] months
-->

---

## Deployment Checklist

<!--
Copy and use before each deployment:

- [ ] Code committed and pushed
- [ ] All tests pass in CI
- [ ] Code reviewed
- [ ] Changelog updated (docs/CHANGELOG.md or RELEASES.md)
- [ ] Version bumped and tagged
- [ ] Staging deployment succeeds
- [ ] Smoke tests pass on staging
- [ ] Approval given for production deployment
- [ ] Production deployment succeeds
- [ ] Health checks pass
- [ ] Metrics are nominal
- [ ] Announcement posted to [team channel]
-->

---

## Environment Configuration

<!--
How are environment variables, secrets, and configuration managed?
Example:
- Secrets are stored in [AWS Secrets Manager / HashiCorp Vault / GitHub Secrets]
- Environment variables are injected at deployment time from [source]
- Configuration file: `.env.production` (not checked in), managed by [process]
- Database connection string: Retrieved from [secrets manager] at startup
-->

---

## Monitoring and Alerts

<!--
What happens after code is deployed? How do you know if it's broken?
Example:
- Error rates are monitored in Datadog; alert if error rate > 1% for 5 min
- Latency p99 is tracked; alert if > [threshold]
- Database connection pool is monitored; alert if exhausted
- Disk space is monitored; alert if < 10% free
- Deployment notifications sent to [Slack channel]
-->

---

## Known Deployment Limitations or Risks

<!--
Document any deployment constraints or gotchas.
Example:
- Database migrations are applied separately; code must be backward-compatible
- Deployment is not atomic: old and new code may run simultaneously for [duration]
- Secrets rotation requires [manual step]
- Large deployments > 100MB take [N] minutes; monitor for timeout
-->
```

### Template: docs/adr.md

When creating a new Architecture Decision Record, use this template:

```markdown
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
- Link it from `docs/ARCHITECTURE.md` or an index
- Update any related ADRs with "Supersedes" or "Related Decisions" links
- If a related decision is reversed, update the old ADR's status to "Superseded"
```

### Template: docs/roadmap.md

```markdown
# Roadmap

Living product roadmap. Each section is one delivered or planned increment.

> A feature moves to **Done** only when its acceptance tests pass and are linked here.
> Source of truth: if a feature is not in Done with a passing test link, it is not considered shipped.

---

## Done

<!--
Each entry follows this pattern:

### [Feature name — short, user-visible]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Acceptance tests:** [path/to/test_file.ext#test_name](path/to/test_file.ext)
- **Use case:** [docs/usecases/use-case-slug.md](docs/usecases/use-case-slug.md) *(if promoted)*
- **Delivered:** [increment slug or YYYY-MM-DD]
-->

---

## Partial

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Acceptance tests:** pending — being written this cycle
- **Increment:** [increment slug]
-->

---

## Planned

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Notes:** [optional — dependencies, open questions, or ordering rationale]
-->

---

## Rules

- Features move left to right: Planned → Partial → Done. Never skip Partial.
- A feature enters Partial when its increment is approved.
- A feature enters Done only when its acceptance test link resolves to a passing test.
- Do not add implementation detail here — link to the use case or ADR for that.
- If a planned feature is no longer needed, remove it and note the removal in `learnings.md` for that cycle.
```