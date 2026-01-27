# Prompt Validation Rules

Use this document to validate the generated prompts against the README.md specification.

---

## 1. Structural Consistency

All prompts MUST have consistent structure and language.

### 1.1 Frontmatter Requirements

| Rule | Check |
|------|-------|
| All prompts have `name` field with `4dc-` prefix | ☐ |
| All prompts have `title` field | ☐ |
| All prompts have `description` field | ☐ |
| All prompts have `version`, `generatedAt`, `source` fields | ☐ |
| Only `increment` has `argument-hint` (user story input) | ☐ |

### 1.2 Section Structure

Each prompt MUST contain these sections in this order:

| Section | Purpose | Required |
|---------|---------|----------|
| `# Prompt: [Title]` | Main heading | ✓ |
| `## Core Purpose` | One-line summary of what prompt does | ✓ |
| `## Persona & Style` | Who the LLM acts as, behavior style | ✓ |
| `## Input Context` | What the LLM should read before starting | ✓ |
| `## Goal` | What outputs are expected | ✓ |
| `## Process` | Step-by-step with STOP gates | ✓ |
| `## Output Structure` | Format of generated artifacts | ✓ |
| `## Anti-Patterns` | What NOT to do | ✓ |
| `## Example Questions` | Sample Socratic questions | ✓ |
| `## Constitutional Self-Critique` | Internal validation loop | ✓ |
| `## Communication Style` | Tone and response format | ✓ |

### 1.3 Language Consistency

| Rule | Check |
|------|-------|
| No mention of "path" as argument (use "current project") | ☐ |
| No hardcoded artifact paths except `.4dc/current/` | ☐ |
| Artifact locations reference "per CONSTITUTION.md" | ☐ |
| No mention of "modes" (lite/medium/heavy) | ☐ |
| No meta-chat about prompts/LLMs in output artifacts | ☐ |
| Consistent use of STOP gates (labeled clearly) | ☐ |

---

## 2. Anti-Hallucination Rules

Prompts MUST provide enough context for LLM to work without guessing.

### 2.1 Input Context Completeness

| Prompt | Required Inputs | Check |
|--------|-----------------|-------|
| constitution | Existing code structure, README, any existing CONSTITUTION.md | ☐ |
| increment | CONSTITUTION.md, user story, existing code | ☐ |
| implement | CONSTITUTION.md, `.4dc/current/increment.md`, existing code + tests | ☐ |
| promote | `.4dc/current/learnings.md`, CONSTITUTION.md, existing ADRs/contracts | ☐ |
| reflect | CONSTITUTION.md, existing ADRs, code + tests, recent commits | ☐ |

### 2.2 When Uncertain, Ask User

Each prompt MUST instruct the LLM to:

| Rule | Check |
|------|-------|
| Ask clarifying questions rather than assume | ☐ |
| Challenge vague answers ("Flexible for what?") | ☐ |
| Confirm understanding at STOP gates before proceeding | ☐ |
| Wait for explicit user approval before writing artifacts | ☐ |
| Never generate solutions before asking discovery questions | ☐ |

### 2.3 STOP Gate Requirements

| Prompt | Required STOP Gates | Check |
|--------|---------------------|-------|
| constitution | STOP 1: context summary, STOP 2: outline approval | ☐ |
| increment | STOP 1: understanding, STOP AC: criteria, STOP UC: use case, STOP 2: deliverables | ☐ |
| implement | STOP after each: test suggestion, red verification, green guidance, refactoring | ☐ |
| promote | STOP for each learning decision, STOP before deletion | ☐ |
| reflect | STOP 1: context, STOP 2: pattern summary | ☐ |

---

## 3. Information Flow (per README.md)

The prompts MUST support this workflow:

```
constitution (one-time) → increment → implement → promote → [reflect periodically]
                              ↑                                      │
                              └──────────────────────────────────────┘
```

### 3.1 Constitution Prompt

| Requirement (from README) | Check |
|---------------------------|-------|
| Creates `CONSTITUTION.md` at project root | ☐ |
| Contains: layering, error handling, testing, artifact layout, delivery | ☐ |
| Asks concrete questions, not abstract values | ☐ |
| Does NOT create: quality lenses, large ADRs, style guides | ☐ |
| Output is permanent, evolves with project | ☐ |

### 3.2 Increment Prompt

| Requirement (from README) | Check |
|---------------------------|-------|
| Takes user story/feature idea as input | ☐ |
| Reads CONSTITUTION.md for alignment | ☐ |
| Creates `.4dc/current/increment.md` | ☐ |
| Contains: user story, acceptance criteria, use case, deliverables | ☐ |
| Slices into small, independently shippable deliverables | ☐ |
| Stays at WHAT/WHY level, no technical HOW | ☐ |
| Output is temporary, deleted after merge | ☐ |

### 3.3 Implement Prompt

| Requirement (from README) | Check |
|---------------------------|-------|
| Reads CONSTITUTION.md and `.4dc/current/increment.md` | ☐ |
| Works one deliverable at a time | ☐ |
| Guides TDD: Red → Green → Refactor | ☐ |
| ONE test at a time, never batch | ☐ |
| Asks promotion questions every 5-10 cycles | ☐ |
| Creates `.4dc/current/learnings.md` | ☐ |
| Creates `.4dc/current/notes.md` (session observations) | ☐ |
| Output: permanent code + tests, temporary notes/learnings | ☐ |

### 3.4 Promote Prompt

| Requirement (from README) | Check |
|---------------------------|-------|
| Reads `.4dc/current/learnings.md` | ☐ |
| For each learning, asks WHERE it should go | ☐ |
| Promotion targets: CONSTITUTION.md, ADRs, API contracts, README, backlog | ☐ |
| Drafts additions, shows exact placement | ☐ |
| Waits for confirmation before writing | ☐ |
| Confirms deletion of `.4dc/current/` | ☐ |
| Output: updates to permanent docs, ephemeral context deleted | ☐ |

### 3.5 Reflect Prompt

| Requirement (from README) | Check |
|---------------------------|-------|
| Reads CONSTITUTION.md and codebase | ☐ |
| Uses quality lenses (defined IN prompt, not constitution) | ☐ |
| 8 lenses: naming, modularity, architecture, testing, duplication, docs, delivery, dependencies | ☐ |
| Identifies concrete refactorings, not reports | ☐ |
| Each refactoring scoped as one increment | ☐ |
| Output: constitution updates, ADRs, new increment ideas, backlog | ☐ |

### 3.6 Artifact Locations

| Artifact | Location | Prompt Creates | Prompt Reads |
|----------|----------|----------------|--------------|
| CONSTITUTION.md | Project root | constitution, promote | all |
| increment.md | `.4dc/current/` | increment | implement |
| notes.md | `.4dc/current/` | implement | implement |
| learnings.md | `.4dc/current/` | implement | promote |
| ADRs | per CONSTITUTION.md | promote | reflect |
| API contracts | per CONSTITUTION.md | promote | - |

### 3.7 Lifecycle Consistency

| Rule | Check |
|------|-------|
| `.4dc/current/` is temporary (deleted after merge) | ☐ |
| CONSTITUTION.md is permanent (evolves) | ☐ |
| ADRs are permanent | ☐ |
| Prompts reference artifacts consistently | ☐ |
| No prompt creates artifacts outside its scope | ☐ |

---

## 4. Validation Checklist

Run through each prompt and verify:

### 4.1 constitution.prompt.md
- [ ] Frontmatter correct (no argument-hint)
- [ ] All required sections present
- [ ] Asks about artifact layout (where ADRs, API docs live)
- [ ] STOP 1 and STOP 2 clearly labeled
- [ ] No hardcoded paths except `.4dc/current/`
- [ ] No mention of modes

### 4.2 increment.prompt.md
- [ ] Frontmatter has argument-hint for user story
- [ ] All required sections present
- [ ] Reads CONSTITUTION.md
- [ ] Creates `.4dc/current/increment.md`
- [ ] Has STOP 1, STOP AC, STOP UC, STOP 2
- [ ] Slices into deliverables
- [ ] No technical implementation details

### 4.3 implement.prompt.md
- [ ] Frontmatter correct (no argument-hint)
- [ ] All required sections present
- [ ] Reads CONSTITUTION.md and increment.md
- [ ] TDD cycle clearly documented
- [ ] One test at a time enforced
- [ ] Promotion checks every 5-10 cycles
- [ ] Creates learnings.md and notes.md

### 4.4 promote.prompt.md
- [ ] Frontmatter correct (no argument-hint)
- [ ] All required sections present
- [ ] Reads learnings.md
- [ ] Decision tree for each learning type
- [ ] Drafts exact content before writing
- [ ] Confirms deletion of .4dc/current/
- [ ] References CONSTITUTION.md for artifact paths

### 4.5 reflect.prompt.md
- [ ] Frontmatter correct (no argument-hint)
- [ ] All required sections present
- [ ] 8 quality lenses defined IN prompt
- [ ] STOP 1 and STOP 2 clearly labeled
- [ ] Outputs increment ideas (not just reports)
- [ ] Refactorings scoped to single increment

---

## Usage

To validate prompts:

1. Generate prompts: `./templates/generate-all.sh`
2. For each prompt, go through section 4 checklist
3. For any failures, trace back to sections 1-3 for specific rule
4. Fix template, regenerate, revalidate
