# Testing

Document on testing practices for this project. Updated when new test patterns emerge or when constitution-level testing rules change.

---

## Test Types and Scope

<!--
Define the types of tests used in this project and what is required for each change.
Example categories:
- **Unit tests** — test single functions or methods in isolation
- **Integration tests** — test components working together
- **Acceptance tests** — test user-visible behavior against job stories
- **Performance tests** — test latency, throughput, or resource constraints
- **Contract tests** — test API or data contracts
-->

---

## What Requires Tests

<!--
State clearly what _must_ have test coverage:
- All behavior changes must have a failing test first (Red → Green → Refactor)
- Bug fixes must have a regression test before the fix
- Refactoring must keep existing tests green
- Structural tidying (`[tidy]` subtasks) must not change observable behavior
-->

---

## Test Location and Naming Conventions

<!--
Example:
- Test files live next to production code or in a `tests/` directory
- Test file naming: `<module>_test.<ext>` or `test_<module>.<ext>`
- Test function naming: `test_<unit>_<scenario>` or `describe('<unit>', () => { it('...')
- Helper functions: `setup_<fixture>()`, `assert_<condition>()`
-->

---

## Running Tests

<!--
How do developers run tests locally?
Example:
```bash
npm test              # All tests
npm test -- --watch  # Watch mode
npm test -- <pattern> # Filter by pattern
```

Or for other languages:
```bash
go test ./...
pytest
python -m unittest discover
```
-->

---

## Test Gate Before Promote

<!--
What must be green before code is promoted to permanent docs?
This is the _constitution-level gate_.
Example:
- All unit and integration tests must pass
- Code coverage must be ≥ 80% for changed files
- No flaky tests
- Performance tests must not regress
-->

---

## Continuous Integration / Testing Automation

<!--
Is there a CI/CD pipeline? What does it test?
Example:
- GitHub Actions / GitLab CI / Jenkins runs tests on every PR
- Coverage reports generated and must meet threshold
- Performance benchmarks tracked
- Linting and static analysis gates the merge
-->

---

## Known Test Gaps or Limitations

<!--
Document any areas where testing is incomplete or infeasible.
Example:
- Real-time features are not tested because [reason]
- External service integrations use mocks, not real endpoints
- UI tests are manual because [reason]
- Performance testing only covers [scenario], not [scenario]
-->

---

## Test Maintenance and Flakiness

<!--
How do you keep tests reliable?
Example:
- Flaky tests are tracked in [location] and fixed before merge
- Tests with external dependencies use retries with [strategy]
- Test data is seeded from [source], reset after each test
- Slow tests are isolated and run separately
-->
