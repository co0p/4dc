# Validate CSV Format

## Job Story
When I provide a CSV file to EpicSum  
I want to know immediately if it's malformed  
So I can fix format issues before attempting to sum hours

**Assumption Being Tested:** Early CSV format validation catches data issues before processing, preventing incorrect time aggregation.

## Acceptance Criteria

- **Given** a CSV file with correct header row and valid semicolon-separated data  
  **When** validation runs  
  **Then** the script proceeds to next processing step without errors

- **Given** a CSV file missing the header row  
  **When** validation runs  
  **Then** script aborts with error "Invalid CSV: Missing or incorrect header row"

- **Given** a CSV file with rows containing wrong number of fields  
  **When** validation runs  
  **Then** script aborts with error showing first problematic line "Invalid CSV: Line X has Y fields, expected 6"

- **Given** a CSV file with a row missing the ticket pattern [E#] in Description column  
  **When** validation runs  
  **Then** script aborts with error "Invalid CSV: Line X missing ticket pattern [<ticketid>] in Description"

- **Given** a CSV file where semicolons are missing or misplaced  
  **When** validation runs  
  **Then** script aborts with error showing first malformed line

## Success Signal
Valid CSV files pass through without errors, and invalid CSV files abort on the first error with clear indication of the problem location and type.

## Out of Scope
- Data type validation (date format, time format)
- Duplicate ticket ID detection
- Empty field validation (e.g., blank Person or Project)
- Full error collection (we fail fast on first error)
