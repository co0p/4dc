const fs = require('fs');
const path = require('path');

function ensureDir(dir) {
  if (!fs.existsSync(dir)) fs.mkdirSync(dir, { recursive: true });
}

function seed() {
  const dataDir = path.join(__dirname, '..', 'data');
  ensureDir(dataDir);
  const file = path.join(dataDir, 'shareit.json');
  const items = [
    { id: 1, name: 'Electric Drill', available: 1 },
    { id: 2, name: 'Ladder', available: 0 },
    { id: 3, name: 'Folding Chair', available: 1 }
  ];
  fs.writeFileSync(file, JSON.stringify(items, null, 2), 'utf8');
  console.log('Seeded', file);
}

if (require.main === module) seed();

module.exports = { seed };
