const fs = require('fs');
const os = require('os');
const path = require('path');
const request = require('supertest');
const { expect } = require('chai');
const { createApp } = require('../server');

describe('integration: GET /api/catalog', () => {
  let tmpDir;
  let dbPath;

  beforeEach(() => {
    tmpDir = fs.mkdtempSync(path.join(os.tmpdir(), 'shareit-int-'));
    dbPath = path.join(tmpDir, 'seed.json');
    const seed = [
      { id: 1, name: 'Electric Drill', available: 1 },
      { id: 2, name: 'Ladder', available: 0 }
    ];
    fs.writeFileSync(dbPath, JSON.stringify(seed), 'utf8');
  });

  afterEach(() => {
    try { fs.rmSync(tmpDir, { recursive: true }); } catch (e) { /* ignore */ }
  });

  it('returns 200 and seeded items', async () => {
    const app = createApp({
      mountRoutes: (app) => {
        const createCatalogRouter = require('../routes/catalog');
        app.use('/api', createCatalogRouter({ db: require('../db'), dbPath }));
      }
    });

    const res = await request(app).get('/api/catalog');
    expect(res.status).to.equal(200);
    expect(res.body).to.be.an('array').with.lengthOf(2);
    expect(res.body[0]).to.include.keys('id', 'name', 'available');
    expect(res.body[0].available).to.equal(true);
  });
});
