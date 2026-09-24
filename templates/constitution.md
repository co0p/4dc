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

{{FOUNDATION:beck-team-agreements}}
{{FOUNDATION:fowler-evolutionary-architecture}}
{{FOUNDATION:deming-systems-thinking}}
{{FOUNDATION:humble-continuous-delivery}}

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
- Must contain a **C4 Level 1 System Context** section (system purpose, users/actors, external systems)
- Must contain a **C4 Level 2 Container** view (runtime containers, responsibilities, communication paths, data stores)
- Includes a **Container Internals (C4 Level 3)** section, present as an on-demand slot even when empty; populated when a Refactor increment promotes internal structure

**`docs/domain.md`** — Domain glossary:
- Required for every project, even when the vocabulary is initially small
- Defines shared concepts, domain events, and system rules in business language
- Must not be replaced by an ADR, README, or implementation-specific notes

**`docs/ui.md`** — Permanent UI decisions (required when the project has a user interface):
- Shared user flows, interaction patterns, visual principles, accessibility rules, and content conventions
- Rationale and consequences of recurring UI decisions
- No component inventory, CSS catalog, or one-off screen notes

### Secondary Artifact

*(none)*

### Documentation Baseline

Before creating or updating the constitution, audit the repository for the complete permanent documentation baseline:

- `CONSTITUTION.md`
- `docs/testing.md`
- `docs/deployment.md`
- `docs/observability.md`
- `docs/architecture.md` with a C4 Level 1 System Context view and a C4 Level 2 Container view
- `docs/domain.md` with the project's glossary
- `docs/ui.md` with the project's UI decisions (if the system has a UI; omit for headless systems)
- `docs/adr/`

Missing baseline documents are constitution outputs; they are not optional follow-up work. Existing documents must be checked for the required content rather than accepted solely because the path exists.

{{SHARED:execution-contract}}

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
3. **Deployment conversation using the Concept Menu** — when the release and deployment category is reached, walk the user through the ten deployment axes below. Present each axis with its options, ask one focused question, and record the chosen option as a guardrail sentence in `docs/deployment.md`. Do not write the concept name (e.g. "continuous", "blue-green") into the durable document — only the guardrail sentence derived from the chosen option.
4. **Boundary review** — present the proposed constitution rules separately from the project-specific documentation updates. Flag every concrete fact and its `docs/` destination. Wait for explicit approval of both lists.
5. **On approval:**
   - Write `CONSTITUTION.md` with references to supporting documents
   - Create `docs/testing.md` with project-specific testing practices
   - Create `docs/deployment.md` with project-specific deployment procedures, section by section, using the answers from the Concept Menu
   - Create `docs/observability.md` with project-specific operational signals and response guidance
    - Create or update `docs/architecture.md` with a current C4 Level 1 System Context view and a C4 Level 2 Container view; add the Container Internals (Level 3) section slot for on-demand use
    - Create `docs/domain.md` with the project's initial glossary, even if only a few concepts are known
    - Create `docs/ui.md` with the project's initial UI decisions when the system has a user interface
   - Create initial `docs/adr/` structure if foundational decisions exist

### Deployment Concept Menu

Use during step 3. Present each axis with its options; the user picks one per axis. The chosen option becomes a guardrail sentence in the corresponding section of `docs/deployment.md`. Concept names are conversation aids only — do not write them into the durable document.

1. **Release cadence** — Continuous / Batched / On-demand
2. **Release trigger** — Automated event / Manual action / Hybrid
3. **Environment progression** — Direct to production / One pre-production environment / Multiple environments with promotion gates
4. **Verification depth** — Automated only / Automated plus manual gate / Automated plus operator sign-off plus staged rollout
5. **Failure response** — Roll back / Roll forward / Both allowed
6. **Deployment safety model** — Atomic / Rolling / Blue-green
7. **Configuration boundary** — Injected at deploy time / Baked into the artifact / Hybrid
8. **Secret handling** — Managed secret store / Operator-injected environment variables / Retrieved at startup
9. **Observability trigger** — Health signal drives release confidence / Time-based confidence / Manual confirmation
10. **Update policy trigger** — On every release model change / On every material incident / Both

Each axis maps to one section of `docs/deployment.md`. The scaffold at `templates/deployment.md` names the sections and states the durable guardrail; the user's choice fills in the specifics.

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
- [ ] `docs/deployment.md` created with project-specific procedures; every deployment axis (release cadence, release trigger, environment progression, verification depth, failure response, deployment safety model, configuration boundary, secret handling, observability trigger, update policy trigger) has an answer recorded as a guardrail sentence
- [ ] `docs/observability.md` created with project-specific signals and operating guidance
- [ ] `docs/architecture.md` created with both required diagrams populated:
  - [ ] C4 Level 1 System Context — system purpose, users/actors, external systems
  - [ ] C4 Level 2 Container view — containers, communication paths, data stores
  - [ ] Container Internals (C4 Level 3) section present but may be empty (populated on demand by Refactor increments)
- [ ] `docs/domain.md` created with the project's initial glossary (concepts, events, rules in domain language) — not just an empty file
- [ ] `docs/ui.md` created with the project's initial UI decisions **if the system has a UI**; explicitly recorded as "not applicable — headless system" otherwise
- [ ] `docs/adr/` directory created with index link from `CONSTITUTION.md` (populate with foundational decisions if identified)
- [ ] Each baseline doc satisfies its content requirement, not merely file existence

---

## Handoff

Terminal artifacts:
- `CONSTITUTION.md` — guardrails and governance
- `docs/testing.md` — testing procedures and practices
- `docs/deployment.md` — deployment and release procedures
- `docs/observability.md` — signals, alerts, health checks, and operating guidance
- `docs/adr/` structure — architectural decisions

Future work should reference the testing, deployment, and architecture documents defined by the project constitution. New ADRs should be added when architectural decisions emerge.

After this skill, continue with the next approved work item using the applicable project workflow.

---

## Appendix: Document Templates

Use these verbatim as the starting content when creating a new document for the first time.

### Template: docs/testing.md

```markdown
{{TEMPLATE:testing}}
```

### Template: docs/deployment.md

```markdown
{{TEMPLATE:deployment}}
```

### Template: docs/observability.md

```markdown
{{TEMPLATE:observability}}
```

