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

## In Progress

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Evidence:** pending — define the verification approach before work starts
-->

---

## Planned

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Why now / ordering:** [dependencies, user value, open questions, or sequencing rationale]
-->

---

## How This List Works

- Features move left to right: Planned → In Progress → Done. Never skip In Progress.
- A feature enters In Progress when work begins.
- A feature enters Done only when its user outcome is verified and the evidence link is present.
- Do not add implementation detail here — link to the use case or ADR for that.
- If a planned feature is no longer needed, remove it and record the reason in a code comment, commit message, or ADR.
