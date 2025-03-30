package repository

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJson(t *testing.T) {
	// Test cases
	tests := []struct {
		name    string
		json    []byte
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "valid json",
			json:    []byte(`{"key": "value"}`),
			want:    map[string]interface{}{"key": "value"},
			wantErr: false,
		},
		{
			name:    "invalid json",
			json:    []byte(`{invalid}`),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty json",
			json:    []byte(`{}`),
			want:    map[string]interface{}{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshalJson[map[string]interface{}](tt.json)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestRepository_GetBytes(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "data"}`))
	}))
	defer server.Close()

	// Test cases
	tests := []struct {
		name       string
		url        string
		setupRepo  func() *Repository
		wantErr    bool
		wantResult []byte
		skip       bool
	}{
		{
			name: "successful request",
			url:  server.URL,
			setupRepo: func() *Repository {
				return &Repository{
					options: &Options{
						ServerUrl: "https://packagist.org",
					},
				}
			},
			wantErr:    false,
			wantResult: []byte(`{"test": "data"}`),
			skip:       false,
		},
		{
			name: "request with proxy",
			url:  server.URL,
			setupRepo: func() *Repository {
				return &Repository{
					options: &Options{
						ServerUrl: "https://packagist.org",
						Proxy:     "http://proxy.example.com", // This will cause an error since it's not a real proxy
					},
				}
			},
			wantErr: true, // Expect an error with an invalid proxy
			skip:    true, // Skip this test in regular runs as it depends on external proxy setup
		},
		{
			name: "invalid url",
			url:  "http://invalid-url-that-doesnt-exist.example",
			setupRepo: func() *Repository {
				return &Repository{
					options: &Options{
						ServerUrl: "https://packagist.org",
					},
				}
			},
			wantErr: true,
			skip:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip tests marked for skipping
			if tt.skip {
				t.Skip("Skipping test that requires external proxy configuration")
			}

			repo := tt.setupRepo()
			result, err := repo.getBytes(context.Background(), tt.url)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.wantResult, result)
				}
			}
		})
	}
}

func TestGetJson(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"key": "value"}`))
	}))
	defer server.Close()

	// Setup repository
	repo := &Repository{
		options: &Options{
			ServerUrl: "https://packagist.org",
		},
	}

	// Test cases
	tests := []struct {
		name    string
		url     string
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "valid response",
			url:     server.URL,
			want:    map[string]interface{}{"key": "value"},
			wantErr: false,
		},
		{
			name:    "invalid url",
			url:     "http://invalid-url-that-doesnt-exist.example",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getJson[map[string]interface{}](context.Background(), repo, tt.url)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.want, got)
				}
			}
		})
	}
}
