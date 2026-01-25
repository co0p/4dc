Perfect question! Here's what you should save as context for each prompt so we can write them later:

---

## `constitution.prompt.md`

### Core Purpose
Help user discover and document their project's specific architectural decisions through Socratic questioning.

### Key Behaviors
- Ask concrete questions about layering, error handling, testing, dependencies
- Challenge vague answers ("make it flexible" → "Flexible for what specific need?")
- Ensure decisions are actionable, not aspirational
- Keep output focused: decisions that guide daily work, NOT values/best-practices/lenses

### Input Context
- Project path (e.g., `.` or `examples/pomodoro`)
- Existing code structure (languages, frameworks, directories)
- Existing docs (README, any ADRs)

### Output
- `CONSTITUTION.md` at project root
- Sections:
  - **Architectural Decisions** (layering, error handling, state management, etc.)
  - **Testing Expectations** (where tests live, speed requirements, mocking strategy)
  - **Artifact Layout** (where ADRs go, where API contracts go, where increment context goes)
  - **Delivery Practices** (PR size, CI requirements, deployment process)

### STOP Gates
- **STOP 1**: Summarize project context, propose initial categorization → wait for confirmation
- **STOP 2**: Present outline of sections and sample decisions → wait for approval
- Only after STOP 2 approval: write full `CONSTITUTION.md`

### Anti-Patterns to Guard Against
- Abstract values ("we value quality") → ask for concrete decision
- Generic best practices ("follow SOLID") → ask how it applies to THIS project
- Mode levels (lite/medium/heavy) → just state actual decisions
- Lenses in constitution → those belong in reflect prompt
- Large ADRs in constitution → those are separate docs

### Example Questions
- "Where should domain logic live relative to UI code?"
- "How do you handle errors? Return codes, exceptions, Result types?"
- "What's your minimum testing expectation? Every function? Critical paths only?"
- "Do you wrap third-party dependencies or use them directly?"
- "The constitution says [X]. Does this new decision conflict? Should we update?"

---

## `increment.prompt.md`

### Core Purpose
Help user slice a feature idea into small, shippable deliverables through discovery questions about WHAT and WHY.

### Key Behaviors
- Start from vague idea ("add password reset") → discover specific outcome
- Challenge scope creep ("Is that THIS increment or follow-up?")
- Challenge vague success criteria ("What specific behavior tells you it worked?")
- Slice into **deliverables** (small, independently shippable pieces)
- Stay at WHAT/WHY level (no technical HOW, no implementation details)

### Input Context
- Project path
- `CONSTITUTION.md` (to align with project decisions)
- Short feature description from user
- Existing code (to understand current capabilities)

### Output
- `.4dc/current/increment.md` (TEMPORARY - will be deleted after merge)
- Sections:
  - **User Story** (As a..., I want..., so that...)
  - **Acceptance Criteria** (observable behaviors that must be true)
  - **Use Case** (actors, preconditions, main flow, alternates)
  - **Context** (why this matters now)
  - **Deliverables** (ordered slices, each shippable independently)
  - **Promotion Checklist** (hints for what might become permanent)

### STOP Gates
- **STOP 1**: Summarize understanding of problem/context → wait for confirmation
- **STOP AC**: Propose acceptance criteria → iterate until user confirms "complete enough"
- **STOP UC**: Propose use case → iterate until user confirms
- **STOP 2**: Propose deliverable slices and full outline → wait for approval
- Only after STOP 2: write full `increment.md`

### Deliverable Slicing Strategy
Each deliverable should:
- Provide value OR learning
- Be shippable (working code, even if feature incomplete)
- Inform the next deliverable
- Example: "add password reset" →
  1. Token generation (foundation + learn about storage)
  2. Email sending (integration + learn about templates)
  3. Reset flow UI (completion + learn about UX)

### Anti-Patterns to Guard Against
- Technical solutions in increment ("use bcrypt") → ask "why does that matter to users?"
- File/class/module names → stay product-level
- Implementation steps → those belong in implement prompt
- Vague deliverables ("backend work") → ask "what specific capability becomes available?"

### Example Questions
- "What's the smallest outcome that would provide value?"
- "What's explicitly OUT of scope for this increment?"
- "How will you know it worked? What metric/behavior changes?"
- "Can we ship deliverable 1 and get feedback before doing 2?"
- "What would you learn from deliverable 1 that informs deliverable 2?"

---

## `implement.prompt.md`

### Core Purpose
Guide user through TDD cycles (Red → Green → Refactor) one deliverable at a time, helping design emerge from code.

### Key Behaviors
- Work one deliverable at a time (informed by previous deliverables)
- Suggest next smallest test (one test at a time, never batch)
- Ask test quality questions (failing for right reason? simplest test?)
- Ask implementation questions (simplest solution? solving just this test?)
- Ask refactoring questions (what smells bad? should we refactor now?)
- Every 5-10 cycles: ask "Have we discovered anything worth promoting?"
- Capture learnings in `.4dc/current/learnings.md`

### Input Context
- `CONSTITUTION.md` (architectural decisions to follow)
- `.4dc/current/increment.md` (what we're building)
- Existing code + tests
- `.4dc/current/notes.md` (previous session observations)
- For deliverable N: learnings from deliverable N-1

### Output
- Working code + tests (PERMANENT)
- `.4dc/current/notes.md` (session observations, TEMPORARY)
- `.4dc/current/learnings.md` (promotion candidates, TEMPORARY)

### TDD Cycle Pattern
```
1. Suggest next test
   STOP → user writes test, shows result

2. Verify red phase
   Ask: "Is this failing for the right reason?"
   Ask: "Is this the simplest test that could fail?"
   STOP → user confirms

3. Guide green phase
   Ask: "What's the simplest implementation?"
   Ask: "Are we solving THIS test or future needs?"
   STOP → user implements, shows result

4. Suggest refactorings
   Ask: "With tests green, what smells bad?"
   Suggest refactorings based on constitution
   Ask: "Should we refactor now or next test?"
   STOP → user decides

5. Every 5-10 cycles: promotion check
   Ask: "Discovered any architectural decisions?"
   Ask: "Discovered any API contracts?"
   Append to learnings.md if yes

6. Deliverable complete check
   Ask: "Is this deliverable shippable?"
   Ask: "What did we learn for next deliverable?"
```

### Learnings.md Format
```markdown
## CONSTITUTION Updates
- [ ] Decision description
      Section: where it belongs

## ADRs to Create
- [ ] Decision description
      Rationale: why it matters

## API Contracts to Add
- [ ] Contract description
      File: docs/api/path/file.yaml

## Backlog Items
- [ ] Future work description
```

### Anti-Patterns to Guard Against
- Suggesting multiple tests at once → ONE test at a time
- Suggesting implementation before test fails → enforce RED first
- Speculative abstractions → ask "Does THIS test require it?"
- Large refactorings with red tests → refactor only when green
- Skipping promotion check → ask every 5-10 cycles

### Example Questions
- "What's the first test for token generation?"
- "Is this test failing for the right reason? (e.g., NameError, not AssertionError)"
- "What's the simplest code that makes this pass? (even if 'wrong')"
- "This duplicates code from UserService—should we extract it per constitution's DRY principle?"
- "We discovered SHA256 is fast enough for tokens—should this go in CONSTITUTION.md Security section?"

---

## `promote.prompt.md`

### Core Purpose
Before merging, ensure important learnings become permanent documentation, then safely delete ephemeral increment context.

### Key Behaviors
- Read `.4dc/current/learnings.md`
- For each learning, ask WHERE it should go
- Draft the additions (show exact placement)
- Wait for confirmation before writing
- After all promotions: confirm deletion of `.4dc/current/`

### Input Context
- `.4dc/current/learnings.md` (populated by implement prompt)
- `CONSTITUTION.md` (to see current structure)
- `docs/adr/` (to see existing ADRs)
- `docs/api/` (to see existing contracts)
- `README.md` (to check if project scope changed)

### Output
- Updates to `CONSTITUTION.md` (if architectural decisions)
- New files in `docs/adr/` (if decisions need explanation)
- New files in `docs/api/` (if public contracts)
- Updates to `README.md` (if project scope changed)
- Confirmation to delete `.4dc/current/`

### Promotion Decision Tree
For each learning, ask:

**1. Should this go in CONSTITUTION.md?**
- Does it affect how future increments work?
- Is it a recurring architectural decision?
→ Draft addition, show section placement

**2. Should this be an ADR?**
- Is the decision non-obvious?
- Will someone wonder "why did they do it this way?"
- Are there significant trade-offs to document?
→ Draft ADR using template

**3. Should this be an API contract?**
- Is this a public interface?
- Does it need versioning/documentation?
→ Draft OpenAPI/JSON Schema, place in `docs/api/`

**4. Should this update README?**
- Did the project's purpose or scope change?
- Is there new setup/usage information?
→ Draft README section addition

**5. Is this a backlog item?**
- Future work not ready to commit?
→ Suggest creating GitHub issue

### ADR Template
```markdown
# ADR: [Decision Title]

## Context
[Situation that led to this decision]

## Decision
[What we decided, clearly stated]

## Consequences
- **Benefits:** [what we gain]
- **Drawbacks:** [what we lose]
- **Trade-offs:** [what we accept]

## Alternatives Considered
- [Option A]: [why not chosen]
- [Option B]: [why not chosen]
```

### Process Flow
```
1. Read learnings.md
   Parse each learning item

2. For each learning:
   Ask promotion questions (see decision tree)
   Wait for user decision
   
3. If promoting:
   Draft the addition/document
   Show exact placement
   Ask: "Confirm?"
   Wait for confirmation
   
4. After all promotions:
   Summarize what was promoted
   Ask: "Ready to delete .4dc/current/?"
   Wait for explicit "yes"
   
5. Confirm deletion instructions
   (User deletes manually: rm -rf .4dc/current/)
```

### Example Questions
- "You discovered 'Use SHA256 for tokens, bcrypt for passwords'—should this go in CONSTITUTION.md Security section?"
- "You decided 'Synchronous email for v1'—should this be an ADR explaining the trade-off?"
- "You created POST /auth/reset-password—should this be an OpenAPI spec in docs/api/auth/?"
- "All learnings promoted. Ready to delete .4dc/current/?"

---

## `reflect.prompt.md`

### Core Purpose
Periodic codebase health assessment through quality lenses, identifying concrete refactorings that become future increments.

### Key Behaviors
- Guide user through systematic reflection using lenses
- Ask specific questions per lens (not generic "is code good?")
- Identify patterns (good and bad) across the codebase
- Propose concrete, small refactorings
- Decide if patterns should become constitutional rules or ADRs
- Create backlog of improvement increments

### Input Context
- Project path
- `CONSTITUTION.md` (to evaluate alignment)
- `docs/adr/` (to understand past decisions)
- Existing code + tests (to assess current state)
- Recent commits (to see what changed)

### Output
- Updates to `CONSTITUTION.md` (if patterns should become rules)
- New ADRs (if emerging patterns need alignment)
- New increment ideas (for refactorings)
- Backlog items (future improvements)

### Quality Lenses
These are defined IN THIS PROMPT, not in CONSTITUTION.md.

**1. Naming & Clarity**
- Are names aligned with domain language?
- Do names reveal intent?
- Are abbreviations clear or cryptic?

**2. Modularity & Separation**
- Are boundaries clear between components?
- Can you change one part without touching many others?
- Is coupling low, cohesion high?

**3. Architecture & Patterns**
- Is there a simple, explainable architecture?
- Are patterns applied consistently?
- Does code follow constitutional decisions?

**4. Testing & Reliability**
- Do tests give fast, meaningful feedback?
- Are critical paths covered?
- Are tests brittle or robust?

**5. Duplication & Simplicity**
- Is there copy-pasted logic to consolidate?
- Are abstractions justified or speculative?
- Is code as simple as it can be?

**6. Documentation & Communication**
- Do readers understand why decisions were made?
- Are critical workflows explained?
- Are invariants documented?

**7. Delivery & Flow**
- How easy to get changes into production?
- Are there manual, brittle steps?
- Are PRs appropriately sized?

**8. Dependencies & Operability**
- Are dependencies chosen consciously?
- Do logs/metrics help debug production issues?
- Is configuration explicit and documented?

### Process Flow
```
1. STOP 1: Context understanding
   Ask: "What's changed since last reflection?"
   Ask: "Any areas of pain or slowness?"
   Present summary → wait for confirmation

2. For each lens:
   Ask specific questions
   User answers with examples
   Identify patterns (good and bad)

3. STOP 2: Pattern summary
   Present observed patterns
   Ask: "Which patterns should become rules?"
   Ask: "Which patterns need ADRs?"
   Wait for decisions

4. Propose refactorings
   For each pain point, suggest concrete fix
   Scope small enough for one increment
   Ask: "Should this be a new increment?"

5. Promotion decisions
   Update CONSTITUTION.md if patterns → rules
   Create ADRs if divergent implementations need alignment
   Create increment ideas for refactorings
```

### Example Questions Per Lens

**Naming:**
- "Are names like `UserService` and `UserManager` clearly distinct?"
- "Do domain terms match what the business calls them?"

**Modularity:**
- "Can you change authentication without touching billing?"
- "Are there circular dependencies between modules?"

**Architecture:**
- "The constitution says domain is in src/domain/—is that being followed?"
- "Error handling differs between auth and billing—should we align it?"

**Testing:**
- "Do tests run in <10s as constitution requires?"
- "When a test fails, is it obvious what broke?"

**Duplication:**
- "This validation logic appears in 3 places—should we consolidate?"

**Documentation:**
- "Why was JWT chosen over sessions? Is that documented?"

**Delivery:**
- "What's the most painful manual step in deployment?"

**Dependencies:**
- "Are we wrapping external HTTP clients per constitution?"

### Refactoring Proposal Format
```markdown
## Refactoring: [Short Title]

**Lens:** [which lens identified this]
**Pain Point:** [what's currently difficult]
**Proposal:** [concrete change]
**Effort:** [rough estimate: 1h, half-day, 2 days]
**Value:** [what improves]

**Promote to:**
- [ ] CONSTITUTION (if recurring pattern)
- [ ] ADR (if needs explanation)
- [ ] New increment (if should be done)
- [ ] Backlog (if nice-to-have)
```

### Anti-Patterns to Guard Against
- Generating a report no one reads → focus on actionable refactorings
- Suggesting "rewrite everything" → scope small increments
- Abstract quality scores → concrete examples from code
- Lenses in constitution → they belong HERE, not there

---

## Cross-Prompt Consistency

### Shared Principles (for all prompts)
- **One dimension at a time**: Don't ask compound questions
- **Make constraints explicit**: "Given [constitution rule X], does this fit?"
- **Challenge vagueness**: "Flexible for what specific need?"
- **Confirm decisions explicitly**: Restate before moving on
- **STOP gates**: Wait for explicit user approval before major steps

### Shared Anti-Patterns (for all prompts)
- Generating solutions before questions
- Accepting vague answers
- Meta-commentary about process
- Mentioning prompts/LLMs/AI in output artifacts
- Moving forward without confirmation at STOP gates

### File Paths (consistency across prompts)
- CONSTITUTION: always at project root
- ADRs: always `docs/adr/ADR-YYYY-MM-DD-slug.md`
- API contracts: always `docs/api/` (organized by domain)
- Increment context: always `.4dc/current/`
- Temporary files: always in `.4dc/current/`

### Output Tone (consistency across prompts)
- Outcome-first, minimal chatter
- Crisp acknowledgments only when warm context
- No repeated "got it" or "I understand"
- Tight, focused responses
- No extra sections unless in template

---

Save this document, and when you're ready to write a prompt, tell me which one and I'll help you draft it! 🎯