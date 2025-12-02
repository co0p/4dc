const { Router } = require('express');
const router = Router();

const { listItems } = require('../domain/items');
const itemsRepo = require('../adapters/itemsRepo');

router.get('/', async (req, res) => {
  try {
    const items = await itemsRepo.findAll();
    const normalized = listItems(items);
    res.json(normalized);
  } catch (err) {
    console.error('GET /api/items failed:', err);
    res.status(500).json({ error: 'Failed to load items' });
  }
});

module.exports = router;