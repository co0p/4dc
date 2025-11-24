# Rule: Generating an Architecture Decision Record (ADR)

## Goal

To guide an AI assistant in creating a focused Architecture Decision Record that documents the key technical decisions for implementing a **small, testable increment**. The ADR explains WHICH technical approaches were chosen and WHY, focusing on component boundaries, data flow, and trade-offs specific to this feature.

## Prerequisites

Before generating an ADR, verify that both prerequisite files exist:

**Required:**
1. **`CONSTITUTION.md`** - Defines technical decisions, frameworks, and development principles
2. **`[feature-name]/feature.md`** - Defines WHAT the user wants to accomplish

**If either file does not exist:**
- Inform the user which prerequisite is missing
- Suggest creating the constitution first (using `create-constitution.md`)
- Then creating the feature (using `create-feature.md`)
- Do not proceed until both files exist

**If both files exist:**
- Read the constitution to understand the broad technical foundation and constraints
- Read the feature to understand what needs to be built (user goals and acceptance criteria)
- Make technical decisions that comply with the constitution while implementing the feature

## Process

1.  **Verify Prerequisites:** Check for `CONSTITUTION.md` and the relevant `feature.md`.
2.  **Receive Initial Prompt:** The user requests an ADR for a specific feature increment.
3.  **Analyze Context:** Review constitution and feature to understand constraints and requirements.
4.  **Ask Clarifying Questions:** Ask only 2-3 essential technical questions not answered by constitution or feature.
5.  **Generate ADR:** Create a focused decision record for this small increment.
6.  **Save ADR:** Save as `[feature-name]/adr.md`.

## Clarifying Questions (Guidelines)

Ask only 2-3 technical questions not already answered by the constitution or use case. Focus on **this specific increment's** architecture.

### Focus Areas:
- **Component boundaries:** How to split responsibilities for this feature
- **Data flow:** How information moves through the system
- **Integration patterns:** How to connect to external services
- **Storage decisions:** What persists vs. what's ephemeral

**Important:** Don't ask about tech stack, frameworks, or principles—those are in the constitution. Focus on this feature's specific technical decisions.

### Formatting Requirements

- **Number all questions** (1, 2, 3)
- **List options as A, B, C, D, etc.**
- **Always include option X: "Skip this question / I don't know yet"**
- Keep it simple for responses like "1A, 2C"

### Example Questions

```
1. How should data flow for this feature?
   A. Client → API → Database → Client
   B. Client → API (async job) → Client polls
   C. Client → Third-party API directly
   X. Skip this question / I don't know yet

2. Where should this feature's state live?
   A. Client-side only (component state)
   B. Backend session/cache
   C. Database (persistent)
   D. Hybrid (client + backend)
   X. Skip this question / I don't know yet

3. How should this integrate with external services?
   A. Direct API calls
   B. Queue-based async processing
   C. Webhook callbacks
   X. Skip this question / I don't know yet
```

## ADR Structure

Keep ADRs focused on **technical decisions for small increments**. Aim for one screen or less.

```markdown
# ADR: [Feature Name]

**Date:** [YYYY-MM-DD]  
**Status:** Draft | Accepted | Superseded

## Decision Summary
[2-3 sentences: What we're building and the core technical approach]

## Technical Decisions

### [Decision Name]
**Choice:** [What was decided]  
**Rationale:** [Why - technical reasoning tied to this increment]  
**Trade-offs:** [What was gained vs. what was sacrificed]  
**Alternatives Rejected:** [What we didn't choose and why]

### [Decision Name]
**Choice:** [What was decided]  
**Rationale:** [Why]  
**Trade-offs:** [Gains vs. costs]  
**Alternatives Rejected:** [What wasn't chosen]

[Repeat for 2-5 key decisions only]

## Architecture Overview

**Components:**
- [Component name]: [Responsibility for this feature]
- [Component name]: [Responsibility]

**Data Flow:**
[Brief description or simple diagram of how data moves]

**Integration Points:**
- [External service/API]: [How we integrate]

**State Management:**
[Where state lives and why for this feature]

## Implementation Constraints
- [Technical constraint or limitation from these decisions]
- [Another constraint developers need to know]

## Open Questions
- [Technical unknown to resolve during implementation]
- [Deferred decision with reason]
```

---

## Template Example

```markdown
# ADR: Upload Profile Picture

**Date:** 2025-11-24  
**Status:** Accepted

## Decision Summary
Client-side image upload with backend validation and storage in cloud object storage. Prioritizes simple implementation over advanced features like cropping or compression.

## Technical Decisions

### Image Upload Pattern
**Choice:** Client uploads directly to backend API, backend stores in S3-compatible storage  
**Rationale:** Simple request/response pattern fits constitutional "speed over complexity". No presigned URLs or direct-to-S3 needed for MVP.  
**Trade-offs:** Backend handles file upload (bandwidth cost) but simpler to implement and secure. Accept slower uploads for cleaner architecture.  
**Alternatives Rejected:** Direct client→S3 upload (presigned URLs) - rejected as over-engineering for small user base.

### Image Validation
**Choice:** Backend validates file type and size before storage  
**Rationale:** Server-side validation is trustworthy; client-side can be bypassed. Prevents malicious uploads.  
**Trade-offs:** Upload completes then fails validation (wasted bandwidth) but more secure. Accept this for MVP.  
**Alternatives Rejected:** Client-only validation - rejected due to security; dual validation - rejected as unnecessary complexity.

### Storage Location
**Choice:** Cloud object storage (S3/compatible), not database  
**Rationale:** Images are large binary data; database storage would bloat backup size and slow queries. Object storage is designed for this.  
**Trade-offs:** External dependency on S3 but appropriate tool for the job. Cost is negligible at current scale.  
**Alternatives Rejected:** Database blob storage - rejected due to performance; filesystem storage - rejected due to scaling/backup complexity.

## Architecture Overview

**Components:**
- ProfileUploadForm: Handles file selection and upload UI
- ProfileAPI: Validates and stores image, returns URL
- ObjectStorage: Persists image files

**Data Flow:**
User selects file → Client validates size/type → POST /api/profile/picture → Backend validates → Store in S3 → Return public URL → Client displays

**Integration Points:**
- AWS S3 (or compatible): Image storage with public read access

**State Management:**
- Profile picture URL stored in user profile (database)
- Actual image in S3 with predictable key: `profile-pictures/{userId}.jpg`

## Implementation Constraints
- Images must be JPG/PNG only (validation at API boundary)
- Max file size: 5MB (enforced backend, suggested frontend)
- No image transformation (cropping/resizing) - users prepare images before upload
- Public read access required (no authentication on image URLs)

## Open Questions
- CDN needed? Defer until traffic shows need
- Image compression? Defer to post-MVP if users upload large files
- Multiple image formats (WebP)? Defer until browser support is issue
```

## Target Audience

ADRs are read by **developers implementing the feature** who need to understand:
- Which technical approach to take
- Why this approach was chosen over alternatives
- What constraints to respect during implementation

Keep ADRs **technical and tactical**—focused on this specific increment's architecture, not general principles (those are in the constitution).

## No Code in ADRs

**Important:** ADRs document decisions and trade-offs, not implementations.
- Actual code belongs in the project repository
- ADRs guide developers on what to build and why, not exact syntax
- Focus on architectural choices, component boundaries, and data flow
- Leave implementation details to the tasks and code itself

## Output

*   **Format:** Markdown (`.md`)
*   **Location:** `[feature-name]/` directory
*   **Filename:** `adr.md`

## Final Instructions

1. **Verify `CONSTITUTION.md` and `feature.md` exist** before proceeding
2. **Read both documents** to understand principles and requirements
3. **Focus on this increment's technical decisions** - not the entire system
4. **Ask 2-3 clarifying questions** only about architectural gaps
5. **Document 2-5 key decisions** - component boundaries, data flow, integration, storage
6. **Make trade-offs explicit** - what was gained vs. sacrificed
7. **Keep it concise** - aim for one screen, small increment focus
8. **No implementation code** - only architectural decisions and rationale
