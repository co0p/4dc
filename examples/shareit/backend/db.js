const fs = require('fs');

// Fallback JSON-backed DAO for lite demo and CI-friendly runs.
// If `dbPath` points to a .json file, read items from it. The JSON file
// should contain an array of records with fields: id, name, available.
// This keeps tests CI-friendly without native sqlite builds.

function findAllItems(dbPath) {
  if (!dbPath || dbPath.endsWith('.json')) {
    // Read JSON file if present, otherwise return empty array
    try {
      if (!fs.existsSync(dbPath)) return [];
      const raw = fs.readFileSync(dbPath, 'utf8');
      const data = JSON.parse(raw);
      if (!Array.isArray(data)) return [];
      return data.map(r => ({ id: String(r.id), name: r.name, available: Boolean(r.available) }));
    } catch (err) {
      throw err;
    }
  }

  // If better-sqlite3 is installed and dbPath is a sqlite file, attempt to use it.
  try {
    // lazy require so missing native dep doesn't break the module until used
    // eslint-disable-next-line global-require
    const Database = require('better-sqlite3');
    const db = new Database(dbPath, { readonly: true });
    try {
      const rows = db.prepare('SELECT id, name, available FROM items').all();
      return rows.map(r => ({ id: String(r.id), name: r.name, available: Boolean(r.available) }));
    } finally {
      try { db.close(); } catch (e) { /* ignore close errors */ }
    }
  } catch (e) {
    // If native sqlite isn't available, prefer a clear error to help debugging.
    throw new Error('SQLite not available and DB path is not a JSON file');
  }
}

module.exports = { findAllItems };
