# Roadmap

Living product roadmap. Each section is one delivered or planned increment.

> A feature moves to **Done** only when its acceptance tests pass and are linked here.
> Source of truth: if a feature is not in Done with a passing test link, it is not considered shipped.

---

## Done

<!--
Each entry follows this pattern:

### [Feature name — short, user-visible]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Acceptance tests:** [path/to/test_file.ext#test_name](path/to/test_file.ext)
- **Use case:** [docs/usecases/use-case-slug.md](docs/usecases/use-case-slug.md) *(if promoted)*
- **Delivered:** [increment slug or YYYY-MM-DD]
-->

---

## Partial

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Acceptance tests:** pending — being written this cycle
- **Increment:** [increment slug]
-->

---

## Planned

<!--
### [Feature name]
- **Job story:** When [situation], I want to [action], so that [outcome].
- **Notes:** [optional — dependencies, open questions, or ordering rationale]
-->

---

## Rules

- Features move left to right: Planned → Partial → Done. Never skip Partial.
- A feature enters Partial when its increment is approved.
- A feature enters Done only when its acceptance test link resolves to a passing test.
- Do not add implementation detail here — link to the use case or ADR for that.
- If a planned feature is no longer needed, remove it and note the removal in `learnings.md` for that cycle.
