package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloadIndexToFile(t *testing.T) {
	err := DownloadIndexToFile(context.Background(), "data/index.json")
	assert.Nil(t, err)
}
