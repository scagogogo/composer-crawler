package repository

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestMain is the entry point for running all tests in the repository package
func TestMain(m *testing.M) {
	// Setup code before running tests
	fmt.Println("Starting repository package tests...")

	// Run all tests
	exitCode := m.Run()

	// Cleanup code after running tests
	fmt.Println("Completed repository package tests")

	// Exit with the test status code
	os.Exit(exitCode)
}

// createMockServer is a helper function to create a test server that returns a fixed response
func createMockServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
}

// createMockServerWithRoutes is a helper function to create a test server with route-specific responses
func createMockServerWithRoutes(routes map[string]string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if response, ok := routes[r.URL.Path]; ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
}

// newTestRepository creates a new Repository instance for testing
func newTestRepository(serverURL string) *Repository {
	return &Repository{
		options: &Options{
			ServerUrl: serverURL,
		},
	}
}

// newTestRepositoryWithProxy creates a new Repository instance with proxy for testing
func newTestRepositoryWithProxy(serverURL, proxy string) *Repository {
	return &Repository{
		options: &Options{
			ServerUrl: serverURL,
			Proxy:     proxy,
		},
	}
}
