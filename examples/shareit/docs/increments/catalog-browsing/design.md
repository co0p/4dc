# Design: Catalog Browsing (SPA with Images + Loading)
**Date:** 2025-12-02  
**Status:** Initial Technical Design

## Design Summary
Deliver a minimal SPA, served by Express, that lists items with name, availability, and an optional thumbnail image. Show a loading indicator while fetching. Keep routes thin, isolate domain logic, and wrap the data source via an adapter to allow later DB/API swaps. This follows the constitution’s principles: Thin Routes, Isolated Domain; Three-Strikes Before Abstraction; Wrap and Pin; Small, Frequent Releases.

## Technical Decisions
- **SPA Served by Express (Static Middleware)**
  - Rationale: Simple mobile-first UX; avoids SSR complexity; keeps backend focused on API.
  - Trade-offs: No SEO; acceptable for app use case.
  - Alternatives: Server-rendered HTML templates; heavier SSR framework.

- **API Contract: GET /api/items (minimal JSON)**
  - Shape: `[{ id, name, availability: 'available'|'borrowed', thumbnailUrl? }]`
  - Rationale: Thin contract for list view; thumbnail optional.
  - Trade-offs: Enforces a specific availability vocabulary.

- **Domain Function: listItems(repo)**
  - Rationale: Normalize data/availability; easy to unit test; framework-agnostic.
  - Trade-offs: Small indirection vs doing mapping inline.

- **Repository Adapter: In-memory to start**
  - Rationale: Enables quick start; easy swap to DB/API later.
  - Trade-offs: Stub maintenance.

- **Loading Indicator + Error Handling in SPA**
  - Rationale: Clear user feedback; aligns with Small, Frequent Releases.
  - Trade-offs: Slight extra code for UI states.

## Architecture Overview
- **Components**
  - `domain/items`: `listItems(repo)` normalizes `{id, name, availability, thumbnailUrl?}`.
  - `adapters/itemsRepo`: in-memory repo with `findAll()`.
  - `routes/items`: Express GET `/api/items` serializes domain output.
  - `server/app`: Express app (static + API routes).
  - `public/index.html`: SPA shell (root div + loading element).
  - `public/app.js`: fetch `/api/items`, manage loading/error states, render list with images.
  - `public/styles.css`: minimal CSS for loading indicator and layout.
- **Data Flow**: client → GET /api/items → route → repo.findAll → domain.listItems → JSON → client renders.
- **Integration Points**: Express routing; future DB/API via repo adapter.
- **State**: Stateless request; SPA holds transient loading/error state.

## Implementation Constraints
- Show only name, availability, and optional thumbnail.
- No pagination, advanced filters, sorting, or detail page in this increment.
- Provide friendly empty and retry on errors.

## Guardrails & Scope
**In Scope**
- API route returning minimal item info + optional thumbnail.
- SPA rendering list with images, availability chips, and loading/error states.

**Out of Scope**
- Pagination/infinite scroll; advanced filters; sorting; item detail/borrow action.

**Planned Files Summary (initial)**
- `server/app.js` — new — Express app (static + API router mount)
- `server/routes/items.js` — new — Express `GET /api/items`
- `server/domain/items.js` — new — `listItems(repo)` normalization
- `server/adapters/itemsRepo.js` — new — in-memory repo `findAll()`
- `public/index.html` — new — SPA shell (root + loading container)
- `public/app.js` — new — SPA logic: fetch, loading indicator, render, retry
- `public/styles.css` — new — minimal styles for loading spinner and cards
- `tests/items.domain.test.js` — new — unit tests for `listItems`
- `tests/items.route.test.js` — new — route contract test (JSON fields)

**Drift Policy**
- STOP and propose update if additional files/modules are needed beyond this list.

**Suggested Branch**
- `feature/catalog-browsing`

## Open Questions
- Thumbnail fallback: use a static placeholder image if `thumbnailUrl` missing?
- Availability label color scheme (e.g., green available, gray borrowed)?
- Adopt Tailwind now vs. plain CSS for this increment?

## Save Location
`examples/shareit/docs/increments/catalog-browsing/design.md`
