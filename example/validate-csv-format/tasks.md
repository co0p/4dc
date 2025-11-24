# Tasks: Validate CSV Format

## Relevant Files

- `epicsum.sh` - Main script with CSV validation logic added after file existence check

## Progress Tracking

**As you complete each task, check it off by changing `- [ ]` to `- [x]`.**

## Implementation Tasks

- [x] **Setup**
  - [x] Create feature branch: `git checkout -b feature/validate-csv-format`

- [x] **Header validation**
  - [x] Read first line of CSV file
  - [x] Compare against expected header: `Date; Person; Project; Aufgabe; Description; Time`
  - [x] If mismatch, write error to stderr and exit 1
  - [x] Verify: Test with valid header (passes), invalid header (aborts with error)

- [x] **Field count validation**
  - [x] Loop through CSV data rows (skip header)
  - [x] Use awk with semicolon delimiter to count fields per row
  - [x] If row has != 6 fields, report line number and field count, exit 1
  - [x] Verify: Test with 6-field rows (passes), 4-field row (aborts at line X)

- [x] **Ticket pattern validation**
  - [x] Extract Description field (field 5) from each row using awk
  - [x] Check for pattern `[<id>]` using grep regex
  - [x] If pattern missing, report line number and exit 1
  - [x] Verify: Test with `[E1]` pattern (passes), missing pattern (aborts at line X)

- [x] **Integration & Verification**
  - [x] Test complete validation flow with valid input.csv (all validations pass)
  - [x] Test with missing header (aborts immediately)
  - [x] Test with wrong field count on line 3 (aborts at line 3)
  - [x] Test with missing ticket pattern on line 5 (aborts at line 5)
  - [x] Verify all acceptance criteria from use case pass

- [ ] **Deploy**
  - [ ] Commit changes: "Add CSV format validation"
  - [ ] Merge to main: `git checkout main && git merge feature/validate-csv-format`
