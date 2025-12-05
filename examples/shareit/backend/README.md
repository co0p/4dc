Shareit backend (minimal)

Quick start (local)

1. Install dependencies:

```bash
cd examples/shareit/backend
npm ci
```

2. Seed a demo DB (creates `data/shareit.json`):

```bash
cd examples/shareit/backend
npm run seed
```

The seed script writes `data/shareit.json` with a few demo items. The app's DB accessor accepts a JSON file as a data source for demo/CI-friendly runs.

3. Start server:

```bash
npm start
# then visit http://localhost:3000/api/catalog (after routes implemented)
```

Running tests:

```bash
npm test
```

Notes:
- This example prefers `better-sqlite3` for a simple sync API. If native sqlite build issues appear in CI, use the in-memory DAO fallback described in the increment notes.
