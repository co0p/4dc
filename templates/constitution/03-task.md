## Task

Create a CONSTITUTION that:

- Describes how the team balances speed, safety, quality, and sustainability.
- Makes the 6 pillars concrete enough to guide everyday decisions.
- Is structured so that later 4dc prompts can:
  - Refer to sections by name.
  - Extract constraints and trade-offs.
  - Understand how to prioritize between pillars when they are in tension.

You MUST:

- First infer as much context as possible from the repository itself (README, docs, code, config).
- Then ask targeted clarifying questions where your inferences are uncertain or ambiguous.
- Only then generate the final CONSTITUTION.

Before writing your final answer, follow these steps **internally** (do NOT include these steps in your output):

1. **Infer project context from the repository**
   - Scan README, docs, and key code/config files.
   - Populate internal notes:
     - `team_and_product_context`
     - `team_values_and_constraints`
     - `existing_practices_and_examples`
     - `inspirations_and_references`
     - `non_negotiables` (as far as you can infer them)

2. **Summarize and validate with the user**
   - Present concise summaries of these internal notes.
   - Highlight any assumptions or uncertainties.
   - Ask a small number of targeted questions to:
     - Confirm or correct your understanding.
     - Fill obvious gaps (especially around non-negotiables and priorities across the 6 pillars).
   - Incorporate the user’s answers back into your internal notes.

3. **Anchor each pillar in this environment**
   - For each of the 6 pillars, decide:
     - What it means specifically for this team.
     - How to tell when they are living up to it.
     - How to recognize when they are violating it.

4. **Define trade-off rules**
   - For common tensions (e.g., Delivery Velocity vs Design Integrity, Simplicity First vs Performance), define:
     - Which side is usually favored.
     - When and how to deliberately override the default.

5. **Make it operational for the 4dc loop**
   - Add practical guidance for:
     - **increment** (WHAT): how big increments should be, how to slice them.
     - **design** (HOW): what “good enough design up front” means.
     - **implement** (DO): how small steps should be, how to think about tests.
     - **improve** (GOOD/FAST): when and how to refactor, pay down debt, or optimize.

6. **Keep it editable and extensible**
   - Leave room for future amendments.
   - Highlight open questions the team should refine over time.

You MUST NOT show these steps or your intermediate reasoning in the final CONSTITUTION; only output the final document itself.