# ADR: Flexible Ticket Pattern Matching

**Date:** 2025-11-24  
**Status:** Accepted

## Decision Summary
Extend ticket ID pattern matching from `[E#]` to support alphanumeric characters and hyphens `[A-Za-z0-9-]+`. Pattern defined once as a bash variable and reused across validation and aggregation to ensure consistency and maintainability.

## Technical Decisions

### Pattern Definition Strategy
**Choice:** Define regex pattern once as a bash variable at the top of the script and reuse in both validation and aggregation sections  
**Rationale:** Single source of truth prevents inconsistencies between validation and aggregation. If pattern changes in future, only one location needs updating. Aligns with "Simplicity Over Flexibility" principle—no functions needed, just a variable.  
**Trade-offs:** Slightly less encapsulated than a function, but simpler to understand and maintain. Accept this for better readability in a single-file bash script.  
**Alternatives Rejected:** 
- Hardcode pattern in two places - rejected due to maintenance risk (easy to update one but forget the other)
- Create validation function - rejected as over-engineering for this scope; variable is sufficient

### Regex Pattern Design
**Choice:** `\[([A-Za-z0-9-]+)\]` - matches brackets containing one or more alphanumeric characters or hyphens  
**Rationale:** Covers all specified acceptance criteria (`[E1]`, `[JIRA-123]`, `[BUG-42]`, `[ABC123]`) while rejecting empty brackets `[]` and patterns with spaces. The `+` quantifier ensures at least one character is required.  
**Trade-offs:** Accepts hyphen anywhere in pattern (including `[---]`), but this is acceptable for v1 simplicity. Users are unlikely to use pathological patterns. Accept this edge case to avoid regex complexity.  
**Alternatives Rejected:**
- More restrictive pattern requiring letter-first - rejected as unnecessary constraint
- Pattern allowing spaces - rejected per acceptance criteria
- Pattern allowing underscores/dots - rejected as explicitly out of scope

### Error Message Format
**Choice:** Simple format: `"Invalid CSV: Line X has invalid ticket pattern [...]"`  
**Rationale:** Follows "Fail Fast, Fail Loud" principle with clear, actionable error. Shows exactly what's wrong (which line, which pattern). User can inspect the CSV and fix it. Detailed format with regex explanation would confuse non-technical users.  
**Trade-offs:** Doesn't teach user the valid format, but keeps error messages clean and consistent with existing validation errors. Accept this for simplicity—users can refer to documentation if needed.  
**Alternatives Rejected:**
- Detailed error with regex explanation - rejected as too technical for typical users
- Error with examples - rejected as too verbose; users can check documentation

## Architecture Overview

**Components:**
- Pattern Variable: `TICKET_PATTERN` defined at script start (near line 10-15)
- CSV Validation: Uses `$TICKET_PATTERN` to validate Description column (around line 46)
- Ticket Extraction: Uses same pattern in awk aggregation logic (around line 83)

**Data Flow:**
1. Script defines `TICKET_PATTERN='\[([A-Za-z0-9-]+)\]'` at initialization
2. Validation phase: grep/awk checks each Description line matches pattern
3. Aggregation phase: awk extracts ticket ID using same pattern
4. Output: Aggregated time per ticket ID

**Integration Points:**
- No external integrations—pure bash with standard utilities (grep, awk)

**State Management:**
- Pattern stored as script-level variable (not configuration)
- No persistent state—stateless processing of CSV input

## Implementation Constraints
- Pattern must be escaped properly for both grep and awk contexts
- Empty brackets `[]` must be explicitly rejected (regex `+` quantifier handles this)
- Spaces in ticket IDs must be rejected (regex excludes space character)
- Backward compatibility: existing `[E1]`, `[E2]`, etc. must continue working
- No configuration file—pattern is hardcoded in script (per constitution)

## Open Questions
- Should we validate hyphen position (e.g., disallow leading/trailing hyphens)? Defer until users report issues.
- Should we support case-insensitive matching? Current design is case-sensitive; defer until user feedback shows need.
