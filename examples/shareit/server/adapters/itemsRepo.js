const items = [
  { id: 1, name: 'Cordless Drill', availability: 'available', thumbnailUrl: '' },
  { id: 2, name: 'Ladder', availability: 'borrowed' },
  { id: 3, name: 'Camping Stove', available: true }
];

async function findAll() {
  return items;
}

async function findById(id) {
  const num = Number(id);
  return items.find((it) => Number(it.id) === num) || null;
}

async function update(id, partial) {
  const num = Number(id);
  const idx = items.findIndex((it) => Number(it.id) === num);
  if (idx === -1) return null;
  const updated = Object.assign({}, items[idx], partial);
  // persist back into in-memory array
  items[idx] = updated;
  return updated;
}

module.exports = { findAll, findById, update };