# Implementation: Catalog Browsing (SPA with Images + Loading)

## 0. Drift Guardrails (Declare Up Front)
- Increment: Catalog Browsing — List Items with Availability
- Acceptance Criteria Reference: see `increment.md` in this folder
- DRIFT ALERT policy: If any step requires files/modules beyond the confirmed list, STOP and propose a scope update (or split follow-up increment).

## 1. Planned Files Summary (Confirmed)
- `examples/shareit/server/app.js` — new — Express app (static + API router mount)
- `examples/shareit/server/routes/items.js` — new — Express `GET /api/items`
- `examples/shareit/server/domain/items.js` — new — `listItems(repo)` normalization
- `examples/shareit/server/adapters/itemsRepo.js` — new — in-memory repo `findAll()`
- `examples/shareit/public/index.html` — new — SPA shell (root + loading container)
- `examples/shareit/public/app.js` — new — Vue SPA logic: fetch, loading indicator, render, retry
- `examples/shareit/public/styles.css` — new — minimal styles (Tailwind handles most)
- `examples/shareit/tests/items.domain.test.js` — new — unit tests for `listItems`
- `examples/shareit/tests/items.route.test.js` — new — route contract placeholder

Note: Tailwind used via CDN; Vue via ESM CDN for this increment.

## 2. Implementation Tasks & Subtasks (Living Plan)
- [ ] Setup
  - [ ] Create increment branch: run `git checkout -b feature/catalog-browsing` (verify branch exists)
  - [ ] Install server deps: `npm --prefix examples/shareit i express` (verify install)
  - [ ] Add `start:catalog` script in `examples/shareit/package.json` (verify script runs)

- [ ] API Route
  - [ ] Ensure `server/adapters/itemsRepo.js` returns sample items (verify export)
  - [ ] Implement `server/domain/items.js:listItems` normalization (verify by unit test)
  - [ ] Implement `server/routes/items.js` GET `/api/items` (verify curl output)

- [ ] SPA Shell & Rendering
  - [ ] Build `public/index.html` shell with Tailwind CDN and root containers (verify page loads)
  - [ ] Implement `public/app.js` Vue app: load, error, list rendering with images (verify in browser)
  - [ ] Include thumbnail fallback (SVG data URI) when missing (verify images render)

- [ ] Loading/Error/Empty States
  - [ ] Show spinner and text while loading (verify visible before data)
  - [ ] Friendly error with Retry action (verify by forcing 500)
  - [ ] Empty-state message when no items (verify by clearing repo)

- [ ] Verification (Manual-first for this increment)
  - [ ] Manual API check: `curl http://localhost:3000/api/items | jq` (verify JSON shape)
  - [ ] Manual UI check: load http://localhost:3000 (verify loading, list, images, availability chip)
  - [ ] Note: Automated tests deferred to a follow-up increment.

- [ ] Stabilization
  - [ ] Update `CONSTITUTION.md` references if needed (verify not required)
  - [ ] Document run steps in `README.md` (Shareit) (verify instructions added)
  - [ ] Minimal formatting/linting on touched files (verify clean)

## 3. Verification Methods
Manual-first for this increment:
  - Start server: `node examples/shareit/server/app.js`
  - API: `curl http://localhost:3000/api/items | jq` (see array with `id,name,availability,thumbnailUrl`)
  - UI: Open http://localhost:3000 in browser; observe loading → list with name, availability chip, thumbnail; retry on forced error; empty-state works.

## 4. Code Implementation (Essentials)
- API: `GET /api/items` returns normalized list from repo via domain function.
- Domain: `listItems(repo)` maps `available:boolean|availability:string` to `available|borrowed`; injects SVG data-URI thumbnail when missing.
- SPA: Vue app via ESM CDN; Tailwind via CDN; loading/error states and list UI.

## 5. Validation
- Criteria 1 (list shows name + availability): satisfied by SPA render bound to `/api/items` JSON.
- Criteria 2 (borrowed reflects): changing repo item and refreshing shows updated availability.
- Criteria 3 (empty-state): SPA shows helpful message when repo empty.
- Criteria 4 (error + retry): SPA shows error banner and retry triggers re-fetch.

## 6. Decision Log (Final Decisions Only)
- 2025-12-02 | Use Vue SPA + Tailwind CDN | Fast iteration; avoids build tooling | minimal
- 2025-12-02 | API shape `{id,name,availability,thumbnailUrl?}` | Keeps UI minimal yet useful | minimal
- 2025-12-02 | In-memory repo stub | Enables quick start; swap later | minimal
- 2025-12-02 | Thumbnail fallback (SVG data URI) | Ensures visual consistency without assets | minimal
- 2025-12-02 | Manual-first verification (tests deferred) | Matches choice B for test strategy | minimal

## 7. Open Questions
- Placeholder style OK or prefer static placeholder asset?
- Tailwind CDN acceptable vs local build pipeline?

## 8. Post-Implementation Stabilization & Merge
- Docs updated (Shareit README usage for server + SPA).
- `.gitignore` hygiene verified; reproducible start steps documented.
- Merge `feature/catalog-browsing` to default branch; tag if needed; delete branch.
