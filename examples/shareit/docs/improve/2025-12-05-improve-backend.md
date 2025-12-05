# Improve: Backend clarity, error visibility, and DB behaviour

## 1. Assessment

- **Constitution Alignment:**
  - Principle: Small, safe steps – ★★★★★ (The codebase is small, modular, and tests are in place to allow incremental, low-risk changes; see router factory in `routes/catalog.js` and tests in `backend/tests`.)
  - Principle: Refactoring as everyday work – ★★★★☆ (Tests accompany key flows and use DI; refactors are practical but a few areas lack explicit tests for edge cases, see `db.js` behavior.)
  - Principle: Pragmatic DRY & simplicity – ★★★★☆ (Implementation favors straightforward, readable code; `db.js` mixes fallback logic that could be clarified.)
  - Principle: Testing, CI & Observability – ★★★☆☆ (Focused unit and integration tests exist; logging is structured in `logger.js`, but startup and DB errors can be silent or produce generic responses.)

- **Design Alignment:**  
  The backend follows a lightweight, testable design: router factories accept injected dependencies, `db.js` uses a JSON fallback for CI, and logging is centralized. This matches the project's `lite` mode well. The architecture favors small increments and easy local runs.

- **Quality:**  
  Code is readable, uses lazy requires to avoid pulling native deps into test runs, and has clear unit/integration tests (`backend/tests`). Some functions (notably `db.findAllItems`) conflate multiple responsibilities (file fallback vs sqlite path detection) which can be clarified with small refactors and extra tests.

- **Risks:**  
  - Silent or generic failures: mounting failures are logged to console but may be missed; error responses return a generic `{"error":"internal error"}` hiding useful diagnostics in development (`server.js`).
  - Confusing DB fallback: when `dbPath` is provided but sqlite is unavailable, `db.findAllItems` throws a generic error that may be unclear to users or CI (`db.js`).
  - Health visibility: `/health` returns `ok`, but the app can silently run without mounted routes (degraded), making failures hard to detect in deployments.
  - Test gaps: no dedicated test for the default runtime path mounting with an empty `data/shareit.json` file, risking regressions.

## 2. Lessons

- **Worked Well:**
  - Dependency injection and router factory pattern in `routes/catalog.js` makes testing simple and isolates side effects.
  - Structured logging with `requestLogger` and `errorLogger` provides useful baseline observability; logs are JSON-friendly (`logger.js`).
  - Tests cover the primary success and failure paths for the catalog route and the JSON DB fallback (`backend/tests`).

- **To Improve:**
  - Make startup faults and DB configuration problems more visible (clear logs, health endpoint signaling).
  - Clarify `db.findAllItems` semantics and error messages for non-JSON DB paths when sqlite is absent.
  - Improve error responses in development to aid debugging (include `request_id`, message, and minimal stack info only when safe).
  - Add an integration test for the default empty JSON DB path to lock in expected behavior.

- **Emerging Patterns:**
  - Favor runtime-friendly fallbacks (JSON file) to keep CI and local runs reliable.
  - Use lazy requires to avoid pulling native dependencies into test runs.
  - Logging emphasizes traceability via `request_id`, but propagation to error responses is partial.

## 3. Improvements

#### Improvement 1: Expose helpful error info in development
- **Lens:** Architecture / Observability  
- **Priority:** M  
- **Effort:** 30–60 min  
- **Files:** `examples/shareit/backend/server.js`, `examples/shareit/backend/logger.js`  
- **Change:** Modify the global error handler in `server.js` so that in non-production environments (`NODE_ENV !== 'production'`) the JSON response includes `request_id` and `message` (and optionally a short `stack` string truncated to a reasonable length). Ensure `errorLogger` always receives and records `request_id` (populate it if missing). Keep the current single-line `{"error":"internal error"}` response for production to avoid leaking internals. Add a focused unit test that injects an error from a route and asserts the presence of `request_id` and `message` in the response when `NODE_ENV` is not `production`.  
- **Increment Hint (optional):** "Expose helpful error info in development and ensure request_id in error responses"

#### Improvement 2: Harden and document `db.findAllItems` semantics
- **Lens:** Naming / Documentation / Dependencies  
- **Priority:** H  
- **Effort:** 1–2 h  
- **Files:** `examples/shareit/backend/db.js`, `examples/shareit/backend/tests/db.test.js`, `examples/shareit/backend/server.js`  
- **Change:** Make `findAllItems(dbPath)` explicit and well-documented:
  - If `dbPath` is falsy: return `[]` (document this as "no DB configured" behavior).
  - If `dbPath` ends with `.json`: attempt to read and parse the file and return an array (current behavior).
  - If `dbPath` is provided and not a `.json` file: attempt to use sqlite; if sqlite (better-sqlite3) is not available or opening fails, throw an Error with a clear message that includes the `dbPath` and guidance (e.g., "SQLite not available; provide a JSON file or install better-sqlite3"). Avoid rethrowing raw errors from native libs; wrap them for clarity.
  - Add unit tests for: missing `dbPath` (returns []), `.json` parse success and empty array, and provided non-JSON path when sqlite is unavailable (assert clear error message). Update `server.js` mounting logic to validate `dbPath` before use and log a clear warning when default `data/shareit.json` is missing.  
- **Increment Hint (optional):** "Harden DB fallback behavior and add tests for missing or invalid DB paths"

#### Improvement 3: Make route mounting failures visible at startup and via health
- **Lens:** Delivery / Observability  
- **Priority:** M  
- **Effort:** 30–60 min  
- **Files:** `examples/shareit/backend/server.js`, `examples/shareit/backend/tests/integration.test.js`  
- **Change:** Replace the current silent `try/catch` around route mounting with a clearer startup policy:
  - If `createApp()` is called without `mountRoutes`, attempt to mount default routes; if mounting fails, either rethrow the error (so the process fails fast) or set an `app.locals.degraded = true` flag and write an explicit structured log entry explaining the degraded state and error. If opting for degraded startup, update `/health` to return HTTP 503 with a JSON body describing degraded status and an optional `reason` when `app.locals.degraded` is true. Add a test that simulates a mounting error and asserts that `/health` reflects degraded startup (or that the app fails fast if you choose that policy).  
- **Increment Hint (optional):** "Make route mounting failures visible at startup and in health checks"

#### Improvement 4: Add consistent `service` and `env` fields to structured logs
- **Lens:** Duplication / Documentation / Observability  
- **Priority:** L  
- **Effort:** 30 min  
- **Files:** `examples/shareit/backend/logger.js`, `examples/shareit/backend/server.js`  
- **Change:** Extend `jsonLog` and log entry structures to include `service: 'shareit-backend'` and `env: process.env.NODE_ENV || 'development'`. Centralize `jsonLog` as the single helper used by both `requestLogger` and `errorLogger`. Update tests that assert log output (where applicable) to tolerate and/or assert the presence of these fields. This creates consistent log records for downstream aggregation.  
- **Increment Hint (optional):** "Add consistent `service` and `env` fields to structured logs"

#### Improvement 5: Integration test for default empty JSON DB behavior
- **Lens:** Testing / Reliability  
- **Priority:** M  
- **Effort:** 30–60 min  
- **Files:** `examples/shareit/backend/tests/integration.test.js`, `examples/shareit/backend/server.js`, `examples/shareit/backend/routes/catalog.js`  
- **Change:** Add an integration test that runs the app using the default mounting behavior (no `mountRoutes` override), but configures a temporary `data/shareit.json` at the default path (or sets `SHAREIT_DB` env var to a temp file) containing an empty array. Assert that `GET /api/catalog` returns `200` with an empty array. This guards against regressions where default path behavior or file expectations change. Clean up the temporary file after the test.  
- **Increment Hint (optional):** "Integration test: default empty JSON DB returns empty catalog"

---

Next steps you can ask for:
- I can open a PR with the new file (if you want commits/PRs made).
- I can implement any of the proposed improvements as small increments (one at a time) and run tests locally.

If you want the file edited or renamed, tell me and I'll update it.
