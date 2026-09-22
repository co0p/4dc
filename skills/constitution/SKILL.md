---
name: 4dc-constitution
description: "Use when CONSTITUTION.md is missing or needs updating. Reads project context, asks focused questions, and produces project guardrails and SDLC standards."
---

# Constitution Skill

## One Responsibility

Create or update `CONSTITUTION.md` — the project's durable engineering guardrails.

The constitution states **principles and boundaries**, not implementations. It answers: what rules do we agree to operate under, what boundaries must not be crossed, and where does each category of concrete detail live? It does not prescribe tools, libraries, frameworks, file names, commands, or configuration — those belong in `docs/` or ADRs.

The generated constitution must describe only the application. Do not include this repository's internal workflow name, phase sequence, agent names, or transient artifact paths.

Before writing it, check that it contains no internal workflow names, skill names, phase names, orchestrator terms, `.agent/` paths, or `.agents/` paths.

---

## Foundations

- **Kent Beck: team agreements before code.** Establish explicit rules before implementation so disagreements are resolved by shared constraints rather than assumptions.
- **Martin Fowler: evolutionary architecture.** Define durable boundaries and fitness constraints while allowing implementation details to evolve.
- **W. Edwards Deming: systems thinking.** Optimize and verify the whole delivery system rather than treating local activity as proof of value.
- **Jez Humble: continuous delivery.** Treat releasability, deployment safety, and production verification as properties of every change.

**What belongs in CONSTITUTION.md (guardrail categories only):**
- Architecture: general dependency, coupling, and system-boundary rules
- Testing: required confidence levels by risk and the minimum release gate
- Observability: which behavior changes require operational signals and the qualities those signals must have
- Security and privacy: general handling, least-privilege, and review boundaries
- Reliability and resilience: failure, recovery, compatibility, and data-integrity expectations
- Performance and efficiency: when budgets are required and how regressions are treated
- Delivery and deployment: releasability, rollback or recovery, migration safety, and production verification rules
- Documentation and ADR governance: where concrete knowledge lives and which decisions require records

**What does NOT belong in CONSTITUTION.md (belongs in `docs/` or ADRs):**
- Test commands, CI scripts, tooling configuration
- Deployment runbooks, environment lists, secret-handling procedures
- Framework choices, library names, file naming conventions
- Coverage numbers, specific thresholds, or tool-specific configuration
- Architecture diagrams or container inventories
- Domain rules, user journeys, endpoint names, event names, concrete performance budgets, environment names, or product examples
- Project rationale, current topology, current risks, or any fact likely to change as the product evolves

---

## Expected Input

- Existing `CONSTITUTION.md` (if present)
- `README.md`
- Current project structure and any existing docs

---

## Concrete Output

### Primary Artifact: `CONSTITUTION.md`

`CONSTITUTION.md` containing only principles and boundaries — no commands, tooling, or concrete procedures:

1. **Engineering principles** — durable decision rules, not project history or methodology exposition.
2. **Architecture boundaries** — generic dependency and coupling rules. Current containers and communication paths live in `docs/architecture.md`.
3. **Testing and quality** — required test depth by risk and the release gate. Commands, tools, suites, and concrete thresholds live in `docs/testing.md`.
4. **Observability** — when operational signals are required and what qualities they must preserve. Concrete signals and alerting live in `docs/observability.md`.
5. **Security and privacy** — review and handling guardrails. Threat models, classifications, controls, and procedures live in project docs or ADRs.
6. **Reliability, performance, and efficiency** — general expectations for failure handling, data integrity, measurable budgets, and regression evidence. Concrete budgets and topology live in docs.
7. **Documentation and ADR policy** — what is permanent, where concrete project knowledge lives, and what decisions trigger an ADR.
8. **Release and deployment** — general releasability, migration, recovery, and production-verification rules. Procedures live in `docs/deployment.md`.

Required `CONSTITUTION.md` headings:
- `## Engineering Principles`
- `## Architecture Boundaries`
- `## Testing And Quality`
- `## Observability`
- `## Security And Privacy`
- `## Reliability, Performance, And Efficiency`
- `## Documentation And ADR Policy`
- `## Release And Deployment`

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

**`docs/observability.md`** — operational visibility for this project:
- Concrete logs, metrics, traces, audit events, health signals, dashboards, and alert ownership
- Signal names and fields, including privacy and cardinality constraints
- Expected operator response and known blind spots

**`docs/adr/`** — Architecture Decision Records:
- Decisions with rationale and consequences
- The `docs/adr/` directory is the index; do not duplicate that list in `CONSTITUTION.md`
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

**`docs/ui.md`** — Permanent UI decisions (required when the project has a user interface):
- Shared user flows, interaction patterns, visual principles, accessibility rules, and content conventions
- Rationale and consequences of recurring UI decisions
- No component inventory, CSS catalog, or one-off screen notes

### Secondary Artifact

`docs/roadmap.md` (created from the template in the Appendix if it does not exist yet)

### Documentation Baseline

Before creating or updating the constitution, audit the repository for the complete permanent documentation baseline:

- `CONSTITUTION.md`
- `docs/testing.md`
- `docs/deployment.md`
- `docs/observability.md`
- `docs/architecture.md` with a C4 Level 2 container view
- `docs/domain.md` with the project's glossary
- `docs/ui.md` with the project's UI decisions (if the system has a UI; omit for headless systems)
- `docs/adr/`
- `docs/roadmap.md`

Missing baseline documents are constitution outputs; they are not optional follow-up work. Existing documents must be checked for the required content rather than accepted solely because the path exists.

## Language and Interaction Rules

- Use plain, direct language. Keep output scannable.
- Prefer short sentences and bullets.
- State decisions, actions, blockers, and evidence. Omit motivational, decorative, and generic advice.
- Do not repeat the user's request, loaded instructions, or handoff contents.
- Explain a choice only when it affects scope, risk, verification, or the next handoff.
- Ask one focused question at a time. Do not use broad questionnaires.
- State assumptions explicitly. If required evidence is missing or contradictory, ask rather than inventing an answer.
- Use the project's domain language in product artifacts. Keep internal workflow and agent terminology out of permanent product documentation.
- Refer to files, symbols, commands, states, and evidence precisely. Avoid vague terms such as "works", "correct", or "should be fine".
- End a phase response with the decision needed, the blocker, or the next handoff. During approved autonomous implementation, continue without routine confirmation.

## Execution Contract

- Never copy internal workflow names, skill names, phase names, orchestrator terms, `.agent/` paths, or `.agents/` paths into permanent product artifacts.
- Before writing a permanent artifact, scan it for internal workflow references and remove them.
- Produce only the artifact for this phase. Do not leak work from a later phase into this one.
- Treat tests, architecture notes, ADRs, and user-facing docs as first-class communication artifacts.
- Gather only enough context to identify the governing constraints, the target artifact, and the cheapest validation step. Then act.
- Resolve conflicts in this order: explicit user approval, approved prior-phase artifacts, `CONSTITUTION.md`, this skill.
- Low-risk actions: reads, searches, diffs, and local validation commands.
- Medium-risk actions: local reversible edits to phase artifacts.
- High-risk actions: destructive file operations, external side effects, or skipping a stop gate. Require explicit approval first.
- If a required input is missing or contradictory, ask one focused question or stop and wait for explicit approval. Do not invent missing facts.
- Before finishing, run the phase checklist and confirm every required section is present.

---

<HARD-GATE>
Do NOT write `CONSTITUTION.md` until the user explicitly approves the proposed guardrails.
Do NOT ask more than one focused question at a time. Continue until every required guardrail category has an explicit decision or explicit deferral.
Do NOT include project specifics — CONSTITUTION.md contains guardrails, not product facts or recipes. Except for pointers to governed documents, if a sentence contains a tool name, command, file path, framework, library, configuration value, domain term, component name, environment, concrete threshold, or example, it belongs in `docs/` or an ADR.
Do NOT justify rules with project-specific examples in `CONSTITUTION.md`; preserve that rationale in the supporting document or ADR.
Do NOT add a `## Delivery and Documentation` section — documentation policy belongs under `## Documentation And ADR Policy`.
</HARD-GATE>

---

## Process

1. **Read project context** — scan `README.md`, existing `CONSTITUTION.md`, directory structure, ADRs, and current SDLC documentation. Use specifics to identify decisions, but route those specifics to `docs/`.
2. **Conversation: Decide guardrails one category at a time** — architecture; testing and quality; observability; security and privacy; reliability; performance and efficiency; documentation and ADRs; release and deployment. Ask one focused question, summarize the decision, then continue. Do not silently choose defaults when evidence is missing.
3. **Boundary review** — present the proposed constitution rules separately from the project-specific documentation updates. Flag every concrete fact and its `docs/` destination. Wait for explicit approval of both lists.
4. **On approval:**
   - Write `CONSTITUTION.md` with references to supporting documents
   - Create `docs/testing.md` with project-specific testing practices
   - Create `docs/deployment.md` with project-specific deployment procedures
   - Create `docs/observability.md` with project-specific operational signals and response guidance
    - Create or update `docs/architecture.md` with the current C4 Level 2 container view
    - Create `docs/domain.md` with the project's initial glossary, even if only a few concepts are known
    - Create `docs/ui.md` with the project's initial UI decisions when the system has a user interface
   - Create initial `docs/adr/` structure if foundational decisions exist
   - Create `docs/roadmap.md` if not present

---

## Checklist

- [ ] Existing docs read (including any deployment or testing practices)
- [ ] Foundational ADRs identified (if any exist)
- [ ] Every required guardrail category explicitly decided or deferred through focused questions
- [ ] Proposed constitution contains no project-specific facts or examples
- [ ] Project-specific decisions have a named `docs/` or ADR destination
- [ ] User approval received
- [ ] `CONSTITUTION.md` written with all required guardrail headings populated
- [ ] `docs/testing.md` created with project-specific practices
- [ ] `docs/deployment.md` created with project-specific procedures
- [ ] `docs/observability.md` created with project-specific signals and operating guidance
- [ ] `docs/adr/` directory created with index link from `CONSTITUTION.md` (populate with foundational decisions if identified)
- [ ] `docs/roadmap.md` created if not present

---

## Handoff

Terminal artifacts:
- `CONSTITUTION.md` — guardrails and governance
- `docs/testing.md` — testing procedures and practices
- `docs/deployment.md` — deployment and release procedures
- `docs/observability.md` — signals, alerts, health checks, and operating guidance
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

## Evidence Required Before Merge

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

### Template: docs/observability.md

```markdown
# Observability

Guide to understanding the system in operation. Keep concrete signals, ownership, interpretation, and response guidance here rather than in the constitution.

## Operational Outcomes

Describe the user-visible and operator-visible outcomes that must be distinguishable in production, including success, expected rejection, degradation, and failure.

## Signals

Document the logs, metrics, traces, audit records, and health checks used by this project. For each signal, state its purpose, important fields or dimensions, privacy and cardinality constraints, and how to interpret it.

## Alerts And Response

Document actionable alert conditions, ownership, first response, verification, and escalation or recovery guidance. Avoid alerts without an expected operator action.

## Release Verification

Describe the signals used to confirm a release is healthy and the conditions that trigger rollback, roll-forward, or further investigation.

## Known Blind Spots

Record material observability gaps, their risk, and the evidence that would justify improving them.

## Maintenance

Update this guide when behavior, architecture, data sensitivity, operational ownership, or failure modes change. Remove stale signals rather than preserving an inventory of obsolete telemetry.
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

## In Progress

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Evidence:** pending — define the verification approach before work starts
-->

---

## Planned

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Why now / ordering:** [dependencies, user value, open questions, or sequencing rationale]
-->

---

## How This List Works

- Features move left to right: Planned → In Progress → Done. Never skip In Progress.
- A feature enters In Progress when work begins.
- A feature enters Done only when its user outcome is verified and the evidence link is present.
- Do not add implementation detail here — link to the use case or ADR for that.
- If a planned feature is no longer needed, remove it and record the reason in a code comment, commit message, or ADR.
```