package datagovsg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// The base URL for Data.gov.sg API.
	baseURL = "https://api.data.gov.sg/v1"
)

var (
	// ErrBadRequest is raised when there are errors in request.
	ErrBadRequest = errors.New("bad request")
)

// APIInfo contains information about the API (from Data.gov.sg)
type APIInfo struct {
	Status string `json:"status"`
}

// ErrorResponse represents the error response returned
// by the API.
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

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

	// Check error
	switch resp.StatusCode {
	case http.StatusBadRequest:
		var e ErrorResponse
		json.Unmarshal(body, &e)
		return nil, fmt.Errorf("%w: %v", ErrBadRequest, e.Message)
	}

	return body, nil
}
