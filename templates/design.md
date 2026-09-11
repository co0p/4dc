# Design — [Project Name]

User interface and interaction design decisions. Updated when visual principles, component patterns, or user workflows change.

> This is a durable design guide for consistency, not a component catalog or style guide. It captures the rationale behind key design choices, user flows, and interaction patterns so contributors and designers can maintain coherence and reason about changes safely.

---

## Design Principles

Core values that guide all design decisions:

- [e.g. "Mobile-first: design for small screens, enhance for larger ones"]
- [e.g. "Fail fast and visible: errors appear immediately, not after form submission"]
- [e.g. "Minimal viable interface: show only what the user needs to act now"]

---

## User Flows

### [Flow Name]

**User goal:** [What does the user want to accomplish?]

**Steps:**
1. [Entry point — how user gets here]
2. [Key interaction or decision point]
3. [Outcome — what changes or is confirmed]

**Design rationale:** [Why this sequence? What makes it effective for this user?]

---

## Key Interaction Patterns

Recurring design patterns that define the experience:

| Pattern | Description | When Used | Example |
|---------|-------------|-----------|---------|
| [Pattern Name] | [What the user does and sees] | [In which workflows] | [e.g. "Start button, timer counts down, stop button"] |
| [Pattern Name] | [description] | [when used] | [example] |

---

## Visual and Layout Decisions

### Color and Typography

- [e.g. "Single accent color (#0073E6) for all interactive elements"]
- [e.g. "System fonts only: no external font dependencies"]
- [e.g. "Text contrast minimum 4.5:1 for accessibility"]

### Layout Grid and Spacing

- [e.g. "8px base unit for all spacing and sizing"]
- [e.g. "Center-aligned for single-column layouts under 800px"]
- [e.g. "Left-aligned for multi-column layouts on desktop"]

### Responsive Behavior

- [e.g. "Stack on mobile, side-by-side on tablet and above"]
- [e.g. "Touch targets minimum 44x44px"]

---

## Component and State Behaviors

### [Component Name]

**Purpose:** [What does this component do?]

**States:**
- [e.g. "Default: ready for input"]
- [e.g. "Active: user is interacting"]
- [e.g. "Disabled: action not available, reason shown"]
- [e.g. "Error: invalid input, message clarifies what went wrong"]
- [e.g. "Loading: operation in progress, user can wait or cancel"]

**Behavior:** [How does it respond to interaction? When does it change state?]

---

## Accessibility

Standards and approach for inclusive design:

- [e.g. "WCAG 2.1 AA compliance target"]
- [e.g. "Keyboard navigation: all actions accessible via Tab and Enter"]
- [e.g. "No color-only information: use icons, text, or pattern"]
- [e.g. "Form labels explicit and associated: `<label for='input-id'>`"]
- [e.g. "Screen reader testing with VoiceOver / NVDA"]

---

## Voice and Microcopy

Language and tone that reflect the product:

- [e.g. "Positive: say what will happen, not what won't"]
- [e.g. "Clear errors: describe the problem and how to fix it"]
- [e.g. "Tone: professional and direct, not cute or jargon-heavy"]

**Example:**
- ✅ "Enter a date before tomorrow"
- ❌ "Date must not be in the future"

---

## Writing Rules

- Record design decisions that recur or affect multiple features — not every button color.
- Explain *why* a design pattern or constraint exists (user goal, accessibility, or performance reason).
- Link to or name specific components only when their behavior or state is non-obvious and affects implementation.
- Do not list every possible state variation, icon set, or font size; that belongs in the style guide or component library.
- If a design choice is still being tested or debated, record the open question in the project decision record instead of committing it here.

---

## Related

- [Architecture](./architecture.md) — System structure and runtime containers
- [Domain](./domain.md) — Shared business language and rules
- [Testing Decisions](./testing-decisions.md) — How to validate user interactions

---

## Update Policy

Update this document when:
- A new user flow or interaction pattern becomes standard across features
- A core visual principle or layout rule changes
- Accessibility or component state behavior changes in ways that affect multiple features

Do NOT update this document for:
- One-off button styling or minor icon changes
- Individual feature-specific screens or dialogs
- Implementation details (CSS class names, component props, etc.)
