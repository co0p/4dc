# Increment: Serve Frontend from Backend as Single Deployable

## Job Story
**When** I build and run the ShareIt app
**I want to** serve the frontend static assets from the Go backend
**So I can** deploy and run the entire app as a single unit

**Assumption Being Tested:**
This approach works for MVP and small teams, and all features remain accessible via the backend server only.

## Acceptance Criteria
- **Given** the Go backend is built and running
  **When** I access the root URL in my browser
  **Then** the frontend app is served and fully functional

- **Given** the backend is running
  **When** I access `/api/catalog` and other API endpoints
  **Then** API responses are returned as expected

- **Given** the build process is complete
  **When** I deploy the app
  **Then** only one deployable unit (container or binary) is required

- **Given** a build or deploy error occurs
  **When** I run the build or start the server
  **Then** a clear error message is shown

## Success Signal
Single deployable unit is built and runs locally; all app features work via backend server only

## Out of Scope
- Independent deployment of frontend and backend
- Advanced scaling or load balancing
- Multi-environment configuration
