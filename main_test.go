package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearchAddressHandler(t *testing.T) {
	reqBody := `{"query": "Moscow"}`
	req := httptest.NewRequest("POST", "/api/address/search", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	SearchAddressHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, but %d", rec.Code)
	}

	var response ResponseAddress
	err := json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Errorf("Error of json decoding: %v", err)
	}

	expectedCity := "Moscow"
	if len(response.Addresses) != 1 || response.Addresses[0].City != expectedCity {
		t.Errorf("Wrong data. Expected %s, but %v", expectedCity, response)
	}
}

func TestGeocodeAddressHandler(t *testing.T) {
	reqBody := `{"lat": "55.755825", "lng": "37.617634"}`
	req := httptest.NewRequest("POST", "/api/address/geocode", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	GeocodeAddressHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, but %d", rec.Code)
	}

	var response ResponseAddress
	err := json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Errorf("Error of json decoding: %v", err)
	}

	expectedCity := "Moscow"
	if len(response.Addresses) != 1 || response.Addresses[0].City != expectedCity {
		t.Errorf("Wrong data. Expected %s, but %v", expectedCity, response)
	}
}
