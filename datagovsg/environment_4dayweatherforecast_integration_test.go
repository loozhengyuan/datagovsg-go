// +build integration

package datagovsg

import (
	"testing"
)

func TestFourDayWeatherForecast(t *testing.T) {
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
			_, err := c.GetFourDayWeatherForecast()
			if err != nil {
				t.Errorf("error executing request: %v", err)
			}
		})
	}
}
