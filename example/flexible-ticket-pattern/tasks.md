# Tasks: Flexible Ticket Pattern Matching

## Relevant Files

- `epicsum.sh` - Main bash script with validation and aggregation logic

## Progress Tracking

**As you complete each task, check it off by changing `- [ ]` to `- [x]`.**

## Implementation Tasks

- [x] **Setup**
  - [x] Create feature branch: `git checkout -b feature/flexible-ticket-pattern`

- [x] **Define pattern variable**
  - [x] Add `TICKET_PATTERN='\[([A-Za-z0-9-]+)\]'` near top of script (after shebang, around line 3)
  - [x] Verify: Variable is accessible throughout script

- [x] **Update validation logic**
  - [x] Replace hardcoded pattern `\[E[0-9]\+\]` with `$TICKET_PATTERN` in validation function (around line 46)
  - [x] Update error message to show invalid pattern when validation fails
  - [x] Verify: Script rejects empty brackets `[]` with error
  - [x] Verify: Script rejects patterns with spaces `[Feature X]` with error

- [x] **Update aggregation logic**
  - [x] Replace hardcoded pattern `\[E[0-9]+\]` in awk match() call (around line 67)
  - [x] Verify: Pattern extraction works with new regex

- [x] **Test backward compatibility**
  - [x] Run script with existing test CSV containing `[E1]`, `[E2]` patterns
  - [x] Verify: Output matches previous behavior, no regressions

- [x] **Test new ticket formats**
  - [x] Create test CSV with `[JIRA-123]` pattern
  - [x] Verify: Script processes successfully, outputs aggregated time for `[JIRA-123]`
  - [x] Create test CSV with `[BUG-42]` pattern
  - [x] Verify: Script processes successfully, outputs aggregated time for `[BUG-42]`
  - [x] Create test CSV with `[ABC123]` pattern (no hyphen)
  - [x] Verify: Script processes successfully, outputs aggregated time for `[ABC123]`

- [x] **Test validation edge cases**
  - [x] Create test CSV with empty brackets `[]` in Description
  - [x] Verify: Script aborts with error "Invalid CSV: Line X missing ticket pattern [<ticketid>] in Description"
  - [x] Create test CSV with `[Feature X]` (space) in Description
  - [x] Verify: Script aborts with error "Invalid CSV: Line X missing ticket pattern [<ticketid>] in Description"

- [x] **Integration & Verification**
  - [x] Test complete flow with mixed patterns: `[E1]`, `[JIRA-123]`, `[BUG-42]` in one CSV
  - [x] Verify: All patterns aggregate correctly
  - [x] Verify: All 6 acceptance criteria from feature.md pass

- [x] **Verify ADR Compliance**
  - [x] Confirm pattern defined once as variable (not duplicated)
  - [x] Confirm pattern used in both validation and aggregation sections
  - [x] Confirm error messages follow simple format from ADR
  - [x] Remove any dead code or commented-out old patterns

- [x] **Deploy**
  - [x] Commit changes: "Add flexible ticket pattern matching support"
  - [x] Merge to main (direct commit per constitution)
