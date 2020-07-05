package datagovsg

import (
	"encoding/json"
	"net/url"
)

// Rainfall is the resource representing the rainfall information.
type Rainfall struct {
	APIInfo  APIInfo          `json:"api_info"`
	Metadata RainfallMetadata `json:"metadata"`
	Items    []RainfallItem   `json:"items"`
}

// RainfallMetadata represents metadata information about the rainfall data.
type RainfallMetadata struct {
	// Metadata about a weather station
	Stations []RainfallMetadataStation `json:"stations"`

	// Information about the reading
	ReadingType string `json:"reading_type"`

	// Measurement unit for the reading
	ReadingUnit string `json:"reading_unit"`
}

// RainfallMetadataStation represents metadata information specific to a
// weather station.
type RainfallMetadataStation struct {
	// ID of a station
	ID string `json:"id"`

	// ID of the device (usually the same as the station)
	DeviceID string `json:"device_id"`

	// Name of the station
	Name string `json:"name"`

	// Location of the station
	Location RainfallMetadataStationLocation `json:"location"`
}

// RainfallMetadataStationLocation represents the geographical coordindates
// of a weather station.
type RainfallMetadataStationLocation struct {
	// Longitude of the station
	Longitude float64 `json:"longitude"`

	// Latitude of the station
	Latitude float64 `json:"latitude"`
}

// RainfallItem represents all rainfall readings at a point in time.
type RainfallItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Data readings
	Readings []RainfallItemReading `json:"readings"`
}

// RainfallItemReading represents a single rainfall reading at
// a specific station at a point in time.
type RainfallItemReading struct {
	// ID of the station
	StationID string `json:"station_id"`

	// Value of the reading
	Value float64 `json:"value"`
}

// GetRainfall returns the rainfall information.
func (c *Client) GetRainfall(options ...*QueryOption) (*Rainfall, error) {
	// Parse URL
	path := "/v1/environment/rainfall/"
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
	data := &Rainfall{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
