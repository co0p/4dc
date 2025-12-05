const express = require('express');
const { requestLogger, errorLogger } = require('./logger');

function createApp(options = {}) {
  const app = express();
  app.use(express.json());
  app.use(requestLogger());

  // Simple route health check
  app.get('/health', (req, res) => res.json({ status: 'ok' }));

  // Mount API routes. Tests can inject routes via options.mountRoutes(app).
  if (typeof options.mountRoutes === 'function') {
    options.mountRoutes(app);
  } else {
    // Default runtime mounting using the local DB implementation and env var
    try {
      // lazy require to keep test harness lightweight
      const createCatalogRouter = require('./routes/catalog');
      const dbPath = process.env.SHAREIT_DB || require('path').join(__dirname, '..', 'data', 'shareit.json');
      app.use('/api', createCatalogRouter({ db: require('./db'), dbPath }));
    } catch (e) {
      // If mounting fails (missing modules), skip — tests will inject routes.
      // Log minimal info for debugging.
      console.error('Failed to mount default routes:', e && e.message);
    }
  }

  // Error handler: logs and returns 500 JSON
  // eslint-disable-next-line no-unused-vars
  app.use((err, req, res, next) => {
    try {
      errorLogger(err, req);
    } catch (e) {
      console.error('errorLogger failed', e);
    }
    res.status(500).json({ error: 'internal error' });
  });

  return app;
}

// If run directly, start server on port 3000
if (require.main === module) {
  const app = createApp();
  const port = process.env.PORT || 3000;
  app.listen(port, () => console.log(`Shareit backend listening on ${port}`));
}

module.exports = { createApp };
