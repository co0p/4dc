---
name: 4dc-promote
description: "Use after implementation.md is marked complete. Reviews all .agent/ artifacts, proposes promotions to permanent docs, and closes the cycle."
---

# Promote Skill

## One Responsibility

Promote durable outcomes to permanent project artifacts, run a final tidy pass on the branch, then land the increment as a squash-merge to `main`. Capture landing evidence, then delete transient artifacts.

---

## Foundations

{{FOUNDATION:poppendieck-eliminate-waste}}
{{FOUNDATION:cockburn-reflective-improvement}}
{{FOUNDATION:deming-systems-thinking}}
{{FOUNDATION:humble-continuous-delivery}}
{{FOUNDATION:farley-continuous-delivery}}

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
- Updated `docs/ui.md` (if shared UI, interaction, visual, accessibility, or content decisions changed)
- Updated `README.md` or other docs (for changed behavior or usage)
- Updated `docs/roadmap.md` — feature moved from In Progress to Done, acceptance test link added
- Acceptance-test evidence — linked for every acceptance criterion; exceptions require explicit prior approval and rationale
- Deleted `.agent/` files after landing evidence is verified (never archived; keeps `.agent/` clean for the next cycle)

Required outputs:
- Promotion candidates are listed individually with destination path, rationale, and approval status.
- The promotion explicitly states whether architecture, domain language, testing guidance, and performance documentation changed or stayed unchanged.

### Permanent Documentation Baseline

Every promotion must verify that the project's permanent documentation baseline exists and is usable:

- `CONSTITUTION.md`
- `docs/testing.md`
- `docs/deployment.md`
- `docs/observability.md` containing current operational signals, ownership, response guidance, and known blind spots
- `docs/architecture.md` containing a current C4 Level 1 System Context view and a C4 Level 2 Container view; a Container Internals (C4 Level 3) section is present, empty or populated
- `docs/domain.md` containing the current domain glossary
- `docs/ui.md` containing current UI decisions (if the system has a UI; omit for headless systems)
- `docs/adr/`
- `docs/roadmap.md`

The check is semantic, not just a file-existence check. A generic architecture narrative does not satisfy the C4 requirement, a glossary hidden in an ADR or README does not satisfy the domain-document requirement, and a deployment health checklist does not replace maintained observability guidance. Missing or inadequate documents become promotion candidates and must be created or corrected before the cycle can close.

{{SHARED:execution-contract}}

---

<HARD-GATE>
Do NOT write permanent docs until each promotion candidate has been individually approved.
Do NOT promote guesses or plans — only promote what was actually built and verified.
Do NOT delete .agent/ files until all promotions are written and confirmed.
Present each candidate separately with destination path and rationale.
Do NOT leave permanent docs stale when the implementation changed architecture, domain language, or performance-critical behavior.
Do NOT close promotion while any required permanent documentation baseline item is missing or inadequate. If runtime structure, domain vocabulary, and UI decisions are unchanged, still verify that `docs/architecture.md`, `docs/domain.md`, and `docs/ui.md` (when applicable) exist and satisfy their requirements.
Do NOT squash-merge until the final tidy pass is complete and all tests are green.
Do NOT write the squash commit message without reading implementation.md to list actual delivered subtasks.
Do NOT land the increment before proving it integrates cleanly with the current `main` tip.
Do NOT archive `.agent/` files. Delete them, and only after landing evidence is captured and verified.
Do NOT delete `.agent/` or the increment branch before `LANDING_COMMIT` is captured and verified reachable from `main`.
On any failure between the squash-merge step and the evidence check, stop and report. Do not touch `.agent/` or the branch.
</HARD-GATE>

---

## Process

1. **Read all `.agent/` artifacts** — full review of increment, plan, implementation, and learnings.
2. **Audit the permanent documentation baseline** — inspect each required path and verify the architecture document contains both a C4 Level 1 System Context view and a C4 Level 2 Container view, has a Container Internals (Level 3) section slot present, and the domain document contains the glossary. Add missing or inadequate documents to the candidate list.
3. **Conversation: Propose promotions** — identify candidates and state what each is, its destination, and why it is durable. Iterate until the user says to proceed.
4. **On approval** — write each approved permanent artifact.
5. **Re-audit the baseline** — confirm every required document exists and satisfies its content requirement before cleanup.
6. **Final tidy pass** — on the increment branch, run the full test suite, then ask: is there any structural cleanup (rename, extract, inline) that would make the branch cleaner before it lands? Apply only behavior-preserving changes. Commit each as `tidy: <what>`. Tests must stay green throughout.
7. **Main-fit review** — present the intended fetch and integration strategy and wait for explicit approval. Then fetch the latest `main` without changing it, inspect the complete diff from its merge base, and check for conflicts, duplicated work, stale assumptions, accidental files, migration ordering, public-contract drift, and documentation inconsistency. Integrate the latest `main` into the increment branch using the approved strategy, resolve conflicts on the increment branch, then rerun the full release gate and every required acceptance test. Present the diff summary and evidence; wait for explicit approval that the increment fits `main`.
8. **Draft the squash commit message** from `implementation.md`:

   ```
   feat: <one-sentence goal from increment.md>

   Increment: <branch name>
   Acceptance criteria:
   - AC-1: <criterion>
   - AC-2: <criterion>

   Subtasks delivered:
   - tidy: <what>
   - feat: <what>
   - refactor: <what>
   [list each subtask commit message from implementation.md]

   Evidence: <test suite result — N passing, 0 failing>
   ```

9. **Land the increment (squash-merge to main):**

   ```
   git checkout main
   git merge --squash <branch>
   git commit
   LANDING_COMMIT=$(git rev-parse HEAD)
   ```

10. **Verify landing evidence:**
    - `LANDING_COMMIT` is non-empty.
    - `git merge-base --is-ancestor "$LANDING_COMMIT" main` succeeds.
    - `main` moved forward by exactly this commit (its previous tip is `LANDING_COMMIT`'s parent).

    If any check fails, treat as landing failure. Do NOT proceed to cleanup. Go to Recovery.

11. **On verified landing evidence — cleanup:**
    - Delete every file under `.agent/` for this cycle. Never archive.
    - Delete the increment branch: `git branch -D <branch>`.
    - Report `LANDING_COMMIT` and the deleted paths.

12. **Handoff** — cycle complete. Next action is `4dc-increment` for the next cycle.

---

## Recovery

If landing fails at any step in 9 or 10:

- Do NOT delete `.agent/` files.
- Do NOT delete the increment branch.
- Report which step failed and what tree state resulted.

Common failure paths:

- **`git checkout main` failed** → tree unchanged; resolve working-tree issues (dirty state, permissions, missing branch), then retry from step 9.
- **`git merge --squash` failed** → `main` unchanged; conflict or hook rejection. Re-run the Main-fit review (step 7), resolve on the increment branch, then retry from step 9.
- **`git commit` failed** → `main` may hold a staged merge with no commit. Run `git status`, address the failure (commit hook, GPG signing, empty commit), then retry `git commit` and re-capture `LANDING_COMMIT`.
- **`LANDING_COMMIT` empty or not reachable from `main`** → treat as commit failure; do not proceed to cleanup.
- **`git branch -D` failed** after landing succeeded → `LANDING_COMMIT` is already on `main`; the branch remains. Delete it manually when convenient; the cycle is still complete.

---

## Promotion Categories

| Type | Trigger | Destination |
|------|---------|-------------|
| Architecture decision | Non-obvious choice with lasting impact | `docs/adr/ADR-<date>-<slug>.md` |
| Guardrail update | A durable, cross-project engineering boundary needs revision | `CONSTITUTION.md` |
| Architecture sync | Runtime containers, dependency direction, or performance-critical paths changed | `docs/architecture.md` |
| Behavior change | Public API, CLI, or user-facing behavior changed | `README.md` |
| Feature shipped | Acceptance tests pass; feature complete | `docs/roadmap.md` — move to Done, add acceptance test link |
| Acceptance evidence | Required feature-level test was run | `docs/roadmap.md` or implementation evidence, linked to every covered criterion |
| Test pattern | New project-specific testing approach worth standardizing | `docs/testing.md` |
| Observability change | Signals, health criteria, alert response, or blind spots changed | `docs/observability.md` |
| Performance contract | A concrete latency, throughput, cost, or scaling expectation changed | `docs/architecture.md` or another project-specific performance document |
| Known issue | Found but not fixed this cycle | `docs/known-issues.md` |
| New domain concept | A concept, event, or rule used in code/tests that has no shared definition | `docs/domain.md` (create using the template in the Appendix if absent) |
| UI decision | Shared UI, interaction, visual, accessibility, or content decision | `docs/ui.md` |
| Structural change | A container added, removed, or re-wired | `docs/architecture.md` (create using the template in the Appendix if absent) |
| Component restructuring | A Refactor increment reshaped components inside a container in a way worth preserving durably | `docs/architecture.md` — Container Internals (C4 Level 3) section |

---

## Checklist

- [ ] All `.agent/` artifacts read
- [ ] Promotion candidates identified and categorized
- [ ] Permanent documentation baseline audited for existence and required content
- [ ] User approval received per candidate
- [ ] Each approved artifact written to permanent location
- [ ] Final baseline audit passes: glossary, C4 architecture view, and UI decisions (when applicable) are present and current
- [ ] Architecture, domain language, testing guidance, and performance documentation either updated or explicitly marked unchanged
- [ ] Required acceptance tests have results recorded for every criterion; approved exceptions include rationale
- [ ] Final tidy pass run on the increment branch: behavior-preserving cleanup committed as `tidy: <what>`, tests green
- [ ] Latest `main` fetched and integrated into the increment branch using the approved repository strategy
- [ ] Full branch diff reviewed against the current merge base for conflicts, stale assumptions, accidental scope, migrations, public contracts, and docs
- [ ] Release gate and required acceptance tests pass after integration
- [ ] User explicitly approved the main-fit review and evidence
- [ ] Squash commit message drafted from `implementation.md`
- [ ] `git merge --squash <branch>` and `git commit` succeeded on `main`
- [ ] `LANDING_COMMIT` captured via `git rev-parse HEAD` and verified reachable from `main`
- [ ] `.agent/` files deleted (not archived) — only after evidence verified
- [ ] Increment branch deleted with `git branch -D <branch>` — only after evidence verified

---

## Handoff

Terminal artifacts: permanent docs updated, `LANDING_COMMIT` on `main`, `.agent/` deleted, increment branch deleted.
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
