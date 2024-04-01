package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestReturnsErrorOnNotHttpOK(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Fatalf("Expected to request '/', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	heders := map[string]string{
		"Authorization": "token",
	}

	_, err := sendApiRequest(server.URL, "GET", heders, nil, true)
	if err == nil {
		t.Fatalf("Request didn't return error")
	}
}

func TestGetRequest(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Fatalf("Expected to request '/', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	heders := map[string]string{
		"Authorization": "token",
	}

	_, err := sendApiRequest(server.URL, "GET", heders, nil, true)
	if err != nil {
		t.Fatalf("Get failed with error\n%s", err)
	}
}
