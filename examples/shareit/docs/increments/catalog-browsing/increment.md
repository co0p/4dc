# Catalog Browsing: List Items with Availability

## Job Story
**When** I open the catalog
**I want to** see a simple list of items and their availability
**So I can** quickly decide what to borrow

## Assumption Being Tested
Users primarily need a simple list to decide quickly.

## Acceptance Criteria
- **Given** items exist with various statuses  
  **When** I open the catalog  
  **Then** I see each item’s name and availability (available/borrowed).

- **Given** an item becomes borrowed  
  **When** I refresh the catalog  
  **Then** its availability shows as borrowed/unavailable.

- **Given** no items exist  
  **When** I open the catalog  
  **Then** I see an empty-state message explaining how to add items.

- **Given** an API/load error occurs  
  **When** I open the catalog  
  **Then** I see a friendly error and can retry.

## Success Signal
The catalog page displays at least item name and availability for all items without errors.

## Out of Scope
- Pagination/infinite scroll
- Advanced filters (tags, owner)
- Sorting

---

## Implementation Guardrails & Handoff
- Branching: work on `feature/catalog-browsing`; no direct commits to default.
- Planned Files Summary: will be proposed and confirmed in the Implement phase (STOP gate there).
- Drift Policy: if implementation needs new files/scope beyond the planned list, raise a DRIFT ALERT and request confirmation before proceeding.
- Verification Expectation: use a test-first loop where feasible; otherwise provide explicit manual checks mapped to acceptance criteria.
