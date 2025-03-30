package repository

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_List(t *testing.T) {
	// Create a test server with a mock package list response
	mockResponse := `{
		"packageNames": [
			"vendor1/package1",
			"vendor2/package2",
			"vendor3/package3"
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/packages/list.json" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockResponse))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Create repository with test server URL
	repo := &Repository{
		options: &Options{
			ServerUrl: server.URL,
		},
	}

	// Test successful request
	t.Run("successful request", func(t *testing.T) {
		packages, err := repo.List(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, packages)
		assert.Len(t, packages, 3)

		// Verify package names
		assert.Equal(t, "vendor1/package1", packages[0].Name)
		assert.Equal(t, "vendor2/package2", packages[1].Name)
		assert.Equal(t, "vendor3/package3", packages[2].Name)
	})

	// Test empty response
	t.Run("empty package list", func(t *testing.T) {
		emptyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer emptyServer.Close()

		emptyRepo := &Repository{
			options: &Options{
				ServerUrl: emptyServer.URL,
			},
		}

		packages, err := emptyRepo.List(context.Background())
		assert.NoError(t, err)
		assert.Empty(t, packages)
	})

	// Test invalid server URL
	t.Run("invalid server URL", func(t *testing.T) {
		invalidRepo := &Repository{
			options: &Options{
				ServerUrl: "http://invalid-url-that-doesnt-exist.example",
			},
		}

		packages, err := invalidRepo.List(context.Background())
		assert.Error(t, err)
		assert.Nil(t, packages)
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

		packages, err := malformedRepo.List(context.Background())
		assert.Error(t, err)
		assert.Nil(t, packages)
	})
}
