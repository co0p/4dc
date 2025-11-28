# Prompt Process for Improve Step

## 1. Receive Initial Prompt
Inform the user: "You have requested to improve the codebase and document architectural decisions."

## 2. Verify Prerequisites
Check for the existence of both `CONSTITUTION.md` and relevant increment and design documents. These define the project's principles, user-focused WHAT, and technical HOW.

## 3. Analyze Project Context
Review the constitution, increment, design, and existing code to identify patterns, trade-offs, and architectural decisions that have emerged from implementation.

## 4. Ask Clarifying Questions (STOP)
Inform the user: "I will ask 2-3 essential questions about recurring patterns, trade-offs, or decisions that need to be documented."
- What patterns have emerged in the code?
- What trade-offs were made during implementation?
- What decisions should be codified for future increments?

**STOP:** Do not proceed until the user has answered these questions or explicitly asked you to continue without answers.

## 5. Generate ADRs
Based on the answers and context, propose new Architecture Decision Records (ADRs) to document recurring patterns and decisions.

## 6. Save ADRs
Save the generated ADRs in the design folder, following the established ADR format.

## 7. Final Validation
Before saving, validate that the ADRs:
- Clearly document the decision, context, consequences, and alternatives
- Are concise and focused
- Are accessible and useful to future contributors
