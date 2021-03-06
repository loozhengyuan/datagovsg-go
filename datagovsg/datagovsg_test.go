package datagovsg

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_Get(t *testing.T) {
	// Create test cases
	cases := []struct {
		name   string
		status int
		body   string
		err    error
	}{
		{"http200", http.StatusOK, "OK", nil},

		// Returns error for non-2xx http status codes
		{"http301_responseNotOk", http.StatusPermanentRedirect, `{"message":"permanent redirect"}`, ErrResponseNotOk},
		{"http400_responseNotOk", http.StatusBadRequest, `{"message":"bad request"}`, ErrResponseNotOk},

		// Returns error when parsing error messages for non-2xx http status codes
		{"http301_parseErrorMessageFailure", http.StatusPermanentRedirect, `"message":"permanent redirect"`, ErrParseErrorMessageFailure},
		{"http400_parseErrorMessageFailure", http.StatusBadRequest, `{message: bad request}`, ErrParseErrorMessageFailure},
	}

	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Mock HTTP server
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.status)
				w.Write([]byte(tc.body))
			})
			server := httptest.NewServer(handler)

			// Parse URL
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Errorf("error parsing url: %v", err)
			}

			// Execute request
			client := NewClient()
			got, err := client.Get(u)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error '%v' but got: %v", tc.err, err)
			}

			// Assert response body if non-error responses
			want := tc.body
			if got != nil && string(got) != want {
				t.Errorf("got %+v want %+v", string(got), want)
			}
		})
	}
}
