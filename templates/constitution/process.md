# Prompt Process for Constitution Generation

## Operating Rules and Guardrails
- Act only with information provided or discovered in the repository. If context is missing, ask targeted questions instead of assuming.
- Enforce the STOP gate at Step 3 until the user answers or explicitly waives.
- Keep principles explicit, testable, and specific; avoid vague phrasing.
- Use today’s date for "Last Updated" in YYYY-MM-DD format.
- Human-first interaction: never surface raw JSON to the user unless explicitly requested.
- When emitting structured outputs for automation, do so internally only (tooling/CI) and emit JSON exactly as specified—no extra prose inside JSON blocks.
- Avoid harmful, hateful, lewd, or violent content; refuse such requests.

Human-first interaction: Always ask questions and present summaries in plain language. JSON emissions are internal-only for tooling and should not be shown to the user unless they explicitly request to see them.


## 1. Receive Initial Prompt
Inform the user: "You have requested to create or update the project constitution."

## 2. Analyze Project Context
Inform the user: "I will now review your project files (especially README.md and any existing CONSTITUTION.md) to understand the technical landscape and infer appropriate principles."

### Summary of Findings
After context analysis, provide a brief summary to the user outlining the project's purpose, tech stack, and any notable architectural patterns or constraints found.


## 3. Ask Pillar Questions & Increment Location (STOP)
Inform the user: "Before we begin, I will ask you explicit questions about each of the 6 pillars to understand your priorities and philosophies."
- What is your philosophy or priority for Delivery Velocity?
- What is your approach to Test Strategy?
- What are your rules for Design Integrity?
- How do you apply Simplicity First?
- What are your boundaries for Technical Debt?
- What is your Dependency Discipline?

Additionally, please specify where increments should be stored in your project. The recommended location is `docs/increments/`, with each increment in its own folder (e.g., `docs/increments/tray-menu/increment.md`). You may choose a different location if preferred.

**STOP:** Do not proceed until the user has answered these questions or explicitly asked you to continue without answers.

Internally emit a ClarificationRequest JSON (see Output section for schema) to capture the questions and the increments location recommendation. Do not display this JSON to the user; only show human-friendly questions.

## 4. Suggest Principles
Inform the user: "Based on your answers and project context, I will propose 3-5 core principles, each mapped to a pillar, with clear rationale."

### Summary of Findings
After suggesting principles, provide a concise summary listing the proposed principles, their mapped pillars, and the rationale for each.

Internally emit a PrinciplesProposal JSON (see Output section for schema). Do not display this JSON to the user; provide a concise human-readable summary instead.

## 5. Ask Clarifying Questions
Inform the user: "If any critical information is missing or the suggested principles need refinement, I will ask targeted follow-up questions."

## 6. Planned Sections Summary (STOP)
Present a human-readable checklist of sections you plan to generate and request confirmation:
- Vision, Mission, Core Values
- Architectural Principles (3–5; mapped to pillars)
- Pillar Coverage
- Update Process
- Technical Decisions (Languages, Frameworks, Deployment)
- Workflow: LLM Collaboration & Increment Workflow; Scope Drift Management; Testing & Verification Policy; Post-Implementation Stabilization; Merge & Release; Documentation & Traceability; Roles & Decision Gates
- Increments Location
- Last Updated (YYYY-MM-DD)

**STOP:** Ask the user to confirm this Planned Sections Summary or request changes before generation.

## 7. Generate Constitution
Inform the user: "Once you confirm the Planned Sections Summary (and provide any additional answers), I will generate the constitution document following the output format. The constitution will always include a section specifying where increments should be stored, using your answer or the recommended location (`docs/increments/`)."

Include the following additional sections in the generated document to guide LLM-supported incremental delivery:
- LLM Collaboration & Increment Workflow (STOP gates, feature branch, commit cadence, non-interactive phases)
- Scope Drift Management (DRIFT ALERT rules, scope containment, design escalation path)
- Testing & Verification Policy (test-first loop, manual verification allowances)
- Post-Implementation Stabilization (docs, hygiene, reproducible build, packaging)
- Merge & Release (branch merge, tagging, cleanup)
- Documentation & Traceability (required increment artifacts, decision logging)
- Roles & Decision Gates (sponsor, implementer, reviewer responsibilities)

## 8. Save Constitution
Inform the user: "I will save the generated constitution as CONSTITUTION.md in the project root."

### Summary of Findings
Provide a brief summary confirming the constitution was saved, listing the included sections and pillars covered.

Internally emit a ConstitutionSummary JSON (see Output section for schema) confirming sections, coverage, counts, paths, and date. Do not display this JSON to the user; provide a human-readable confirmation.

## 9. Final Validation
Inform the user: "Before saving, I will validate that all requirements are met: 3-5 principles, at least 3 pillars covered, each principle labeled, pillar coverage summary, declarative/testable/specific principles, and technical decisions section. If anything is missing, I will STOP and ask for clarification or fixes."

Also validate the presence of new workflow sections (LLM Collaboration, Scope Drift, Testing, Stabilization, Merge & Release, Documentation & Roles). If any are missing, STOP and add them.

---

## Structured JSON Outputs (Internal Only)

To enable automation and validation, emit a concise JSON block at specific steps. Place each JSON in a fenced block marked with `json` and do not include non-JSON prose inside the block.

Visibility: These JSON blocks are for internal tooling/CI only. Do not surface them to users unless explicitly requested.

### ClarificationRequest (Step 3)
Emit internally when asking the pillar questions and the increments location. Do not proceed until answered.

Schema (informal):
```json
{
	"step": "questions",
	"questions": [
		{ "id": 1, "pillar": "Delivery Velocity", "question": "...", "options": [ {"key": "A", "label": "..."}, {"key": "B", "label": "..."}, {"key": "C", "label": "..."}, {"key": "X", "label": "Skip"}, {"key": "_", "label": "Custom"} ] },
		{ "id": 2, "pillar": "Test Strategy", "question": "...", "options": [ ... ] },
		{ "id": 3, "pillar": "Design Integrity", "question": "...", "options": [ ... ] },
		{ "id": 4, "pillar": "Simplicity First", "question": "...", "options": [ ... ] },
		{ "id": 5, "pillar": "Technical Debt Boundaries", "question": "...", "options": [ ... ] },
		{ "id": 6, "pillar": "Dependency Discipline", "question": "...", "options": [ ... ] }
	],
	"increments_location_recommendation": "docs/increments/<increment-folder>/increment.md",
	"instructions": "Answer with letter keys (A, B, C, ...), X to skip, _ for custom text. Provide increments location or accept the recommendation."
}
```

### PrinciplesProposal (Step 4)
Emit internally when proposing 3–5 core principles mapped to pillars.

Schema (informal):
```json
{
	"step": "proposal",
	"principles": [
		{
			"id": 1,
			"name": "Short, Actionable Name",
			"pillar": "Delivery Velocity",
			"statement": "Declarative, testable, specific rule.",
			"rationale": "Why this exists for this project.",
			"implications": ["Notable consequence #1", "Notable consequence #2"]
		}
	],
	"coverage": ["Delivery Velocity", "Test Strategy", "Design Integrity"]
}
```

### ConstitutionSummary (Save Step)
Emit internally after generating and saving the constitution to confirm coverage and metadata.

Schema (informal):
```json
{
	"step": "summary",
	"sections_included": ["Vision", "Mission", "Core Values", "Architectural Principles", "Update Process", "Pillar Coverage", "Technical Decisions", "Last Updated"],
	"pillars_covered": ["Delivery Velocity", "Test Strategy", "Design Integrity"],
	"counts": { "principles": 4, "pillars": 3 },
	"paths": { "constitution": "CONSTITUTION.md", "increments": "docs/increments/<increment-folder>/increment.md" },
	"last_updated": "YYYY-MM-DD"
}
```
