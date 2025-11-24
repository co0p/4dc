# ADR: Validate CSV Format

**Date:** 2025-11-24  
**Status:** Draft

## Decision Summary
Implement CSV format validation using bash for header verification and awk for row-level field counting and ticket pattern detection. Validation occurs immediately after file existence check and aborts on first error, aligning with "Fail Fast, Fail Loud" principle.

## Technical Decisions

### CSV Parsing Strategy
**Choice:** Hybrid approach—bash string operations for header check, awk for row validation  
**Rationale:** Header is a one-time check where bash string comparison is simple and clear. Row-by-row validation benefits from awk's field-splitting capabilities (`-F';'`) which handles semicolon delimiters naturally. This keeps both operations readable.  
**Trade-offs:** Two different parsing tools (bash + awk) but leverages each tool's strength. Accept slight complexity for clearer code versus forcing everything through one tool.  
**Alternatives Rejected:** Pure bash (too verbose for field counting), pure awk (header string comparison less clear)

### Ticket Pattern Detection
**Choice:** Grep/regex on the 5th field (Description column) of each row  
**Rationale:** Description column is field 5 in the CSV. Extract field with awk, pipe to grep with pattern `\[E[0-9]+\]` to detect ticket format. Simple pipeline that leverages bash's text processing strength.  
**Trade-offs:** Two-step process (awk extract + grep match) but each step has single responsibility. Accept pipeline overhead for code clarity.  
**Alternatives Rejected:** Awk pattern matching (more complex regex in awk), bash regex (less efficient for pattern matching at scale)

### Validation Flow Integration
**Choice:** Separate validation function called after file existence check  
**Rationale:** Validation is a distinct concern from argument checking. Separate function (`validate_csv_format`) can be tested independently and keeps main flow readable: check args → check file exists → validate CSV → process. Follows existing pattern from input validation.  
**Trade-offs:** Additional function call overhead (negligible) but better separation of concerns and testability.  
**Alternatives Rejected:** Inline validation (clutters main flow), combined function (mixes argument and format concerns)

### Error Reporting Strategy
**Choice:** Fail on first error with line number and specific problem description  
**Rationale:** Aligns with "Fail Fast, Fail Loud" principle. User fixes one issue at a time, re-runs, and discovers next issue. Keeps error handling simple—no error collection or aggregation needed.  
**Trade-offs:** Multiple run cycles to find all errors but simpler implementation and clearer error messages. Users prefer specific "Line 5 has 4 fields" over generic "multiple format errors."  
**Alternatives Rejected:** Full error collection (violates fail-fast principle, adds complexity), silent failure (violates constitution)

## Architecture Overview

**Components:**
- `validate_csv_format()`: Main validation function that orchestrates header and row checks
- Header validation: Bash string comparison against expected column names
- Row validation: Awk field counting + grep pattern matching per row

**Data Flow:**
File existence confirmed → Read first line → Check header columns → Loop through data rows → For each row: count fields (awk) → extract Description field (awk) → check ticket pattern (grep) → Abort on first error or proceed to processing

**Integration Points:**
- Called from main script flow after file existence check
- Errors write to stderr (>&2) and exit with code 1
- Success returns control to main processing logic

**State Management:**
- No persistent state—pure validation pass/fail
- Line number tracked during iteration for error reporting
- Validation results are ephemeral (error or success)

## Implementation Constraints
- Must use bash built-ins and standard utilities only (awk, grep)
- Header must match exactly: `Date; Person; Project; Aufgabe; Description; Time`
- All data rows must have exactly 6 semicolon-separated fields
- Ticket pattern must be `[E<number>]` format in Description column (field 5)
- Error messages must include line number for user correction
- First error aborts entire validation (no error collection)

## Open Questions
- Should we validate header is case-sensitive or case-insensitive? (Defer: assume exact match for v1)
- Should we trim whitespace around semicolons before field counting? (Defer: assume clean CSV for v1)
- What if Description has multiple ticket patterns like "[E1] and [E2]"? (Defer: current pattern check accepts this)
