

# Design Output Format

The generated design document should include the following sections:

## 1. Design Summary
2-3 sentences describing what is being built and the initial technical approach.
Include short references to relevant principles from the constitution and any existing ADRs if applicable.

## 2. Initial Approach
Document 2-5 key technical considerations, each with:
- Approach
- Rationale
- Trade-offs
- Alternatives to consider
Reference constitution principles or ADRs for rationale or constraints where relevant.

## 3. Architecture Overview
- Components and their responsibilities
- Data flow (brief description or diagram)
- Integration points (external services/APIs)
- State management (where state lives and why)
Include references to constitution or ADRs for architectural patterns or decisions if applicable.

## 4. Implementation Constraints
List any technical constraints or limitations from these decisions.
Reference constitution or ADRs for constraints if relevant.

## 5. Open Questions
Technical unknowns or deferred decisions to resolve during implementation.
Reference constitution or ADRs for open questions or areas needing future decisions if applicable.

---

**Example Structure:**

```markdown
# Design: [Increment Name]

**Date:** [YYYY-MM-DD]  
**Status:** Initial Technical Design

## Design Summary
[2-3 sentences: What we're building and the initial technical approach to try]
[Reference: Constitution Principle X, ADR-001]

## Technical Decisions
- **Framework:** [e.g., React]
	- **Rationale:** [Why this framework fits the increment and constitution]
	- **Trade-offs:** [Pros/cons for this increment]
	- **Alternatives Considered:** [Other frameworks and why not chosen]
- **Language:** [e.g., Go]
	- **Rationale:** [Why this language fits]
	- **Trade-offs:** [Pros/cons]
	- **Alternatives Considered:** [Other languages]
[Repeat for other major technology choices]

## Initial Approach

### [Design Consideration]
**Approach:** [What we'll try first]  
**Rationale:** [Why this seems reasonable - technical reasoning]  
**Trade-offs:** [What we're accepting vs. what we're gaining]  
**Alternatives to Consider:** [Other approaches we might pivot to]
[Reference: Constitution Principle Y, ADR-002]

[Repeat for 2-5 key considerations only]

## Architecture Overview

**Components:**
- [Component name]: [Responsibility for this increment]
- [Component name]: [Responsibility]

**Data Flow:**
[Brief description or simple diagram of how data moves]
[Reference: ADR-003]

**Integration Points:**
- [External service/API]: [How we integrate]

**State Management:**
[Where state lives and why for this increment]

## Implementation Constraints
- [Technical constraint or limitation from these decisions]
- [Another constraint developers need to know]
[Reference: Constitution Principle Z]

## Open Questions
- [Technical unknown to resolve during implementation]
- [Deferred decision with reason]
[Reference: ADR-004]
```
