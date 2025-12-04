# 4dc Prompt Authoring Guide (Flat Rules)

Use this file as a checklist when **authoring or editing prompts and templates** for 4dc.

Each rule describes how prompts should be structured so that runtime LLMs behave consistently.  
Prompts themselves MUST be **self-contained** (no dependency on this file at runtime).

---

1. **Each prompt must define a clear target scope.**  
   - Prompts MUST take an explicit argument that identifies the subject project or component (e.g. a directory path).
   - The instructions MUST say that this scope is the **subject** of the prompt.

2. **Prompts must constrain the LLM to that scope when reading files.**  
   - Instructions MUST tell the LLM to:
     - Read only files and directories **inside** the given path.
     - NOT rely on content from parent directories, sibling projects, or other repositories as primary context.
   - Surrounding repo content (frameworks, vendor dirs, examples) MAY be mentioned only as background, never as the subject.

3. **Prompts must be self-contained.**  
   - A generated `.prompt.md` MUST contain **all information the LLM needs at runtime**.
   - It MUST NOT:
     - Reference `RULES.md`.
     - Assume knowledge of other prompts.
     - Tell the LLM to “see another file” for essential rules.

4. **No cross-prompt references for behavior.**  
   - Prompts MUST NOT say:
     - “Follow the same rules as the design prompt.”
     - “See the constitution prompt for STOP behavior.”
   - If shared behavior is needed (e.g., STOP gates, scoping), it MUST be restated briefly in each prompt that needs it.

5. **Every main-phase prompt must include a human-in-the-loop flow.**  
   - Constitution, Increment, Design, Implement, and Improve prompts MUST:
     - Include at least two explicit STOP points.
     - STOP 1: After summarizing findings/context.
     - STOP 2: After summarizing proposed decisions/plan or outline.
   - The LLM MUST be instructed to:
     - Wait for user feedback at STOP 1.
     - Wait for explicit “yes” at STOP 2 before writing final artifacts.

6. **Prompts must tell the LLM to write artifacts only after explicit confirmation.**  
   - Instructions MUST say:
     - Do NOT write or overwrite the target artifact before STOP 2 approval.
     - Only generate the final document after an explicit “yes / go ahead / looks good”.
   - If the user says “no” or asks for changes:
     - The LLM must update the outline and re-ask for confirmation.

7. **Final artifacts must not mention prompts or assistants.**  
   - Prompts MUST instruct the LLM that:
     - Output artifacts (e.g. `CONSTITUTION.md`, `increment.md`, `design.md`) MUST NOT:
       - Mention prompts, LLMs, or assistants.
       - Contain meta-comments about how they were generated.
   - Artifacts must read as if written directly by the team.

8. **Final artifacts must follow a clear, fixed structure.**  
   - Each prompt MUST define a specific output structure (headings/sections or a schema).
   - Instructions MUST tell the LLM to:
     - Produce all required sections.
     - Not add unrelated top-level sections.
   - Optional extensions must be explicitly allowed if needed.

9. **Prompts may define acceptance criteria for the artifact.**  
   - It is recommended to include an “Acceptance” section in templates that says:
     - When an artifact is considered “good enough”.
     - What must be true for scope, alignment with goals, and clarity.
   - Acceptance criteria MUST be written in terms of the **artifact**, not the assistant.

10. **Prompts must describe persona and style briefly and concretely.**  
    - Each prompt SHOULD define:
      - A clear persona (e.g. Principal Engineer, Product Manager).
      - A concise style (direct, outcome-oriented, no fluff).
    - Persona/style text SHOULD be short and specific, not pages of prose.

11. **Prompts should separate goal (WHAT) from task/process (HOW).**  
    - Each main-phase template SHOULD:
      - Have a section that explains the **Goal** (what artifact we want and why).
      - Have a section that explains the **Task/Process** (how the LLM should behave step-by-step).
    - The Goal SHOULD be understandable without reading the Task; the Task SHOULD operationalize the Goal.

12. **Prompts should require focused, minimal clarifying questions.**  
    - Instructions SHOULD tell the LLM to:
      - Ask **targeted** clarifying questions only when critical information is missing or ambiguous.
      - Avoid long generic questionnaires.
    - Questions SHOULD mostly appear between STOP 1 and STOP 2.

13. **Prompts must keep the focus on the subject project, not the framework.**  
    - Instructions MUST emphasize:
      - The subject is the project/component under the given path.
      - The hosting repo, 4dc itself, or other tooling is NOT the subject unless explicitly part of the domain.
    - Any mention of frameworks or 4dc should be minimal and only when relevant to the project’s architecture.

14. **Prompts should be concise and opinionated, not verbose and vague.**  
    - Templates SHOULD:
      - Prefer concise, concrete rules over long, generic explanations.
      - State clear defaults and recommended practices.
    - Avoid:
      - “It depends” without guidance.
      - Huge walls of text for persona or style.

15. **Prompts must not ask the LLM to perform file I/O beyond what the host can handle.**  
    - If the environment supports file writes:
      - Prompts MUST still respect STOP 2 before writing.
    - If the environment is read-only:
      - Prompts must talk about generating content, not actually writing files.

---

When you (the human author) or a meta-assistant edits templates:

- Use this flat list as a **checklist**.
- Ensure each `.prompt.md`:
  - Has clear scoping instructions in its Inputs section (Rules 1–2).
  - Is self-contained and stop-gated (Rules 3–7).
  - Defines structure, persona, and goal/process cleanly (Rules 8–11).
  - Encourages focused clarifications and subject focus (Rules 12–14).

---

## 4dc Phases and Artifact Relationships

In addition to the flat rules above, 4dc templates follow a shared model of **four main phases** and how their artifacts relate to the codebase and to each other.

This section is for **prompt authors**; runtime prompts MUST NOT reference it.

### 1. Phases and Artifacts

4dc recognizes four main phases:

1. **Increment** – define the next small, valuable change in product terms.
2. **Design** – define the technical approach to realize that increment.
3. **Implement** – define a concrete, ordered plan of steps to apply the design.
4. **Improve** – analyze and improve the overall health of the codebase and delivery.

Each phase has:

- A distinct **persona**.
- A clear **scope** (what part of the project it looks at).
- A defined **artifact** it produces.

Typical artifacts:

- Increment: `increments/<slug>/increment.md`
- Design:   `increments/<slug>/design.md`
- Implement:`increments/<slug>/implement.md`
- Improve:  `improve.md` (at a project or subproject root)

### 2. Paths and Scope

Prompts MUST treat the `path` argument as their **subject root**:

- **Increment**:
  - `path` points at a **project or subproject root** (e.g. `.` or `examples/pomodoro`).
  - The prompt may read:
    - `CONSTITUTION.md`.
    - Readme / docs under that root.
    - Existing increment, design, implement, improve docs under that root.
  - It MUST NOT treat other repositories or parents as primary context.

- **Design** and **Implement**:
  - `path` points at an **increment folder**, for example:
    - `<project>/increments/<slug>`
  - They read:
    - `increment.md` in that folder.
    - `design.md` in that folder (for Implement).
    - `CONSTITUTION.md`, ADRs, and other docs under the project root.
    - Relevant code and tests under the project root.

- **Improve**:
  - `path` points at a **project or subproject root**.
  - It may examine any code, tests, and docs under that root.
  - It does not use slug folders for its own output, but may refer to them.

Every prompt must:

- Be explicit about where it expects `path` to point.
- Constrain itself to files and folders under that `path` as its subject.

### 3. Slug Folders for Increment, Design, and Implement

Increments, designs, and implementation plans share a common directory structure under a project root:

- `increments/<slug>/increment.md`
- `increments/<slug>/design.md`
- `increments/<slug>/implement.md`

Where `<slug>` is derived from an increment title by:

- Lowercasing.
- Replacing any sequence of non-alphanumeric characters with a single `-`.
- Collapsing repeated `-`s into one.
- Trimming leading/trailing `-`.

Rules for prompts:

- The **Increment** prompt:
  - Proposes a title and corresponding slug.
  - Proposes the folder name `<slug>`.
  - Produces `increment.md` in `increments/<slug>/`.
- The **Design** prompt:
  - Runs with `path` pointing at `increments/<slug>/`.
  - Produces `design.md` in the same folder.
- The **Implement** prompt:
  - Runs with `path` pointing at `increments/<slug>/`.
  - Produces `implement.md` in the same folder.

The **Improve** prompt:

- Does not create or depend on slug folders.
- Treats its `path` as the analysis root.

### 4. Responsibilities and Boundaries

Prompts for each phase must keep clear responsibilities:

- **Increment (WHAT, product)**:
  - Defines **what outcome** should be true when the increment is complete.
  - Operates in user/business language:
    - Context, goal, non-goals, tasks, risks, success criteria, observability.
  - Must NOT:
    - Specify code-level details (files, functions, classes, schemas).
    - Redesign architecture.
    - Describe concrete implementation steps.

- **Design (HOW, technical)**:
  - Turns the increment’s WHAT into a **technical design** for this codebase.
  - Grounded in:
    - `increment.md`.
    - `CONSTITUTION.md` and ADRs.
    - The existing code and architecture.
  - Describes:
    - Components, modules, services, and their responsibilities.
    - Interfaces, contracts, data models, and flows.
    - Testing strategy, CI/CD and rollout, observability, and risks.
  - Must NOT:
    - Redefine the product goal or tasks.
    - Contradict the increment’s non-goals without flagging it as a risk / follow-up.

- **Implement (plan, ordered steps)**:
  - Turns the combination of increment (WHAT) and design (HOW) into an **ordered plan of steps**.
  - Produces `implement.md` as:
    - Workstreams (e.g. backend, frontend, data, observability).
    - Concrete steps within each workstream.
    - Testing, CI/CD, rollout, and validation notes.
  - Intended usage:
    - Humans (and their coding tools) take steps from `implement.md` one at a time as small units of work.
  - Must NOT:
    - Silently redefine the design or increment.
    - Claim to be performing code changes itself.

- **Improve (system health)**:
  - Analyzes the overall health of the system under `path`.
  - Identifies:
    - Architectural and design issues.
    - Operational and observability gaps.
    - Developer experience and delivery friction.
  - Suggests:
    - Refactorings, cleanups, and systemic improvements.
    - Possible new increments to capture specific improvement work.
  - Must NOT:
    - Directly write or modify increment, design, or implement documents.
    - Redefine the project’s constitution.

### 5. Inputs vs. Behavioral Coupling

Artifacts may be used as **inputs** by other prompts:

- `increment.md` → input to Design and Implement.
- `design.md`    → input to Implement.
- Any of them    → input to Improve.

However:

- Prompts MUST NOT rely on other prompts’ **instructions** for their behavior.
- Each prompt must define:
  - Its own persona.
  - Its own inputs, process (including STOP gates), and output structure.
- Prompts must be executable in isolation given appropriate files under `path`.

### 6. Execution Loop (How These Are Used)

4dc is about authoring artifacts that guide real work. The intended loop is:

1. **Increment**:
   - At a project root, define a small, outcome-focused increment in `increments/<slug>/increment.md`.

2. **Design**:
   - In that increment folder, define a technical design in `design.md` grounded in the actual codebase.

3. **Implement**:
   - In that increment folder, define an implementation plan in `implement.md`:
     - Workstreams and ordered steps.
     - Testing, rollout, and validation hooks.

4. **Improve**:
   - At any time, analyze the codebase under a given root and suggest systemic improvements.

Execution of code changes is **deliberately human-driven**:

- Engineers (and their tools) use `implement.md` as a **backlog for the increment**.
- They take steps one by one, make code/test changes, run checks, and keep reality aligned with the plan.
- If the plan or design becomes out of date, new increments or updates can be created following the same phases.

All templates under `templates/` for Increment, Design, Implement, and Improve must be aligned with this model.