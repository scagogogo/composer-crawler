#!/bin/bash

# Script to run all tests for the composer-crawler project

set -e  # Exit on error

# Print header
echo "===================================="
echo "Running composer-crawler tests"
echo "===================================="

# Run tests with verbose flag
go test -v ./pkg/...

# If all tests pass, run coverage report
if [ $? -eq 0 ]; then
  echo "===================================="
  echo "Generating code coverage report"
  echo "===================================="
  
  # Create a directory for coverage reports if it doesn't exist
  mkdir -p coverage
  
  # Generate coverage report
  go test -coverprofile=coverage/coverage.out ./pkg/...
  
  # Display coverage statistics
  go tool cover -func=coverage/coverage.out
  
  # Generate HTML report (optional)
  go tool cover -html=coverage/coverage.out -o coverage/coverage.html
  
  echo "===================================="
  echo "HTML coverage report saved to coverage/coverage.html"
  echo "===================================="
  echo "All tests passed successfully!"
  echo "===================================="
  exit 0
else
  echo "===================================="
  echo "Tests failed!"
  echo "===================================="
  exit 1
fi 