// +build integration

package datagovsg

import (
	"testing"
)

func TestTwoHourWeatherForecast(t *testing.T) {
	// Create test cases
	cases := []struct {
		name string
	}{
		{"default"},
	}

	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Execute request
			c := NewClient()
			_, err := c.GetTwoHourWeatherForecast()
			if err != nil {
				t.Errorf("error executing request: %v", err)
			}
		})
	}
}
