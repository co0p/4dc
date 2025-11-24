# Tasks: Sum Time by Ticket

## Relevant Files

- `epicsum.sh` - Main script with time aggregation logic added after CSV validation

## Progress Tracking

**As you complete each task, check it off by changing `- [ ]` to `- [x]`.**

## Implementation Tasks

- [x] **Setup**
  - [x] Create feature branch: `git checkout -b feature/sum-time-by-ticket`

- [x] **Ticket ID extraction**
  - [x] Add function to extract ticket ID from Description field using grep -oE
  - [x] Pattern: `\[E[0-9]+\]` from field 5
  - [x] Verify: Test extraction with sample Description values, confirm correct ticket ID returned

- [x] **Time conversion helpers**
  - [x] Implement H:MM to minutes conversion (split on `:`, calculate hours*60 + minutes)
  - [x] Implement minutes to H:MM conversion (divide by 60 for hours, modulo 60 for minutes)
  - [x] Verify: Test with 2:00 (120 mins), 4:15 (255 mins), 0:30 (30 mins) conversions work both ways

- [x] **Time aggregation logic**
  - [x] Use awk associative arrays for bash 3.2 compatibility
  - [x] Loop through CSV rows (skip header), extract ticket ID and time
  - [x] Convert time to minutes, accumulate in array
  - [x] Increment entry counts
  - [x] Verify: Test with sample CSV, check arrays contain correct accumulated values

- [x] **Table output formatting**
  - [x] Print table header: "Ticket ID | Total Time | Entries"
  - [x] Loop through ticket array keys
  - [x] Convert accumulated minutes back to H:MM format
  - [x] Use printf with fixed widths: `"%-12s | %-12s | %-8s\n"`
  - [x] Verify: Output shows aligned columns with ticket ID, formatted time, entry count

- [x] **Integration & Verification**
  - [x] Test with input.csv: [E1] appears twice (2:00 + 4:15 = 6:15), [E2] once (3:30), [E3] once (5:00)
  - [x] Verify table shows all three tickets with correct totals and entry counts
  - [x] Test time carry-over: manually add row with 2:50, another with 0:30, confirm total 3:20
  - [x] Verify all acceptance criteria from use case pass

- [x] **Deploy**
  - [x] Commit changes: "Add time aggregation by ticket ID"
  - [x] Merge to main: `git checkout main && git merge feature/sum-time-by-ticket`
