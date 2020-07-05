package datagovsg

import (
	"encoding/json"
	"net/url"
)

// WindDirection is the resource representing the wind direction information.
type WindDirection struct {
	APIInfo  APIInfo               `json:"api_info"`
	Metadata WindDirectionMetadata `json:"metadata"`
	Items    []WindDirectionItem   `json:"items"`
}

// WindDirectionMetadata represents metadata information about the wind direction data.
type WindDirectionMetadata struct {
	// Metadata about a weather station
	Stations []WindDirectionMetadataStation `json:"stations"`

	// Information about the reading
	ReadingType string `json:"reading_type"`

	// Measurement unit for the reading
	ReadingUnit string `json:"reading_unit"`
}

// WindDirectionMetadataStation represents metadata information specific to a
// weather station.
type WindDirectionMetadataStation struct {
	// ID of a station
	ID string `json:"id"`

	// ID of the device (usually the same as the station)
	DeviceID string `json:"device_id"`

	// Name of the station
	Name string `json:"name"`

	// Location of the station
	Location WindDirectionMetadataStationLocation `json:"location"`
}

// WindDirectionMetadataStationLocation represents the geographical coordindates
// of a weather station.
type WindDirectionMetadataStationLocation struct {
	// Longitude of the station
	Longitude float64 `json:"longitude"`

	// Latitude of the station
	Latitude float64 `json:"latitude"`
}

// WindDirectionItem represents all wind direction readings at a point in time.
type WindDirectionItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Data readings
	Readings []WindDirectionItemReading `json:"readings"`
}

// WindDirectionItemReading represents a single wind direction reading at
// a specific station at a point in time.
type WindDirectionItemReading struct {
	// ID of the station
	StationID string `json:"station_id"`

	// Value of the reading
	Value float64 `json:"value"`
}

// GetWindDirection returns the wind direction information.
func (c *Client) GetWindDirection(options ...*QueryOption) (*WindDirection, error) {
	// Parse URL
	path := "/v1/environment/wind-direction/"
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
	data := &WindDirection{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
