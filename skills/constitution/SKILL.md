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

`CONSTITUTION.md` containing:
1. Engineering principles grounded in XP, Tidy First, and lean software development
2. Architectural boundaries and dependency direction
3. **Testing strategy** — test types in scope (unit / integration / e2e), what must have tests, the gate that must be green before promote, naming conventions, Red→Green→Refactor discipline
4. **Release and deployment** — how a release is triggered (manual, CI, tag), versioning scheme, deployment target(s) and steps, rollback procedure
5. Tidy First policy: when to tidy before changing behavior, what constitutes a valid tidying
6. Documentation rules and ADR policy
7. SDLC artifact expectations: the `.agent/` contract (which files, what lifecycle)
8. AI collaboration rules: constrain context per task, preserve optionality, maintain human judgment over architectural decisions

`docs/roadmap.md` (created from `templates/roadmap.md` if it does not exist yet)

---

<HARD-GATE>
Do NOT write CONSTITUTION.md until the HTML review in `.agent/constitution-review.html` has been explicitly approved.
Do NOT ask more than 5 questions per round.
Do NOT include implementation details — CONSTITUTION.md contains guardrails, not recipes.
Do NOT copy generic principles from the internet. Every rule must be justified by this project's specific context.
</HARD-GATE>

---

## Process

1. **Read project context** — scan `README.md`, existing `CONSTITUTION.md`, directory structure, any ADRs or docs
2. **Ask 3–5 focused questions** — surface constraints, pain points, and non-negotiables one round at a time
3. **Generate `.agent/constitution-review.html`** — present proposed guardrails in review format
4. **STOP** — wait for explicit approval or revision requests
5. **On approval** — write `CONSTITUTION.md`

---

## Checklist

- [ ] Existing docs read
- [ ] 3–5 questions asked and answered — include: test strategy, deployment target, versioning scheme
- [ ] HTML review generated and shown
- [ ] User approval received
- [ ] `CONSTITUTION.md` written with Testing and Release sections populated
- [ ] `docs/roadmap.md` created if not present

---

## Handoff

Terminal artifacts: `CONSTITUTION.md`, `docs/roadmap.md`
Next skill: `4dc-increment` — load `skills/increment/SKILL.md`
