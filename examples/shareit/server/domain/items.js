function normalizeAvailability(item) {
  if (typeof item.availability === 'string') {
    return item.availability === 'borrowed' ? 'borrowed' : 'available';
  }
  if (typeof item.available === 'boolean') {
    return item.available ? 'available' : 'borrowed';
  }
  return 'available';
}

function normalizeThumbnail(item) {
  if (item.thumbnailUrl && typeof item.thumbnailUrl === 'string') {
    return item.thumbnailUrl;
  }
  const label = encodeURIComponent((item.name || 'Item').slice(0, 1));
  const svg = `<?xml version="1.0" encoding="UTF-8"?>\n<svg xmlns='http://www.w3.org/2000/svg' width='120' height='120'>\n  <rect width='100%' height='100%' fill='#e5e7eb'/>\n  <text x='50%' y='55%' dominant-baseline='middle' text-anchor='middle' fill='#6b7280' font-size='48' font-family='ui-sans-serif, system-ui'>${label}</text>\n</svg>`;
  return `data:image/svg+xml;charset=UTF-8,${encodeURIComponent(svg)}`;
}

function listItems(items) {
  return (items || []).map((item) => ({
    id: String(item.id ?? ''),
    name: String(item.name ?? ''),
    availability: normalizeAvailability(item),
    thumbnailUrl: normalizeThumbnail(item)
  }));
}

async function borrowItem(repo, id) {
  const item = await repo.findById(id);
  if (!item) {
    const err = new Error('Item not found');
    err.status = 404;
    throw err;
  }
  const current = normalizeAvailability(item);
  if (current === 'borrowed') {
    const err = new Error('Item already borrowed');
    err.status = 409;
    throw err;
  }
  const updated = await repo.update(id, { availability: 'borrowed' });
  return listItems([updated])[0];
}

async function returnItem(repo, id) {
  const item = await repo.findById(id);
  if (!item) {
    const err = new Error('Item not found');
    err.status = 404;
    throw err;
  }
  const current = normalizeAvailability(item);
  if (current === 'available') {
    const err = new Error('Item not borrowed');
    err.status = 409;
    throw err;
  }
  const updated = await repo.update(id, { availability: 'available' });
  return listItems([updated])[0];
}

module.exports = { listItems, borrowItem, returnItem };