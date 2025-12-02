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

module.exports = { listItems };