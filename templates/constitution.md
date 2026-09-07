---
name: 4dc-constitution
description: "Use when CONSTITUTION.md is missing or needs updating. Reads project context, asks focused questions, and produces project guardrails and SDLC standards."
---

# Constitution Skill

## One Responsibility

Create or update `CONSTITUTION.md` — the project's durable engineering guardrails.

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
7. SDLC artifact expectations: the `.agent/` contract (which files, what lifecycle)

Required `CONSTITUTION.md` headings:
- `## Engineering Principles`
- `## Architecture Boundaries`
- `## Testing Strategy`
- `## Performance Envelope`
- `## Documentation And ADR Policy`
- `## Release And Deployment`
- `## Artifact Lifecycle`

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
- Created using the ADR template during constitution phase if foundational decisions exist
- Updated during promote phases when architectural decisions emerge

### Secondary Artifact

`docs/roadmap.md` (created from the template in the Appendix if it does not exist yet)

{{SHARED:execution-contract}}

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

All future increments will reference the testing and deployment documents from their constitution. New ADRs are added during promote phases.

Next skill: `4dc-increment` — load `skills/increment/SKILL.md`

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
