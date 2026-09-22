# Deployment

Guide to how this project releases and operates. Each section states a durable guardrail — the rule that survives specific technology choices. Fill in the project-specific details during constitution creation. Update this document when the release model or operational shape changes.

---

## Deployment Shape

State the release unit and where it runs. All later deployment decisions inherit from this shape and must be re-evaluated when it changes.

<!--
Answer before completing this section:
- What is the release unit? (application, service, library, artifact, script)
- Where does it run? (managed platform, self-hosted runtime, distributed to end users, embedded)
- Who is the operator responsible for a live release?
- What operational assumptions does this shape make? (network, uptime, scaling model, data locality)
-->

---

## Release Trigger and Cadence

State how releases start and how often they happen. Releases are predictable events driven by an explicit trigger, not by mood or availability.

<!--
Answer before completing this section:
- What starts a release? (automated event, manual action, hybrid)
- How often are releases expected? (per merge, batched on a schedule, on demand)
- Who is authorized to initiate a release?
- What evidence must exist before a release starts?
-->

---

## Environments and Verification

State each environment's purpose and the evidence required to promote from one to the next. Each environment either verifies something the previous did not, or it should not exist.

<!--
Answer before completing this section:
- Which environments exist between development and production?
- What does each environment verify that the previous did not?
- What evidence is required for promotion between environments?
- Who decides that evidence is sufficient?
-->

---

## Configuration and Secret Handling

State where configuration lives, where secrets come from, and how safe defaults are guaranteed. Configuration is external to source; secrets are never committed.

<!--
Answer before completing this section:
- Where does deploy-time configuration come from? (injected, baked in, hybrid)
- Where do secrets come from? (managed store, operator-injected, retrieved at startup)
- What are the safe defaults when configuration is missing or invalid?
- Who is authorized to change deployed configuration?
-->

---

## Rollback, Roll-Forward, and Recovery

State how a release is reversed or fixed forward. Every release states its recovery approach before it lands.

<!--
Answer before completing this section:
- Is rollback to a previous release supported, or is the strategy forward-fix only?
- Who decides to roll back or roll forward?
- What data implications does each approach have? (schema changes, migrations, external state)
- What is the maximum time from problem detection to mitigation?
-->

---

## Deployment Safety Constraints

State the release properties that must hold. These are the risks a release is not allowed to take.

<!--
Answer before completing this section:
- Is deployment atomic, rolling, or blue-green?
- Can old and new code run concurrently? For how long?
- Which operations during release are irreversible? (data migrations, external calls, credential rotations)
- Which release-time actions require an approval gate?
-->

---

## Health, Signals, and Operator Action

State how release health is observed and what action each signal triggers. Every alert has an owner and an expected response.

<!--
Answer before completing this section:
- Which signals confirm a release is healthy? (defined in docs/observability.md)
- How long is a release watched before it is considered stable?
- What operator action does each alert trigger?
- What automated action, if any, happens on signal breach?
-->

---

## Update Policy

State when this document changes. This document adapts to evidence from real deployments, not once and never again.

<!--
Answer before completing this section:
- What triggers an update to this document? (release model change, material incident, operator model change)
- Who is responsible for keeping this document current?
- What is removed when it becomes obsolete? (do not preserve stale procedures as historical reference)
-->
