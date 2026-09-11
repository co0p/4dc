---
name: 4dc-constitution
description: "Use when CONSTITUTION.md is missing or needs updating. Reads project context, asks focused questions, and produces project guardrails and SDLC standards."
---

# Constitution Skill

## One Responsibility

Create or update `CONSTITUTION.md` — the project's durable engineering guardrails.

The generated constitution must describe only the application. Do not include this repository's internal workflow name, phase sequence, agent names, or transient artifact paths.

Before writing it, check that it contains no internal workflow names, skill names, phase names, orchestrator terms, `.agent/` paths, or `.agents/` paths.

---

## Foundations

- **Beck — team agreements before code.** The constitution is the set of rules the team agrees to operate under. It is written before implementation, not retrofitted after.
- **Poppendieck — eliminate ambiguity upstream.** Decide the guardrails early so later phases do not rediscover the same constraints. For reversible choices, leave the decision late; for structural rules, fix them now.
- **Fowler — evolutionary architecture.** Guardrails, not blueprints. The constitution sets the boundaries that let the design evolve safely, not a fixed architecture that must be followed verbatim.

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
- `docs/architecture.md` with a C4 Level 2 container view
- `docs/domain.md` with the project's glossary
- `docs/ui.md` with the project's UI decisions (if the system has a UI; omit for headless systems)
- `docs/adr/`
- `docs/roadmap.md`

Missing baseline documents are constitution outputs; they are not optional follow-up work. Existing documents must be checked for the required content rather than accepted solely because the path exists.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT write `CONSTITUTION.md` until the user explicitly approves the proposed guardrails.
Do NOT ask more than 5 questions per round.
Do NOT include implementation details — CONSTITUTION.md contains guardrails, not recipes.
Do NOT copy generic principles from the internet. Every rule must be justified by this project's specific context.
</HARD-GATE>

---

## Process

1. **Read project context** — scan `README.md`, existing `CONSTITUTION.md`, directory structure, any ADRs or docs, and existing deployment or testing practices
2. **Conversation: Propose the guardrails** — summarize the proposed testing and deployment strategies, identify foundational ADR candidates, and ask whether the direction feels right. Iterate until the user says “looks good” or “proceed.”
3. **On approval:**
   - Write `CONSTITUTION.md` with references to supporting documents
   - Create `docs/testing.md` with project-specific testing practices
   - Create `docs/deployment.md` with project-specific deployment procedures
    - Create or update `docs/architecture.md` with the current C4 Level 2 container view
    - Create `docs/domain.md` with the project's initial glossary, even if only a few concepts are known
    - Create `docs/ui.md` with the project's initial UI decisions when the system has a user interface
   - Create initial `docs/adr/` structure if foundational decisions exist
   - Create `docs/roadmap.md` if not present

---

## Checklist

- [ ] Existing docs read (including any deployment or testing practices)
- [ ] Foundational ADRs identified (if any exist)
- [ ] Proposed testing strategy, deployment strategy, and ADR candidates discussed
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
{{TEMPLATE:testing}}
```

### Template: docs/deployment.md

```markdown
{{TEMPLATE:deployment}}
```

### Template: docs/adr.md

When creating a new Architecture Decision Record, use this template:

```markdown
{{TEMPLATE:adr}}
```

### Template: docs/roadmap.md

```markdown
{{TEMPLATE:roadmap}}
```
