#!/bin/bash

# Check if the filename argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <filename>"
    exit 1
fi

# Assign the provided filename to a variable
filename=$1

# Specify the Git repository URL
git_repo_url="https://raw.githubusercontent.com/aB34R/PSGO/main"

# Use curl to download the file from the Git repository
curl -O "$git_repo_url/$filename"

# Check if the download was successful
if [ $? -eq 0 ]; then
    echo "Download successful: $filename"
else
    echo "Error downloading: $filename"
fi

