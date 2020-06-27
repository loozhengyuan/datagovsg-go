package datagovsg

import (
	"encoding/json"
	"net/url"
)

// PM25 is the resource representing the PM2.5 information.
type PM25 struct {
	APIInfo APIInfo    `json:"api_info"`
	Items   []PM25Item `json:"items"`
}

// PM25RegionMetadata represents metadata about each region.
type PM25RegionMetadata struct {
	// Name of the region
	Name string `json:"name"`

	// Geographical coordinates of the region
	LabelLocation PM25RegionMetadataLabelLocation `json:"label_location"`
}

// PM25RegionMetadataLabelLocation represents the geographical
// coordinates of the region.
type PM25RegionMetadataLabelLocation struct {
	// Longitude of the region
	Longitude float64 `json:"longitude"`

	// Latitude of the region
	Latitude float64 `json:"latitude"`
}

// PM25Item represents a PM2.5 reading at a point in time.
type PM25Item struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Timestamp of data acquisition
	UpdateTimestamp string `json:"update_timestamp"`

	// Data readings
	Readings PM25ItemReadings `json:"readings"`
}

// PM25ItemReadings represents all PM2.5 data readings.
type PM25ItemReadings struct {
	// One-hourly reading of the PM2.5 concentration.
	PM25OneHourly map[string]int `json:"pm25_one_hourly"`
}

// GetPM25 returns the PM2.5 information.
func (c *Client) GetPM25(options ...*QueryOption) (*PM25, error) {
	// Parse URL
	path := "/v1/environment/pm25/"
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
	data := &PM25{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
