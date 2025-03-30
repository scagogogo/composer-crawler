# Composer Crawler - Unit Testing Documentation

This document describes the testing approach for the Composer Crawler project.

## Overview

The Composer Crawler project has a comprehensive suite of unit tests to ensure code quality and reliability. The tests are organized by package and functionality, covering core repository operations, API interactions, and response models.

## Test Coverage

Current test coverage: **58.5%** of all statements.

Note: Some methods show 0% coverage because they involve actual HTTP calls to external services, and their tests are skipped by default to avoid external dependencies during automated testing.

## Running Tests

You can run the tests using the provided `run_tests.sh` script:

```bash
./run_tests.sh
```

This script will:
1. Run all tests in the project
2. Generate a code coverage report
3. Create an HTML coverage report in the `coverage/` directory

## Manual Testing

Some tests that require actual HTTP calls to external services are skipped by default. You can run these tests manually by:

1. Uncomment the test code in the relevant test files (look for the `// Skip actual HTTP calls in unit tests` comments)
2. Run the specific test or test file with:
   ```bash
   go test -v ./pkg/repository/index_api_test.go
   ```

## Test Structure

The tests are organized as follows:

### Core Repository Tests
- `repository_test.go`: Tests for core functions like JSON handling and HTTP requests
- `repository_test_main.go`: Test utilities and setup

### API-specific Tests
- `get_statistics_test.go`: Tests for statistics retrieval
- `list_security_advisories_test.go`: Tests for security advisories endpoints
- `listing_package_names_test.go`: Tests for package listing
- `index_api_test.go`: Tests for index download functionality (mostly skipped)

### Response Model Tests
- `statistics_test.go`: Tests for statistics response structures
- `advisory_test.go`: Tests for security advisory response structures

## Test Design Principles

1. **Independence**: Tests can run independently and don't rely on external services
2. **Comprehensiveness**: Multiple test cases cover happy paths, edge cases, and error scenarios
3. **Mocking**: HTTP servers are mocked to avoid actual API calls
4. **Clarity**: Tests are structured with clear setup, execution, and verification phases

## Adding New Tests

When adding new functionality, please follow these guidelines:
1. Create tests for all new methods
2. Use the existing patterns for HTTP mocking when testing API calls
3. Include test cases for success scenarios, error handling, and edge cases
4. Skip tests that require external services, but provide commented code for manual testing

## Troubleshooting

If tests are failing:
1. Check if there are connectivity issues if you've enabled the skipped tests
2. Verify that the mock responses match the expected structure
3. Ensure dependencies are correctly installed 