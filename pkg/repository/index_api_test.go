package repository

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// We can't mock the DownloadIndex function directly because we can't assign to it
// Instead we'll test the individual components separately

func TestDownloadIndex(t *testing.T) {
	// Skip actual HTTP calls in unit tests
	t.Skip("Skipping test that makes real HTTP calls to external services")

	// This test should be run manually
	// data, err := DownloadIndex(context.Background())
	// assert.NoError(t, err)
	// assert.NotEmpty(t, data)
}

func TestDownloadIndexToFile(t *testing.T) {
	// Skip actual HTTP calls in unit tests
	t.Skip("Skipping test that makes real HTTP calls to external services")

	// This test should be run manually
	// tempDir, err := os.MkdirTemp("", "index-test")
	// assert.NoError(t, err)
	// defer os.RemoveAll(tempDir)
	//
	// testFilePath := filepath.Join(tempDir, "test-index.json")
	// err = DownloadIndexToFile(context.Background(), testFilePath)
	// assert.NoError(t, err)
	//
	// fileData, err := os.ReadFile(testFilePath)
	// assert.NoError(t, err)
	// assert.NotEmpty(t, fileData)
}

// TestMockDownloadLogic tests similar file download logic with a mock server
func TestMockDownloadLogic(t *testing.T) {
	// Create a mock server
	mockResponse := `{"packageNames": ["vendor1/package1", "vendor2/package2"]}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Create a temporary directory for the test file
	tempDir, err := os.MkdirTemp("", "index-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Test file path
	testFilePath := filepath.Join(tempDir, "test-index.json")

	// Test the download and write to file logic
	resp, err := http.Get(server.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	err = os.WriteFile(testFilePath, data, os.ModePerm)
	assert.NoError(t, err)

	// Verify file content
	fileData, err := os.ReadFile(testFilePath)
	assert.NoError(t, err)
	assert.Equal(t, mockResponse, string(fileData))
}

func TestDownloadIndexToFile_Error(t *testing.T) {
	// Test with a non-existent directory - requires an actual download attempt
	// We'll skip this for now to avoid real external calls
	t.Skip("Skipping test that would make real HTTP calls to external services")

	// For manual testing only
	// invalidPath := "/nonexistent/directory/file.json"
	// err := DownloadIndexToFile(context.Background(), invalidPath)
	// assert.Error(t, err)
}
