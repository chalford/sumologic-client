package sumologic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestNewSumologic(t *testing.T) {
	client := NewSumologic("a", "b", "c", "d")

	if client == nil {
		t.Fatalf("Sumlogic client not created.")
	}

	assert.Equal(t, "a", client.AccessID)
	assert.Equal(t, "b", client.AccessKey)
	assert.Equal(t, "c", client.BaseURL)
	assert.Equal(t, "d", client.APIPath)
}
