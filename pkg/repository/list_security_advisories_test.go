package repository

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepository_ListSecurityAdvisories(t *testing.T) {
	// Sample response for testing
	mockAdvisoriesResponse := `{
		"advisories": {
			"vendor/package1": [
				{
					"advisoryId": "GHSA-1234-5678-9012",
					"packageName": "vendor/package1",
					"remoteId": "GHSA-1234-5678-9012",
					"title": "Critical vulnerability in package1",
					"link": "https://example.com/advisory/123",
					"cve": "CVE-2023-1234",
					"affectedVersions": "<2.0.0",
					"source": "GitHub",
					"reportedAt": "2023-01-15",
					"composerRepository": "https://packagist.org",
					"sources": [
						{
							"name": "GitHub",
							"remoteId": "GHSA-1234-5678-9012"
						}
					]
				}
			]
		}
	}`

	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/security-advisories/" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockAdvisoriesResponse))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Test repository
	repo := &Repository{
		options: &Options{
			ServerUrl: server.URL,
		},
	}

	// Test successful request
	t.Run("successful request", func(t *testing.T) {
		// Use a fixed timestamp for testing
		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

		result, err := repo.ListSecurityAdvisories(context.Background(), testTime)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// Verify the response content
		assert.Contains(t, result.Advisories, "vendor/package1")
		advisories := result.Advisories["vendor/package1"]
		assert.Len(t, advisories, 1)

		advisory := advisories[0]
		assert.Equal(t, "GHSA-1234-5678-9012", advisory.AdvisoryID)
		assert.Equal(t, "vendor/package1", advisory.PackageName)
		assert.Equal(t, "Critical vulnerability in package1", advisory.Title)
		assert.Equal(t, "CVE-2023-1234", advisory.Cve)
	})

	// Test invalid server URL
	t.Run("invalid server URL", func(t *testing.T) {
		invalidRepo := &Repository{
			options: &Options{
				ServerUrl: "http://invalid-url-that-doesnt-exist.example",
			},
		}

		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		result, err := invalidRepo.ListSecurityAdvisories(context.Background(), testTime)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	// Test malformed response
	t.Run("malformed response", func(t *testing.T) {
		malformedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{malformed json}`))
		}))
		defer malformedServer.Close()

		malformedRepo := &Repository{
			options: &Options{
				ServerUrl: malformedServer.URL,
			},
		}

		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		result, err := malformedRepo.ListSecurityAdvisories(context.Background(), testTime)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestRepository_ListAdvisories(t *testing.T) {
	// Sample package name for testing
	testPackage := "vendor/package1"

	// Sample response for testing
	mockAdvisoriesResponse := `{
		"advisories": {
			"vendor/package1": [
				{
					"advisoryId": "GHSA-1234-5678-9012",
					"packageName": "vendor/package1",
					"remoteId": "GHSA-1234-5678-9012",
					"title": "Critical vulnerability in package1",
					"link": "https://example.com/advisory/123",
					"cve": "CVE-2023-1234",
					"affectedVersions": "<2.0.0",
					"source": "GitHub",
					"reportedAt": "2023-01-15",
					"composerRepository": "https://packagist.org",
					"sources": [
						{
							"name": "GitHub",
							"remoteId": "GHSA-1234-5678-9012"
						}
					]
				},
				{
					"advisoryId": "GHSA-5678-9012-3456",
					"packageName": "vendor/package1",
					"remoteId": "GHSA-5678-9012-3456",
					"title": "Another vulnerability in package1",
					"link": "https://example.com/advisory/456",
					"cve": "CVE-2023-5678",
					"affectedVersions": "<1.5.0",
					"source": "GitHub",
					"reportedAt": "2023-02-20",
					"composerRepository": "https://packagist.org",
					"sources": [
						{
							"name": "GitHub",
							"remoteId": "GHSA-5678-9012-3456"
						}
					]
				}
			]
		}
	}`

	// Create test server with conditional response based on request URL
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := fmt.Sprintf("/api/security-advisories/")
		if r.URL.Path == expectedPath {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockAdvisoriesResponse))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Test repository
	repo := &Repository{
		options: &Options{
			ServerUrl: server.URL,
		},
	}

	// Test successful request
	t.Run("successful request", func(t *testing.T) {
		advisories, err := repo.ListAdvisories(context.Background(), testPackage)
		assert.NoError(t, err)
		assert.NotNil(t, advisories)
		assert.Len(t, advisories, 2)

		// Verify first advisory
		assert.Equal(t, "GHSA-1234-5678-9012", advisories[0].AdvisoryID)
		assert.Equal(t, "vendor/package1", advisories[0].PackageName)
		assert.Equal(t, "Critical vulnerability in package1", advisories[0].Title)
		assert.Equal(t, "CVE-2023-1234", advisories[0].Cve)

		// Verify second advisory
		assert.Equal(t, "GHSA-5678-9012-3456", advisories[1].AdvisoryID)
		assert.Equal(t, "vendor/package1", advisories[1].PackageName)
		assert.Equal(t, "Another vulnerability in package1", advisories[1].Title)
		assert.Equal(t, "CVE-2023-5678", advisories[1].Cve)
	})

	// Test package with no advisories
	t.Run("package with no advisories", func(t *testing.T) {
		emptyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer emptyServer.Close()

		emptyRepo := &Repository{
			options: &Options{
				ServerUrl: emptyServer.URL,
			},
		}

		advisories, err := emptyRepo.ListAdvisories(context.Background(), "no-advisories/package")
		assert.NoError(t, err)
		assert.Nil(t, advisories) // Should be nil since the package doesn't exist in response
	})

	// Test invalid server URL
	t.Run("invalid server URL", func(t *testing.T) {
		invalidRepo := &Repository{
			options: &Options{
				ServerUrl: "http://invalid-url-that-doesnt-exist.example",
			},
		}

		advisories, err := invalidRepo.ListAdvisories(context.Background(), testPackage)
		assert.Error(t, err)
		assert.Nil(t, advisories)
	})

	// Test malformed response
	t.Run("malformed response", func(t *testing.T) {
		malformedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{malformed json}`))
		}))
		defer malformedServer.Close()

		malformedRepo := &Repository{
			options: &Options{
				ServerUrl: malformedServer.URL,
			},
		}

		advisories, err := malformedRepo.ListAdvisories(context.Background(), testPackage)
		assert.Error(t, err)
		assert.Nil(t, advisories)
	})
}
