# Domain Vocabulary

Shared language for this project. Update when the product's domain vocabulary changes.

---

## Concepts

### [ConceptName]

**Definition:** [One sentence in domain terms].

**State:**
- `[field]` — [what it represents and valid values]

**Rules:**
- [Invariant or constraint in domain language]

**Related:**
- [Other concepts this connects to]
- [Events it raises]

---

## Domain Events

### [EventName]

**When:** [What triggers this event in domain terms]
**Payload:** [Key fields involved]
**Consumers:** [Who or what reacts in the domain]

---

## Rules and Constraints

[System-wide invariants, state transitions, cross-concept rules described in domain language, not implementation details]

---

## Example

### User

**Definition:** A person who has authenticated and is interacting with the system.

**State:**
- `email` — contact address; must be unique
- `status` — `active`, `suspended`, or `deleted`

**Rules:**
- Email uniqueness enforced across the system
- Status can transition: `active` ↔ `suspended` → `deleted` (one-way)
- Cannot be deleted if currently has active sessions

**Related:**
- Session (one User has many Sessions)
- UserCreated, UserSuspended, UserDeleted (events raised)


### Session

**Definition:** An authenticated connection between a User and the system.

**State:**
- `user_id` — reference to owning User
- `status` — `active`, `expired`, or `revoked`
- `expires_at` — when the session becomes invalid

**Rules:**
- Expiry time is set at creation and cannot be changed
- Session cannot outlive the User who owns it
- Multiple Sessions can exist for one User

**Related:**
- User (many Sessions per User)
- SessionCreated, SessionExpired, SessionRevoked (events raised)

### UserCreated

**When:** A new User completes authentication for the first time.
**Payload:** `user_id`, `email`, timestamp
**Consumers:** Email service, audit log, user preference initialization

### SessionExpired

**When:** A Session's expiry time elapses or is explicitly revoked.
**Payload:** `session_id`, `user_id`, reason (timeout, logout, or user deletion)
**Consumers:** Cache invalidation, session cleanup, session analytics

---

## System Rules

**User Deletion:** When a User is deleted, all their Sessions are revoked in the same operation.

**Session Lifetime:** A Session cannot outlast the User who created it. If a User is deleted while Sessions exist, those Sessions are revoked.

**Email Uniqueness:** No two active Users can share an email address.
