---
name: 4dc-implement
description: "Use after plan.md is approved. Executes the plan in Red→Green→Refactor order, tracking progress and capturing learnings after each subtask."
---

# Implement Skill

## One Responsibility

Execute `.agent/plan.md` in controlled steps, maintain continuous progress state, and record every decision and deviation as it happens.

---

## Expected Input

- `CONSTITUTION.md`
- `.agent/increment.md` (approved)
- `.agent/plan.md` (approved)
- Current test baseline (run tests before touching anything)

---

## Concrete Output

- **`.agent/implementation.md`** — live progress log, updated after each subtask; final status either `status: complete` or `status: blocked`
- **`.agent/learnings.md`** — decisions made, deviations from plan, and lessons for `promote`

Required `implementation.md` headings:
- `## Baseline`
- `## Subtasks`
- `## Final Verification`

Required `learnings.md` headings:
- `## Decisions`
- `## Deviations`
- `## Surprises`
- `## Promote Candidates`

## Execution Contract

- Produce only the artifact for this phase. Do not leak work from a later phase into this one.
- Treat tests, architecture notes, ADRs, and user-facing docs as first-class communication artifacts.
- Gather only enough context to identify the governing constraints, the target artifact, and the cheapest validation step. Then act.
- Resolve conflicts in this order: explicit user approval, approved prior-phase artifacts, `CONSTITUTION.md`, this skill.
- Low-risk actions: reads, searches, diffs, and local validation commands.
- Medium-risk actions: local reversible edits to phase artifacts.
- High-risk actions: destructive file operations, external side effects, or skipping a stop gate. Require explicit approval first.
- If a required input is missing or contradictory, ask one focused question or stop at the review gate. Do not invent missing facts.
- Before finishing, run the phase checklist and confirm every required section is present.

---

<HARD-GATE>
Do NOT write production code for a subtask before a failing test exists for it (Red→Green→Refactor).
Do NOT mark a subtask complete without objective evidence (test output, command result, or observable behavior).
Do NOT skip or reorder subtasks without documenting the reason in learnings.md.
Do NOT proceed to the next subtask if the current one fails verification.
Do NOT mix structural tidying and behavior change in the same subtask.
</HARD-GATE>

---

## Process

1. **Establish baseline** — run existing tests; record pass/fail state in `implementation.md`
2. **Execute `[research]` subtasks first when present** — resolve only the blocking unknown, record the evidence, and return to the approved plan
3. **Execute `[tidy]` subtasks before any `[behavior]` subtask they support**
   a. Make the structural change
   b. Run the relevant tests — confirm they stay green
   c. Update `implementation.md` with evidence
   d. Record any design or architecture implications in `learnings.md`
4. **For each `[behavior]` subtask in plan.md:**
   a. Write the failing test (Red)
   b. Run it — confirm it fails for the right reason
   c. Write minimal production code (Green)
   d. Run the narrowest relevant tests, then the broader suite needed by the constitution
   e. Refactor if needed — run tests again
   f. Update `implementation.md`: mark subtask complete with evidence
   g. Append any decisions, deviations, surprises, or promote candidates to `learnings.md`
5. **Final verification** — run the full test suite or constitution-defined release gate; confirm all acceptance criteria from `increment.md` are met
6. **Mark complete** — set `status: complete` in `implementation.md`

## Tidy First Rule

When the approved plan includes `[tidy]` subtasks, use Kent Beck's Tidy First rule:

- Finish the structural preparation before behavior work that depends on it.
- Keep every `[tidy]` subtask behavior-preserving and test-green.
- If a tidy step would change observable behavior, it was mislabeled and must move to `[behavior]`.

---

## implementation.md Structure

```markdown
# Implementation: <increment goal>

status: in-progress  <!-- or: complete | blocked -->
started: <ISO date>

## Baseline
Tests before: X passing, Y failing

## Subtasks

### 1. <subtask name>
type: tidy | behavior | research
status: complete
evidence: `npm test` — 12 passing, 0 failing (added test: <name>)

### 2. <subtask name>
status: in-progress
```

---

## learnings.md Structure

```markdown
# Learnings: <increment goal>

## Decisions
- <decision>: <rationale>

## Deviations
- Subtask N: <what changed and why>

## Surprises
- <unexpected finding>

## Promote Candidates
- <ADR, architecture update, domain-language update, test pattern, or performance contract worth keeping>
```

## HTML Review Contract

Before writing final Markdown artifacts, generate a reviewable HTML file in `.agent/` and pause for approval.

**Workflow Order (MANDATORY):**
1. Generate HTML review file in `.agent/`
2. Present HTML to user for review
3. STOP and wait for explicit approval
4. Only after approval: write final Markdown artifacts

**File Naming Convention:**
- All `.agent/` artifacts MUST use lowercase filenames
- Examples: `increment.md`, `plan.md`, `implementation.md`, `learnings.md`
- Review files: `constitution-review.html`, `increment-review.html`, `plan-review.html`, `implementation-review.html`, `promotion-review.html`

Required report sections:
1. Objective
2. Inputs Reviewed
3. Proposed Output Summary
4. Risks and Trade-offs
5. Open Questions
6. Approval Decision

HTML requirements:
- Human-readable headings and table(s) where useful.
- Include a timestamp and phase name.
- Include a clear line: `Status: Pending Approval` until approved.
- Use a two-column layout with a left sidebar for quick section navigation.
- Sidebar must contain anchor links to all required report sections.
- Apply a pleasing no-fuzz CSS theme:
	- clear spacing scale and typography hierarchy
	- strong contrast and legible colors
	- simple surfaces and borders without noisy effects
	- mobile-friendly responsive behavior
- Include code highlighting support for snippets:
	- provide semantic classes for tokens (`kw`, `str`, `fn`, `cm`, `id`)
	- style `pre` and `code` blocks for readability
- Support inline SVG rendering for diagrams when useful:
	- allow dedicated diagram sections with embedded `<svg>`
	- style SVG text, lines, and nodes for visual consistency with the theme
- Tailor the phase label, reviewed inputs, and artifact filenames to the current phase. Never leave sample placeholders such as `Phase: Implement` in the final review file.
- Do not write final Markdown artifacts until approval is explicit.

Minimum HTML skeleton (adapt as needed for the current phase):

```html
<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>[Phase] Review</title>
	<style>
		:root {
			--bg: #f4f3ef;
			--panel: #ffffff;
			--ink: #111111;
			--accent: #1a56db;
			--muted: #6b7280;
			--border: #e5e7eb;
			--radius: 4px;
		}
	</style>
</head>
<body>
	<div class="layout">
		<nav class="sidebar" aria-label="Review Sections">
			<h2>Navigate</h2>
			<a href="#objective">Objective</a>
			<a href="#inputs-reviewed">Inputs Reviewed</a>
			<a href="#proposed-output-summary">Proposed Output Summary</a>
			<a href="#risks-and-trade-offs">Risks and Trade-offs</a>
			<a href="#open-questions">Open Questions</a>
			<a href="#approval-decision">Approval Decision</a>
		</nav>

		<main class="content">
			<section class="card" id="objective">
				<h1>[Phase] Review</h1>
				<p>Phase: [phase-name] | Generated: YYYY-MM-DD HH:MM UTC</p>
				<p><span class="status">Status: Pending Approval</span></p>
			</section>

			<section class="card" id="inputs-reviewed">
				<h2>Inputs Reviewed</h2>
				<ul>
					<li>[list only the inputs actually reviewed for this phase]</li>
				</ul>
			</section>

			<section class="card" id="proposed-output-summary">
				<h2>Proposed Output Summary</h2>
				<p>[Summarize the artifact to be written after approval.]</p>
			</section>

			<section class="card" id="risks-and-trade-offs">
				<h2>Risks and Trade-offs</h2>
				<table>
					<thead>
						<tr><th>Risk</th><th>Trade-off</th><th>Mitigation</th></tr>
					</thead>
					<tbody>
						<tr><td>Example</td><td>Example</td><td>Example</td></tr>
					</tbody>
				</table>
			</section>

			<section class="card diagram" id="open-questions">
				<h2>Open Questions</h2>
				<p>[List unresolved questions, or say None.]</p>
			</section>

			<section class="card" id="approval-decision">
				<h2>Approval Decision</h2>
				<p>Do not write final Markdown artifacts until explicit approval.</p>
			</section>
		</main>
	</div>
</body>
</html>
```

---

## Checklist

- [ ] Baseline test run recorded
- [ ] `[tidy]` subtasks, if any, stayed behavior-preserving and test-green
- [ ] Each subtask follows Red→Green→Refactor
- [ ] Each subtask has objective completion evidence
- [ ] All acceptance criteria met
- [ ] `implementation.md` status set to `complete`
- [ ] `learnings.md` has promote candidates listed

---

## Handoff

Terminal artifacts: `.agent/implementation.md` (status: complete) + `.agent/learnings.md`
Next skill: `4dc-promote` — load `skills/promote/SKILL.md`