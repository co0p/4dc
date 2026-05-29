# Product Requirements Document (PRD)

## Product
Mobile Website Todo App

## Purpose
Build a simple, fast mobile-first todo app that runs fully in the browser, stores data in localStorage, and has zero external runtime dependencies.

## Goals
- Let users quickly capture and manage personal tasks on mobile.
- Keep the app usable offline after first load.
- Keep implementation lightweight for testing the 4dc prompt phases.

## Non-Goals
- Multi-user collaboration
- Cloud sync or accounts
- Push notifications
- Native mobile app packaging

## Target Users
- Individual users who want a basic personal task list on their phone browser.

## Core User Stories
- As a user, I can add a todo with a title.
- As a user, I can mark a todo as complete or incomplete.
- As a user, I can edit a todo title.
- As a user, I can delete a todo.
- As a user, I can filter todos by all, active, and completed.
- As a user, my todos remain available after closing and reopening the browser.

## Functional Requirements
1. Add todo
	- User can create a task with a non-empty title.
2. Toggle completion
	- User can switch a task between active and completed.
3. Edit todo
	- User can update the title of an existing task.
4. Delete todo
	- User can remove a task permanently.
5. Filter list
	- User can view all, active, or completed tasks.
6. Persist data
	- App stores all task state in localStorage.
7. Empty states
	- App shows meaningful text when no tasks exist or no tasks match a filter.

## Non-Functional Requirements
- Mobile-first layout for common phone widths (320px and above).
- No external dependencies at runtime (no framework, CDN, or third-party service).
- Fast interactions on low-end mobile devices.
- Basic accessibility: semantic HTML, visible focus state, touch-friendly controls.

## Constraints
- Technology: HTML, CSS, and vanilla JavaScript only.
- Storage: browser localStorage only.
- Network: must work without network after initial load.

## Success Criteria
- A new user can add, complete, edit, delete, and filter tasks without guidance.
- Data persists across browser reload and tab close/open.
- App remains fully functional with airplane mode enabled after initial load.

## Acceptance Criteria
- Given an empty app, when the user enters a valid title and submits, then a new todo appears in the list.
- Given an existing todo, when the user toggles completion, then the visual state and stored state update.
- Given an existing todo, when the user edits and saves a new valid title, then the updated title is displayed and persisted.
- Given an existing todo, when the user deletes it, then it is removed from UI and localStorage.
- Given multiple todos with mixed states, when the user selects a filter, then only matching items are shown.
- Given existing todos in localStorage, when the page reloads, then all todos and their states are restored.

## Suggested Delivery Slice
1. Add + list + persist
2. Toggle + filter
3. Edit + delete + empty states
