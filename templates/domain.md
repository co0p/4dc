# Domain Vocabulary

Shared business language for this project. Use this as a reference when requirements, code, tests, and conversations use the same concept. Update when meaning, rules, or relationships change; do not use it as a data dictionary or implementation catalog.

---

## Concepts

### [ConceptName]

**Definition:** [One sentence in domain terms].

**Meaningful state:**
- [Only state that changes how people understand or use the concept]

**Rules:**
- [Invariant or constraint in domain language]

**Related:**
- [Concepts or events that clarify the meaning]

---

## Domain Events

### [EventName]

**When:** [What triggers this event in domain terms]
**Information carried:** [Business information, not an implementation payload schema]
**Consumers:** [Who or what reacts in the domain]

---

## Rules and Constraints

[System-wide invariants, state transitions, and cross-concept rules described in domain language.]

---

## Writing Rules

- Define terms in language a product owner and implementer can both use.
- Record rules and relationships only when they affect decisions or outcomes.
- Do not list database columns, code paths, test cases, APIs, or historical introductions.
- If a term is still uncertain, record the ambiguity in the project decision record instead of inventing a definition.
