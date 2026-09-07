---
name: 4dc-promote
description: "Use after implementation.md is marked complete. Reviews all .agent/ artifacts, proposes promotions to permanent docs, and closes the cycle."
---

# Promote Skill

## One Responsibility

Merge durable outcomes from the `.agent/` working set into permanent project artifacts before the branch merges.

---

## Expected Input

- `CONSTITUTION.md`
- `.agent/increment.md`
- `.agent/plan.md`
- `.agent/implementation.md` (status: complete)
- `.agent/learnings.md`
- Existing `docs/`, ADR log, and any project documentation

---

## Concrete Output

One or more of the following, per approval:
- Updated `CONSTITUTION.md` (if guardrails need revision)
- New ADR in `docs/adr/` (for significant architectural decisions)
- Updated `docs/architecture.md` (if runtime structure, dependencies, or performance-critical paths changed)
- Updated `docs/domain.md` (if domain language changed or new concepts appeared)
- Updated `README.md` or other docs (for changed behavior or usage)
- Updated `docs/roadmap.md` — feature moved from Partial to Done, acceptance test link added
- Deleted or archived `.agent/` files after promotion (keeping `.agent/` clean for next cycle)

Required review outputs:
- Promotion candidates are listed individually with destination path, rationale, and approval status.
- The promotion review explicitly states whether architecture, domain language, testing guidance, and performance documentation changed or stayed unchanged.

### Permanent Documentation Baseline

Every promotion must verify that the project's permanent documentation baseline exists and is usable:

- `CONSTITUTION.md`
- `docs/testing.md`
- `docs/deployment.md`
- `docs/architecture.md` containing a current C4 Level 2 container view (or an explicitly labeled equivalent)
- `docs/domain.md` containing the current domain glossary
- `docs/adr/`
- `docs/roadmap.md`

The check is semantic, not just a file-existence check. A generic architecture narrative does not satisfy the C4 requirement, and a glossary hidden in an ADR or README does not satisfy the domain-document requirement. Missing or inadequate documents become promotion candidates and must be created or corrected before the cycle can close.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT write permanent docs until each promotion candidate has been individually approved.
Do NOT promote guesses or plans — only promote what was actually built and verified.
Do NOT delete .agent/ files until all promotions are written and confirmed.
Present each candidate separately with destination path and rationale.
Do NOT leave permanent docs stale when the implementation changed architecture, domain language, or performance-critical behavior.
Do NOT close promotion while any required permanent documentation baseline item is missing or inadequate. If runtime structure and domain vocabulary are unchanged, still verify that `docs/architecture.md` and `docs/domain.md` exist and satisfy their C4 and glossary requirements.
</HARD-GATE>

---

## Process

1. **Read all `.agent/` artifacts** — full review of increment, plan, implementation, and learnings.
2. **Audit the permanent documentation baseline** — inspect each required path and verify the architecture document contains a C4 Level 2 container view and the domain document contains the glossary. Add missing or inadequate documents to the candidate list.
3. **Conversation: Propose promotions** — identify candidates and state what each is, its destination, and why it is durable. Iterate until the user says to proceed.
4. **Generate `.agent/promotion-review.md`** — include the required Markdown review sections and an individual approval checkbox for every candidate.
5. **STOP** — present the review and wait for explicit per-candidate approval.
6. **On approval** — write each approved permanent artifact.
7. **Re-audit the baseline** — confirm every required document exists and satisfies its content requirement before cleanup.
8. **Clean up** — archive or delete `.agent/` files for this cycle.

## Markdown Review Contract

Use `.agent/promotion-review.md`. Include **Objective**, **Inputs Reviewed**, **Proposed Output Summary**, **Promotion Candidates** (type, destination, rationale, approval), **Risks and Trade-offs**, **Open Questions**, and **Approval Decision**. Record each explicit conversational approval in the review.

---

## Promotion Categories

| Type | Trigger | Destination |
|------|---------|-------------|
| Architecture decision | Non-obvious choice with lasting impact | `docs/adr/ADR-<date>-<slug>.md` |
| Guardrail update | Constitution rule violated, needs clarification | `CONSTITUTION.md` |
| Architecture sync | Runtime containers, dependency direction, or performance-critical paths changed | `docs/architecture.md` |
| Behavior change | Public API, CLI, or user-facing behavior changed | `README.md` |
| Feature shipped | Acceptance tests pass; feature complete | `docs/roadmap.md` — move to Done, add acceptance test link |
| Test pattern | New testing approach worth standardizing | `CONSTITUTION.md` testing section |
| Performance contract | A latency, throughput, cost, or scaling expectation changed | `CONSTITUTION.md` or `docs/architecture.md` |
| Known issue | Found but not fixed this cycle | `docs/known-issues.md` |
| New domain concept | A concept, event, or rule used in code/tests that has no shared definition | `docs/domain.md` (create using the template in the Appendix if absent) |
| Structural change | A container added, removed, or re-wired | `docs/architecture.md` (create using the template in the Appendix if absent) |

---

## Checklist

- [ ] All `.agent/` artifacts read
- [ ] Promotion candidates identified and categorized
- [ ] Permanent documentation baseline audited for existence and required content
- [ ] Markdown review generated covering all candidates
- [ ] User approval received per candidate
- [ ] Each approved artifact written to permanent location
- [ ] Final baseline audit passes: glossary and C4 architecture view are present and current
- [ ] Architecture, domain language, testing guidance, and performance documentation either updated or explicitly marked unchanged
- [ ] `.agent/` files cleaned up

---

## Handoff

Terminal artifacts: permanent docs updated, `.agent/` clean
Cycle complete. Next action: `4dc-increment` for the next cycle — load `skills/increment/SKILL.md`

---

## Appendix: Document Templates

Use these verbatim as the starting content when creating a new document for the first time.

### Template: docs/domain.md

```markdown
{{TEMPLATE:domain}}
```

### Template: docs/architecture.md

```markdown
{{TEMPLATE:architecture}}
```
