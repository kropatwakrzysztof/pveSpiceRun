package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiUrl(t *testing.T) {
	tests := map[string]struct {
		hostUrl  string
		expected string
	}{
		"full url":       {hostUrl: "https://lcalhost:4345", expected: "https://lcalhost:4345/api2/json"},
		"trailing slash": {hostUrl: "https://lcalhost:4345/", expected: "https://lcalhost:4345/api2/json"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			apiUrl, err := makeApiUrl(tc.hostUrl)
			if err != nil {
				t.Fatalf("In \"%s\" error while parsing host url \"%s\"", name, err)
			}

			if apiUrl != tc.expected {
				t.Fatalf("In \"%s\" expected api url \"%s\" got \"%s\"", name, tc.expected, apiUrl)
			}
		})
	}
}

func TestMakeProxyAddress(t *testing.T) {
	tests := map[string]struct {
		proxy    string
		hostUrl  string
		expected string
	}{
		"proxy set":      {proxy: "10.0.0.1", hostUrl: "https://localhost:345", expected: "10.0.0.1"},
		"no proxy":       {proxy: "", hostUrl: "http://10.0.1.0", expected: "10.0.1.0"},
		"trailing slash": {proxy: "", hostUrl: "https://localhost:394/", expected: "localhost"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			proxyAddr, err := makeProxyAddr(tc.proxy, tc.hostUrl)
			if err != nil {
				t.Fatalf("In \"%s\" error while parsing proxy address \"%s\"", name, err)
			}

			if proxyAddr != tc.expected {
				t.Fatalf("In \"%s\" expected proxy address \"%s\" got \"%s\"", name, tc.expected, proxyAddr)
			}
		})
	}
}

func TestNoAuthentication(t *testing.T) {
	// Prepare mock setup
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api2/json/access/permissions" {
			t.Fatalf("Expected to request '/api2/json/access/permissions', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	// Call the application code
	err := ConnectSpiceTarget(server.URL, "", true, 203, "pveUser", "pveUserToken", "pveUserSecret", "")

	// Assert that os.Exit gets called
	expectedError := `ERROR: Authentication failed: 401 Unauthorized`
	if err.Error() != expectedError {
		t.Fatalf("Expected error message \"%s\" got \"%s\"", expectedError, err)
	}
}
