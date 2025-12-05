Implement: List Catalog API

Context
- Goal: Add `GET /api/catalog` that returns a JSON array of item records (`id`, `name`, `available`) so clients can fetch the canonical catalog. See `increment.md` for acceptance criteria.
- Design: Minimal Express backend under `examples/shareit/backend` with a thin SQLite wrapper exposing `findAllItems()`, structured JSON logging, unit tests, and one integration test. See `design.md`.
- Constitution: `examples/shareit/CONSTITUTION.md` is `lite` mode — prefer small, pragmatic steps, unit tests plus one integration test, and simple stdout JSON logging.
- Links: `increment.md`, `design.md`, `CONSTITUTION.md`

## 1. Workstreams
- Workstream A – Backend scaffold & routing (Express bootstrap, route mounting)
- Workstream B – Data layer & seeding (SQLite wrapper `findAllItems()` + seed script)
- Workstream C – Tests & fixtures (unit tests + single integration test)
- Workstream D – Observability, docs & QA (structured logs, QA checklist)

## 2. Steps

### Step 1: Initialize backend package and scripts
- Workstream: A
- Based on Design: "CI/CD and Rollout" / "Testing and Safety Net"
- Files: `examples/shareit/backend/package.json`, `.gitignore`
- Actions:
  - Create `package.json` with minimal metadata and scripts:
    - `start`: `node server.js`
    - `test`: runs the test runner (e.g. `npm test` -> `mocha` or `jest`).
  - Add dev/runtime dependencies: `express`, `supertest`, a sqlite library (see note), and a test runner (`mocha` + `chai` or `jest`).
  - Add `.gitignore` entries for `data/shareit.db`, `node_modules/` and test DB files.
- Tests / Verification:
  - Run `npm test` (sanity) and ensure the command runs (no failing tests expected yet).

### Step 2: Add DB wrapper `db.js` with `findAllItems()` and unit tests for mapping
- Workstream: B, C
- Based on Design: "Data layer (SQLite)" and "Mapping rules"
- Files: `examples/shareit/backend/db.js`, `examples/shareit/backend/tests/db.test.js`
- Actions:
  - Implement `findAllItems(dbPath)` which opens the SQLite DB at `dbPath` and returns an array of JS objects: `{ id: String(id), name, available: Boolean(available) }`.
  - Keep API of `db.js` small and test-friendly; accept `dbPath` so tests can pass a temp DB.
  - In the unit test, create a temporary SQLite DB (seed two rows), call `findAllItems()` and assert returned items and correct boolean mapping.
- Tests / Verification:
  - `db.test.js` asserts mapping and empty DB handling. Run `npm test` and confirm tests pass.

### Step 3: Add Express bootstrap `server.js` (export app) and logger middleware
- Workstream: A, D
- Based on Design: "Express app: route registration, request/response plumbing, error handling, and structured request logging"
- Files: `examples/shareit/backend/server.js`, `examples/shareit/backend/logger.js` (optional)
- Actions:
  - Implement Express app that:
    - Adds JSON body parser.
    - Generates a simple `request_id` per request and attaches it to `req`.
    - Adds a request-completion logger middleware that writes a JSON line with `timestamp`, `request_id`, `method`, `path`, `status`, `duration_ms`.
    - Adds a top-level error handler that logs errors with `request_id` and returns `500` JSON `{ "error": "internal error" }`.
  - Export the app (e.g. `module.exports = app`) so tests can import it without starting a network listener.
- Tests / Verification:
  - Manually require the exported app in tests. Start via `node server.js` in dev to smoke-check startup.

### Step 4: Implement `routes/catalog.js` and wire into `server.js`
- Workstream: A, B
- Based on Design: "Catalog route handler" and "Short request flow"
- Files: `examples/shareit/backend/routes/catalog.js`, update `server.js` to `app.use('/api', catalogRouter)`
- Actions:
  - Implement router registering `GET /api/catalog`.
  - Handler should call `db.findAllItems(dbPath)` (injected via env var or constructor), respond `res.json(items)` with 200 on success.
  - On error, `next(err)` so the top-level error handler returns 500 and logs.
  - Optionally set `res.locals.result_count` to help the logger include `result_count`.
- Tests / Verification:
  - Unit test `tests/catalog-route.test.js` that stubs `db.findAllItems()` and asserts handler returns 200 + expected JSON when DB returns rows, and 500 + `{ error }` when DB throws.

### Step 5: Add integration test using `supertest` and a seeded ephemeral DB
- Workstream: C, B, A
- Based on Design: "Integration test (single) ... supertest" and "Fixtures and test data"
- Files: `examples/shareit/backend/tests/integration.test.js`, `examples/shareit/backend/data/seed.sql` or inline seed in test
- Actions:
  - Integration test creates a temp SQLite DB file, seeds it with deterministic rows, imports the Express app configured to use that DB, and performs `GET /api/catalog` with `supertest`.
  - Assert response status `200` and response body is a non-empty array; assert each item contains `id`, `name`, `available` (boolean).
  - Clean up temp DB file after test.
- Tests / Verification:
  - Run `npm test` and ensure the integration test passes in CI-local environment.

### Step 6: Add lightweight structured logging tests and QA checklist
- Workstream: D
- Based on Design: "Observability and Operations" / "Logging"
- Files: `examples/shareit/backend/tests/logging.test.js`, `examples/shareit/docs/increments/list-catalog-endpoint/implement.md`
- Actions:
  - Ensure request and error logs are emitted as JSON lines to stdout. Add a test that simulates an error path and captures stdout to assert presence of `request_id` and `error.message`.
  - Add a short QA checklist at the bottom of this file (or update `increment.md`) describing manual validation steps (seed DB, run server, curl endpoint, check logs).
- Tests / Verification:
  - Run `npm test` and verify logging test passes; manually run `npm start` and curl endpoint to confirm logs appear.

### Step 7: Seed script, README snippet and small polish
- Workstream: D
- Based on Design: "DB schema" and "Documentation"
- Files: `examples/shareit/backend/scripts/seed.js` or `data/seed.sql`, `examples/shareit/backend/README.md`
- Actions:
  - Add a seed script or SQL file to create the `items` table and insert demo rows.
  - Add a short README describing how to run the server and tests locally.
  - Note the sqlite native dependency and fallback plan (in-memory fallback) in README.
- Tests / Verification:
  - Manual: run seed script, `npm start`, then `curl http://localhost:3000/api/catalog` and assert the JSON array.

## 3. Rollout & Validation Notes
- Suggested PR grouping:
  - PR 1: Steps 1–3 (package.json, `db.js`, `server.js`, router skeleton, basic unit tests). Small and reviewable.
  - PR 2: Steps 4–5 (route implementation, integration test, seed helpers).
  - PR 3: Steps 6–7 (logging tests, seed script, README and QA checklist).
- CI checks:
  - `npm ci && npm test` in `examples/shareit/backend` should pass. Keep tests fast and deterministic.
- Validation checkpoints:
  - After PR 1: confirm `npm test` runs and new files follow project layout.
  - After PR 2: run integration test and confirm `GET /api/catalog` returns 200 + expected fields.
  - After PR 3: manual QA checklist (seed DB, start server, curl endpoint, inspect logs).

## Notes and Decisions
- SQLite library: prefer `better-sqlite3` for its simple, synchronous API and deterministic tests; if CI shows native build issues, fall back to a minimal in-memory JS DAO as a follow-up.
- Test strategy: follow `lite` constitution — unit tests + one integration test only. Keep tests isolated, create/tear down ephemeral DB files during tests.
- Files to create under `examples/shareit/backend`:
  - `package.json`, `server.js`, `logger.js` (optional), `routes/catalog.js`, `db.js`, `scripts/seed.js` (or `data/seed.sql`), `tests/*.test.js`, `README.md`.

## QA Checklist (quick)
- Run seed and start server:
  - `node scripts/seed.js` (or run SQL seed) to create `data/shareit.db`.
  - `npm start` (in `examples/shareit/backend`) and verify server binds.
- Request the endpoint:
  - `curl -sS http://localhost:3000/api/catalog | jq '.'` and confirm an array of items with `id`, `name`, `available` (boolean).
- Inspect logs:
  - Confirm request log JSON line appears with `request_id`, `method`, `path`, `status`.

---
Implementer: follow these steps in order, creating small PRs. If any native sqlite issues appear in CI, run the in-memory DAO fallback and open a short follow-up increment to replace with the SQLite implementation.
