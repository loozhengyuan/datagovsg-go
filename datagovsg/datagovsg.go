package datagovsg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// The base URL for Data.gov.sg API.
	baseURL = "https://api.data.gov.sg"
)

var (
	// ErrResponseNotOk is returned by Client.Get calls when the
	// the API returns a response with a non-200 HTTP status code.
	ErrResponseNotOk = errors.New("datagovsg: response not ok")

	// ErrParseErrorMessageFailure is returned by Client.Get calls
	// when the API call is not successful but there the error
	// message could not be successfully parsed.
	ErrParseErrorMessageFailure = errors.New("datagovsg: error parsing error message")
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
func (c *Client) Get(u *url.URL) ([]byte, error) {
	// Create request
	req, err := http.NewRequest("GET", u.String(), nil)
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

	// Handle non-success HTTP responses
	if resp.StatusCode != http.StatusOK {
		var e ErrorResponse
		if err := json.Unmarshal(body, &e); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrParseErrorMessageFailure, string(body))
		}
		return nil, fmt.Errorf("%w: %v", ErrResponseNotOk, e.Message)
	}

	return body, nil
}
