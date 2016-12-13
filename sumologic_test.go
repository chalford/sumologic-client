package sumologic

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestNewDefaultSumologic(t *testing.T) {
	oldAccessID := os.Getenv("SUMOLOGIC_ACCESS_ID")
	oldAccessKey := os.Getenv("SUMOLOGIC_ACCESS_KEY")

	os.Setenv("SUMOLOGIC_ACCESS_ID", "a")
	os.Setenv("SUMOLOGIC_ACCESS_KEY", "b")

	client, err := NewDefaultSumologic()

	if err != nil || client == nil {
		t.Fatalf("Sumlogic client not created.")
	}

	assert.Equal(t, "a", client.AccessID)
	assert.Equal(t, "b", client.AccessKey)

	os.Setenv("SUMOLOGIC_ACCESS_ID", oldAccessID)
	os.Setenv("SUMOLOGIC_ACCESS_KEY", oldAccessKey)
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
