package datagovsg

import (
	"encoding/json"
	"net/url"
)

// TwoHourWeatherForecast is the resource representing the two-hourly weather
// forecast information.
type TwoHourWeatherForecast struct {
	APIInfo      APIInfo                              `json:"api_info"`
	AreaMetadata []TwoHourWeatherForecastAreaMetadata `json:"area_metadata"`
	Items        []TwoHourWeatherForecastItem         `json:"items"`
}

// TwoHourWeatherForecastAreaMetadata represents metadata about each area.
type TwoHourWeatherForecastAreaMetadata struct {
	// Name of the area
	Name string `json:"name"`

	// Geographical coordinates of the area
	LabelLocation TwoHourWeatherForecastAreaMetadataLabelLocation `json:"label_location"`
}

// TwoHourWeatherForecastAreaMetadataLabelLocation represents the geographical
// coordinates of the area.
type TwoHourWeatherForecastAreaMetadataLabelLocation struct {
	// Longitude of the area
	Longitude float64 `json:"longitude"`

	// Latitude of the area
	Latitude float64 `json:"latitude"`
}

// TwoHourWeatherForecastItem represents a forecast reading at a point in time.
type TwoHourWeatherForecastItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Timestamp of data acquisition
	UpdateTimestamp string `json:"update_timestamp"`

	// Validity of the forecast
	ValidPeriod TwoHourWeatherForecastItemValidity `json:"valid_period"`

	// Forecast for each area
	Forecasts []TwoHourWeatherForecastItemForecast `json:"forecasts"`
}

// TwoHourWeatherForecastItemValidity represents the valid period of a forecast.
type TwoHourWeatherForecastItemValidity struct {
	// Starting timestamp of the valid period
	Start string `json:"start"`

	// Ending timestamp of the valid period
	End string `json:"end"`
}

// TwoHourWeatherForecastItemForecast represents a single forecast for a specific area.
type TwoHourWeatherForecastItemForecast struct {
	// Geographical area
	Area string `json:"area"`

	// Value of the forecast
	Forecast string `json:"forecast"`
}

// GetTwoHourWeatherForecast returns the two-hourly weather forecast information.
func (c *Client) GetTwoHourWeatherForecast(options ...*QueryOption) (*TwoHourWeatherForecast, error) {
	// Parse URL
	path := "/v1/environment/2-hour-weather-forecast/"
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
	data := &TwoHourWeatherForecast{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
