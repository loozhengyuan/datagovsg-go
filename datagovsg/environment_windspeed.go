package datagovsg

import (
	"encoding/json"
	"net/url"
)

// WindSpeed is the resource representing the wind speed information.
type WindSpeed struct {
	APIInfo  APIInfo           `json:"api_info"`
	Metadata WindSpeedMetadata `json:"metadata"`
	Items    []WindSpeedItem   `json:"items"`
}

// WindSpeedMetadata represents metadata information about the wind speed data.
type WindSpeedMetadata struct {
	// Metadata about a weather station
	Stations []WindSpeedMetadataStation `json:"stations"`

	// Information about the reading
	ReadingType string `json:"reading_type"`

	// Measurement unit for the reading
	ReadingUnit string `json:"reading_unit"`
}

// WindSpeedMetadataStation represents metadata information specific to a
// weather station.
type WindSpeedMetadataStation struct {
	// ID of a station
	ID string `json:"id"`

	// ID of the device (usually the same as the station)
	DeviceID string `json:"device_id"`

	// Name of the station
	Name string `json:"name"`

	// Location of the station
	Location WindSpeedMetadataStationLocation `json:"location"`
}

// WindSpeedMetadataStationLocation represents the geographical coordindates
// of a weather station.
type WindSpeedMetadataStationLocation struct {
	// Longitude of the station
	Longitude float64 `json:"longitude"`

	// Latitude of the station
	Latitude float64 `json:"latitude"`
}

// WindSpeedItem represents all wind speed readings at a point in time.
type WindSpeedItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Data readings
	Readings []WindSpeedItemReading `json:"readings"`
}

// WindSpeedItemReading represents a single wind speed reading at
// a specific station at a point in time.
type WindSpeedItemReading struct {
	// ID of the station
	StationID string `json:"station_id"`

	// Value of the reading
	Value float64 `json:"value"`
}

// GetWindSpeed returns the wind speed information.
func (c *Client) GetWindSpeed(options ...*QueryOption) (*WindSpeed, error) {
	// Parse URL
	path := "/v1/environment/wind-speed/"
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
	data := &WindSpeed{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
