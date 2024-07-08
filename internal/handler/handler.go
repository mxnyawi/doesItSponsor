package handler

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/mxnyawi/doesItSponsor/internal/db"
)

// Define your handler struct with database dependency
type Handler struct {
	DB *db.Database
}

// GetOrganisation function to handle HTTP requests to get organisations by name
func (h *Handler) GetOrganisation(w http.ResponseWriter, r *http.Request) {
	// Get parameters from request
	vars := mux.Vars(r)
	// Build query
	orgName := vars["organisation_name"]

	if !h.isValidInput(orgName) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := db.BuildQuery(orgName, "organisation_name")

	orgs, err := h.DB.GetDocument(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode organisations as JSON and send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orgs); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// GetRoute function to handle HTTP requests to get organisations by route
func (h *Handler) GetRoute(w http.ResponseWriter, r *http.Request) {
	// Get parameters from request
	vars := mux.Vars(r)
	// Build query
	route := vars["route"]

	if !h.isValidInput(route) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := db.BuildQuery(route, "route")

	orgs, err := h.DB.GetDocument(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode organisations as JSON and send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orgs); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// GetType function to handle HTTP requests to get organisations by type
func (h *Handler) GetType(w http.ResponseWriter, r *http.Request) {
	// Get parameters from request
	vars := mux.Vars(r)
	// Build query

	typeRating := vars["type"]

	if !h.isValidInput(typeRating) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := db.BuildQuery(typeRating, "type")

	orgs, err := h.DB.GetDocument(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode organisations as JSON and send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orgs); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// GetCity function to handle HTTP requests to get organisations by city
func (h *Handler) GetCity(w http.ResponseWriter, r *http.Request) {
	// Get parameters from request
	vars := mux.Vars(r)
	// Build query
	city := vars["city"]

	if !h.isValidInput(city) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := db.BuildQuery(city, "city")

	orgs, err := h.DB.GetDocument(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode organisations as JSON and send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orgs); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// GetCounty function to handle HTTP requests to get organisations by county
func (h *Handler) GetCounty(w http.ResponseWriter, r *http.Request) {
	// Get parameters from request
	vars := mux.Vars(r)
	// Build query
	county := vars["county"]

	if !h.isValidInput(county) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := db.BuildQuery(county, "county")

	orgs, err := h.DB.GetDocument(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode organisations as JSON and send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orgs); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (h *Handler) isValidInput(input string) bool {
	validStrings := `^[a-zA-Z0-9\s\.,'-]+$`
	matched, _ := regexp.MatchString(validStrings, input)
	return matched
}
