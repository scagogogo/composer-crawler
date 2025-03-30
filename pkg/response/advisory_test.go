package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdvisoriesResponse_UnmarshalJSON(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		jsonData string
		want     *AdvisoriesResponse
		wantErr  bool
	}{
		{
			name: "valid advisories response",
			jsonData: `{
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
			}`,
			want: &AdvisoriesResponse{
				Advisories: map[string][]*Advisory{
					"vendor/package1": {
						{
							AdvisoryID:         "GHSA-1234-5678-9012",
							PackageName:        "vendor/package1",
							RemoteID:           "GHSA-1234-5678-9012",
							Title:              "Critical vulnerability in package1",
							Link:               "https://example.com/advisory/123",
							Cve:                "CVE-2023-1234",
							AffectedVersions:   "<2.0.0",
							Source:             "GitHub",
							ReportedAt:         "2023-01-15",
							ComposerRepository: "https://packagist.org",
							Sources: []*Sources{
								{
									Name:     "GitHub",
									RemoteID: "GHSA-1234-5678-9012",
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:     "empty advisories response",
			jsonData: `{"advisories": {}}`,
			want: &AdvisoriesResponse{
				Advisories: map[string][]*Advisory{},
			},
			wantErr: false,
		},
		{
			name:     "completely empty response",
			jsonData: `{}`,
			want:     &AdvisoriesResponse{},
			wantErr:  false,
		},
		{
			name:     "invalid json",
			jsonData: `{invalid json}`,
			want:     nil,
			wantErr:  true,
		},
		{
			name: "multiple packages with advisories",
			jsonData: `{
				"advisories": {
					"vendor/package1": [
						{
							"advisoryId": "GHSA-1234-5678-9012",
							"packageName": "vendor/package1",
							"title": "Critical vulnerability in package1",
							"cve": "CVE-2023-1234"
						}
					],
					"vendor/package2": [
						{
							"advisoryId": "GHSA-5678-9012-3456",
							"packageName": "vendor/package2",
							"title": "Critical vulnerability in package2",
							"cve": "CVE-2023-5678"
						}
					]
				}
			}`,
			want: &AdvisoriesResponse{
				Advisories: map[string][]*Advisory{
					"vendor/package1": {
						{
							AdvisoryID:  "GHSA-1234-5678-9012",
							PackageName: "vendor/package1",
							Title:       "Critical vulnerability in package1",
							Cve:         "CVE-2023-1234",
						},
					},
					"vendor/package2": {
						{
							AdvisoryID:  "GHSA-5678-9012-3456",
							PackageName: "vendor/package2",
							Title:       "Critical vulnerability in package2",
							Cve:         "CVE-2023-5678",
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got *AdvisoriesResponse
			err := json.Unmarshal([]byte(tt.jsonData), &got)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestAdvisory_UnmarshalJSON(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		jsonData string
		want     *Advisory
		wantErr  bool
	}{
		{
			name: "full advisory data",
			jsonData: `{
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
			}`,
			want: &Advisory{
				AdvisoryID:         "GHSA-1234-5678-9012",
				PackageName:        "vendor/package1",
				RemoteID:           "GHSA-1234-5678-9012",
				Title:              "Critical vulnerability in package1",
				Link:               "https://example.com/advisory/123",
				Cve:                "CVE-2023-1234",
				AffectedVersions:   "<2.0.0",
				Source:             "GitHub",
				ReportedAt:         "2023-01-15",
				ComposerRepository: "https://packagist.org",
				Sources: []*Sources{
					{
						Name:     "GitHub",
						RemoteID: "GHSA-1234-5678-9012",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "minimal advisory data",
			jsonData: `{
				"advisoryId": "GHSA-1234-5678-9012",
				"packageName": "vendor/package1",
				"title": "Critical vulnerability in package1"
			}`,
			want: &Advisory{
				AdvisoryID:  "GHSA-1234-5678-9012",
				PackageName: "vendor/package1",
				Title:       "Critical vulnerability in package1",
			},
			wantErr: false,
		},
		{
			name:     "empty advisory",
			jsonData: `{}`,
			want:     &Advisory{},
			wantErr:  false,
		},
		{
			name:     "invalid json",
			jsonData: `{invalid json}`,
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got *Advisory
			err := json.Unmarshal([]byte(tt.jsonData), &got)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestSources_UnmarshalJSON(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		jsonData string
		want     *Sources
		wantErr  bool
	}{
		{
			name: "complete sources data",
			jsonData: `{
				"name": "GitHub",
				"remoteId": "GHSA-1234-5678-9012"
			}`,
			want: &Sources{
				Name:     "GitHub",
				RemoteID: "GHSA-1234-5678-9012",
			},
			wantErr: false,
		},
		{
			name: "partial sources data - only name",
			jsonData: `{
				"name": "GitHub"
			}`,
			want: &Sources{
				Name: "GitHub",
			},
			wantErr: false,
		},
		{
			name: "partial sources data - only remoteId",
			jsonData: `{
				"remoteId": "GHSA-1234-5678-9012"
			}`,
			want: &Sources{
				RemoteID: "GHSA-1234-5678-9012",
			},
			wantErr: false,
		},
		{
			name:     "empty sources",
			jsonData: `{}`,
			want:     &Sources{},
			wantErr:  false,
		},
		{
			name:     "invalid json",
			jsonData: `{invalid json}`,
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got *Sources
			err := json.Unmarshal([]byte(tt.jsonData), &got)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
