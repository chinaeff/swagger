package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"net/http"
)

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type ResponseAddress struct {
	Addresses []*Address `json:"addresses"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	City string `json:"city"`
}

func SearchAddressHandler(w http.ResponseWriter, r *http.Request) {
	var request RequestAddressSearch
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Wrong data", http.StatusBadRequest)
		return
	}

	response := &ResponseAddress{
		Addresses: []*Address{
			{
				City: "Moscow",
			},
		},
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GeocodeAddressHandler(w http.ResponseWriter, r *http.Request) {
	var request RequestAddressGeocode
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Wrong data", http.StatusBadRequest)
		return
	}

	response := &ResponseAddress{
		Addresses: []*Address{
			{
				City: "Moscow",
			},
		},
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	r := chi.NewRouter()

	r.Post("/api/address/search", SearchAddressHandler)
	r.Post("/api/address/geocode", GeocodeAddressHandler)

	docgen.PrintRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
