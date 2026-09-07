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
- `VALIDATION.md`
- `templates/`
- `skills/`
- `scripts/`

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
- `skills/plan/SKILL.md`
- `skills/implement/SKILL.md`
- `skills/promote/SKILL.md`

Check command:

```bash
./scripts/generate-4dc.sh
find skills -name SKILL.md | sort
```

## Phase Requirements

### constitution
- Produces `.agent/constitution-review.md` before writing `CONSTITUTION.md`
- Defines engineering guardrails, performance expectations, and SDLC artifact policy

### increment
- Produces `.agent/increment-review.md` before writing `.agent/increment.md`
- Stays at WHAT/WHY and avoids technical design detail

### plan
- Produces `.agent/plan-review.md` before writing `.agent/plan.md`
- Converts requirements to ordered, verifiable technical subtasks with `[research]`, `[tidy]`, and `[behavior]` separation when needed

### implement
- Produces `.agent/implementation-review.md` before marking `.agent/implementation.md` complete
- Maintains `.agent/implementation.md` and `.agent/learnings.md`
- Records objective verification evidence and follows Tidy First plus Red→Green→Refactor

### promote
- Produces `.agent/promotion-review.md`
- Applies only approved updates to permanent artifacts
- Suggests emptying `.agent/` after each promote
- Confirms `.agent` cleanup decision and documentation sync

## Review-First Rule

For every phase, verify:
1. Markdown review exists in `.agent/`
2. Status is pending approval before final write
3. Final Markdown write occurs only after explicit approval

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
- `.agent/plan.md`
- `.agent/implementation.md`
- `.agent/learnings.md`
- phase review Markdown files

## Installer Validation

Run installer in a test repository and confirm:
- skill files are copied into `.agents/skills/<phase>/SKILL.md`
- `.agent/` is created
- `.agent` is present in `.gitignore`

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
- The Markdown review contract is present and phase-specific rather than copied with misleading placeholders.
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
