# Flexible Ticket Pattern Matching

## Job Story
When I process CSV files with different ticket ID formats  
I want the tool to accept any alphanumeric and hyphen pattern between brackets  
So I can track time across multiple projects with different ticketing systems

**Assumption Being Tested:** Supporting flexible ticket ID patterns (alphanumeric and hyphens) enables the tool to work with diverse project naming conventions without restricting users to a single format.

## Acceptance Criteria

- **Given** a CSV with ticket ID `[E1]` in Description  
  **When** validation and aggregation run  
  **Then** the tool processes successfully and outputs `[E1]` with aggregated time

- **Given** a CSV with ticket ID `[JIRA-123]` in Description  
  **When** validation and aggregation run  
  **Then** the tool processes successfully and outputs `[JIRA-123]` with aggregated time

- **Given** a CSV with ticket ID `[BUG-42]` in Description  
  **When** validation and aggregation run  
  **Then** the tool processes successfully and outputs `[BUG-42]` with aggregated time

- **Given** a CSV with ticket ID `[ABC123]` (alphanumeric, no hyphen) in Description  
  **When** validation and aggregation run  
  **Then** the tool processes successfully and outputs `[ABC123]` with aggregated time

- **Given** a CSV with empty brackets `[]` in Description  
  **When** validation runs  
  **Then** script aborts with error "Invalid CSV: Line X has empty ticket pattern []"

- **Given** a CSV with ticket ID `[Feature X]` (contains space) in Description  
  **When** validation runs  
  **Then** script aborts with error "Invalid CSV: Line X has invalid ticket pattern [Feature X]"

## Success Signal
CSV files with various ticket ID formats (`[E1]`, `[JIRA-123]`, `[BUG-42]`, `[ABC123]`) process successfully. Empty brackets `[]` and patterns with spaces are rejected with clear error messages.

## Out of Scope
- Multiple ticket IDs in one Description field
- Ticket IDs without brackets (e.g., plain "JIRA-123")
- Special characters beyond alphanumeric and hyphen (e.g., `[TICKET_123]`, `[PROJ.456]`)
- Configurable pattern rules (fixed regex for v1)
