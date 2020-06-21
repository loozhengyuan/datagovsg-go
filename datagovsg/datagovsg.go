package datagovsg

import (
	"io/ioutil"
	"net/http"
)

const (
	// The base URL for Data.gov.sg API.
	baseURL = "https://api.data.gov.sg/v1"
)

// Client is a simple http.Client wrapper.
// TODO: Use http.DefaultClient?
type Client struct {
	Client  *http.Client
	BaseURL string
}

// NewClient returns a new Client object.
func NewClient() *Client {
	return &Client{
		Client:  http.DefaultClient,
		BaseURL: baseURL,
	}
}

// Get executes a HTTP GET request.
func (c *Client) Get(url string) ([]byte, error) {
	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Execute request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// APIInfo contains information about the API (from Data.gov.sg)
type APIInfo struct {
	Status string `json:"status"`
}
