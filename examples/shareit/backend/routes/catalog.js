const express = require('express');

// Factory: createCatalogRouter({ db, dbPath }) -> Router
// - db: optional module implementing findAllItems(dbPath)
// - dbPath: optional path passed to db.findAllItems
module.exports = function createCatalogRouter({ db = require('../db'), dbPath } = {}) {
  const router = express.Router();

  router.get('/catalog', async (req, res, next) => {
    try {
      const items = await Promise.resolve(db.findAllItems(dbPath));
      if (Array.isArray(items)) res.locals.result_count = items.length;
      return res.json(items);
    } catch (err) {
      return next(err);
    }
  });

  return router;
};
