# Design: Backend Serves Frontend as Single Deployable
**Date:** 2025-12-01  
**Status:** Initial Technical Design

## Design Summary
We will package and serve the Svelte frontend as static assets from the Go backend, allowing deployment as a single unit. The Go backend will handle both API requests and serve the built frontend from a static directory. This approach supports rapid iteration and aligns with the constitution’s principles of simplicity and maintainability. [Reference: Constitution Principle 1, ADR-serve-frontend-from-backend]

## Technical Decisions
- **Frameworks:** Go (net/http), Svelte
	- **Rationale:** Go is reliable and simple for serving APIs and static files; Svelte is already used for the frontend.
	- **Trade-offs:** No independent scaling or deployment; all assets must be copied into backend static directory.
	- **Alternatives Considered:** Go embed for assets (adds complexity), separate static server (not needed for MVP).
- **Build/Deploy:** Copy built frontend assets into backend static directory
	- **Rationale:** Simple and robust for MVP and small teams.
	- **Trade-offs:** Requires build step and asset copy; not suitable for large-scale deployments.
	- **Alternatives Considered:** Use Go embed, custom build scripts, or external CDN.

## Initial Approach
### Serve Frontend from Backend Static Directory
**Approach:** Build Svelte frontend, copy output to Go backend `/static` directory, serve with `http.FileServer`.
**Rationale:** Fast to implement, easy to maintain, minimal infrastructure.
**Trade-offs:** No asset versioning or CDN; all assets must be rebuilt and copied for updates.
**Alternatives to Consider:** Use Go embed for assets, external CDN for frontend.

## Architecture Overview
**Components:**
- Go backend: Serves API endpoints and static frontend assets
- Svelte frontend: Built and served as static files

**Data Flow:**
Client → Go backend (`/api/*` for API, `/` and `/static/*` for frontend) → Client renders app

**Integration Points:**
- `/api/*`: API endpoints
- `/`: Serves frontend index.html
- `/static/*`: Serves frontend assets

**State Management:**
- Frontend state lives in client (browser)
- Backend state lives in Go server

## Implementation Constraints
- All frontend assets must be copied into backend static directory after build
- No independent deployment or scaling of frontend/backend
- Backend must serve both API and static files
[Reference: ADR-serve-frontend-from-backend]

## Open Questions
- How will asset updates be managed for future increments?
- Should we add asset versioning or cache busting?
- When should we consider splitting deployables for scaling?
[Reference: ADR-serve-frontend-from-backend]
