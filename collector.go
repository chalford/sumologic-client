package sumologic

import (
	"encoding/json"
	"strconv"
)

const (
	collectorsURL   = "/collectors?limit=:limit&offset=:offset"
	collectorURL    = "/collectors/:id"
	newCollectorURL = "/collectors/"
)

// Collector is a representation of a Sumo Logic log collector
type Collector struct {
	ID               int                 `json:"id"`
	Name             string              `json:"name,omitempty"`
	CollectorType    string              `json:"collectorType,omitempty"`
	LastSeenAlive    int                 `json:"lastSeenAlive,omitempty"`
	Alive            bool                `json:"alive,omitempty"`
	Links            []map[string]string `json:"links,omitempty"`
	CollectorVersion string              `json:"collectorVersion,omitempty"`
	Ephemeral        bool                `json:"ephemeral,omitempty"`
	Description      string              `json:"description,omitempty"`
	OSName           string              `json:"osName,omitempty"`
	OSArch           string              `json:"osArch,omitempty"`
	OSVersion        string              `json:"osVersion,omitempty"`
	Category         string              `json:"category,omitempty"`
}

// Collectors returns a list of collectors, limited by `limit`, from an offset of `offset`
func (s *Sumologic) Collectors(offset, limit int) ([]*Collector, error) {
	url := s.ResourceURL(collectorsURL, map[string]string{":limit": strconv.Itoa(limit), ":offset": strconv.Itoa(offset)})

	var collectorWrapper struct {
		Collectors []*Collector `json:"collectors"`
	}

	contents, err := s.execRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &collectorWrapper)
	}

	return collectorWrapper.Collectors, err
}

// Collector returns a single collector of the specified ID
func (s *Sumologic) Collector(id int) (*Collector, error) {
	url := s.ResourceURL(collectorURL, map[string]string{":id": strconv.Itoa(id)})

	var collectorWrapper struct {
		Collector *Collector `json:"collector"`
	}

	contents, err := s.execRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &collectorWrapper)
	}

	return collectorWrapper.Collector, err
}

// DeleteCollector deletes a collector and returns an error if the deletion was not
// successful
func (s *Sumologic) DeleteCollector(id int) error {
	url := s.ResourceURL(collectorURL, map[string]string{":id": strconv.Itoa(id)})

	_, err := s.execRequest("DELETE", url, nil)

	return err
}

// CreateCollector creates a new Sumo Logic collector
func (s *Sumologic) CreateCollector(newCollector Collector) error {
	url := s.ResourceURL(newCollectorURL, nil)

	collectorBody, err := json.Marshal(newCollector)

	if err == nil {
		_, err = s.execRequest("POST", url, collectorBody)
	}

	return err
}
