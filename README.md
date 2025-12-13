<p align="center">
  <picture>
    <source srcset="assets/4dc-logo-traced-dark.svg" media="(prefers-color-scheme: dark)">
    <img src="assets/4dc-logo-traced.svg" alt="4DC logo" width="120" />
  </picture>
</p>

<p align="center"><strong>4dc – four document cascade</strong></p>
<p align="center"><em>Prompt-driven 4-document workflow for incremental, evolutionary software development.</em></p>

---

4dc is a set of **LLM‑friendly prompt files** you add to your project so that you and your AI tools work in **small, incremental, evolutionary steps** instead of big risky changes.

You describe **what and why** in an increment, get a **design** for how the system will change, break it down into **concrete implementation steps**, and periodically run **improve passes** that feed back into future increments.

All of this is guided by a **project constitution** (`CONSTITUTION.md`) that encodes your values, expectations, and layout so your tools behave more like a consistent teammate than a loose cannon.

### Installation

This assumes:

- You already have a Git repo for your project.
- You want the 4dc prompt files under `.github/prompts/4dc` in **your** project.

From the root of your project, run:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/co0p/4dc/main/scripts/install-4dc-prompts.sh)"
```

This script:

- Downloads the installer from this repo.
- Runs it in your current project directory.
- Copies the assembled `*.prompt.md` files into `.github/prompts/4dc/`.

Afterwards, you’ll have:

- `.github/prompts/4dc/create-constitution.prompt.md`
- `.github/prompts/4dc/increment.prompt.md`
- `.github/prompts/4dc/design.prompt.md`
- `.github/prompts/4dc/implement.prompt.md`
- `.github/prompts/4dc/improve.prompt.md`

Point Copilot Chat or your LLM at these prompt files as sources, and you’re ready to start using 4dc.

---

## Details

### The cycle (four document cascade)

4dc keeps four kinds of artifacts separate so humans and tools don’t blur concerns. In practice, the cascade looks like this:

```text
      ┌──────────────────────────────────────────┐
      │          CONSTITUTION                    │
      │  values, expectations, layout            │
      └──────────────────────────────────────────┘
       ↑            ↑            ↑            ↑
       │            │            │            │
[Increment]  →  [Design]  →  [Implement]  →  [Improve]
  WHAT & WHY     TECH HOW     STEPS          LEARN

  Small,         How the      Concrete       Codebase-wide
  outcome        system       changes and    assessment,
  focused        should       testable       lessons, and
  change         evolve       tasks          next ideas
```

In words:

1. **Increment – WHAT & WHY**  
   You define the outcome you want and why it matters, without jumping into solutions or task lists.

2. **Design – HOW (technically)**  
   You describe how the system will change to achieve that increment: components, data flows, interfaces, and trade‑offs.  
   Design stays at the **architecture and contract level**: it defines what components/modules do, how they interact, and what data/interface shapes they use.  
   It specifies testing/CI/observability as **constraints and targets** (e.g., "must have integration tests covering this flow", "must log request IDs"), not as step-by-step instructions.  
   Design **does not** list file paths, per-file actions, chronological sequences, or task lists; those belong in Implement.

3. **Implement – STEPS (concrete tasks)**  
   You break the design into small, testable tasks and suggested PR groupings you can actually execute.  
   Implement **consumes** the design's contracts, interfaces, and data shapes without redefining them. It turns the design into **ordered, file-specific work items** that name modules, functions, and tests to add or change. If the design seems incomplete or problematic, Implement flags gaps as risks or follow-up increments rather than silently inventing new contracts.

4. **Improve – LEARN & REFINE (codebase‑wide)**  
   You periodically scan the codebase against your constitution, capture lessons, and propose refactors and ADRs that become future increments.

Everything is guided by your **Constitution** (`CONSTITUTION.md`), which encodes:

- Values and principles (e.g., small, incremental, evolutionary changes; refactoring as everyday work).
- Testing and observability expectations.
- Dependency and architecture rules.
- Where docs and artifacts live in the repo.

### Principles & values

Your constitution is where you make your implicit “this is how we work” rules explicit – for humans and for the LLM. Typical themes:

- **Small, safe steps**  
  Prefer many small, reversible changes over a few large, risky ones. This echoes incremental and evolutionary design: keep the system always in a working state, learn from each step, and avoid “big bang” rewrites.

- **Refactoring as everyday work**  
  Don’t treat refactoring as a separate phase. Improving names, extracting functions, clarifying boundaries, and simplifying flows are just part of doing the work. 4dc makes space for this by:
  - Capturing refactor ideas in Improve docs.
  - Turning them into future increments instead of “someday/maybe” items.
  - Letting your tools propose refactors in a controlled way.

- **Clear contracts and boundaries**  
  Design documents focus on the seams between components: interfaces, data shapes, and invariants. This is aligned with domain‑driven thinking: reflect the domain language in boundaries and make cross‑component communication explicit.

- **Evidence‑friendly delivery**  
  Testing, observability, and deployment are not afterthoughts:
  - Increments describe acceptance criteria and what “good” looks like in production.
  - Designs call out where to add or improve tests, metrics, and logs.
  - Implementation plans include checks that support fast, confident delivery.

- **Sustainable flow over heroics**  
  The goal is steady, predictable change: short feedback cycles, low rework, and fewer surprises. The Improve docs help you spot bottlenecks and sources of pain so you can adjust how you work, not just what you build.

When you write `CONSTITUTION.md`, you’re essentially answering:

> “If this codebase had a senior teammate who cared about long‑term health, what rules would they keep repeating to us and to the AI?”

Those principles then drive how each 4dc prompt behaves when it generates increments, designs, implementation plans, and improve passes.

### Lenses (how to look at the codebase)

When you run an Improve pass, you don’t just ask “is the code good?” – you look at it through a set of **lenses**. Each lens is a way to inspect the system, its design, and how it’s delivered:

- **Naming & Clarity**  
  Are names (types, functions, modules, tests) clear and aligned with the domain language?  
  Does reading a file feel like reading a coherent explanation, or a puzzle?

- **Modularity & Separation**  
  Are boundaries between components clear, or are responsibilities tangled?  
  Can you change one part of the system without touching half the repo?

- **Architecture & Patterns**  
  Is there a simple, explainable architecture behind the code?  
  Are patterns (layering, message flows, aggregates, adapters, etc.) applied consistently where they help – and avoided where they would add accidental complexity?

- **Testing & Reliability**  
  Do tests give fast, meaningful feedback, or are they flaky / slow / rare?  
  When something breaks, can you find it quickly and with confidence?

- **Duplication & Simplicity**  
  Is there copy‑pasted logic or repeated patterns that could be clarified or extracted?  
  Are abstractions justified by real use, or speculative and fragile?

- **Documentation & Communication**  
  Do readers understand why key decisions were made (design docs, ADRs, comments)?  
  Are the most important workflows and invariants explained somewhere you can find them?

- **Delivery & Flow**  
  How easy is it to get a change from “idea” to “running in production”?  
  Are there manual, brittle steps in build, test, or deployment that could be simplified or automated?

- **Dependencies & Operability**  
  Are external dependencies (libraries, services, infra) chosen consciously?  
  When the system misbehaves, do logs, metrics, and traces help you understand what’s going on?

An Improve run uses these lenses to:

1. **Assess** – Give a rough, shared sense of where you are today (often with simple 1–5 star ratings and short rationales per lens).
2. **Capture lessons** – What’s working well, what keeps hurting, and what patterns are emerging.
3. **Propose improvements** – Concrete, scoped improvements (refactors, tests, documentation, infrastructure changes) that can become future increments or ADRs.

Because these lenses are written down in your constitution, your LLM can apply them consistently – not just once, but over and over as the system and team evolve.

### The prompts

All assembled prompt entrypoints live at the repo root (and get copied into `.github/prompts/4dc/` in your project):

- `create-constitution.prompt.md`  
  Helps you create or refine `CONSTITUTION.md` for a repo or subproject. Encodes values, principles, test/observability expectations, dependency rules, and doc layout.

- `increment.prompt.md`  
  Defines a **small, outcome‑focused change**:
  - Context & goal.
  - Scope and non‑goals.
  - Acceptance criteria & observability.
  - Risks and open questions.  
  Stays strictly at the WHAT/WHY level; no solution or task list.

- `design.prompt.md`  
  Takes an increment folder (with `increment.md`) and generates `design.md`:
  - How the system will change.
  - Interfaces, data flows, components.
  - Key decisions and trade‑offs.  
  Does not produce tasks.

- `implement.prompt.md`  
  Takes an increment folder (with `increment.md` and `design.md`) and generates `implement.md`:
  - Steps grouped into workstreams (e.g., backend, UI, tests).
  - Small, testable tasks with concrete file actions.
  - Suggested grouping into phases/PRs.

- `improve.prompt.md`  
  Runs a **codebase‑wide improve pass** for a project rooted at `path` and produces `docs/improve/YYYY-MM-DD-improve.md` with:
  - Assessment vs. constitution (1–5 star ratings + rationales).
  - Lessons (worked well / to improve / emerging patterns).
  - Concrete improvements (refactoring proposals) plus ADR candidates.

You can wire these prompt files into Copilot Chat “Custom Instructions”, or any LLM integration that can load prompts from disk.
