package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kheyssper/go-saas-microservice/pkg/platform/models"
	"github.com/kheyssper/go-saas-microservice/pkg/platform/services"

	"github.com/gorilla/mux"
)

// PlatformController lida com as requisições HTTP relacionadas às plataformas.
type PlatformController struct {
	PlatformService *services.PlatformService
}

// NewPlatformController cria uma nova instância do PlatformController.
func NewPlatformController(service *services.PlatformService) *PlatformController {
	return &PlatformController{
		PlatformService: service,
	}
}

// CreatePlatform cria uma nova plataforma.
func (pc *PlatformController) CreatePlatform(w http.ResponseWriter, r *http.Request) {
	var platform models.Platform
	err := json.NewDecoder(r.Body).Decode(&platform)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdPlatform, err := pc.PlatformService.CreatePlatform(&platform)
	if err != nil {
		http.Error(w, "Error creating platform", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPlatform)
}

// ListPlatforms lista todas as plataformas.
func (pc *PlatformController) ListPlatforms(w http.ResponseWriter, r *http.Request) {
	platforms, err := pc.PlatformService.ListPlatforms()
	if err != nil {
		http.Error(w, "Error listing platforms", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(platforms)
}

// RunPlatform ativa uma plataforma específica.
func (pc *PlatformController) RunPlatform(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid platform ID", http.StatusBadRequest)
		return
	}

	err = pc.PlatformService.RunPlatform(id)
	if err != nil {
		http.Error(w, "Error running platform", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// StopPlatform desativa uma plataforma específica.
func (pc *PlatformController) StopPlatform(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid platform ID", http.StatusBadRequest)
		return
	}

	err = pc.PlatformService.StopPlatform(id)
	if err != nil {
		http.Error(w, "Error stopping platform", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
