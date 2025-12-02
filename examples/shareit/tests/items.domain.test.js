const assert = require('assert');

// lightweight fake repo for unit tests
function makeRepo(initialItems) {
  const items = initialItems.map((it) => ({ ...it }));
  return {
    async findById(id) {
      const num = Number(id);
      return items.find((it) => Number(it.id) === num) || null;
    },
    async update(id, partial) {
      const num = Number(id);
      const idx = items.findIndex((it) => Number(it.id) === num);
      if (idx === -1) return null;
      items[idx] = Object.assign({}, items[idx], partial);
      return items[idx];
    }
  };
}

async function run() {
  const { borrowItem, returnItem } = require('../server/domain/items');

  // Test: borrow available item
  const repo1 = makeRepo([{ id: 10, name: 'Test', availability: 'available' }]);
  const borrowed = await borrowItem(repo1, 10);
  assert.strictEqual(borrowed.availability, 'borrowed', 'availability should be borrowed after borrow');

  // Test: borrow already borrowed -> 409
  const repo2 = makeRepo([{ id: 11, name: 'Busy', availability: 'borrowed' }]);
  let saw409 = false;
  try {
    await borrowItem(repo2, 11);
  } catch (err) {
    saw409 = (err && err.status === 409);
  }
  assert.ok(saw409, 'should throw 409 when borrowing already-borrowed item');

  // Test: return borrowed item
  const repo3 = makeRepo([{ id: 12, name: 'R', availability: 'borrowed' }]);
  const ret = await returnItem(repo3, 12);
  assert.strictEqual(ret.availability, 'available', 'availability should be available after return');

  // Test: return when available -> 409
  const repo4 = makeRepo([{ id: 13, name: 'R2', availability: 'available' }]);
  let saw409return = false;
  try {
    await returnItem(repo4, 13);
  } catch (err) {
    saw409return = (err && err.status === 409);
  }
  assert.ok(saw409return, 'should throw 409 when returning an available item');

  console.log('items.domain tests passed');
}

if (require.main === module) {
  run().catch((err) => {
    console.error('Test run failed:', err);
    process.exit(1);
  });
}

module.exports = { run };
