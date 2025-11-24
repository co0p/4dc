# ADR: Sum Time by Ticket

**Date:** 2025-11-24  
**Status:** Accepted

## Decision Summary
Aggregate time entries by ticket ID using awk's built-in associative arrays and minute-based arithmetic. Extract ticket IDs and process data in a single awk pass for maximum portability and simplicity. Format output as a readable table to stdout.

## Technical Decisions

### Data Structure for Aggregation
**Choice:** Awk associative arrays with ticket ID as key, accumulated minutes as value  
**Rationale:** Awk's associative arrays work on all systems (no bash version dependency). Single awk pass processes entire CSV without shell loops or multiple tool invocations. Built-in string processing and arithmetic. Aligns with constitutional simplicity—no external dependencies, runs everywhere.  
**Trade-offs:** All logic in awk script (less familiar to pure bash developers) but eliminates bash 4+ requirement. macOS ships with bash 3.2, making awk the portable choice. Accept awk learning curve for universal compatibility.  
**Alternatives Rejected:** Bash associative arrays (requires bash 4+, not available on macOS by default), temporary files per ticket (excessive I/O, harder to clean up), multiple grep/cut passes (inefficient, complex pipeline)

### Ticket ID Extraction
**Choice:** Awk's `match()` function with regex to extract `[E#]` pattern from Description field  
**Rationale:** Awk's `match($5, /\[E[0-9]+\]/)` sets RSTART/RLENGTH, then `substr()` extracts exact pattern. Single tool handles parsing, extraction, and aggregation. Pattern already validated during CSV format check, so extraction is safe.  
**Trade-offs:** Slightly more complex than grep pipeline but keeps all logic in one awk invocation (faster, cleaner). Accept awk regex syntax over bash familiarity.  
**Alternatives Rejected:** Grep with capture group (requires separate subprocess per row), sed (more complex for extraction), pure bash regex (bash 4+ features not portable)

### Time Arithmetic Strategy
**Choice:** Convert H:MM to total minutes in awk, accumulate, then convert back to H:MM for display  
**Rationale:** Minute-based arithmetic avoids carry-over complexity (60 minutes = 1 hour). Awk handles this naturally: `split(time_value, t, ":")` then `hours * 60 + mins`. Sum totals, then `int(total_mins / 60)` and `total_mins % 60` for output. Clear, testable logic.  
**Trade-offs:** Two conversions (to/from minutes) but eliminates edge cases like minute overflow. Awk's integer arithmetic is precise and fast. Accept conversion overhead for bug-free arithmetic and code clarity.  
**Alternatives Rejected:** Direct H:MM addition (complex overflow handling), decimal hours (loses H:MM format requirement), bash arithmetic with arrays (portability issues)

### Output Formatting
**Choice:** Awk printf-based table with fixed column widths for alignment  
**Rationale:** Awk `printf` provides clean columnar output identical to bash printf. Format string `"%-12s | %-12s | %-8s\n"` ensures consistent alignment. Single tool handles aggregation and formatting—no separate display logic needed.  
**Trade-offs:** Fixed column widths may truncate very long ticket IDs but keeps implementation simple. Accept this for v1 (constitutional "good enough" principle).  
**Alternatives Rejected:** Column utility (external dependency), dynamic width calculation (over-engineering), CSV output (less human-readable), bash printf (requires bash arrays for data)

## Architecture Overview

**Components:**
- `sum_time_by_ticket()`: Bash function wrapper that invokes awk script
- Awk script: Single-pass processor handling extraction, conversion, aggregation, and formatting
- Associative arrays (awk): `ticket_minutes` (total mins), `ticket_counts` (entry count)

**Data Flow:**
CSV validation passes → Invoke awk with semicolon delimiter → Skip header (NR==1) → For each row: extract ticket ID (match/substr) → extract time (field 6, trim whitespace) → split time on `:`, convert to minutes → accumulate in `ticket_minutes[ID]` → increment `ticket_counts[ID]` → END block: convert minutes to H:MM, format table, print to stdout

**Integration Points:**
- Called after `validate_csv_format` completes successfully
- Reads same CSV file passed as $1 argument
- Outputs directly to stdout (no file writes)
- Pure awk solution—no bash loops or external tool dependencies

**State Management:**
- Awk associative arrays hold aggregated state in memory during processing
- No persistent state—pure transformation of CSV to summary output
- Arrays automatically cleared when awk terminates

## Implementation Constraints
- Requires awk (universally available on Unix systems)
- Works with bash 3.2+ (no modern bash features needed)
- Input CSV must already be validated (depends on `validate_csv_format` function)
- Time format must be H:MM (single digit or double digit hours, colon, two digit minutes)
- Ticket IDs must match `[E#]` pattern (validated in CSV check)
- All processing in memory—CSV size limited by available RAM (acceptable for typical use cases)
- Output column widths fixed (ticket ID ≤12 chars, time ≤12 chars for display)

## Open Questions
- Should output be sorted by ticket ID or total time? (Defer: natural order from awk array iteration for v1)
- Handle hours >99 in output (e.g., 125:30)? (Defer: unlikely in typical usage, accept potential display width issue)
- Grand total row at bottom? (Defer: not in acceptance criteria, add if users request)
