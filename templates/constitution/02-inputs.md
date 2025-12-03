## Inputs

You have access to:

- This repository’s contents as exposed by the tools that call you.
- Any answers the user provides during this interaction.
- A **project root path argument** that identifies the TARGET project within this repository hierarchy.

The calling tool passes you a **project root path argument** (for example `"."` or `"examples/pomodoro"`).

You MUST apply the following rules to this argument:

1. **Definition of project root folder**

   For the purpose of this prompt, the **project root folder** is defined as:

   - The directory at that project root path argument.
   - Files that live **directly** in that directory, such as:
     - `README.md`
     - `CONSTITUTION.md`
     - `LICENSE`
     - Other top-level markdown or configuration files.

2. **Scope of the constitution**

   - You are writing the constitution for the **project located at this project root path**, not for the repository as a whole.
   - Files under this root path are the only files that may define the project’s product description and domain.
   - Files outside this root path belong to other projects or tooling; they MUST NOT define the target project’s product description.

> Critical constraints:
> - All **subdirectories under this project root** (for example: `src/`, `docs/`, `.github/`, `examples/`, `templates/`, `tests/`, etc. inside this target root) are **not** considered part of the root folder for product description purposes.
> - You MUST NOT use files from subdirectories to override or define the primary product description.
> - Files **outside** this project root path (for example, a parent framework repo’s README, prompts, or other example projects) may be used only to understand general engineering values and workflows, NOT as the subject project’s description.
> - The generated constitution MUST NOT describe or name external framework repositories unless they are explicit runtime dependencies or architectural elements of the target project.

When inferring context, you MUST respect this source hierarchy and scoping:

1. **Primary project README in the TARGET root folder**
   - Look for `README.md` that lives **directly** in the project root path.
   - Use this file as the **authoritative source** of:
     - Product description
     - Target users
     - High-level goals.
   - Do NOT:
     - Combine this with other READMEs from subdirectories under the project root.
     - Mix example or sample project descriptions from outside the project root into the main product description.
   - If there is no `README.md` at the project root, ask the user to describe the project in their own words instead of inferring from nested docs.

2. **Other root-level documents in the TARGET project**
   - You may use other files directly in the project root folder (e.g., `CONSTITUTION.md`, `ARCHITECTURE.md`) to refine:
     - High-level architecture
     - Existing principles
     - Non-negotiable constraints.
   - These documents refine or extend the project root `README.md`; they do not describe separate products.

3. **Subdirectories under the TARGET project root (NOT for primary product description)**
   - Treat all files under subdirectories of the project root (e.g., `src/`, `docs/`, `.github/`, `examples/`, `templates/`, `tests/`, etc. inside the project) as:
     - Implementation details
     - Engineering practices
     - Internal tooling
     - Examples or sample content.
   - Use them ONLY to understand:
     - Engineering practices (tests, CI, workflows, scripts)
     - Code structure and architectural patterns
     - Tooling and prompts used in this project.
   - You MUST NOT:
     - Treat docs, READMEs, or prompts under these subdirectories as the main product description.
     - Merge their text into the root product description.

4. **Files outside the TARGET project root**
   - Treat files that live **outside** the project root path (for example, a parent framework repo’s `README.md`, `.github/prompts`, `templates/`, or other example projects) as:
     - Tooling, frameworks, and general engineering philosophy.
   - You may use them to:
     - Infer preferred workflows or general engineering values.
   - You MUST NOT:
     - Treat any of these as the target project’s product description.
     - Copy their text directly into the target project’s product narrative.
     - Include their names in the constitution unless they are explicit runtime dependencies or architectural elements of the target project.

From these sources, you MUST build and maintain the following **internal notes**. You may show them to the user for confirmation and refinement:

1. **Team / product context** (`team_and_product_context`)
   - Explain:
     - What this target project is about.
     - Who it serves.
     - The main problem space.
   - **Derived ONLY from:**
     - The target project’s root `README.md` (if present).
     - Root-level constitution/architecture docs (e.g., `CONSTITUTION.md`, `ARCHITECTURE.md`) that live directly in the target project root folder.
   - If these are missing or ambiguous, ask the user to clarify.

2. **Team values, preferences, and constraints** (`team_values_and_constraints`)
   - How the team appears to balance speed vs safety.
   - Any explicit or implicit quality bars.
   - **Derived from:**
     - The target project’s root `README.md`
     - Root-level `CONSTITUTION.md` (if present) in the target project
     - ADRs (wherever stored for the target project)
     - Tooling/workflow docs (e.g., CI config, `.github/` workflows, prompts) as additional evidence.

3. **Existing engineering practices / examples** (`existing_practices_and_examples`)
   - How the team currently reviews, tests, deploys, refactors, and documents.
   - **Derived from:**
     - CI/workflow files within or clearly associated with the target project
     - `docs/` or `examples/` directories under the target project
     - Folder structure and scripts under the target project root
     - Conventions visible in the target project’s codebase and tests.

4. **Inspirations / reference materials** (`inspirations_and_references`)
   - Any frameworks, books, or methodologies explicitly referenced in the target project.
   - Any implicit influences you can reasonably infer (e.g., hexagonal architecture, domain-driven design patterns).

5. **Known non-negotiables** (`non_negotiables`)
   - Compliance, security, regulatory or uptime constraints, if discoverable in the target project’s docs.
   - Otherwise, a list of questions you must ask the user to clarify.