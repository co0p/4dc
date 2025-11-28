# Design: List Catalog

**Date:** 2025-11-28  
**Status:** Initial Technical Design

## Design Summary
We will build a mobile-first catalog list screen for ShareIt using Svelte, loading sample data from a local JSON file, managing state within the component, and using URLs to placeholder images for each item. Svelte is chosen for its simplicity, minimal boilerplate, and suitability for junior developers. This approach aligns with the constitution's principles of rapid iteration, simplicity, and best practice design, and meets the increment's acceptance criteria for usability and responsiveness.

## Technical Decisions
- **Framework/Library:** Svelte
  - **Rationale:** Svelte offers a simple, readable syntax and compiles to efficient JavaScript, making it easy for junior developers to contribute. Minimal boilerplate and direct reactivity simplify implementation and maintenance.
  - **Trade-offs:** Smaller ecosystem and fewer UI libraries than Vue/React, but ideal for small, focused increments. May require more custom code for advanced features.
  - **Alternatives Considered:** React (too complex for this use case), Vue (more boilerplate, larger ecosystem), vanilla JS (less structure, harder to scale).
- **Language:** JavaScript (ES6+)
  - **Rationale:** Universally known, easy for juniors, and well supported in browsers and tooling.
  - **Trade-offs:** No type safety (unless using TypeScript), but simplicity is prioritized for this increment.
  - **Alternatives Considered:** TypeScript (adds complexity), Python/Go (not suitable for frontend/mobile web).

## Initial Approach

### Data Loading
**Approach:** Client loads sample data from a local JSON file.
**Rationale:** Enables fast prototyping and testing without backend dependencies, supporting rapid iteration.
**Trade-offs:** No real data integration; limited to static sample data. Acceptable for MVP and usability testing.
**Alternatives to Consider:** Mock API endpoint or in-memory data generation for future extensibility.

### State Management
**Approach:** State lives in the catalog list component only.
**Rationale:** Keeps implementation simple and focused; avoids unnecessary global state for a single screen.
**Trade-offs:** Not scalable for multi-screen or shared state, but ideal for isolated MVP features.
**Alternatives to Consider:** Global store if catalog state needs to be shared or persisted later.

### Image Handling
**Approach:** Use URLs to placeholder images (e.g., Unsplash) for catalog items.
**Rationale:** Simplifies asset management and ensures visually appealing sample data without bundling images.
**Trade-offs:** Reliance on external image service; may require fallback for offline use. Acceptable for MVP.
**Alternatives to Consider:** Local static assets or generated avatars/icons for offline or custom branding needs.

## Architecture Overview

**Components:**
- CatalogListScreen (Svelte): Loads and displays catalog items
- CatalogItem (Svelte): Renders image, last used date, and user name for each item

**Data Flow:**
CatalogListScreen loads local JSON → Maps data to CatalogItem components → Renders responsive list on mobile

**Integration Points:**
- Unsplash (or similar): Provides placeholder images via URLs

**State Management:**
- State is managed within CatalogListScreen; no global or backend state required

## Implementation Constraints
- No backend integration; all data is local and static
- Must be responsive and readable on mobile devices
- Images must be accessible via URLs; fallback required for offline scenarios
- "Never"/"N/A" must display for unused items

## Open Questions
- Should we support offline image fallback for MVP?
- How will sample data be updated or extended for future usability tests?
- What is the best way to handle error states (e.g., image load failure)?
