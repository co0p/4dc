## Task

Create a CONSTITUTION that:

- Describes how the team balances speed, safety, quality, and sustainability for **this project**.
- Makes the 6 pillars concrete enough to guide everyday decisions on this project.
- Can be used by people and tools as a stable reference when making engineering trade-offs.

You MUST:

- First infer as much context as possible from the **target project root path** and its files, with a strict separation between:
  - Files in the target project root folder (for product description and high-level constraints), and
  - Files in subdirectories or outside that path (for practices, tooling, and values only).
- Then ask targeted clarifying questions where your inferences are uncertain or ambiguous.
- Then propose a brief plan for the constitution and present it to the user.
- Only after an explicit **yes** from the user, generate the full CONSTITUTION and write/update `CONSTITUTION.md` under the target project root.
- You MUST follow this “infer → clarify → plan → confirm → write” sequence exactly. Do not skip the confirmation step before writing the file.

Before writing your final answer, follow these steps **internally** (do NOT include these steps in your output):

1. **Infer project context from the TARGET root**
   - Focus on files that live **directly** in the target project root folder, especially:
     - `README.md` (authoritative for product description and users)
     - `CONSTITUTION.md` (if present)
     - Other root-level docs (e.g., `ARCHITECTURE.md`).
   - Respect the scoping rules from the Inputs section:
     - Use the target root `README.md` as the **only** source for the primary product description.
     - Do NOT incorporate product descriptions from subdirectories, examples, or prompts inside or outside the target root.
   - Populate internal notes:
     - `team_and_product_context`
     - Any obvious `non_negotiables` mentioned in root-level docs.

2. **Infer engineering practices from subdirectories and external tooling**
   - Look at subdirectories under the target project root such as `src/`, `docs/`, `.github/`, `examples/`, `templates/`, `tests/`, etc.
   - Optionally, look at surrounding framework/tooling repos only to infer general engineering values.
   - From these, enrich:
     - `team_values_and_constraints`
     - `existing_practices_and_examples`
     - `inspirations_and_references`
   - Treat any product-like descriptions here as *examples or supporting material*, not as the main product.
   - Do **not** bring in names of meta-frameworks or hosting repositories into the final constitution text.

3. **Summarize and validate with the user**
   - Present concise summaries of your internal notes, clearly labeled by source, for example:
     - Product/Team context (from the target root `README.md` / root-level docs)
     - Engineering practices (from CI, docs, code under the target root)
     - Values and constraints (from constitution-like docs, ADRs, prompts).
   - Ask the user to confirm or correct:
     - The **project description** (what this project is and is not), based only on the target root understanding.
     - The **main audience/users**.
   - Highlight any assumptions or uncertainties.
   - Ask a small number of targeted questions to:
     - Confirm or correct your understanding.
     - Fill obvious gaps (especially around non-negotiables and priorities across the 6 pillars).
   - Incorporate the user’s answers back into your internal notes.

4. **Plan the constitution and present a brief summary (STOP)**
   - Based on your refined understanding, construct a **short, human-readable plan** for the constitution, including:
     - Target project name and root path.
     - Main themes for:
       - Vision and Mission
       - Core Values
       - Architectural Principles (mapped to pillars)
       - Update Process
       - Pillar Coverage
       - Technical Decisions
       - Any notable non-negotiables.
   - Present this plan to the user along with a checklist of sections you intend to include.
   - Then ask the user for an explicit **yes/no** confirmation to proceed with generating and writing `CONSTITUTION.md` under the target project root. Example:

     > I plan to create or update `CONSTITUTION.md` under `<target-project-root>` with the sections and themes described above.  
     > Would you like me to generate and write this constitution file now? (yes/no)

   - This is a **STOP** point:
     - If the user answers **no**, adjust the plan based on their feedback and repeat this step.
     - If the user answers **yes**, proceed to the next steps.
     - You MUST NOT generate or write `CONSTITUTION.md` until the user answers **yes**.

5. **Anchor each pillar in this environment**
   - For each of the 6 pillars, decide:
     - What it means specifically for this project.
     - How to tell when they are living up to it.
     - How to recognize when they are violating it.

6. **Define trade-off rules**
   - For common tensions (e.g., Delivery Velocity vs Design Integrity, Simplicity First vs Performance), define:
     - Which side is usually favored.
     - When and how to deliberately override the default.

7. **Make it operational for day-to-day work**
   - Add practical guidance for:
     - How to choose and slice work (e.g., what "small, safe changes" look like).
     - How much design is expected up front vs evolved over time.
     - Expectations around tests and verification before merging.
     - When and how to refactor or pay down technical debt.

8. **Generate and write the CONSTITUTION after confirmation**
   - Once the user has answered **yes** in step 4:
     - Generate the full constitution document following the Output structure.
     - Write or update a file named `CONSTITUTION.md` in the **target project root** (the exact path passed as the project root argument):
       - If the file does not exist, create it.
       - If it exists, overwrite its contents with the new constitution.
     - Do NOT ask again whether the file should be created; the earlier **yes** confirmation already granted permission.
     - Inform the user that `CONSTITUTION.md` has been written, including the target root path used.

9. **Keep it editable and extensible**
   - Leave room for future amendments.
   - Highlight open questions the team should refine over time.

You MUST NOT:

- Show these steps or your intermediate reasoning in the final output.
- Offer to perform additional actions (for example, "If you'd like, I can also add...", "I can create a workflow for you", or similar).
- Include implementation tasks, checklists of what you (the assistant) could do next, or proposals for CI/CD setup in the constitution text itself.

Your final output MUST consist only of:

- The human-facing summaries and questions described above (when you are in the summarize/clarify/plan steps), and
- The final constitution document itself, formatted exactly according to the Output structure, **without** extra commentary, offers, or follow-up actions.