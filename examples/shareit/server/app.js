const path = require('path');
const express = require('express');

const app = express();
app.use(express.json());

// Static SPA assets
app.use(express.static(path.join(__dirname, '..', 'public')));

// API routes
const itemsRouter = require('./routes/items');
app.use('/api/items', itemsRouter);

const port = process.env.PORT || 3000;

if (require.main === module) {
  app.listen(port, () => {
    console.log(`Shareit server listening on http://localhost:${port}`);
  });
}

module.exports = app;