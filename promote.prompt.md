---
name: 4dc-promote
title: Promote learnings to trunk artifacts
description: Merge durable outcomes into permanent documentation and architecture records
version: "15d786a"
generatedAt: "2026-05-26T09:01:55Z"
source: https://github.com/co0p/4dc
---

# Prompt: Promote Phase

You are promoting durable outcomes from the 4dc working set into permanent project artifacts.

## Purpose
Convert transient implementation learnings into stable documentation and governance updates before merge.

## Execution Rules
- Review all `.4dc` artifacts before proposing promotions.
- Present each promotion candidate with destination and rationale.
- Require explicit approval before writing permanent docs.

## Language and Interaction Rules

- Use plain, direct language and keep output scannable.
- Ask focused questions; avoid broad questionnaires.
- For non-trivial work (more than three meaningful tasks, unknown dependencies, or cross-cutting changes), switch to plan mode before execution:
  1. Publish a short task plan.
  2. Execute in small verified steps.
  3. Update progress after each step.
- Keep assumptions explicit. If evidence is missing, ask one clarifying question.
- Treat source code and committed docs as the source of truth.
- Never claim work is complete without objective evidence.
- Default to forward-only change: do not preserve backward compatibility unless explicitly requested.
- Do not include rollback planning; define the next safe forward step instead.
## HTML Review Contract

Before writing final Markdown artifacts, generate a reviewable HTML file in `.4dc/` and pause for approval.

**Workflow Order (MANDATORY):**
1. Generate HTML review file in `.4dc/`
2. Present HTML to user for review
3. STOP and wait for explicit approval
4. Only after approval: write final Markdown artifacts

**File Naming Convention:**
- All `.4dc/` artifacts MUST use lowercase filenames
- Examples: `increment.md`, `plan.md`, `implementation.md`, `promote.md`
- Review files: `increment-review.html`, `plan-review.html`, `promotion-report.html`

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
- Do not write final Markdown artifacts until approval is explicit.

Canonical HTML skeleton (recommended):

```html
<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Phase Review</title>
	<style>
		:root {
			--bg: #f4f3ef;
			--panel: #ffffff;
			--ink: #111111;
			--muted: #4d4c45;
			--line: #1b1b1b;
			--accent: #234642;
			--codeBg: #151515;
			--codeInk: #ececec;
			--kw: #f4c95d;
			--str: #8dd694;
			--fn: #8cb4ff;
			--cm: #9aa0a6;
			--id: #ff9f7a;
		}
		* { box-sizing: border-box; }
		body {
			margin: 0;
			background: var(--bg);
			color: var(--ink);
			font-family: "IBM Plex Sans", "Segoe UI", -apple-system, sans-serif;
			line-height: 1.45;
		}
		.layout {
			max-width: 1180px;
			margin: 0 auto;
			display: grid;
			grid-template-columns: 250px minmax(0, 1fr);
			gap: 14px;
			padding: 16px;
		}
		.sidebar {
			position: sticky;
			top: 12px;
			align-self: start;
			border: 2px solid var(--line);
			background: var(--panel);
			padding: 12px;
		}
		.sidebar h2 {
			margin: 0 0 8px;
			font-size: 0.98rem;
			letter-spacing: 0.02em;
		}
		.sidebar a {
			display: block;
			color: var(--accent);
			text-decoration: none;
			margin: 7px 0;
			font-weight: 700;
		}
		.content .card {
			border: 2px solid var(--line);
			background: var(--panel);
			padding: 14px;
			margin-bottom: 12px;
		}
		.status {
			display: inline-block;
			padding: 6px 10px;
			border: 2px solid var(--line);
			font-weight: 800;
			background: #fff6dd;
		}
		table {
			width: 100%;
			border-collapse: collapse;
			margin-top: 10px;
		}
		th, td {
			border: 1px solid #2f2f2f;
			text-align: left;
			padding: 8px;
			vertical-align: top;
		}
		pre {
			margin: 10px 0 0;
			padding: 12px;
			border: 1px solid #2d2d2d;
			background: var(--codeBg);
			color: var(--codeInk);
			overflow-x: auto;
		}
		code {
			font-family: "JetBrains Mono", Menlo, Monaco, monospace;
			font-size: 0.92rem;
		}
		.kw { color: var(--kw); }
		.str { color: var(--str); }
		.fn { color: var(--fn); }
		.cm { color: var(--cm); }
		.id { color: var(--id); }
		.diagram svg {
			width: 100%;
			height: auto;
			border: 1px solid #2f2f2f;
			background: #f9f8f4;
		}
		.diagram text {
			fill: #111111;
			font-size: 12px;
			font-family: "IBM Plex Sans", "Segoe UI", sans-serif;
		}
		.diagram .node {
			fill: #ffffff;
			stroke: #1b1b1b;
			stroke-width: 2;
		}
		.diagram .edge {
			stroke: #1b1b1b;
			stroke-width: 2;
			fill: none;
			marker-end: url(#arrow);
		}
		@media (max-width: 920px) {
			.layout { grid-template-columns: 1fr; }
			.sidebar {
				position: static;
				margin-bottom: 8px;
			}
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
				<h1>Phase Review</h1>
				<p>Phase: Implement | Generated: YYYY-MM-DD HH:MM UTC</p>
				<p><span class="status">Status: Pending Approval</span></p>
			</section>

			<section class="card" id="inputs-reviewed">
				<h2>Inputs Reviewed</h2>
				<ul>
					<li>CONSTITUTION.md</li>
					<li>.4dc/increment.md</li>
					<li>.4dc/plan.md</li>
					<li>.4dc/implementation.md</li>
					<li>.4dc/promote.md</li>
				</ul>
			</section>

			<section class="card" id="proposed-output-summary">
				<h2>Proposed Output Summary</h2>
				<pre><code><span class="kw">const</span> <span class="id">status</span> = <span class="str">"Pending Approval"</span>;
<span class="cm">// Example token classes for code highlighting</span></code></pre>
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
				<svg viewBox="0 0 520 140" role="img" aria-label="Optional flow diagram">
					<defs>
						<marker id="arrow" markerWidth="10" markerHeight="7" refX="9" refY="3.5" orient="auto">
							<polygon points="0 0, 10 3.5, 0 7" fill="#1b1b1b"></polygon>
						</marker>
					</defs>
					<rect class="node" x="20" y="35" width="130" height="54"></rect>
					<text x="45" y="67">Input</text>
					<path class="edge" d="M150,62 L250,62"></path>
					<rect class="node" x="250" y="35" width="130" height="54"></rect>
					<text x="273" y="67">Review</text>
					<path class="edge" d="M380,62 L500,62"></path>
					<rect class="node" x="500" y="35" width="0" height="0" style="display:none"></rect>
				</svg>
				<p>Replace this sample SVG with any relevant architecture or flow graph when needed.</p>
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

## Required Inputs
- `CONSTITUTION.md`
- `.4dc/increment.md`
- `.4dc/plan.md`
- `.4dc/implementation.md`
- `.4dc/promote.md`
- Existing permanent docs and ADRs

## Required Output
- HTML report: `.4dc/promotion-report.html`
- Approved updates to permanent artifacts (for example CONSTITUTION, ADRs, design docs, OpenAPI, observability docs)
- Cleanup summary for `.4dc` transient files
- Retrospective delta in `.4dc/promote.md` including:
	- what improved delivery speed
	- what caused defects or friction
	- one concrete process adjustment for the next cycle

Promotion destinations to evaluate:
1. `CONSTITUTION.md`
2. ADR collection
3. OpenAPI specification
4. Visual design guide
5. Personas
6. Deployment strategy
7. Testing decisions
8. Observability documentation
9. C4 diagrams (system, container, component)
10. README and other canonical docs

## Process
1. Extract promotion candidates and classify durable vs transient. Label STOP PR1.
2. Draft exact target updates and traceability mapping. Label STOP PR2.
3. Produce `.4dc/promotion-report.html` for deep review. Label STOP PR3.
4. Apply only approved updates.
5. Suggest that the user empties `.4dc` after each promote and confirm the cleanup decision.
6. Record one next-cycle process adjustment in `.4dc/promote.md`.

## Anti-Patterns
- Promoting without traceability to evidence.
- Dropping unresolved decisions.
- Deleting transient artifacts before approval.

## Done When
- Approved promotions are written.
- `.4dc/promotion-report.html` is available and reviewed.
- A post-promote suggestion to empty `.4dc` was given.
- Retrospective delta is captured in `.4dc/promote.md`.
- Cleanup actions are explicitly confirmed.