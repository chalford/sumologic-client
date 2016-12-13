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
	if *integration {
		fmt.Println("Running integration tests")
	}
	var err error
	liveSumoClient, err = NewDefaultSumologic()
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

	assert.Equal(t, 100111448, collector.ID)
	assert.Equal(t, "Academy", collector.Name)
	assert.Equal(t, "BBC Academy service layer and main web site (COSMOS):\n\nContactEmail: academy-development-owner@lists.forge.bbc.co.uk", collector.Description)
	assert.Equal(t, true, collector.Alive)
	assert.Equal(t, "Hosted", collector.CollectorType)
}

func TestCreateCollector(t *testing.T) {
	var sumoClient *Sumologic
	if *integration {
		sumoClient = liveSumoClient
	} else {
		_, sumoClient = Stub("stubs/nil-response.json")
	}

	newCollector := Collector{
		ID: 1234,
	}

	err := sumoClient.CreateCollector(newCollector)

	if err != nil {
		t.Fatalf("Failed to create collector: %s", err)
	}
}

func TestDeleteCollector(t *testing.T) {
	_, sumoClient := Stub("stubs/nil-response.json")

	err := sumoClient.DeleteCollector(1234)

	if err != nil {
		t.Fatalf("Failed to delete collector: %s", err)
	}
}
