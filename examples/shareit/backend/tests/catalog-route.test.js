const request = require('supertest');
const { expect } = require('chai');
const { createApp } = require('../server');

describe('GET /api/catalog', () => {
  it('returns items from injected db', async () => {
    const fakeDb = {
      findAllItems: (dbPath) => [
        { id: '1', name: 'Hammer', available: true },
        { id: '2', name: 'Saw', available: false }
      ]
    };

    const app = createApp({
      mountRoutes: (app) => {
        const createCatalogRouter = require('../routes/catalog');
        app.use('/api', createCatalogRouter({ db: fakeDb }));
      }
    });

    const res = await request(app).get('/api/catalog');
    expect(res.status).to.equal(200);
    expect(res.body).to.be.an('array').with.lengthOf(2);
    expect(res.body[0]).to.include({ id: '1', name: 'Hammer', available: true });
  });

  it('returns 500 when db throws', async () => {
    const badDb = {
      findAllItems: () => { throw new Error('boom'); }
    };

    const app = createApp({
      mountRoutes: (app) => {
        const createCatalogRouter = require('../routes/catalog');
        app.use('/api', createCatalogRouter({ db: badDb }));
      }
    });

    const res = await request(app).get('/api/catalog');
    expect(res.status).to.equal(500);
    expect(res.body).to.have.property('error');
  });
});
