# Foundation Catalog

Each file contains one canonical influence statement. Skill templates select ideas with `{{FOUNDATION:<id>}}`; generated skills contain the rendered text and have no runtime dependency on this directory.

| ID | Influence | Core idea | Good fit |
|----|-----------|-----------|----------|
| `beck-team-agreements` | Kent Beck | Explicit team rules before code | Constitution |
| `beck-planning-game` | Kent Beck | Planning as a revisable conversation | Plan |
| `beck-small-steps` | Kent Beck | Small, reversible changes | Tidy, implementation |
| `beck-tidy-first` | Kent Beck | Prepare structure before behavior | Plan, tidy |
| `beck-red` | Kent Beck | A failing test specifies missing behavior | TDD Red |
| `beck-green` | Kent Beck | Minimum code to pass | TDD Green |
| `beck-refactor-pass` | Kent Beck | Improve design after green | Refactor |
| `beck-spike` | Kent Beck | Time-boxed learning experiment | Prototype |
| `cockburn-communication` | Alistair Cockburn | Conversation and close feedback | Increment, subtask plan |
| `cockburn-reflective-improvement` | Alistair Cockburn | Adapt the process from evidence | Promote |
| `cockburn-information-radiators` | Alistair Cockburn | Make work state visible | Implement |
| `deming-systems-thinking` | W. Edwards Deming | Optimize and verify the whole system | Constitution, promote |
| `farley-continuous-delivery` | David Farley | Keep change releasable through feedback | Constitution, promote |
| `fowler-evolutionary-architecture` | Martin Fowler | Guardrails that permit safe evolution | Constitution, ADR |
| `fowler-behavior-preserving-refactoring` | Martin Fowler | Improve structure without changing behavior | Tidy, refactor |
| `fowler-two-hats` | Martin Fowler | Separate behavior change from refactoring | Subtask plan, TDD Green, refactor |
| `freeman-pryce-tests-guide-design` | Steve Freeman and Nat Pryce | Tests provide design feedback | Plan, TDD Red |
| `humble-continuous-delivery` | Jez Humble | Releasability and deployment evidence | Constitution, promote |
| `jeffries-card-conversation-confirmation` | Ron Jeffries | Story, conversation, and proof | Increment, subtask plan, TDD Red |
| `nygard-architecture-decisions` | Michael Nygard | Preserve consequential decisions and trade-offs | ADR |
| `poppendieck-decide-late` | Mary and Tom Poppendieck | Delay reversible commitments | Increment, plan, ADR |
| `poppendieck-eliminate-waste` | Mary and Tom Poppendieck | Build only verified value | TDD Green, refactor, promote |
| `poppendieck-pull-small-batches` | Mary and Tom Poppendieck | Pull small batches from value | Plan, implement, subtask plan, tidy |
| `poppendieck-cheapest-experiment` | Mary and Tom Poppendieck | Resolve uncertainty with minimal waste | Prototype |
| `wake-invest` | Bill Wake | Independent, valuable, small, testable stories | Increment |

## Maintenance Rules

- One idea per file.
- Name files `<surname>-<idea>.md` using lowercase kebab-case.
- Render one Markdown list item containing the attribution, idea name, and operational meaning.
- Keep wording generic enough to reuse, but specific enough to guide a decision.
- Add the fragment to this catalog before referencing it from a skill.
- Do not place phase process, project facts, or tool instructions in a foundation fragment.
