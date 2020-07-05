package datagovsg

import (
	"encoding/json"
	"net/url"
)

// UVIndex is the resource representing the UVIndex information.
type UVIndex struct {
	APIInfo APIInfo       `json:"api_info"`
	Items   []UVIndexItem `json:"items"`
}

// UVIndexItem represents a collection of all UV Index readings within
// the day, up until a specific point in time.
type UVIndexItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Timestamp of data acquisition
	UpdateTimestamp string `json:"update_timestamp"`

	// List of all readings within the day
	Index []UVIndexItemReading `json:"index"`
}

// UVIndexItemReading represents a single Ultra-Violet Index reading at a
// specific point in time.
type UVIndexItemReading struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Value of the index
	Value int `json:"value"`
}

// GetUVIndex returns the UVIndex information.
func (c *Client) GetUVIndex(options ...*QueryOption) (*UVIndex, error) {
	// Parse URL
	path := "/v1/environment/uv-index/"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}

	// Set query parameters
	v := url.Values{}
	for _, option := range options {
		v.Add(option.Key, option.Value)
	}
	u.RawQuery = v.Encode()

	// Execute request
	b, err := c.Get(u)
	if err != nil {
		return nil, err
	}

	// Handle response
	data := &UVIndex{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
