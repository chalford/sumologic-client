package sumologic

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	defaultBaseURL = "https://api.eu.sumologic.com"
	defaultAPIPath = "/api/v1"
)

// Sumologic struct for connecting to Sumo Logic services
type Sumologic struct {
	BaseURL   string
	APIPath   string
	AccessID  string
	AccessKey string
	Client    *http.Client
}

// NewSumologic creates a new Sumologic client. You may pass nil to baseURL and apiPath
// to get default values
func NewSumologic(accessID, accessKey, baseURL, apiPath string) *Sumologic {
	client := &http.Client{}

	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	if apiPath == "" {
		apiPath = defaultAPIPath
	}

	return &Sumologic{
		BaseURL:   baseURL,
		APIPath:   apiPath,
		AccessID:  accessID,
		AccessKey: accessKey,
		Client:    client,
	}
}

// NewDefaultSumologic returns a new Sumologic client, using credentials set in environment variables
func NewDefaultSumologic() (*Sumologic, error) {
	envAccessID := os.Getenv("SUMOLOGIC_ACCESS_ID")
	envAccessKey := os.Getenv("SUMOLOGIC_ACCESS_KEY")

	if envAccessID == "" || envAccessKey == "" {
		return nil, errors.New("SUMOLOGIC_ACCESS_ID or SUMOLOGIC_ACCESS_KEY environment variables not set")
	}

	client := NewSumologic(envAccessID, envAccessKey, "", "")

	return client, nil
}

// ResourceURL returns a URL with parameters substituted
func (s *Sumologic) ResourceURL(url string, params map[string]string) string {
	if params != nil {
		for key, val := range params {
			url = strings.Replace(url, key, val, -1)
		}
	}

	url = s.BaseURL + s.APIPath + url

	return url
}

func (s *Sumologic) execRequest(method, url string, body []byte) ([]byte, error) {
	var req *http.Request
	var err error

	if body != nil {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(method, url, reader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		panic("Error while building Sumologic request")
	}

	// Sumo Logic's API authentication uses basic auth...
	req.SetBasicAuth(s.AccessID, s.AccessKey)

	// Sumo Logic's API endpoints accept JSON when PUTing and POSTing
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client.Do error: %q", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("*Sumologic.execRequest() failed: <%d> %s", resp.StatusCode, req.URL)
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return contents, err
}
