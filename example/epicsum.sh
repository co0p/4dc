#!/bin/bash

# Ticket pattern for validation and extraction
TICKET_PATTERN='\[([A-Za-z0-9-]+)\]'

# Date format (auto-detected from first data row)
DATE_FORMAT=""

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

# Sum time by ticket ID with monthly grouping using awk for aggregation
sum_time_by_ticket() {
    local file="$1"
    
    # Use awk to process and aggregate by month
    awk -F';' '
    NR == 1 { next }  # Skip header
    {
        # Extract and parse date (field 1)
        date_field = $1
        gsub(/^[ \t]+|[ \t]+$/, "", date_field)
        
        # Detect date format and extract year/month
        if (index(date_field, ".") > 0) {
            # DD.MM.YY format
            split(date_field, d, ".")
            day = d[1]
            month = d[2]
            year = "20" d[3]  # Convert YY to 20YY
        } else {
            # YYYY-MM-DD format
            split(date_field, d, "-")
            year = d[1]
            month = d[2]
            day = d[3]
        }
        
        # Create numeric month key for sorting (YYYYMM)
        month_key = year sprintf("%02d", month)
        
        # Create display format (YYYY/MM)
        month_display[month_key] = year "/" sprintf("%02d", month)
        
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
        
        # Accumulate by month and ticket (use concatenated key)
        key = month_key SUBSEP ticket_id
        month_ticket_minutes[key] += total_mins
        month_ticket_counts[key] += 1
        
        # Track tickets per month
        if (!(month_key in month_tickets_seen)) {
            month_tickets_seen[month_key] = ticket_id
        } else if (index(month_tickets_seen[month_key], ticket_id) == 0) {
            month_tickets_seen[month_key] = month_tickets_seen[month_key] "," ticket_id
        }
        
        # Accumulate monthly totals
        month_totals[month_key] += total_mins
        month_counts[month_key] += 1
    }
    END {
        # Collect and sort months in descending order
        n = 0
        for (month in month_totals) {
            months[n++] = month
        }
        
        # Simple bubble sort descending
        for (i = 0; i < n; i++) {
            for (j = i + 1; j < n; j++) {
                if (months[i] < months[j]) {
                    temp = months[i]
                    months[i] = months[j]
                    months[j] = temp
                }
            }
        }
        
        # Print results grouped by month
        for (m = 0; m < n; m++) {
            month_key = months[m]
            
            # Calculate monthly total time
            month_mins = month_totals[month_key]
            month_hours = int(month_mins / 60)
            month_min_part = month_mins % 60
            
            # Print month header
            printf "\n=== %s - Total: %d:%02d (%d entries) ===\n\n", \
                month_display[month_key], month_hours, month_min_part, month_counts[month_key]
            
            # Print ticket table header
            printf "%-12s | %-12s | %-8s\n", "Ticket ID", "Total Time", "Entries"
            printf "%-12s-+-%-12s-+-%-8s\n", "------------", "------------", "--------"
            
            # Print tickets for this month
            split(month_tickets_seen[month_key], tickets, ",")
            for (idx in tickets) {
                ticket_id = tickets[idx]
                if (ticket_id == "") continue
                
                key = month_key SUBSEP ticket_id
                ticket_mins = month_ticket_minutes[key]
                ticket_hours = int(ticket_mins / 60)
                ticket_min_part = ticket_mins % 60
                printf "%-12s | %d:%02d       | %-8d\n", \
                    ticket_id, ticket_hours, ticket_min_part, month_ticket_counts[key]
            }
        }
    }
    ' "$file"
}

sum_time_by_ticket "$1"
