const { randomUUID } = require('crypto');

function jsonLog(obj) {
  try {
    console.log(JSON.stringify(obj));
  } catch (err) {
    console.log('{"error":"log-serialization-failed"}');
  }
}

function requestLogger() {
  return (req, res, next) => {
    const start = Date.now();
    req.request_id = req.headers['x-request-id'] || randomUUID();

    res.on('finish', () => {
      const duration_ms = Date.now() - start;
      const entry = {
        timestamp: new Date().toISOString(),
        request_id: req.request_id,
        method: req.method,
        path: req.originalUrl || req.url,
        status: res.statusCode,
        duration_ms
      };
      if (res.locals && typeof res.locals.result_count === 'number') {
        entry.result_count = res.locals.result_count;
      }
      jsonLog(entry);
    });

    next();
  };
}

function errorLogger(err, req) {
  const entry = {
    timestamp: new Date().toISOString(),
    request_id: req && req.request_id,
    path: req && (req.originalUrl || req.url),
    error: err && err.message,
    stack: err && err.stack
  };
  jsonLog(entry);
}

module.exports = { requestLogger, errorLogger };
