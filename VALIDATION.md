# 4dc Validation Prompt

Use this file in two ways:

1. As a human checklist for manual review.
2. As an attached instruction file for an LLM: attach this file and ask the model to validate your changes.

If you are the LLM reading this file, treat it as your validation contract.

## LLM Validation Contract

When a user attaches this file and asks whether their changes validate, do this:

1. Validate the current repository state against the rules in this file.
2. Default to review-only mode. Do not edit files unless the user explicitly asks for fixes.
3. Ignore the `examples/` directory unless the user explicitly asks you to include it.
4. Prefer running the documented checks directly from this file:

```bash
./scripts/generate-4dc.sh
find skills -name SKILL.md | sort
rg -n "\\.4dc|promotion-report|<p>Phase: Implement \\| Generated:" AGENTS.md README.md scripts templates skills
```

5. If a documented check fails, inspect the referenced files and explain which rules failed and why.
6. If a documented check cannot run because of missing tools or environment limits, perform a manual validation using the rules below and say that the result is manual.
7. Treat literal search patterns in this validation document as documentation, not stale active-contract references.
8. Report findings in this order:
	- Overall verdict: `Validates` or `Does not validate`
	- Check summary: documented commands run, passed checks, failed checks
	- Findings: each failed rule with file evidence
	- Gaps: checks this file cannot prove
9. If there are no findings, say that explicitly.
10. Do not claim alignment with vendor guidance unless the rules below pass.

## Scope

Validate these active instruction surfaces, including:
- `AGENTS.md`
- `templates/`
- `skills/`
- `scripts/`

`VALIDATION.md` is a repository-maintainer document. It is not installed into consuming projects and is included here only as the validation contract for changes to 4dc itself.

Permanent documentation validation also checks `docs/ui.md` when the target project has a user interface. It must contain recurring UI, interaction, visual, accessibility, and content decisions rather than a component or CSS inventory.

Do not treat generated drift inside `examples/` as a failure for this validation prompt.

## Quick Start

Recommended user prompt:

```text
Use VALIDATION.md as the validation contract. Review my current changes, do not modify files, ignore examples/, run the documented checks if possible, and tell me whether the repo validates.
```

## Direct Check Procedure

Use these commands when shell access is available:

```bash
./scripts/generate-4dc.sh
find skills -name SKILL.md | sort
rg -n "\\.4dc|promotion-report|<p>Phase: Implement \\| Generated:" AGENTS.md README.md scripts templates skills
```

Interpretation:
- `generate-4dc.sh` must complete without error.
- `find skills -name SKILL.md | sort` must list the five phase skill files.
- The consistency sweep must return no matches.
- Then inspect the files manually against the phase and vendor rules below.

## Validation Checklist

Use this checklist to validate the fresh 4dc prompt suite.

## Prompt Set

Expected generated files:
- `skills/constitution/SKILL.md`
- `skills/increment/SKILL.md`
- `skills/prototype/SKILL.md`
- `skills/plan/SKILL.md`
- `skills/adr/SKILL.md`
- `skills/tidy/SKILL.md`
- `skills/tdd-red/SKILL.md`
- `skills/tdd-green/SKILL.md`
- `skills/refactor/SKILL.md`
- `skills/promote/SKILL.md`

Check command:

```bash
./scripts/generate-4dc.sh
find skills -name SKILL.md | sort
```

## Phase Requirements

### constitution
- Writes `CONSTITUTION.md` only after the user explicitly approves the proposed guardrails
- Defines engineering guardrails, performance expectations, and SDLC artifact policy

### increment
- Writes `.agent/increment.md` only after the user explicitly approves the proposed increment
- Stays at WHAT/WHY and avoids technical design detail

### plan
- Writes `.agent/plan.md` only after the user explicitly approves the proposed plan
- Converts requirements to ordered, verifiable technical subtasks with `[research]`, `[tidy]`, and `[behavior]` separation when needed
- May define optional acceptance scenarios for larger increments; scenarios are advisory by default and must not become implicit blockers

### tdd-red
- Writes exactly one failing test for the current `[behavior]` subtask
- Confirms the test fails for the right reason before handing off
- Does not write production code

### tdd-green
- Makes the failing test pass with minimal code (no refactoring)
- Sets `state: green` and hands off to `refactor`
- Handles `[research]` subtasks (which skip Red and Refactor, set `state: complete` directly)
- Commits behavior work as `feat:` or `fix:`, research as `research:`

### tidy
- Executes one `[tidy]` subtask: behavior-preserving structural change
- Tests must stay green before and after
- Commits as `tidy: <what changed>`

### refactor
- Improves design without changing behavior (Fowler two-hats)
- Tests must stay green throughout
- Commits as `refactor: <what changed>` (or skips commit if no refactoring needed)
- Sets `implementation.md` status to `complete` only after final verification and user approval

### prototype
- Builds a throwaway spike to resolve one named unknown
- No production code merged; finding recorded in `.agent/prototype.md`
- Time-boxed; disposed after the finding is captured

### adr
- Writes one ADR per decision with context, alternatives, rationale, and consequences
- Does not write ADRs for implementation details or decisions already in `CONSTITUTION.md`
- User confirms the decision before the ADR is written

### promote
- Writes permanent artifacts only after each candidate is individually approved
- Applies only approved updates to permanent artifacts
- Suggests emptying `.agent/` after each promote
- Confirms `.agent` cleanup decision and documentation sync
- Records optional acceptance-scenario evidence without treating advisory scenarios as default promotion gates

## Artifact Policy

Permanent artifacts to evaluate every cycle:
- ADRs
- Architecture docs
- Domain model docs
- Deployment strategy
- Testing decisions
- Observability docs
- C4 diagrams or equivalent architecture views
- Roadmap (`docs/roadmap.md`)

Transient artifacts:
- `.agent/increment.md`
- `.agent/prototype.md`
- `.agent/plan.md`
- `.agent/implementation.md`
- `.agent/learnings.md`

## Consistency Sweep

Run a targeted text check after generation:

```bash
rg -n "\\.4dc|promotion-report|<p>Phase: Implement \\| Generated:" AGENTS.md README.md scripts templates skills
```

Expected result:
- No `.4dc` references in active repo instructions
- No `promotion-report` references in active repo instructions
- No rendered `Phase: Implement` placeholder lines left in generated skill contracts

## Vendor Alignment Rules

Use these rules after manual prompt changes to estimate whether the repo still aligns with current guidance from OpenAI, Anthropic, and Google.

Interpretation:
- `Pass` means the expected rule is present in the active orchestration or generated skills.
- `Fail` means the rule is missing or contradicted and should be reviewed.
- This is a heuristic alignment check, not a guarantee of runtime quality.

Suggested manual scoring:
- Count each alignment bullet below as one check.
- Mark it `Pass`, `Fail`, or `Not proven`.
- Use the count of `Pass` over total applicable checks as the alignment score.

### OpenAI Alignment

These rules reflect current OpenAI guidance for AGENTS.md, skills, and agentic coding prompts.

- The orchestrator defines instruction precedence so the model does not waste effort reconciling contradictions.
- The orchestrator defines safe versus high-risk actions and requires approval for destructive or externally visible operations.
- Each skill has a clear output contract and stop condition before the next phase.
- Each skill includes a concise execution contract so the model can act without reopening the whole repo.
- Generated skills contain no unrendered template markers.
- Generated skills avoid stale contract paths such as `.4dc/`.

### Anthropic Alignment

These rules reflect current Anthropic guidance for clear instructions, structured prompts, examples, and agentic state.

- Each skill stays focused on one job and names that responsibility explicitly.
- Each skill uses a stable section schema: responsibility, inputs, outputs, hard gate, process, checklist, handoff.
- Each skill contains explicit hard gates instead of relying on implied behavior.
- Each skill pauses for explicit conversational approval before writing its final artifact, rather than generating an intermediate review file.
- Implement explicitly encodes Red→Green→Refactor and Tidy First behavior.
- Promote explicitly requires durable documentation sync, not just code completion.

### Google Alignment

These rules reflect current Gemini guidance for direct instructions, consistent prompt structure, decomposition, and constraints.

- The prompts keep a consistent structure across phases.
- Constraints and required output headings are stated explicitly rather than inferred.
- Complex work is decomposed into explicit subtask types where relevant.
- Architectural and performance-sensitive concerns are surfaced before implementation.
- Validation is runnable from documented commands and supports specific failure reporting.

## Scoring Guide

Use the validation result as a fast alignment signal:
- 100% pass rate: strong textual alignment with the current 4dc contracts and the targeted vendor guidance.
- 85% to 99%: acceptable, but inspect the failed checks before merging.
- Below 85%: the prompt set is drifting and should be corrected before relying on it.

## What The Validator Does Not Check

- Whether the prompts are actually optimal for a specific model snapshot.
- Whether the wording is too verbose or too terse in practice.
- Whether agents will always obey the rules under long-context pressure.
- Whether a changed prompt produces better outcomes without evals.

Use this validator as a guard rail. Use real prompt evals and manual review for final judgment.
