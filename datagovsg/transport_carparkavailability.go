package datagovsg

import (
	"encoding/json"
	"net/url"
)

// CarparkAvailability is the resource representing the Carpark Availability.
type CarparkAvailability struct {
	APIInfo APIInfo                   `json:"api_info"`
	Items   []CarparkAvailabilityItem `json:"items"`
}

// CarparkAvailabilityItem represents a set of carpark data and their
// corresponding lot availability.
type CarparkAvailabilityItem struct {
	// Time of acquisition of data
	Timestamp string `json:"timestamp"`

	// Carpark availability information
	CarparkData []CarparkAvailabilityCarpark `json:"carpark_data"`
}

// CarparkAvailabilityCarpark represents a single carpark.
type CarparkAvailabilityCarpark struct {
	// Identifier string of the carpark
	CarparkNumber string `json:"carpark_number"`

	// Timestamp of last update
	UpdateDateTime string `json:"update_datetime"`

	// Availability information of carpark
	CarparkInfo []CarparkAvailabilityCarparkInfo `json:"carpark_info"`
}

// CarparkAvailabilityCarparkInfo represents the availability information
// of a carpark at a point in time.
type CarparkAvailabilityCarparkInfo struct {
	// Total number of lots
	TotalLots string `json:"total_lots"`

	// Type of the carpark lot
	LotType string `json:"lot_type"`

	// Number of lots available
	LotsAvailable string `json:"lots_available"`
}

// GetCarparkAvailability returns the lot availability across all carparks
// in Singapore.
func (c *Client) GetCarparkAvailability() (*CarparkAvailability, error) {
	// Parse URL
	path := "/v1/transport/carpark-availability/"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}

	// Execute request
	b, err := c.Get(u)
	if err != nil {
		return nil, err
	}

	// Handle response
	data := &CarparkAvailability{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
