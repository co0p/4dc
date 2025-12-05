# ADR Output Template

Use this format for architectural decisions that emerge from the improve phase and are broadly relevant to the project. Reference this template from improvement plans when an ADR is required.

## ADR: [Decision Title]

### Context

Describe the situation, problem, or pattern that led to this decision.

### Decision

State the architectural decision clearly and concisely.

### Consequences

- List the benefits, drawbacks, and trade-offs resulting from this decision.
- Note any impacts on maintainability, extensibility, or performance.

### Alternatives Considered

- [Alternative approach]: Reason not chosen
- [Another alternative]: Reason not chosen

---

**Example:**

```markdown
# ADR: Centralize Error Handling in Catalog Module

## Context
Error handling was previously scattered across multiple components, leading to inconsistent behavior and duplicated logic.

## Decision
Centralize all error handling for catalog features in a dedicated module, with standardized error messages and handling routines.

## Consequences
- Improved consistency and maintainability
- Easier to test and extend error handling
- Minor refactoring required for existing components

## Alternatives Considered
- Keep error handling decentralized: Simpler now, but harder to maintain
- Use a third-party error handling library: Adds complexity and dependencies
```