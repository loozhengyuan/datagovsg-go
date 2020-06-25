package datagovsg

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get_basic(t *testing.T) {
	// Create test cases
	cases := []struct {
		name   string
		status int
		body   string
	}{
		{"http200", http.StatusOK, "OK"},
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

			// Execute request
			client := NewClient()
			got, err := client.Get(server.URL)
			if err != nil {
				t.Errorf("expected no errors but got: %v", err)
			}
			want := tc.body
			if string(got) != want {
				t.Errorf("got %+v want %+v", string(got), want)
			}
		})
	}
}

func TestClient_Get_errorInResponse(t *testing.T) {
	// Create test cases
	cases := []struct {
		name   string
		status int
		err    string
	}{
		{"http400_1", http.StatusBadRequest, "invalid datetime format"},
		{"http400_2", http.StatusBadRequest, "invalid request format"},
	}

	// Run test cases
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Format body
			body := fmt.Sprintf(`{"message":"%s"}`, tc.err)

			// Mock HTTP server
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.status)
				w.Write([]byte(body))
			})
			server := httptest.NewServer(handler)

			// Execute request
			client := NewClient()
			_, err := client.Get(server.URL)
			got := err.Error()
			want := fmt.Sprintf("%v: %v", ErrBadRequest, tc.err)
			if got != want {
				t.Errorf("got %+v want %+v", string(got), want)
			}
		})
	}
}
