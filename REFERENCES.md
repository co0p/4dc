Best high-signal sources to refine 4dc prompts for GPT-5 Codex and Claude Sonnet:

1. OpenAI GPT-5 Prompting Guide (most practical, agentic-focused)
- https://developers.openai.com/cookbook/examples/gpt-5/gpt-5_prompting_guide  
Why: Covers controlling agent eagerness, tool preambles, reasoning effort, instruction conflict cleanup, and coding-specific prompt patterns.

2. OpenAI Prompt Engineering Guide
- https://developers.openai.com/api/docs/guides/prompt-engineering  
Why: Canonical guide for instruction hierarchy, role usage, XML/Markdown structuring, few-shot patterns, context planning, and eval-driven iteration.

3. OpenAI Text Generation Guide
- https://developers.openai.com/api/docs/guides/text  
Why: Practical API-level guidance for Responses API, role authority, reusable prompts, and model behavior differences.

4. OpenAI Code Generation Guide (GPT-5 + Codex positioning)
- https://developers.openai.com/api/docs/guides/code-generation  
Why: How to choose between GPT-5 family models for coding and when to use Codex workflows.

5. Codex Best Practices
- https://developers.openai.com/codex/learn/best-practices  
Why: Direct guidance on prompt structure (Goal, Context, Constraints, Done-when), AGENTS.md patterns, planning mode, testing/review loops, and common mistakes.

6. Codex AGENTS.md Guide
- https://developers.openai.com/codex/guides/agents-md  
Why: Exactly how instruction layering/precedence works and how to organize persistent repo guidance.

7. Codex Skills Guide
- https://developers.openai.com/codex/skills  
Why: How to package repeatable workflows into skills, write triggerable descriptions, and keep skills scoped.

8. OpenAI API Skills Guide
- https://developers.openai.com/api/docs/guides/tools-skills  
Why: Versioned skill packaging, invocation behavior, and safety constraints when skills are used with tools.

9. Anthropic Prompting Best Practices (Claude Sonnet/Opus)
- https://platform.claude.com/docs/en/docs/build-with-claude/prompt-engineering/claude-prompting-best-practices  
Why: The primary Anthropic reference for clear instructions, XML-tag structuring, examples, tool behavior, thinking controls, and agentic system tuning.

10. Anthropic Prompt Engineering Overview
- https://platform.claude.com/docs/en/docs/build-with-claude/prompt-engineering/overview  
Why: Strong process framing: define success criteria, build evaluations, then iterate prompts.

11. Anthropic XML Tags Guidance
- https://platform.claude.com/docs/en/docs/build-with-claude/prompt-engineering/use-xml-tags  
Why: Best source for Claude-specific section/tag design and nested structure conventions.

12. GitHub Copilot Customization (for repo instruction architecture)
- https://docs.github.com/en/copilot/concepts/prompting/response-customization  
- https://docs.github.com/en/copilot/how-tos/configure-custom-instructions/add-repository-instructions  
Why: Instruction precedence, file types, and practical rules for repository/path-specific instruction design that map well to 4dc prompt orchestration.

---

## Structural Improvements To Apply Across All `*prompt.md`

These are cross-cutting improvements derived from the references above and validated against the current 4dc prompt set.

### 1. Add a consistent "Execution Contract" section near the top

Purpose: reduce ambiguity and over/under-eager behavior.

Include in every prompt:
- **Autonomy policy**: when to act directly vs ask for confirmation.
- **Tool policy**: investigate first, no guessing, no placeholder parameters.
- **Stop conditions**: explicit criteria for when the prompt is "done."
- **Safety boundaries**: destructive actions require explicit confirmation.

### 2. Standardize instruction hierarchy and conflict handling

Purpose: improve adherence for highly steerable models.

Add one short subsection in every prompt:
- "If two instructions conflict: prioritize user-confirmed scope, then constitution constraints, then this prompt's defaults."
- "If conflict remains unresolved: surface one concise clarification question."

### 3. Use a shared section schema in the same order

Purpose: improve reliability, maintainability, and generation quality.

Recommended order for every prompt:
1. Core Purpose
2. Execution Contract
3. Persona & Style
4. Input Context
5. Output Contract (required artifacts + format)
6. Process (phases + STOP gates)
7. Quality Checks (self-check rubric)
8. Anti-Patterns
9. Communication Style

### 4. Add explicit Output Contract with machine-checkable completion criteria

Purpose: avoid partial completion and ambiguous done state.

Require each prompt to define:
- Required files/artifacts
- Required sections/headings
- Required status fields/checklists (where applicable)
- A final completion checklist (all items must be true)

### 5. Normalize status/progress semantics across prompts

Purpose: fix inconsistent work-item completion tracking.

Adopt common status vocabulary everywhere:
- `Not started`
- `In progress`
- `Done`

And require:
- Status line in working artifacts
- Checkbox updates for completed items
- "Next step" pointer after each major transition

### 6. Add reusable "context gathering budget" guidance

Purpose: prevent over-searching and latency bloat.

Use a compact policy in every prompt:
- Start broad, then one focused pass.
- Stop searching once exact target paths are identified.
- Search again only if validation fails or uncertainty remains material.

### 7. Require preamble and phase-transition summaries

Purpose: improve traceability in long agentic flows.

In every prompt, require:
- Brief upfront plan before actions
- Short transition summary at each STOP gate
- Final delta summary: what changed, what remains, risks/open questions

### 8. Add examples as structured few-shot blocks

Purpose: improve formatting and decision consistency.

For each prompt, add 2-3 short examples with:
- Input situation
- Expected behavior
- Expected output snippet

Keep examples tagged and clearly separated from instructions.

### 9. Add "no contradiction" linting pass in constitutional self-critique

Purpose: prevent internally conflicting instructions that hurt GPT-5 performance.

In each self-critique block, add:
- Check for conflicting MUST/SHOULD statements.
- Resolve by choosing one canonical rule and removing duplicates.
- Ensure each STOP gate has one clear proceed condition.

### 10. Add lightweight eval hooks for prompt iteration

Purpose: support measurable prompt improvements.

Each prompt should define 3-5 eval checks, e.g.:
- Completeness: all required sections present
- Determinism: same input yields same structure
- Actionability: each finding maps to next step
- Scope control: no out-of-scope expansions

### 11. Harmonize markup strategy

Purpose: reduce parse ambiguity across model families.

Use consistent delimiters in all prompts:
- Markdown headings for top-level structure
- Optional XML-style tags only for complex embedded blocks (examples/contracts)
- Keep tag names stable across all prompt files

### 12. Factor shared boilerplate into reusable template fragments

Purpose: avoid drift between prompt files.

Promote shared blocks into common fragments:
- Execution Contract
- Output Contract skeleton
- Quality Checks rubric
- Anti-Patterns core list

Then include prompt-specific deltas only.

---

## High-Impact First Pass (recommended)

Apply these first across all five prompt files:
1. Add `Execution Contract` + stop conditions.
2. Add explicit `Output Contract` + completion checklist.
3. Standardize status semantics (`Not started / In progress / Done`).
4. Add contradiction linting to self-critique.
5. Add one structured few-shot example per prompt.
