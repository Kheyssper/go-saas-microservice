package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/models"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/services"
)

// PlatformController lida com as requisições HTTP relacionadas às plataformas.
type PlatformController struct {
	service *services.PlatformService
}

// NewPlatformController cria uma nova instância de PlatformController.
func NewPlatformController(service *services.PlatformService) *PlatformController {
	return &PlatformController{service: service}
}

// CreatePlatformHandler trata a criação de uma nova plataforma.
func (c *PlatformController) CreatePlatformHandler(w http.ResponseWriter, r *http.Request) {
	var platform models.Platform

	// Decodifica o corpo da requisição diretamente para o struct Platform
	if err := json.NewDecoder(r.Body).Decode(&platform); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Chama o serviço para criar a plataforma
	newPlatform, err := c.service.CreatePlatform(r.Context(), platform.PlatformName, platform.PlatformSlug, platform.CreatorID)
	if err != nil {
		http.Error(w, "Failed to create platform", http.StatusInternalServerError)
		return
	}

	// Retorna a resposta em JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPlatform)
}

// ListPlatformsHandler lida com a listagem de todas as plataformas.
func (c *PlatformController) ListPlatformsHandler(w http.ResponseWriter, r *http.Request) {
	platforms, err := c.service.ListPlatforms(r.Context())
	if err != nil {
		http.Error(w, "Failed to list platforms", http.StatusInternalServerError)
		return
	}

	// Retorna as plataformas em JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(platforms)
}

// RunOrStopPlatformHandler lida com a atualização do status da plataforma (executar ou parar).
func (c *PlatformController) RunOrStopPlatformHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid platform ID", http.StatusBadRequest)
		return
	}

	var platform models.Platform

	// Decodifica o corpo da requisição diretamente para o struct Platform
	if err := json.NewDecoder(r.Body).Decode(&platform); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Atualiza o status da plataforma
	err = c.service.RunOrStopPlatform(r.Context(), id, platform.Status)
	if err != nil {
		http.Error(w, "Failed to update platform status", http.StatusInternalServerError)
		return
	}

	// Resposta de sucesso
	w.WriteHeader(http.StatusNoContent)
}

// DeletePlatform deleta uma plataforma por ID.
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

	w.WriteHeader(http.StatusNoContent)
}

// GetPlatformByID busca uma plataforma por ID.
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
