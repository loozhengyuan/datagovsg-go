package datagovsg

import (
	"encoding/json"
	"net/url"
)

// PSI is the resource representing the PSI information.
type PSI struct {
	APIInfo APIInfo   `json:"api_info"`
	Items   []PSIItem `json:"items"`
}

// PSIRegionMetadata represents metadata about each region.
type PSIRegionMetadata struct {
	// Name of the region
	Name string `json:"name"`

	// Geographical coordinates of the region
	LabelLocation PSIRegionMetadataLabelLocation `json:"label_location"`
}

// PSIRegionMetadataLabelLocation represents the geographical
// coordinates of the region.
type PSIRegionMetadataLabelLocation struct {
	// Longitude of the region
	Longitude float64 `json:"longitude"`

	// Latitude of the region
	Latitude float64 `json:"latitude"`
}

// PSIItem represents a PSI reading at a point in time.
type PSIItem struct {
	// Timestamp of the reading
	Timestamp string `json:"timestamp"`

	// Timestamp of data acquisition
	UpdateTimestamp string `json:"update_timestamp"`

	// Data readings
	Readings PSIItemReadings `json:"readings"`
}

// PSIItemReadings represents PSI and its sub-component readings.
type PSIItemReadings struct {
	// PSI overall reading
	PSITwentyFourHourly map[string]int `json:"psi_twenty_four_hourly"`

	// PM10 readings
	PM10SubIndex         map[string]int `json:"pm10_sub_index"`
	PM10TwentyFourHourly map[string]int `json:"pm10_twenty_four_hourly"`

	// PM2.5 readings
	PM25SubIndex         map[string]int `json:"pm25_sub_index"`
	PM25TwentyFourHourly map[string]int `json:"pm25_twenty_four_hourly"`

	// Ozone readings
	O3SubIndex     map[string]int `json:"o3_sub_index"`
	O3EightHourMax map[string]int `json:"o3_eight_hour_max"`

	// Carbon Monoxide readings
	// NOTE: COEightHourMax reading is the only non-integer data point.
	COSubIndex     map[string]int     `json:"co_sub_index"`
	COEightHourMax map[string]float64 `json:"co_eight_hour_max"`

	// Sulphur Dioxide readings
	SO2SubIndex         map[string]int `json:"so2_sub_index"`
	SO2TwentyFourHourly map[string]int `json:"so2_twenty_four_hourly"`

	// Nitrogen Dioxide readings
	// NOTE: Sub index reading for NO2 is not available.
	NO2OneHourMax map[string]int `json:"no2_one_hour_max"`
}

// GetPSI returns the PSI information.
func (c *Client) GetPSI(options ...*QueryOption) (*PSI, error) {
	// Parse URL
	path := "/v1/environment/psi/"
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
	data := &PSI{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
