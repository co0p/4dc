<p align="center">
  <picture>
    <source srcset="assets/4dc-logo-traced-dark.svg" media="(prefers-color-scheme: dark)">
    <img src="assets/4dc-logo-traced.svg" alt="4DC logo" width="120" />
  </picture>
</p>

<p align="center"><strong>4dc – four document cascade</strong></p>
<p align="center"><em>Discovery-driven prompts for emergent, test-driven development.</em></p>

---

> [!CAUTION]
> DO NOT TRUST AI, and do not trust others' prompts. Running prompts from untrusted sources is a security risk. Watch https://media.ccc.de/v/39c3-agentic-probllms

## Executive Summary

4dc is a set of **discovery-driven prompts** for building software through Socratic dialogue with an LLM. The prompts ask the right questions at the right time—helping you discover what to build and letting design emerge from TDD. You stay in control; the LLM acts as a pair-programming navigator.

**Ephemeral context, permanent knowledge.** Feature work lives in `.4dc/current/` as temporary scaffolding. Before merging, you promote learnings to permanent docs—`CONSTITUTION.md` for decisions, `DESIGN.md` for emergent architecture, ADRs for trade-offs. After merge, increment context is deleted.

**Extreme programming principles:** small increments, emergent design, continuous testing, working code as truth. Design doesn't happen in a separate phase; it emerges from TDD cycles.

---

## Installation

From the root of your project, run:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install-4dc-prompts.sh)"
```

This copies the prompt files into `.github/prompts/` with a `4dc-` prefix:

- `.github/prompts/4dc-constitution.prompt.md`
- `.github/prompts/4dc-increment.prompt.md`
- `.github/prompts/4dc-implement.prompt.md`
- `.github/prompts/4dc-promote.prompt.md`
- `.github/prompts/4dc-reflect.prompt.md`

It also creates the `.4dc/` working directory and adds it to `.gitignore`.

Reference these prompts in GitHub Copilot Chat, Cursor, or any LLM that supports prompt files.

---

## Quick Example

### 1. Create Your Constitution (One-Time Setup)

```bash
# Reference the constitution prompt and answer questions about your project
#4dc-constitution.prompt.md .
```

The LLM asks questions about your architectural decisions:
- "Where should domain logic live relative to UI code?"
- "How do you handle errors?"
- "What's your minimum testing expectation?"

You answer based on your project's reality. The result is `CONSTITUTION.md` with concrete decisions:

```markdown
# CONSTITUTION.md

## Layering
- Domain logic: `src/domain/`
- API routes: `src/api/`  
- Infrastructure: `src/infra/`

## Error Handling
- Domain returns `Result<T, Error>`
- API converts to HTTP status codes

## Testing
- Unit tests colocated with code
- Integration tests in `tests/integration/`
- All tests run in <10s locally
```

### 2. Start a Feature

```bash
# Create working context
mkdir -p .4dc/current

# Run increment prompt with your feature idea
#4dc-increment.prompt.md "add password reset"
```

The LLM asks discovery questions:
- "What's the smallest outcome that provides value?"
- "What's explicitly out of scope?"
- "How will you know it worked?"

Together you develop a user story, acceptance criteria, and slice the work into deliverables. The result is `.4dc/current/increment.md` (temporary LLM context).

### 3. Implement via TDD

```bash
# Run implement prompt
#4dc-implement.prompt.md
```

The LLM guides you through test-driven cycles:

```
LLM: "What's the first test for token generation?"
You: "Test that tokens are 32 bytes"
LLM: "Write a test that calls generate_token() and asserts length"

You: [writes test, runs it - fails]

LLM: "Good - failing for the right reason. What's the simplest 
      implementation that makes it pass?"
You: [implements, test passes]

LLM: "With tests green, what smells bad? Should we refactor?"
You: "The function name is unclear"
LLM: "What would be clearer?"
You: "generate_reset_token"
LLM: "Refactor the name, verify tests still pass"
```

As you work, the LLM appends important discoveries to `.4dc/current/learnings.md`.

### 4. Promote Learnings Before Merging

```bash
#4dc-promote.prompt.md
```

The LLM reads `.4dc/current/learnings.md` and asks what should be permanent:

```
LLM: "You discovered: 'Use SHA256 for tokens, bcrypt for passwords'
      Should this go in CONSTITUTION.md?"
You: "Yes"

LLM: [drafts addition to CONSTITUTION.md Security section]
     "Confirm?"
You: "Yes"

LLM: "You decided: 'Synchronous email delivery for v1'
      Should this be an ADR?"
You: "Yes"

LLM: [creates docs/adr/ADR-2025-01-26-sync-email.md]

LLM: "All learnings promoted. Ready to delete .4dc/current/?"
You: "Yes"
```

```bash
# Clean up and merge
rm -rf .4dc/current/
git add CONSTITUTION.md docs/ src/ tests/
git commit -m "Add password reset"
git push
```

---

## The Prompts

### constitution

**Purpose:** Document your project's architectural decisions.

**When to use:** At project start, or when adopting 4dc.

**How it works:** Asks concrete questions: Where should layers be separated? How do you handle errors? What's your testing strategy? You answer based on your project's reality. The LLM synthesizes answers into `CONSTITUTION.md`.

**What it creates:** Layering rules, error handling patterns, testing expectations, artifact layout. This becomes the reference for all other prompts.

**What it does NOT create:** Abstract values, style guides (use linters), or quality lenses (those belong in reflect).

**Output:** `CONSTITUTION.md` (permanent, evolves with project)

---

### increment

**Purpose:** Slice a feature idea into shippable deliverables with testable acceptance criteria.

**When to use:** When starting new work—a feature, bug fix, refactoring, or exploration.

**How it works:** Turns a vague idea into a concrete plan through Socratic questioning. Challenges scope creep and vague criteria. Produces a user story, acceptance criteria, use case, and deliverable slices.

**Acceptance test stubs:** For each acceptance criterion, generates a greppable test name:

```
Test<Feature>_Given<Context>_When<Action>_Then<Result>
```

Example:
| Acceptance Criterion | Test Stub |
|---------------------|----------|
| Given idle, when Start clicked, then countdown begins | `TestPomodoro_GivenIdle_WhenStartClicked_ThenCountdownBegins` |

This creates 1:1 traceability between ACs and tests. Find all tests for a feature: `grep -r "TestPomodoro_Given" pkg/`

**Deliverable slicing:** Breaks work into small, independently shippable pieces. Each deliverable goes through its own TDD cycle; learnings from one inform the next.

**Output:** `.4dc/current/increment.md` (temporary, deleted after merge)

---

### implement

**Purpose:** Guide TDD cycles, one deliverable at a time.

**When to use:** After defining an increment, when ready to write code.

**How it works:** Pair-programming navigator. Suggests the next smallest test. After you write a failing test: "Is this failing for the right reason?" After green: "What's the simplest implementation?" During refactor: suggests improvements aligned with constitution.

**Design emerges through questioning:** Every 5-10 cycles, asks: "Have we discovered anything to promote?" The code and tests ARE the design; the prompt helps you recognize when insights should become permanent.

**Output:** Working code + tests (permanent), `.4dc/current/learnings.md` (promotion candidates, temporary)

---

### promote

**Purpose:** Promote learnings to permanent docs before deleting increment context.

**When to use:** Before merging, after completing an increment.

**How it works:** Reads `.4dc/current/learnings.md` and asks what should be promoted:

| Learning Type | Destination |
|--------------|-------------|
| Architectural decision | `CONSTITUTION.md` |
| Emergent design pattern | `DESIGN.md` |
| Non-obvious trade-off | `docs/adr/` |
| Public interface | `docs/api/` |
| Future work | GitHub issue |

**DESIGN.md:** Documents architecture that **emerged from TDD**, not planned upfront. Updated each increment with patterns discovered, module structure evolution, and open questions. Retrospective, not prescriptive.

**After promotion:** Confirms all learnings captured, then: delete `.4dc/current/`, commit permanent additions, merge.

**Output:** Updates to `CONSTITUTION.md`, `DESIGN.md`, new ADRs, API contracts (ephemeral context deleted)

---

### reflect

**Purpose:** Periodic codebase health check to identify refactoring opportunities.

**When to use:** After several increments, when velocity slows, or when code feels painful.

**How it works:** Guides assessment using **quality lenses**: naming, modularity, architecture, testing, duplication, documentation, delivery. For each lens: "Are names aligned with domain language?" "Do tests give fast feedback?"

**From assessment to action:** Identifies concrete refactorings scoped small enough to become increments.

**Output:** Updates to `CONSTITUTION.md`, `DESIGN.md`, new ADRs, new increment ideas

---

## Repository Structure

```
my-project/
├── CONSTITUTION.md              # Permanent: architectural decisions
├── DESIGN.md                    # Permanent: emergent architecture (updated each increment)
├── README.md                    # Permanent: project overview
│
├── docs/
│   ├── adr/                     # Permanent: decision records
│   │   └── ADR-2025-01-26-sync-email.md
│   └── api/                     # Permanent: contracts, schemas
│       └── auth/
│           └── password-reset.openapi.yaml
│
├── src/                         # Permanent: code
├── tests/                       # Permanent: tests
│
└── .4dc/                        # Temporary: working context
    └── current/                 # Deleted after merge
        ├── increment.md         # What you're building + test stubs
        └── learnings.md         # Promotion candidates
```

**.gitignore:**
```
.4dc/
```

---

## What Lives Where

| Artifact | Location | Lifecycle | Purpose |
|----------|----------|-----------|---------|
| **CONSTITUTION.md** | Root | Permanent | Architectural decisions |
| **DESIGN.md** | Root | Permanent | Emergent architecture (what TDD discovered) |
| **README.md** | Root | Permanent | Project overview |
| **ADRs** | `docs/adr/` | Permanent | Trade-off explanations |
| **API Contracts** | `docs/api/` | Permanent | Public interfaces |
| **Code + Tests** | `src/`, `tests/` | Permanent | The implementation |
| **Increment context** | `.4dc/current/increment.md` | Temporary | User story, ACs, test stubs |
| **Learnings** | `.4dc/current/learnings.md` | Temporary | Promotion candidates |

**Key principle:** Increment artifacts are ephemeral. After merge, only promoted decisions and working code remain.

---

## Workflow Diagram

```mermaid
graph TB
    Start([Feature Idea])
    
    Constitution[CONSTITUTION.md<br/>Architectural decisions]
    
    Start --> Inc{increment prompt}
    Inc -->|Discover WHAT,<br/>slice deliverables| IncDoc[.4dc/current/<br/>increment.md]
    
    IncDoc --> D1[Deliverable 1:<br/>implement prompt]
    D1 -->|TDD cycles| D1Code[Code + Tests<br/>+ learnings.md]
    D1Code --> D1Done{Deliverable done?}
    D1Done -->|More tests| D1
    D1Done -->|Yes| D2
    
    D2[Deliverable 2:<br/>implement prompt<br/><i>informed by D1</i>]
    D2 -->|TDD cycles| D2Code[Code + Tests<br/>+ learnings.md]
    D2Code --> D2Done{Deliverable done?}
    D2Done -->|More tests| D2
    D2Done -->|Yes| D3
    
    D3[Deliverable 3:<br/>implement prompt<br/><i>informed by D1+D2</i>]
    D3 --> D3Code[Code + Tests<br/>+ learnings.md]
    D3Code --> AllDone{All deliverables<br/>complete?}
    
    AllDone -->|Yes| Promote{promote prompt}
    Promote -->|Update| ConstUpdate[CONSTITUTION.md]
    Promote -->|Create| ADR[docs/adr/]
    Promote -->|Create| API[docs/api/]
    
    Promote --> Cleanup[Delete .4dc/current/]
    Cleanup --> Merge[Merge PR]
    
    Merge --> Reflect{reflect prompt<br/><i>periodic</i>}
    Reflect -->|Identify refactorings| NextInc[Next increments]
    NextInc --> Inc
    
    Constitution -.->|Guides| D1
    Constitution -.->|Guides| D2
    Constitution -.->|Guides| D3
    Constitution -.->|Evaluates against| Reflect
    
    style IncDoc fill:#fff4e1
    style D1Code fill:#e1ffe1
    style D2Code fill:#e1ffe1
    style D3Code fill:#e1ffe1
    style Constitution fill:#e1f5ff
```

---

## Why This Works

**Emergent design:** Design happens during TDD, not before. Each deliverable informs the next.

**Working code as truth:** Code and tests are the design. Docs capture only decisions that guide future work.

**Fast feedback:** Deliverables are small (hours to days). Ship frequently, learn quickly.

**No stale docs:** Increment context deleted after merge. Only promoted learnings remain.

**Traceability:** Greppable test names map 1:1 to acceptance criteria. DESIGN.md tracks what emerged.

**LLM as navigator:** You decide. The LLM questions, challenges, ensures TDD discipline.

---

## Getting Started

1. **Install the prompts** (see Installation above)
2. **Create your constitution**: `#4dc-constitution.prompt.md .`
3. **Start your first increment**: `#4dc-increment.prompt.md "your feature idea"`
4. **Implement via TDD**: `#4dc-implement.prompt.md`
5. **Promote learnings before merge**: `#4dc-promote.prompt.md`
6. **Periodically reflect**: `#4dc-reflect.prompt.md` (after several increments)