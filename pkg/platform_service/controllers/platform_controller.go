package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/models"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/services"
)

// PlatformController handles HTTP requests related to platforms.
type PlatformController struct {
	service *services.PlatformService
}

// NewPlatformController creates a new instance of PlatformController.
func NewPlatformController(service *services.PlatformService) *PlatformController {
	return &PlatformController{service: service}
}

// CreatePlatformHandler handles creating a new platform.
func (c *PlatformController) CreatePlatformHandler(w http.ResponseWriter, r *http.Request) {
	var platform models.Platform

	// Decode the request body directly into the Platform struct
	if err := json.NewDecoder(r.Body).Decode(&platform); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service to create the platform
	newPlatform, err := c.service.CreatePlatform(r.Context(), platform.PlatformName, platform.PlatformSlug, platform.CreatorID)
	if err != nil {
		http.Error(w, "Failed to create platform", http.StatusInternalServerError)
		return
	}

	// Return the JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPlatform)
}

// ListPlatformsHandler handles listing all platforms.
func (c *PlatformController) ListPlatformsHandler(w http.ResponseWriter, r *http.Request) {
	platforms, err := c.service.ListPlatforms(r.Context())
	if err != nil {
		http.Error(w, "Failed to list platforms", http.StatusInternalServerError)
		return
	}

	// Return platforms in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(platforms)
}

// RunOrStopPlatformHandler handles updating the platform status (run or stop).
func (c *PlatformController) RunOrStopPlatformHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid platform ID", http.StatusBadRequest)
		return
	}

	var platform models.Platform

	// Decode the request body directly into the Platform struct
	if err := json.NewDecoder(r.Body).Decode(&platform); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the platform status
	err = c.service.RunOrStopPlatform(r.Context(), id, platform.Status)
	if err != nil {
		http.Error(w, "Failed to update platform status", http.StatusInternalServerError)
		return
	}

	// Success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Platform Status updated successfully"}`))
}

// DeletePlatform deletes a platform by ID.
func (pc *PlatformController) DeletePlatform(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid platform ID", http.StatusBadRequest)
		return
	}

	err = pc.service.DeletePlatform(r.Context(), id)
	if err != nil {
		http.Error(w, "Error deleting platform", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Platform deleted successfully"}`))

}

// GetPlatformByID retrieves a platform by ID.
func (pc *PlatformController) GetPlatformByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid platform ID", http.StatusBadRequest)
		return
	}

	platform, err := pc.service.GetPlatformByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Platform not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(platform)
}
