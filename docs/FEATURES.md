# Feature Registry

Living index of implemented features. Updated during the `promote` phase each cycle.

> Source of truth: if a feature is not listed here, it is not considered implemented.

---

## Features

### FEAT-001 — Phase-aware agent orchestration

- **Description:** The agent detects the current 4dc phase from workspace state and loads the appropriate skill automatically.
- **Added in increment:** increment-2025-01-01
- **Resources:**
  - **Issues:** https://github.com/co0p/4dc/issues/1
  - **Tests:** `scripts/test_orchestration.sh#test_phase_detection`
  - **PRD:** `.agent/increment.md` (increment-2025-01-01)
- **PR:** https://github.com/co0p/4dc/pull/1
- **Notes:** Phase detection relies on presence of handoff contract files in the workspace root and `.agent/` directory.

---

### FEAT-002 — HTML review gate before permanent artifact writes

- **Description:** Every phase generates a human-readable HTML review file in `.agent/` and pauses for explicit user approval before writing any permanent artifact.
- **Added in increment:** increment-2025-01-15
- **Resources:**
  - **Issues:** https://github.com/co0p/4dc/issues/2
  - **Tests:** `scripts/test_review_gate.sh#test_html_generated_before_write`
  - **PRD:** `.agent/increment.md` (increment-2025-01-15)
- **PR:** https://github.com/co0p/4dc/pull/2
- **Notes:** The HTML file must exist and display `Status: Pending Approval` before the final Markdown write is permitted.
