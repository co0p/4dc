# 4dc Validation Contract

Use this document to review `AGENTS.md`, `templates/`, generated `skills/`, and installer or generation scripts. Ignore `examples/` unless explicitly included in the review.

For every rule, report `Pass`, `Fail`, or `Not proven` with file and line evidence. Run applicable repository checks rather than validating wording alone. Finish with one overall verdict: `Validates` only when no required rule fails; otherwise `Does not validate`.

## Anthropic

Reference guidance:
- [Prompt engineering overview](https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/overview)
- [Claude Code subagents](https://code.claude.com/docs/en/sub-agents)

Validate these generic rules for reliable skills and agents:

- Success criteria are explicit, observable, and testable rather than implied by prose.
- Instructions are direct, unambiguous, and ordered in the sequence they must be executed.
- Each skill has one focused responsibility and a description that clearly states when it should be selected.
- Inputs, outputs, constraints, stop conditions, and handoff expectations are explicit.
- Structured headings, lists, tables, and examples clarify complex instructions without duplicating or contradicting the normative rules.
- Examples demonstrate the required shape but cannot be mistaken for project facts or mandatory implementation choices.
- Long-lived context is separated from task-local context; agents load only the context needed for their current responsibility.
- Delegated agents or skills receive enough context to work independently because they cannot be assumed to share the parent agent's complete context.
- Tool access follows least privilege. Read-only investigation does not receive unnecessary write capabilities, and high-impact operations retain explicit permission boundaries.
- Agent and skill metadata is valid, concise, unique, and useful for discovery. Names and descriptions identify both capability and trigger conditions.
- Durable state and progress are written to artifacts rather than relying on chat memory.
- The workflow defines what to do when evidence contradicts the plan, context is missing, or a decision exceeds the current responsibility.
- Prompt changes are evaluated against defined success criteria; stylistic preference alone is not accepted as proof of improvement.

## OpenAI

Reference guidance:
- [Agents guide](https://developers.openai.com/api/docs/guides/agents)
- [Skills guide](https://developers.openai.com/api/docs/guides/tools-skills)

Validate these generic rules for reliable skills and agents:

- Every skill is a self-contained directory with one discoverable `SKILL.md` manifest.
- Skill frontmatter contains a valid name and a description explaining both what the skill does and when to use it.
- Reusable instructions stay in `SKILL.md`; supporting references, scripts, and assets are separated when they would otherwise overload the primary prompt.
- The orchestrator makes skill selection deterministic where required instead of relying entirely on model inference.
- Agent state, phase state, and handoffs are explicit and persisted across steps.
- The workflow distinguishes planning, tool execution, verification, and completion rather than treating a generated answer as completed work.
- Tools have clear purposes, bounded inputs, and predictable outputs. Instructions say which evidence must be read after execution.
- Write, destructive, network, deployment, merge, and other high-impact actions have appropriate approval and policy gates.
- Skills and external instructions are treated as privileged input. Untrusted content cannot silently override governing instructions or trigger sensitive actions.
- Instruction precedence is defined so conflicts between the user, orchestrator, project rules, and skills resolve consistently.
- Long-running work records progress and can resume from artifacts without reconstructing state from conversation history.
- Completion requires objective evidence, including command results or reproducible checks, rather than the model's assertion.
- Generated skills contain no unresolved template markers, stale paths, or source/generated drift.

## Google AI

Reference guidance:
- [Gemini API prompting strategies](https://ai.google.dev/gemini-api/docs/prompting-strategies)
- [Gemini CLI context files](https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/gemini-md.md)

Validate these generic rules for reliable skills and agents:

- Instructions state the task, relevant context, constraints, and expected output explicitly.
- Complex work is decomposed into small ordered steps with clear intermediate outcomes.
- Prompt structure is consistent enough for the model to identify requirements, context, examples, and output format reliably.
- Context is scoped hierarchically: global rules remain general, project rules apply to the repository, and component-specific details load only where relevant.
- Large instruction sets are modularized instead of duplicated across context files and skills.
- Shared language rules originate in `templates/language.md` and are rendered consistently into every generated skill through the shared execution contract.
- Foundation statements originate as one-idea fragments under `templates/foundations/` and skills select them with `{{FOUNDATION:<id>}}` markers.
- More specific instructions do not accidentally contradict higher-level safety, quality, or project constraints.
- Examples are representative, concise, and aligned with the requested output format.
- Tool calls are grounded in available tool definitions; required arguments, side effects, and expected results are clear.
- Multi-step tool use verifies each meaningful result before depending on it in a later step.
- The agent distinguishes missing information from implementation freedom and asks only when the missing answer materially affects the result.
- The workflow limits context growth by reading targeted files and preserving durable findings in artifacts.
- Output requirements are machine-checkable where practical, with explicit headings, states, or schemas.
- Validation includes both static consistency checks and behavioral evidence from realistic execution paths.

## 4dc Consistency

Validate that the complete prompt suite preserves these core 4dc ideas across `AGENTS.md`, templates, generated skills, scripts, and README documentation.

### Repository Checks

Run:

```bash
./scripts/generate-4dc.sh
git diff --check
bash -n scripts/generate-4dc.sh
bash -n scripts/install-4dc.sh
find skills -name SKILL.md | sort
rg -n '\{\{(SHARED|TEMPLATE|FOUNDATION):' skills
rg -n '\.4dc|promotion-report|<p>Phase: Implement \| Generated:' AGENTS.md README.md scripts templates skills
```

Expected evidence:
- Generation completes successfully.
- Generated skills match their source templates.
- Every generated skill contains the shared language rules from `templates/language.md` exactly once.
- Every foundation marker resolves to an existing cataloged fragment, and no skill selects the same foundation twice.
- Shell scripts parse successfully.
- Every installed skill is generated and every generated skill is installed.
- No unresolved template marker remains in generated skills.
- No stale `.4dc`, `promotion-report`, or rendered placeholder reference remains in active instructions.

### Flow And State

- Phase detection is complete and unambiguous from constitution through promote.
- The defined flow is `constitution → increment → optional prototype → plan → implement → subtask-plan → tidy or Red-Green-Refactor → promote`.
- ADR remains an on-demand decision mechanism rather than a mandatory sequential phase.
- Every phase has one owner, one primary responsibility, explicit expected inputs, a concrete output, hard gates, a process, a checklist, and a handoff.
- Phase artifacts are not written before their required conversational approval.
- Silence is never treated as approval.
- `.agent/` artifacts are transient handoff state; product knowledge is promoted to permanent project documentation.
- The todo list and `.agent/implementation.md` agree on current subtask order and state, with exactly one item in progress.
- A subtask is complete only after evidence and its commit hash are recorded.

### Constitution And Documentation

- `CONSTITUTION.md` contains only durable engineering guardrails, never project-specific tools, topology, domain examples, commands, concrete thresholds, or procedures.
- Constitution guardrails cover architecture, testing and quality, observability, security and privacy, reliability, performance and efficiency, documentation and ADR governance, and release and deployment.
- Project-specific testing, deployment, observability, architecture, domain, UI, and operational knowledge is routed to `docs/` or ADRs.
- Constitution creation asks one focused question at a time until every guardrail category is explicitly decided or deferred.
- Permanent documentation is treated as part of the product and is checked semantically, not merely for file existence.

### Increment And Plan

- The increment describes WHAT and WHY without implementation detail.
- Increment discovery explicitly establishes the user, trigger, outcome, failure boundary, and exclusions.
- The subtraction test removes independently releasable outcomes until the smallest useful, testable increment remains.
- Acceptance criteria are binary, observable, and mapped to the job story.
- The technical plan reads the relevant code and names exact files, symbols, boundaries, risks, and verification steps.
- The plan covers data shape, call flow, errors and edge cases, observability intent, and architecture delta when applicable.
- Every acceptance criterion has a required feature-level acceptance test with preconditions, action, exact observable outcome, and evidence; exceptions require explicit approval.
- Plan subtasks are only `[tidy]` or `[behavior]`. Plan-blocking research is resolved before approval, while local implementation uncertainty is handled during subtask mini-planning.

### Implementation Conversation

- `implement` runs once, scaffolds `.agent/implementation.md`, and populates the todo list without changing production code or tests.
- Every pending subtask passes through `subtask-plan` before implementation.
- Subtask planning investigates local unknowns and records findings, exact files and symbols, ordered steps, verification, observability, risks, and non-goals.
- The user explicitly approves one mini-plan for the current subtask.
- After mini-plan approval, the agent executes the complete subtask autonomously and records every transition in `.agent/implementation.md` and the todo list.
- Routine Tidy, Red, Green, Refactor, verification, and commit transitions do not introduce additional user gates.
- Autonomous execution stops only when findings change scope or acceptance criteria, require an unapproved structural decision, invalidate safety, or require destructive or external approval.
- A tidy subtask preserves observable behavior and keeps tests green.
- A behavior subtask executes one active test at a time through confirmed Red, minimal Green, and behavior-preserving Refactor.
- Behavior and structural changes use separate commits and the documented commit prefixes.
- Refactor activates the next test automatically or completes the subtask before returning to mini-planning for the next subtask.

### Promotion And Main Fit

- Final verification proves every acceptance criterion and required acceptance test before implementation is marked complete.
- Promote considers only implemented and verified outcomes, not plans or guesses.
- Each permanent-document promotion candidate is presented separately and approved before writing.
- The final tidy pass is behavior-preserving and keeps the release gate green.
- The latest target branch is integrated into the increment branch using an explicitly approved strategy before landing.
- The complete branch diff is reviewed for conflicts, duplicated work, stale assumptions, accidental scope, migration ordering, public-contract drift, and documentation consistency.
- Release and acceptance gates run again after integration with the latest target branch.
- The user explicitly approves the main-fit evidence before choosing squash-merge or pull request delivery.
- Cleanup happens only after documentation promotion and successful landing or PR creation.
