package datagovsg

import (
	"encoding/json"
	"net/url"
)

// FourDayWeatherForecast is the resource representing the twenty-four-hourly weather
// forecast information.
type FourDayWeatherForecast struct {
	APIInfo APIInfo                      `json:"api_info"`
	Items   []FourDayWeatherForecastItem `json:"items"`
}

// FourDayWeatherForecastItem represents a forecast reading at a point in time.
type FourDayWeatherForecastItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Timestamp of data acquisition
	UpdateTimestamp string `json:"update_timestamp"`

	// Forecast for each day period
	Forecasts []FourDayWeatherForecastItemForecast `json:"forecasts"`
}

// FourDayWeatherForecastItemForecast represents the forecast for a single day period.
type FourDayWeatherForecastItemForecast struct {
	// Date of the forecast
	Date string `json:"date"`

	// Timestamp of the forecast
	Timestamp string `json:"timestamp"`

	// General weather forecast
	Forecast string `json:"forecast"`

	// Relative humidity forecast
	RelativeHumidity FourDayWeatherForecastItemGeneralRelativeHumidity `json:"relative_humidity"`

	// Temperature forecast
	Temperature FourDayWeatherForecastItemGeneralTemperature `json:"temperature"`

	// Wind forecast
	Wind FourDayWeatherForecastItemGeneralWind `json:"wind"`
}

// FourDayWeatherForecastItemGeneralRelativeHumidity represents the forecast information for
// relative humidity.
type FourDayWeatherForecastItemGeneralRelativeHumidity struct {
	// Lowest reading within the period
	Low int `json:"low"`

	// Highest reading within the period
	High int `json:"high"`
}

// FourDayWeatherForecastItemGeneralTemperature represents the forecast information for
// temperature.
type FourDayWeatherForecastItemGeneralTemperature struct {
	// Lowest reading within the period
	Low int `json:"low"`

	// Highest reading within the period
	High int `json:"high"`
}

// FourDayWeatherForecastItemGeneralWind represents the forecast information for wind
// direction and speeds.
type FourDayWeatherForecastItemGeneralWind struct {
	// Wind direction
	Direction string `json:"direction"`

	// Wind speed
	Speed FourDayWeatherForecastItemGeneralWindSpeed `json:"speed"`
}

// FourDayWeatherForecastItemGeneralWindSpeed represents the forecast information for
// wind speeds.
type FourDayWeatherForecastItemGeneralWindSpeed struct {
	// Lowest reading within the period
	Low int `json:"low"`

	// Highest reading within the period
	High int `json:"high"`
}

// GetFourDayWeatherForecast returns the four-daily weather forecast information.
func (c *Client) GetFourDayWeatherForecast(options ...*QueryOption) (*FourDayWeatherForecast, error) {
	// Parse URL
	path := "/v1/environment/4-day-weather-forecast/"
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
	data := &FourDayWeatherForecast{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
