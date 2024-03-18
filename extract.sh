#!/bin/bash

# Check if the script argument (path to the .gresource file) is present
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <path to .gresource file>"
    exit 1
fi

# Define the .gresource file
GRESOURCE_FILE="$1"

# Create a directory for the extracted files
EXTRACT_DIR="./extracted_files"
mkdir -p "$EXTRACT_DIR"

# Get the list of resources from the .gresource file
RESOURCE_LIST=$(gresource list "$GRESOURCE_FILE")

# Extract each resource
echo "Extracting resources from $GRESOURCE_FILE..."
for RESOURCE in $RESOURCE_LIST; do
    # Determine the full path for the extracted resource within EXTRACT_DIR
    RESOURCE_PATH="${EXTRACT_DIR}${RESOURCE}"
    
    # Create necessary directories based on the resource path
    mkdir -p "$(dirname "$RESOURCE_PATH")"
    
    # Extract the resource
    gresource extract "$GRESOURCE_FILE" "$RESOURCE" > "$RESOURCE_PATH"
done

echo "All resources have been extracted to the directory $EXTRACT_DIR."
