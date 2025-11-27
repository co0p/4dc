# 4DC – 4 Document Cascade

A documentation system for clear, incremental, AI-assisted development, enabling robust, explicit architectural learning and evolution of the codebase.  
**Learn fast within safe bounds.**

---

## The Cycle: Increment, Design, Breakdown, Improve

The core workflow in 4DC is an explicit, repeatable learning loop:

1. **Increment (WHAT):**  
   Describe the user-focused value to deliver. Summarize requirements and acceptance criteria in `[increment-name]/increment.md`.

2. **Design (HOW – Initial Sketch):**  
   Outline a technical approach to guide implementation in `[increment-name]/design.md`.

3. **Breakdown (HOW – Detailed):**  
   Break the solution into small, verifiable steps in `[increment-name]/breakdown.md`. Complete each task, checking off progress as you implement.

4. **Improve:**  
   After acceptance criteria pass, refactor and identify patterns. Codify recurring decisions as Architecture Decision Records.

**The Constitution provides guardrails:**  
Define foundational principles and technical constraints once (`CONSTITUTION.md`). Every increment explores and experiments within these safe bounds. Architecture emerges from implementation, not upfront planning.

---

## Architectural Wisdom: ADRs

As you refactor and improve, record recurring patterns and tradeoffs as **Architecture Decision Records** (ADRs):

- **When?**  
  - A pattern appears across increments  
  - Project-wide implications  
  - New convention or trade-off
- **Where?**  
  - `design/[adr-number]-[topic].md`
- **Format:**  
  - Context, Decision, Consequences, Alternatives

Future increments reference both the constitution and accumulated ADRs.

---

## Example Directory

See real workflows in [`examples/`](examples/):

- **Todo App** (`examples/todo/`): Minimal browser-based todo with zero build steps.
- **EpicSum** (`examples/epicsum/`): Summing ticket hours using portable shell scripting.

Each example demonstrates the full documentation cycle, from constitution through implementation.

---

## Using Prompts in Visual Studio Code

4DC is designed for seamless AI-assistance via prompt files:

1. **Find prompt files** in the root or `.github/prompts/` directory (e.g. `create-constitution.prompt.md`, `increment.prompt.md`, etc).
2. **Invoke workflows** by referencing prompt files in Copilot Chat:
   ```
   @workspace /create-constitution
   @workspace /increment add todo item
   @workspace /design
   @workspace /breakdown
   @workspace /improve
   ```

3. **Integration tip:** For repository-wide discoverability, copy prompt files to `.github/prompts/` and reference as:
   ```
   #file:.github/prompts/create-feature.md
   ```

**For more see:**  
- [`prompts/README.md`](prompts/README.md)
- [VS Code Copilot Prompt Files Documentation](https://code.visualstudio.com/docs/copilot/customization/prompt-files)

---

## Inspirations

4DC synthesizes concepts from multiple software engineering thought leaders and methodologies:

### Core Philosophy
- **Kent Beck** – ["Make it work, make it good, make it fast"](https://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast) - The improve phase embodies this three-step cycle, emphasizing working software first, then refactoring for quality
- **Lean Software Development (Mary & Tom Poppendieck)** – Defer commitment, learn through building, eliminate waste - 4DC defers architectural decisions until patterns emerge from real implementation

### Documentation Patterns
- **Architecture Decision Records (Michael Nygard)** – [Documenting architectural decisions](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions) in a lightweight, searchable format - 4DC's `design/` folder captures emergent patterns as ADRs
- **Martin Fowler** – [Evolutionary Architecture](https://martinfowler.com/articles/designDead.html), incremental design, and refactoring discipline - Constitution vs. ADRs separation mirrors upfront vs. emergent decisions

### Requirements & Acceptance Testing
- **Gherkin/BDD (Behavior-Driven Development)** – Given/When/Then format for acceptance criteria - Used in `increment.md` to define testable user outcomes
- **Job Stories (Intercom/JTBD)** – ["When [situation], I want to [action], so I can [outcome]"](https://www.intercom.com/blog/using-job-stories-design-features-ui-ux/) - Captures user motivation more effectively than user stories

### Software Craftsmanship
- **Robert C. Martin (Uncle Bob)** – [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), SOLID principles, separation of concerns - Constitution's 6 Pillars framework maps to these foundational engineering practices
- **The Agile Manifesto** – Working software over comprehensive documentation - 4DC uses just enough documentation to enable autonomous development without bureaucracy