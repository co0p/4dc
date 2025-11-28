

# Prompt Process for Design Step

## 1. Receive Initial Prompt
Inform the user: "You have requested a technical design for a feature increment."

## 2. Verify Prerequisites
Check for the existence of both `CONSTITUTION.md` and `[increment-name]/increment.md`. These documents define the project's principles and the user-focused WHAT.

## 3. Analyze Project Context
Review the constitution, increment, and any existing Architecture Decision Records (ADRs) to understand technical constraints, user goals, and acceptance criteria. Summarize findings: project purpose, tech stack, architectural patterns, constraints, and relevant prior decisions from ADRs.

## 4. Ask Technical Clarifying Questions (STOP)
Inform the user: "I will ask 2-3 essential technical questions about component boundaries, data flow, or integration for this increment."
- How should responsibilities be split for this feature?
- How should data flow through the system?
- How should this feature integrate with external services or storage?

**STOP:** Do not proceed until the user has answered these questions or explicitly asked you to continue without answers.

## 5. Generate Technical Design
Based on the answers and context, propose a lightweight, focused technical design for the increment. Document key technical decisions, trade-offs, and alternatives. Use terms like "initial technical design", "design outline", "design draft", or "design proposal" for clarity and robustness.

## 6. Save Design
Save the generated design as `design.md` in the increment's directory.

## 7. Final Validation
Before saving, validate that the technical design:
- Addresses the increment's acceptance criteria
- Respects constitutional principles and constraints
- Documents 2-5 key technical considerations
- States trade-offs and alternatives
- Is concise and focused (one screen max)

