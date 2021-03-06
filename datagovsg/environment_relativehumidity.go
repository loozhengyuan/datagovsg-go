package datagovsg

import (
	"encoding/json"
	"net/url"
)

// RelativeHumidity is the resource representing the relative humidity information.
type RelativeHumidity struct {
	APIInfo  APIInfo                  `json:"api_info"`
	Metadata RelativeHumidityMetadata `json:"metadata"`
	Items    []RelativeHumidityItem   `json:"items"`
}

// RelativeHumidityMetadata represents metadata information about the relative humidity data.
type RelativeHumidityMetadata struct {
	// Metadata about a weather station
	Stations []RelativeHumidityMetadataStation `json:"stations"`

	// Information about the reading
	ReadingType string `json:"reading_type"`

	// Measurement unit for the reading
	ReadingUnit string `json:"reading_unit"`
}

// RelativeHumidityMetadataStation represents metadata information specific to a
// weather station.
type RelativeHumidityMetadataStation struct {
	// ID of a station
	ID string `json:"id"`

	// ID of the device (usually the same as the station)
	DeviceID string `json:"device_id"`

	// Name of the station
	Name string `json:"name"`

	// Location of the station
	Location RelativeHumidityMetadataStationLocation `json:"location"`
}

// RelativeHumidityMetadataStationLocation represents the geographical coordindates
// of a weather station.
type RelativeHumidityMetadataStationLocation struct {
	// Longitude of the station
	Longitude float64 `json:"longitude"`

	// Latitude of the station
	Latitude float64 `json:"latitude"`
}

// RelativeHumidityItem represents all relative humidity readings at a point in time.
type RelativeHumidityItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Data readings
	Readings []RelativeHumidityItemReading `json:"readings"`
}

// RelativeHumidityItemReading represents a single relative humidity reading at
// a specific station at a point in time.
type RelativeHumidityItemReading struct {
	// ID of the station
	StationID string `json:"station_id"`

	// Value of the reading
	Value float64 `json:"value"`
}

// GetRelativeHumidity returns the relative humidity information.
func (c *Client) GetRelativeHumidity(options ...*QueryOption) (*RelativeHumidity, error) {
	// Parse URL
	path := "/v1/environment/relative-humidity/"
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
	data := &RelativeHumidity{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
