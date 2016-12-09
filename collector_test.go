package sumologic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectors(t *testing.T) {
	_, sumoClient := Stub("stubs/collectors.json")

	collectors, err := sumoClient.Collectors(0, 10)

	if err != nil {
		t.Fatalf("Failed to retrieve collectors: %s", err)
	}

	assert.Equal(t, 10, len(collectors))
	for _, v := range collectors {
		assert.NotNil(t, v.ID)
	}

}

func TestCollector(t *testing.T) {
	_, sumoClient := Stub("stubs/collector.json")

	collector, err := sumoClient.Collector(1234)

	if err != nil {
		t.Fatalf("Failed to retrieve collector: %s", err)
	}

	assert.Equal(t, 100111448, collector.ID)
	assert.Equal(t, "Academy", collector.Name)
	assert.Equal(t, "BBC Academy", collector.Description)
	assert.Equal(t, true, collector.Alive)
	assert.Equal(t, "Hosted", collector.CollectorType)
}

func TestDeleteCollector(t *testing.T) {
	_, sumoClient := Stub("stubs/nil-response.json")

	err := sumoClient.DeleteCollector(1234)

	if err != nil {
		t.Fatalf("Failed to delete collector: %s", err)
	}
}

func TestCreateCollector(t *testing.T) {
	_, sumoClient := Stub("stubs/nil-response.json")

	newCollector := Collector{
		ID: 1234,
	}

	err := sumoClient.CreateCollector(newCollector)

	if err != nil {
		t.Fatalf("Failed to create collector: %s", err)
	}
}
