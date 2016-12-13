package sumologic

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	integration    = flag.Bool("integration", false, "run integration tests")
	liveSumoClient *Sumologic
)

func TestMain(m *testing.M) {
	flag.Parse()
	var err error
	if *integration {
		fmt.Println("Running integration tests")
		liveSumoClient, err = NewDefaultSumologic()
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result := m.Run()
	os.Exit(result)
}

func TestCollectors(t *testing.T) {
	var sumoClient *Sumologic
	if *integration {
		sumoClient = liveSumoClient
	} else {
		_, sumoClient = Stub("stubs/collectors.json")
	}

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
	var sumoClient *Sumologic
	if *integration {
		sumoClient = liveSumoClient
	} else {
		_, sumoClient = Stub("stubs/collector.json")
	}

	collector, err := sumoClient.Collector(100111448)

	if err != nil {
		t.Fatalf("Failed to retrieve collector: %s", err)
	}

	if *integration {
		assert.Equal(t, 100111448, collector.ID)
		assert.Equal(t, "Academy", collector.Name)
		assert.Equal(t, "BBC Academy service layer and main web site (COSMOS):\n\nContactEmail: academy-development-owner@lists.forge.bbc.co.uk", collector.Description)
	} else {
		assert.Equal(t, 100111111, collector.ID)
		assert.Equal(t, "Test Collector", collector.Name)
		assert.Equal(t, "A Test Collector description", collector.Description)
	}
	assert.Equal(t, true, collector.Alive)
	assert.Equal(t, "Hosted", collector.CollectorType)

}

func TestCreateDeleteCollector(t *testing.T) {
	createdCollector := createCollector(t)
	deleteCollector(t, createdCollector.ID)
}

func createCollector(t *testing.T) *Collector {
	var sumoClient *Sumologic
	if *integration {
		sumoClient = liveSumoClient
	} else {
		_, sumoClient = Stub("stubs/collector.json")
	}

	newCollector := Collector{
		CollectorType: "Hosted",
		Name:          "Test Collector",
		Description:   "A Test Collector description",
		Category:      "HTTP Collection",
	}

	createdCollector, err := sumoClient.CreateCollector(newCollector)

	if err != nil || createdCollector == nil {
		t.Fatalf("Failed to create collector: %s", err)
	}

	assert.Equal(t, newCollector.Name, createdCollector.Name)

	return createdCollector
}

func deleteCollector(t *testing.T, id int) {
	var sumoClient *Sumologic
	if *integration {
		sumoClient = liveSumoClient
	} else {
		_, sumoClient = Stub("stubs/nil-response.json")
	}

	err := sumoClient.DeleteCollector(id)

	if err != nil {
		t.Fatalf("Failed to delete collector: %s", err)
	}
}
