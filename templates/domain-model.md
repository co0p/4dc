# Domain Model

Living record of this project's domain concepts. Updated during the `promote` phase each cycle.

> Source of truth for shared language. Entries are typed — each concept is classified as an **Aggregate**, **Event**, **Value Object**, or **Term**. If a concept is not here, its meaning is ambiguous — add it.

---

## Format

### Aggregates

Aggregates are consistency boundaries — a cluster of entities and value objects with a single root that enforces all invariants.

**AggregateName** — One-sentence description of what this aggregate governs.
- *Root entity:* The entity through which all access must go.
- *Invariants:* Rules enforced within this boundary.
- *Raises:* Events this aggregate produces.
- *Added:* `[increment slug]`

### Events

Domain events record something that happened in the domain — past tense, immutable facts.

**EventName** — One-sentence description of what occurred and why it matters.
- *Payload:* Key data carried by this event.
- *Trigger:* What causes this event to be raised.
- *Consumers:* Who reacts to this event.
- *Added:* `[increment slug]`

### Value Objects

Value objects represent a domain concept defined entirely by its attributes — no identity, immutable.

**ValueObjectName** — One-sentence definition of what it represents.
- *Attributes:* Fields that define equality.
- *Invariants:* Rules that must always hold.
- *Related:* Other concepts this depends on.
- *Added:* `[increment slug]`

### Terms

Ubiquitous language terms that appear in code, tests, and conversations but are not events or value objects.

**Term** — One-sentence definition in domain context.
- *Example:* How it appears in code, CLI output, or tests.
- *Related:* Other terms this depends on or contrasts with.
- *Added:* `[increment slug]`

---

## Aggregates

### [ExampleAggregate]

**[ExampleAggregate]** — [One-sentence description of what consistency boundary this aggregate owns].
- *Root entity:* `[RootEntityName]`
- *Invariants:* [e.g. total must never exceed limit, state transitions must follow X]
- *Raises:* [ExampleEventName]
- *Added:* `[increment slug]`

---

## Events

### [ExampleEventName]

**[ExampleEventName]** — [One-sentence description of the fact that occurred].
- *Payload:* `[field1]`, `[field2]`
- *Trigger:* [What user action or system condition raises this]
- *Consumers:* [Who or what reacts]
- *Added:* `[increment slug]`

---

## Value Objects

### [ExampleValueObject]

**[ExampleValueObject]** — [One-sentence definition].
- *Attributes:* `[attribute1]`, `[attribute2]`
- *Invariants:* [e.g. must be non-negative, must match pattern X]
- *Related:* [ExampleEventName]
- *Added:* `[increment slug]`

---

## Terms

### [ExampleTerm]

**[ExampleTerm]** — [One-sentence definition grounded in this project's domain, not a generic definition].
- *Example:* `[concrete example from code, CLI, or test]`
- *Related:* [ExampleValueObject]
- *Added:* `[increment slug]`

---

## Rules

- Entries must be placed in the correct section: Aggregates, Events, Value Objects, or Terms.
- Definitions are written from the perspective of this project's domain, not general software engineering.
- When a term's meaning changes, update the definition in place — do not add a new entry.
- If two entries turn out to mean the same thing, pick one and add a redirect: `**Alias** — See [Canonical Name].`
- Remove entries that are no longer part of the codebase.
