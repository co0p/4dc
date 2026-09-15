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

These three traditions are the source for guardrails. Borrow from them by name — attribute the rule to its author so the team knows why it exists, not just what it says.

- **Beck — team agreements before code.** The constitution is the set of rules the team agrees to operate under. It is written before implementation, not retrofitted after. Rules must be specific enough to resolve disputes; vague principles are not guardrails.
- **Poppendieck — eliminate ambiguity upstream.** Decide the guardrails early so later phases do not rediscover the same constraints. For reversible choices, leave the decision late; for structural rules, fix them now. Pull rules from value; do not mandate process that does not serve delivery.
- **Fowler — evolutionary architecture.** Guardrails, not blueprints. The constitution sets the fitness functions and boundaries that let the design evolve safely. It does not fix the implementation — it defines what must remain true as the implementation changes.

**What belongs in CONSTITUTION.md (principles and boundaries):**
- Engineering principles attributed to their source (Beck, Poppendieck, Fowler, or a project-specific decision)
- Architectural boundaries: which direction dependencies flow, which containers must stay decoupled, what crosses the system boundary
- Testing strategy: what must have tests, what constitutes a green gate before promote, and what kinds of tests are in or out of scope — as rules, not commands
- Performance envelope: latency, throughput, cost, or scale expectations as stated constraints
- Documentation policy: what is permanent, what is transient, where each category lives
- ADR policy: what kinds of decisions require an ADR

**What does NOT belong in CONSTITUTION.md (belongs in `docs/` or ADRs):**
- Test commands, CI scripts, tooling configuration
- Deployment runbooks, environment lists, secret-handling procedures
- Framework choices, library names, file naming conventions
- Coverage numbers, specific thresholds, or tool-specific configuration
- Architecture diagrams or container inventories

---

## Expected Input

- Existing `CONSTITUTION.md` (if present)
- `README.md`
- Current project structure and any existing docs

---

## Concrete Output

### Primary Artifact: `CONSTITUTION.md`

`CONSTITUTION.md` containing only principles and boundaries — no commands, tooling, or concrete procedures:

1. **Engineering principles** — stated as rules, attributed to Beck, Poppendieck, Fowler, or a project decision. Each rule must be justified by this project's context, not copied from generic advice.
2. **Architectural boundaries** — dependency direction, which containers must remain decoupled, performance-critical paths stated as constraints. No container inventory (that lives in `docs/architecture.md`).
3. **Testing strategy** — what must have tests, what constitutes a green gate before promote, and what kinds of tests are in or out of scope — as rules. Points to `docs/testing.md` for procedures and commands.
4. **Performance envelope** — latency, throughput, cost, or scale expectations as stated constraints, or an explicit `N/A` with rationale. No monitoring configuration.
5. **Documentation and ADR policy** — what is permanent, where each category lives, and what decisions trigger an ADR.
6. **Release and deployment** — the release model as a rule (e.g. "every merge to main is releasable"). Points to `docs/deployment.md` for procedures.

Required `CONSTITUTION.md` headings:
- `## Engineering Principles`
- `## Architecture Boundaries`
- `## Testing Strategy`
- `## Performance Envelope`
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
Do NOT include implementation details — CONSTITUTION.md contains guardrails, not recipes. If a sentence contains a tool name, a command, a file path, a framework, a library, or a configuration value, it belongs in docs/ or an ADR, not in CONSTITUTION.md.
Do NOT copy generic principles from the internet. Every rule must be justified by this project's specific context.
Do NOT add a `## Delivery and Documentation` section — documentation policy belongs under `## Documentation And ADR Policy`.
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
