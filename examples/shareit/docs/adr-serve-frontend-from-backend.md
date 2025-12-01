# ADR: Serve Frontend from Backend as a Single Deployable Unit

## Context
Currently, the ShareIt project separates the Go backend and Svelte frontend into distinct deployables. For some deployment scenarios, it may be preferable to serve the frontend directly from the backend, packaging both as a single deployable unit. This can simplify deployment, reduce infrastructure complexity, and ensure tight integration between backend and frontend.

## Decision
Package and serve the Svelte frontend as static assets from the Go backend. The backend will handle both API requests and serve the frontend application, allowing deployment as a single unit (e.g., a Docker container or binary). This approach is suitable for small teams, MVPs, and environments where simplicity is prioritized over independent scaling.

## Consequences
- Simplifies deployment and infrastructure (single container or binary)
- Ensures backend and frontend are always in sync
- Reduces operational overhead for small projects
- Limits independent scaling and updates of backend/frontend
- May require custom build steps to copy frontend assets into backend static directory

## Alternatives Considered
- Keep backend and frontend as separate deployables: Enables independent scaling and updates, but increases deployment complexity
- Use a reverse proxy to route requests: Adds infrastructure overhead, but maintains separation

---
