const { expect } = require('chai');
const fs = require('fs');
const os = require('os');
const path = require('path');
const { findAllItems } = require('../db');

function mktemp() {
  return fs.mkdtempSync(path.join(os.tmpdir(), 'shareit-'));
}

describe('db.findAllItems', () => {
  it('returns mapped items from seeded JSON DB', () => {
    const tmp = mktemp();
    const dbPath = path.join(tmp, 'test.json');
    const seed = [
      { id: 1, name: 'Electric Drill', available: 1 },
      { id: 2, name: 'Ladder', available: 0 }
    ];
    fs.writeFileSync(dbPath, JSON.stringify(seed), 'utf8');

    const items = findAllItems(dbPath);
    expect(items).to.be.an('array').with.lengthOf(2);
    expect(items[0]).to.include.keys('id', 'name', 'available');
    expect(items[0].available).to.equal(true);
    expect(items[1].available).to.equal(false);
  });

  it('returns empty array when file missing or contains empty array', () => {
    const tmp = mktemp();
    const dbPath = path.join(tmp, 'empty.json');
    fs.writeFileSync(dbPath, JSON.stringify([]), 'utf8');

    const items = findAllItems(dbPath);
    expect(items).to.be.an('array').that.is.empty;
  });
});
