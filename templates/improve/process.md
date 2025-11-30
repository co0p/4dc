
# Improvement Process (Codebase-Wide)

## 1. Receive Initial Improvement Request
Inform the user: "You have requested a refactoring and codebase improvement analysis."

## 2. Analyze Project Context and Assess Implementation
Inform the user: "I will now review your project files (especially README.md, CONSTITUTION.md, and recent increments/designs) as well as ADRs to understand the technical landscape and assess the implementation."

### Assessment Tasks
- **Evaluate vs. Constitution:** Assess how well the implementation adheres to the project's core principles and constraints.
- **Evaluate vs. Design Goals:** Assess whether the implementation meets the intended design approach, component boundaries, and data flow.
- **Quality Evaluation:** Assess code quality, readability, maintainability, and testability.
- **Identify Risks:** List technical debt, potential bugs, performance concerns, or security issues.
- **Identify Architectural Opportunities:** Note opportunities for improved structure, patterns, or abstractions.

### Summary of Findings
After context analysis and assessment, provide a brief summary to the user outlining the project's purpose, tech stack, and key assessment findings.

## 3. Analyse Codebase Through Lenses
After analyzing the project context, identify improvement opportunities by examining the codebase through the context-based lenses described in lenses.md:
   - Naming & Clarity
   - Modularity & Separation
   - Architecture & Patterns
   - Testing & Reliability
   - Duplication & Simplicity
   - Documentation & Communication

- Identify and list actionable improvement suggestions relevant to the codebase, grouped by these lens contexts.
- For each suggestion, reference the relevant lens group and provide a clear rationale inspired by industry best practices.
Use these context-based lenses as the organizing principle for improvement suggestions, guiding the analysis and recommendations.
Do not ask the user what to look for; instead, use your analysis and the lenses to recommend actionable improvements.

### Document Lessons Learned
During the lens analysis, identify and document:
- **What Worked Well:** Patterns, approaches, or decisions that proved effective.
- **What Could Be Improved:** Areas where the implementation fell short or could be enhanced.
- **Unexpected Challenges:** Issues that arose during implementation and how they were addressed.
- **Emerging Patterns:** Recurring solutions or approaches that could be standardized.

### Surface ADR Candidates
When divergent implementations or emerging patterns appear (e.g., different approaches to user validation, error handling, or component structure), surface these as potential ADR candidates:
- Describe the observation (divergent approaches or emerging pattern).
- Recommend whether an ADR should be created.
- Provide rationale for the recommendation.

If the user agrees, create the ADR file separately using the ADR Output Template in `adr.md`. ADRs are independent artifacts and should not be listed in the improve.md output.

ADRs should only be created for broadly relevant patterns or architectural decisions, not for trivial or stylistic changes.

## 4. Present Lenses-Based Findings

Inform the user: "Based on my analysis through the lenses, here are the factual findings and recommended improvements."

Present each finding as a factual observation from the codebase:
- State the observation clearly with evidence (file references, code patterns).
- Map the finding to a specific lens context.
- Provide a clear recommendation based on industry best practices.

Do not ask the user what to look for or what to improve. Let the lenses and codebase analysis guide all recommendations.

### Summary of Findings
Provide a concise summary listing all findings, their mapped lens contexts, and the recommended actions.

## 5. Generate Improvement Plan
Inform the user: "I will generate the improvement document strictly following the improve output format. This document serves as a reference for the next increment cycle."

For the improve output, use the concise format described in output.md:
   - **Assessment:** Brief evaluation of Constitution/Design alignment, quality, and risks
   - **Lessons:** What worked well, what to improve, emerging patterns
   - **Improvements:** Each improvement as a separate section with explicit file references, lens, priority, and effort

Each improvement section must include:
   - Lens context (Naming/Modularity/Architecture/Testing/Duplication/Documentation)
   - Priority (H/M/L)
   - Effort estimate
   - Explicit file paths to change
   - Specific change description

This format ensures the output is parsable by the /increment prompt for follow-up implementation in a new increment cycle.

Note: ADRs are created as separate, independent artifacts when the user agrees. They are not part of the improve.md output.

Note: The improve phase does NOT implement the suggested improvements. The user decides whether to implement the suggestions by starting a new increment, design, implement cycle.

## 6. Save Improvement Plan
Inform the user: "I will save the generated improvement plan as improve.md in the project root."

**Action:** Write the improvement plan as `improve.md` in the project root directory. The file must contain:
- **Assessment:** Brief evaluation against Constitution/Design goals
- **Lessons:** Key learnings from the implementation
- **Improvements:** Each improvement as a separate section with file references

### Summary of Findings
Provide a brief summary confirming the plan was saved.

## 7. Final Validation
Inform the user: "Before saving, I will validate that all requirements are met."

**Verification Checklist:**
- Assessment section contains Constitution/Design alignment and risks
- Lessons section documents learnings and emerging patterns
- Each improvement is a separate section with explicit file references
- Improvements include lens, priority, effort, and change description

If anything is missing, STOP and ask for clarification or fixes.
