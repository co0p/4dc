#!/bin/bash

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
    
    # Read and validate header
    local header=$(head -n 1 "$file")
    if [ "$header" != "$expected_header" ]; then
        echo "Invalid CSV: Missing or incorrect header row" >&2
        exit 1
    fi
    
    # Validate data rows
    local line_num=1
    while IFS= read -r line; do
        line_num=$((line_num + 1))
        
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
        if ! echo "$description" | grep -q '\[E[0-9]\+\]'; then
            echo "Invalid CSV: Line $line_num missing ticket pattern [<ticketid>] in Description" >&2
            exit 1
        fi
    done < "$file"
}

validate_csv_format "$1"

# Main logic will go here
echo "Processing file: $1"
