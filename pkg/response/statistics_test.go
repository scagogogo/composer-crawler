package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatisticsResponse_UnmarshalJSON(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		jsonData string
		want     *StatisticsResponse
		wantErr  bool
	}{
		{
			name:     "valid statistics response",
			jsonData: `{"totals": {"downloads": 1000000, "packages": 500, "versions": 2000}}`,
			want: &StatisticsResponse{
				Totals: Totals{
					Downloads: 1000000,
					Packages:  500,
					Versions:  2000,
				},
			},
			wantErr: false,
		},
		{
			name:     "empty statistics response",
			jsonData: `{}`,
			want:     &StatisticsResponse{},
			wantErr:  false,
		},
		{
			name:     "partial statistics response",
			jsonData: `{"totals": {"downloads": 1000000}}`,
			want: &StatisticsResponse{
				Totals: Totals{
					Downloads: 1000000,
					Packages:  0,
					Versions:  0,
				},
			},
			wantErr: false,
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
			var got *StatisticsResponse
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

func TestTotals_UnmarshalJSON(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		jsonData string
		want     *Totals
		wantErr  bool
	}{
		{
			name:     "full totals data",
			jsonData: `{"downloads": 1000000, "packages": 500, "versions": 2000}`,
			want: &Totals{
				Downloads: 1000000,
				Packages:  500,
				Versions:  2000,
			},
			wantErr: false,
		},
		{
			name:     "partial totals data",
			jsonData: `{"downloads": 1000000}`,
			want: &Totals{
				Downloads: 1000000,
				Packages:  0,
				Versions:  0,
			},
			wantErr: false,
		},
		{
			name:     "zero values",
			jsonData: `{"downloads": 0, "packages": 0, "versions": 0}`,
			want: &Totals{
				Downloads: 0,
				Packages:  0,
				Versions:  0,
			},
			wantErr: false,
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
			var got *Totals
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
