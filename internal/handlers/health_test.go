package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nstoker/fictional-pancake/internal/version"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthHandler)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Health{}
	got := Health{}
	expected.OK = true
	expected.Version = version.Version()

	json.Unmarshal(rr.Body.Bytes(), &got)
	if got.OK != expected.OK || got.Version != expected.Version {
		t.Errorf("handler returned wrong OK response: got %v, want %v", got, expected)
	}
}
