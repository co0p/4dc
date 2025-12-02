const items = [
  { id: 1, name: 'Cordless Drill', availability: 'available', thumbnailUrl: '' },
  { id: 2, name: 'Ladder', availability: 'borrowed' },
  { id: 3, name: 'Camping Stove', available: true }
];

async function findAll() {
  return items;
}

module.exports = { findAll };