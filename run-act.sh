#!/bin/bash

# Script to run GitHub Actions locally using act
# This script helps test the GitHub Actions workflow without pushing to GitHub

# Colors for better output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Print header
echo -e "${BLUE}==================================================${NC}"
echo -e "${BLUE}  Testing GitHub Actions Workflow with act${NC}"
echo -e "${BLUE}==================================================${NC}"

# Check if act is installed
if ! command -v act &> /dev/null; then
    echo -e "${RED}Error: 'act' is not installed.${NC}"
    echo -e "Please install it first. See: https://github.com/nektos/act"
    exit 1
fi

# Create artifacts directory if it doesn't exist
mkdir -p /tmp/artifacts

echo -e "${GREEN}Running GitHub Actions workflow locally...${NC}"

# Run act with specific parameters to make it simpler
act -P ubuntu-latest=nektos/act-environments-ubuntu:18.04 -j test

# Check the result
if [ $? -eq 0 ]; then
    echo -e "${GREEN}Success: GitHub Actions workflow completed successfully!${NC}"
    echo -e "Artifacts can be found in /tmp/artifacts"
else
    echo -e "${RED}Error: GitHub Actions workflow failed.${NC}"
    echo -e "Please check the output above for details."
fi 