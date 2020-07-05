package datagovsg

import (
	"encoding/json"
	"net/url"
)

// TwentyFourHourWeatherForecast is the resource representing the twenty-four-hourly weather
// forecast information.
type TwentyFourHourWeatherForecast struct {
	APIInfo APIInfo                             `json:"api_info"`
	Items   []TwentyFourHourWeatherForecastItem `json:"items"`
}

// TwentyFourHourWeatherForecastItem represents a forecast reading at a point in time.
type TwentyFourHourWeatherForecastItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Timestamp of data acquisition
	UpdateTimestamp string `json:"update_timestamp"`

	// Validity of the forecast
	ValidPeriod TwentyFourHourWeatherForecastItemValidity `json:"valid_period"`

	// General forecast for the period
	General TwentyFourHourWeatherForecastItemGeneral `json:"general"`

	// Forecast breakdown over regions and periods
	Periods []TwentyFourHourWeatherForecastItemPeriod `json:"periods"`
}

// TwentyFourHourWeatherForecastItemValidity represents the valid period of a forecast.
type TwentyFourHourWeatherForecastItemValidity struct {
	// Starting timestamp of the valid period
	Start string `json:"start"`

	// Ending timestamp of the valid period
	End string `json:"end"`
}

// TwentyFourHourWeatherForecastItemGeneral represents general information of a forecast.
type TwentyFourHourWeatherForecastItemGeneral struct {
	// General weather forecast
	Forecast string `json:"forecast"`

	// Relative humidity forecast
	RelativeHumidity TwentyFourHourWeatherForecastItemGeneralRelativeHumidity `json:"relative_humidity"`

	// Temperature forecast
	Temperature TwentyFourHourWeatherForecastItemGeneralTemperature `json:"temperature"`

	// Wind forecast
	Wind TwentyFourHourWeatherForecastItemGeneralWind `json:"wind"`
}

// TwentyFourHourWeatherForecastItemGeneralRelativeHumidity represents the forecast information for
// relative humidity.
type TwentyFourHourWeatherForecastItemGeneralRelativeHumidity struct {
	// Lowest reading within the period
	Low int `json:"low"`

	// Highest reading within the period
	High int `json:"high"`
}

// TwentyFourHourWeatherForecastItemGeneralTemperature represents the forecast information for
// temperature.
type TwentyFourHourWeatherForecastItemGeneralTemperature struct {
	// Lowest reading within the period
	Low int `json:"low"`

	// Highest reading within the period
	High int `json:"high"`
}

// TwentyFourHourWeatherForecastItemGeneralWind represents the forecast information for wind
// direction and speeds.
type TwentyFourHourWeatherForecastItemGeneralWind struct {
	// Wind direction
	Direction string `json:"direction"`

	// Wind speed
	Speed TwentyFourHourWeatherForecastItemGeneralWindSpeed `json:"speed"`
}

// TwentyFourHourWeatherForecastItemGeneralWindSpeed represents the forecast information for
// wind speeds.
type TwentyFourHourWeatherForecastItemGeneralWindSpeed struct {
	// Lowest reading within the period
	Low int `json:"low"`

	// Highest reading within the period
	High int `json:"high"`
}

// TwentyFourHourWeatherForecastItemPeriod represents forecast information of regions
// over a time period.
type TwentyFourHourWeatherForecastItemPeriod struct {
	// Valid time period of the forecast
	Time TwentyFourHourWeatherForecastItemPeriodTime `json:"time"`

	// Weather forecast for each region
	Regions map[string]string `json:"regions"`
}

// TwentyFourHourWeatherForecastItemPeriodTime represents the valid period of a forecast.
type TwentyFourHourWeatherForecastItemPeriodTime struct {
	// Starting timestamp of the valid period
	Start string `json:"start"`

	// Ending timestamp of the valid period
	End string `json:"end"`
}

// GetTwentyFourHourWeatherForecast returns the twenty-four-hourly weather forecast information.
func (c *Client) GetTwentyFourHourWeatherForecast(options ...*QueryOption) (*TwentyFourHourWeatherForecast, error) {
	// Parse URL
	path := "/v1/environment/24-hour-weather-forecast/"
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
	data := &TwentyFourHourWeatherForecast{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
