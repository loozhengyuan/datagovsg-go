package datagovsg

import (
	"encoding/json"
	"net/url"
)

// AirTemperature is the resource representing the air temperature information.
type AirTemperature struct {
	APIInfo  APIInfo                `json:"api_info"`
	Metadata AirTemperatureMetadata `json:"metadata"`
	Items    []AirTemperatureItem   `json:"items"`
}

// AirTemperatureMetadata represents metadata information about the air temperature data.
type AirTemperatureMetadata struct {
	// Metadata about a weather station
	Stations []AirTemperatureMetadataStation `json:"stations"`

	// Information about the reading
	ReadingType string `json:"reading_type"`

	// Measurement unit for the reading
	ReadingUnit string `json:"reading_unit"`
}

// AirTemperatureMetadataStation represents metadata information specific to a
// weather station.
type AirTemperatureMetadataStation struct {
	// ID of a station
	ID string `json:"id"`

	// ID of the device (usually the same as the station)
	DeviceID string `json:"device_id"`

	// Name of the station
	Name string `json:"name"`

	// Location of the station
	Location AirTemperatureMetadataStationLocation `json:"location"`
}

// AirTemperatureMetadataStationLocation represents the geographical coordindates
// of a weather station.
type AirTemperatureMetadataStationLocation struct {
	// Longitude of the station
	Longitude float64 `json:"longitude"`

	// Latitude of the station
	Latitude float64 `json:"latitude"`
}

// AirTemperatureItem represents all air temperature readings at a point in time.
type AirTemperatureItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Data readings
	Readings []AirTemperatureItemReading `json:"readings"`
}

// AirTemperatureItemReading represents a single air temperature reading at
// a specific station at a point in time.
type AirTemperatureItemReading struct {
	// ID of the station
	StationID string `json:"station_id"`

	// Value of the reading
	Value float64 `json:"value"`
}

// GetAirTemperature returns the AirTemperature information.
func (c *Client) GetAirTemperature(options ...*QueryOption) (*AirTemperature, error) {
	// Parse URL
	path := "/v1/environment/air-temperature/"
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
	data := &AirTemperature{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
