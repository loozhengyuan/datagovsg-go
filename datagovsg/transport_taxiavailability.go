package datagovsg

import (
	"encoding/json"
	"net/url"
)

// TaxiAvailability is the resource representing the Taxi Availability in Singapore.
type TaxiAvailability struct {
	Type     string                    `json:"type"`
	CRS      TaxiAvailabilityCRS       `json:"crs"`
	Features []TaxiAvailabilityFeature `json:"features"`
}

// TaxiAvailabilityCRS describes the Coordinate Reference System used.
type TaxiAvailabilityCRS struct {
	Type       string                        `json:"type"`
	Properties TaxiAvailabilityCRSProperties `json:"properties"`
}

// TaxiAvailabilityCRSProperties represents properties of the CRS.
type TaxiAvailabilityCRSProperties struct {
	Type string `json:"type"`
	Href string `json:"href"`
}

// TaxiAvailabilityFeature corresponds to a GeoJSON Feature object.
type TaxiAvailabilityFeature struct {
	Type       string                            `json:"type"`
	Geometry   TaxiAvailabilityFeatureGeometry   `json:"geometry"`
	Properties TaxiAvailabilityFeatureProperties `json:"properties"`
}

// TaxiAvailabilityFeatureGeometry corresponds to a GeoJSON Geometry object.
type TaxiAvailabilityFeatureGeometry struct {
	Type        string                                       `json:"type"`
	Coordinates []TaxiAvailabilityFeatureGeometryCoordinates `json:"coordinates"`
}

// TaxiAvailabilityFeatureGeometryCoordinates represents the
// geographical coodinates of a Taxi.
type TaxiAvailabilityFeatureGeometryCoordinates []float64

// TaxiAvailabilityFeatureProperties rer
type TaxiAvailabilityFeatureProperties struct {
	Timestamp string  `json:"timestamp"`
	TaxiCount int     `json:"taxi_count"`
	APIInfo   APIInfo `json:"api_info"`
}

// GetTaxiAvailability returns the taxi availability and the geographical
// coordinates of these available taxis in Singapore.
func (c *Client) GetTaxiAvailability(options ...*QueryOption) (*TaxiAvailability, error) {
	// Parse URL
	path := "/v1/transport/taxi-availability/"
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
	data := &TaxiAvailability{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
