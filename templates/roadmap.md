# Roadmap

Product direction and sequencing guide. Each entry explains the user outcome, current confidence, and ordering rationale. Keep it concise and decision-oriented; detailed implementation status belongs in phase artifacts and code evidence.

> A feature moves to **Done** only when its user outcome is verified and the evidence is linked here.
> Source of truth: if a feature is not in Done with a passing test link, it is not considered shipped.

---

## Done

<!--
Each entry follows this pattern:

### [Feature name — short, user-visible]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Evidence:** [link to the smallest durable verification record](path/to/evidence)
- **Use case:** [docs/usecases/use-case-slug.md](docs/usecases/use-case-slug.md) *(if promoted)*
- **Delivered:** [increment slug or YYYY-MM-DD]
-->

---

## Partial

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Evidence:** pending — define the verification approach in the approved plan
- **Increment:** [increment slug]
-->

---

## Planned

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Why now / ordering:** [dependencies, user value, open questions, or sequencing rationale]
-->

---

## Rules

- Features move left to right: Planned → Partial → Done. Never skip Partial.
- A feature enters Partial when its increment is approved.
- A feature enters Done only when its evidence link resolves to a current verification record.
- Do not add implementation detail here — link to the use case or ADR for that.
- If a planned feature is no longer needed, remove it and note the removal in `learnings.md` for that cycle.
