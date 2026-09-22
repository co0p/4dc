# Observability

Guide to understanding the system in operation. Keep concrete signals, ownership, interpretation, and response guidance here rather than in the constitution.

## Operational Outcomes

Describe the user-visible and operator-visible outcomes that must be distinguishable in production, including success, expected rejection, degradation, and failure.

## Signals

Document the logs, metrics, traces, audit records, and health checks used by this project. For each signal, state its purpose, important fields or dimensions, privacy and cardinality constraints, and how to interpret it.

## Alerts And Response

Document actionable alert conditions, ownership, first response, verification, and escalation or recovery guidance. Avoid alerts without an expected operator action.

## Release Verification

Describe the signals used to confirm a release is healthy and the conditions that trigger rollback, roll-forward, or further investigation.

## Known Blind Spots

Record material observability gaps, their risk, and the evidence that would justify improving them.

## Maintenance

Update this guide when behavior, architecture, data sensitivity, operational ownership, or failure modes change. Remove stale signals rather than preserving an inventory of obsolete telemetry.
