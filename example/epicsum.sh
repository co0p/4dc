#!/bin/bash

# Ticket pattern for validation and extraction
TICKET_PATTERN='\[([A-Za-z0-9-]+)\]'

# Argument validation: Check if input file argument is provided
if [ -z "$1" ]; then
    echo "please specify an input csv" >&2
    exit 1
fi

# File existence check: Verify the file exists
if [ ! -f "$1" ]; then
    echo "Error: File '$1' not found" >&2
    exit 1
fi

# CSV format validation
validate_csv_format() {
    local file="$1"
    local expected_header="Date; Person; Project; Aufgabe; Description; Time"
    
    # Read and validate header (strip carriage returns for Windows line endings)
    local header=$(head -n 1 "$file" | tr -d '\r')
    if [ "$header" != "$expected_header" ]; then
        echo "Invalid CSV: Missing or incorrect header row" >&2
        exit 1
    fi
    
    # Validate data rows
    local line_num=1
    while IFS= read -r line; do
        line_num=$((line_num + 1))
        
        # Strip carriage returns for Windows line endings
        line="${line%$'\r'}"
        
        # Skip header row
        if [ $line_num -eq 2 ]; then
            continue
        fi
        
        # Count fields using awk
        local field_count=$(echo "$line" | awk -F';' '{print NF}')
        if [ "$field_count" -ne 6 ]; then
            echo "Invalid CSV: Line $line_num has $field_count fields, expected 6" >&2
            exit 1
        fi
        
        # Extract Description field (field 5) and check for ticket pattern
        local description=$(echo "$line" | awk -F';' '{print $5}')
        if ! echo "$description" | grep -qE "$TICKET_PATTERN"; then
            echo "Invalid CSV: Line $line_num has invalid ticket pattern in Description" >&2
            exit 1
        fi
    done < "$file"
}

validate_csv_format "$1"

# Sum time by ticket ID using awk for aggregation
sum_time_by_ticket() {
    local file="$1"
    
    # Use awk to process and aggregate
    awk -F';' '
    NR == 1 { next }  # Skip header
    {
        # Extract ticket ID from Description (field 5)
        match($5, /\[([A-Za-z0-9-]+)\]/)
        ticket_id = substr($5, RSTART, RLENGTH)
        
        # Extract and trim time value (field 6)
        gsub(/^[ \t]+|[ \t]+$/, "", $6)
        time_value = $6
        
        # Convert H:MM to minutes
        split(time_value, t, ":")
        hours = t[1]
        mins = t[2]
        total_mins = hours * 60 + mins
        
        # Accumulate
        ticket_minutes[ticket_id] += total_mins
        ticket_counts[ticket_id] += 1
    }
    END {
        # Print header
        printf "%-12s | %-12s | %-8s\n", "Ticket ID", "Total Time", "Entries"
        printf "%-12s-+-%-12s-+-%-8s\n", "------------", "------------", "--------"
        
        # Print results
        for (ticket_id in ticket_minutes) {
            total_mins = ticket_minutes[ticket_id]
            hours = int(total_mins / 60)
            mins = total_mins % 60
            printf "%-12s | %d:%02d       | %-8d\n", ticket_id, hours, mins, ticket_counts[ticket_id]
        }
    }
    ' "$file"
}

sum_time_by_ticket "$1"
