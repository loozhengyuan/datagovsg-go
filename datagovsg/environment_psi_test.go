package datagovsg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestClient_GetPSI(t *testing.T) {
	// Create test cases
	cases := []struct {
		name    string
		fixture string
	}{
		{"default", "testdata/fixtures/environment_psi_default.json"},
	}

	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Load fixtures
			f, err := os.Open(tc.fixture)
			if err != nil {
				t.Errorf("error loading test fixtures: %v", err)
			}
			defer f.Close()

			// Read json
			b, err := ioutil.ReadAll(f)
			if err != nil {
				t.Errorf("error reading file: %v", err)
			}

			// Mock HTTP server
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			})
			server := httptest.NewServer(handler)

			// Execute request
			client := NewClient()
			client.BaseURL = server.URL
			got, err := client.GetPSI()
			if err != nil {
				t.Errorf("expected no errors but got: %v", err)
			}

			// Assert response body
			want := &PSI{}
			if err := json.Unmarshal(b, &want); err != nil {
				t.Errorf("error unmarshalling fixture: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v want %+v", got, want)
			}
		})
	}
}
