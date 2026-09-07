# Testing

Guide to making reliable testing decisions for this project. Explain why the test strategy is shaped this way, how a developer should choose the cheapest test that gives sufficient confidence, and how to run the relevant checks. Update when the architecture, risk profile, or test workflow changes.

---

## Testing Approach and Rationale

Explain the risks the test strategy is designed to control and the boundaries where each kind of test provides confidence. Prefer principles and decision guidance over inventories. For example, explain why domain rules are tested without infrastructure, why persistence boundaries need integration checks, or why an acceptance test exercises a complete user job story.

---

## Choosing Test Depth

<!--
Describe how to decide whether a change needs a focused check, an integration check, an acceptance scenario, a performance measurement, or no new test. Tie the decision to user risk, architectural boundaries, determinism, and failure cost.
Do not maintain a catalog of individual tests or report a coverage percentage here.
-->

---

## Test Design Conventions

<!--
Describe conventions that make tests communicate behavior: naming, fixture ownership, isolation, determinism, test data, and how user-facing assertions should avoid implementation details. Keep examples small and illustrative rather than listing the suite.
-->

---

## Running the Checks

<!--
Document the actual commands for fast local feedback, the complete pre-merge gate, and any setup required for acceptance or environment-dependent checks. Explain when to use each command and how to interpret failures. Commands must be maintained as executable guidance, not illustrative placeholders.
-->

---

## Evidence Required Before Promotion

<!--
State the evidence required before a change is considered complete. Focus on behavior, risk, and reproducibility. Do not use line coverage or a list of passing tests as a substitute for explaining why the evidence is sufficient.
-->

---

## Automation and Feedback Loops

<!--
Explain where checks run (local, CI, release), what feedback each loop provides, and how failures are triaged. Record the rationale for any intentionally manual or environment-specific check.
-->

---

## Known Risks and Gaps

<!--
Document meaningful confidence gaps, why they exist, and what signal would justify changing the approach. Do not turn this section into a test inventory or coverage report.
-->

---

## Maintenance Guidance

<!--
Explain how tests are kept deterministic, how flakiness is handled, when fixtures or helpers should be changed, and how this guide itself is updated when the testing rationale changes.
-->
