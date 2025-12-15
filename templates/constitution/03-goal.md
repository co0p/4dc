## Goal

Generate a concise **CONSTITUTION.md** that:

- Describes **how this project wants to build and evolve software**.
- Scales to the project’s size and criticality via a **Constitution Mode**:
  - `lite` — small tools, demos, scripts; minimal ceremony.
  - `medium` — typical product/service; balanced engineering practices.
  - `heavy` — long-lived, critical, multi-team or regulated systems; strong process.
- Defines a clear **Implementation & Doc Layout**:
  - Where increments, designs, implement plans, improve docs, and ADRs live.
- Captures a small set of **Design & Delivery Principles** inspired by:
  - Fowler, Beck, Poppendiecks, DORA, Martin, Feathers, Wirfs-Brock, Evans, Metz, Thomas & Hunt, BDD.

The constitution will be used by:

- **Increment** prompts to:
  - Frame small, testable increments in line with values and layout.
- **Design** prompts to:
  - Respect architecture guardrails.
  - Choose simple, incremental designs.
- **Implement** prompts to:
  - Break work into small, testable tasks aligned with the design and constitution.
- **Improve** prompts to:
  - Evaluate system health against these principles.
  - Suggest refactors and ADRs consistent with the constitution.

The constitution itself must:

- Be **short enough** to read in minutes.
- Be **specific enough** to influence daily decisions.
- Be **stable** over time, but easy to extend by humans as the project matures.