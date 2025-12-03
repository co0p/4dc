## Output

You MUST:

- Output only the constitution document in Markdown using the structure below.
- NOT include any meta commentary about what you could do next (for example, "If you'd like, I can also add...", "Next, I can create...", "I can generate a workflow").
- NOT include suggestions for additional files, CI workflows, or other automation tasks inside the constitution. Those may be implied by principles, but not offered as actions by you.

Return the result as **Markdown** with the following structure:

```markdown
# Engineering Constitution for {{team_or_product_name}}

## Purpose

Explain in 2–4 sentences:
- Why this CONSTITUTION exists for this project.
- How it should be used in everyday engineering work and decision-making.

## Context

Summarize the environment and constraints:
- Product / domain:
  - ...
- Team:
  - ...
- Non-negotiables:
  - ...

## Our Principles and Trade-offs

Explain the team’s overall philosophy and how it relates to:
- Speed vs safety
- Short-term delivery vs long-term maintainability
- Experimentation vs stability

### Default Trade-off Rules

- When in doubt between **shipping faster** and **polishing the design**, we usually:
  - ...
- When in doubt between **adding a dependency** and **building it ourselves**, we usually:
  - ...
- When in doubt between **adding tests now** and **moving on**, we usually:
  - ...

---

## The 6 Pillars of Our Engineering

### 1. Delivery Velocity

Describe how the team thinks about:
- Desired iteration speed.
- Typical size of changes.
- Release cadence and acceptable risk per release.

Include:

- **We optimize for:**
  - ...
- **We accept the following risks:**
  - ...
- **We avoid:**
  - ...

### 2. Test Strategy

Describe:
- What must be tested.
- How much coverage / confidence is “enough” for this project.
- Preferred testing strategies (e.g., unit vs integration vs end-to-end).

Include:

- **Minimum expectations:**
  - ...
- **When moving fast, we are allowed to:**
  - ...
- **We never skip tests for:**
  - ...

### 3. Design Integrity

Describe:
- How the team structures code and architecture.
- What “good boundaries” mean in this project.
- How to think about modules, responsibilities, and dependencies.

Include:

- **We strive for:**
  - ...
- **We are okay with:**
  - "...some messiness in leaf modules as long as boundaries remain clear."
- **Red flags that trigger redesign or refactoring:**
  - ...

### 4. Simplicity First

Describe:
- How the team avoids premature abstraction and over-engineering.
- How to decide when to introduce patterns, indirection, or generalization.

Include:

- **We prefer:**
  - "The simplest thing that could possibly work, then iterate."
- **We add abstraction only when:**
  - ...
- **We treat complexity as acceptable when:**
  - ...

### 5. Technical Debt Boundaries

Describe:
- When it is acceptable to take shortcuts.
- How debt is recorded and prioritized.
- How and when debt must be paid.

Include:

- **Allowed short-term shortcuts:**
  - ...
- **Debt must be recorded when:**
  - ...
- **We commit to paying down debt when:**
  - ...

### 6. Dependency Discipline

Describe:
- How the team chooses, isolates, and upgrades dependencies (libraries, frameworks, external services).
- What “good” vs “bad” dependency use looks like.

Include:

- **We add a new dependency only when:**
  - ...
- **We isolate dependencies by:**
  - ...
- **We avoid:**
  - "Frameworks bleeding into our domain model", etc.

---

## How We Use This Constitution

Explain briefly how this constitution should influence:

- How work is chosen and sliced.
- How designs are evaluated.
- How implementation and testing decisions are made.
- When to refactor, pay down debt, or revisit architecture.

Keep this section high-level and project-focused. Do not mention specific tooling, frameworks, or meta-processes used to apply the constitution (such as prompt systems, LLM workflows, or framework names).

---

## Amendments and Evolution

Describe:
- How this CONSTITUTION can be updated.
- Under what circumstances you expect to revisit it (e.g., major product shift, team growth, repeated friction).
- How amendments should be documented (e.g., dated changes, versioning).

---

## References and Inspirations

List key references that influenced this CONSTITUTION, such as:

- Books, articles, or talks that inspired your engineering approach.
- Internal documents or prior decisions that shaped these principles.

---

## Open Questions

List questions the team should explicitly revisit, for example:

- "What’s our acceptable MTTR vs MTBF trade-off?"
- "How strict should we be about mutation testing or coverage thresholds?"
- "What performance budgets matter most for our users?"

These should be concrete enough to guide future amendments.
```