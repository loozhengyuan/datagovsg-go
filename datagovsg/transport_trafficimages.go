package datagovsg

import (
	"encoding/json"
	"net/url"
)

// TrafficImages is the resource representing the Traffic Images.
type TrafficImages struct {
	APIInfo APIInfo             `json:"api_info"`
	Items   []TrafficImagesItem `json:"items"`
}

// TrafficImagesItem represents a set of traffic cameras
// and their images at a point in time.
type TrafficImagesItem struct {
	// Time of acquisition of data
	Timestamp string `json:"timestamp"`

	// Camera infromation and images
	Cameras []TrafficImagesCamera `json:"cameras"`
}

// TrafficImagesCamera represents a traffic image retrieved
// from a traffic camera at a point in time.
type TrafficImagesCamera struct {
	// Time of image
	Timestamp string `json:"timestamp"`

	// URL of image
	Image string `json:"image"`

	// Location of the traffic camera
	Location TrafficImagesCameraLocation `json:"location"`

	// Camera ID provided by LTA
	CameraID string `json:"camera_id"`

	// Metadata of image file
	ImageMetadata TrafficImagesCameraImageMetadata `json:"image_metadata"`
}

// TrafficImagesCameraLocation represents the geographical coordinates
// of a traffic camera.
type TrafficImagesCameraLocation struct {
	// Latitude of the traffic camera
	Latitude float64 `json:"latitude"`

	// Longitude of the traffic camera
	Longitude float64 `json:"longitude"`
}

// TrafficImagesCameraImageMetadata represents metadata information
// of a single traffic image.
type TrafficImagesCameraImageMetadata struct {
	// Height of the image in pixels
	Height int `json:"height"`

	// Width of the image in pixels
	Width int `json:"width"`

	// MD5 hash of the image
	MD5 string `json:"md5"`
}

// GetTrafficImages returns the latest images from traffic
// cameras all around Singapore.
func (c *Client) GetTrafficImages(options ...*QueryOption) (*TrafficImages, error) {
	// Parse URL
	path := "/v1/transport/traffic-images/"
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
	data := &TrafficImages{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return data, nil
}
