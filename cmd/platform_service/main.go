package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kheyssper/go-saas-microservice/internal/config"
	"github.com/kheyssper/go-saas-microservice/internal/db"
	"github.com/kheyssper/go-saas-microservice/internal/middleware"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/controllers"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/repositories"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/services"
)

func main() {
	// Carrega a configuração
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar a configuração: %v", err)
	}

	// Conecta ao banco de dados
	dbPool, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer dbPool.Close()

	// Inicializa os repositórios, serviços e controladores
	platformRepo := repositories.NewPlatformRepository(dbPool)
	platformService := services.NewPlatformService(platformRepo)
	platformController := controllers.NewPlatformController(platformService)

	// Cria o roteador e define as rotas
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/platforms", platformController.CreatePlatformHandler).Methods("POST")
	router.HandleFunc("/platforms", platformController.ListPlatformsHandler).Methods("GET")
	router.HandleFunc("/platforms/{id:[0-9]+}", platformController.GetPlatformByID).Methods("GET")
	router.HandleFunc("/platforms/{id:[0-9]+}/status", platformController.RunOrStopPlatformHandler).Methods("PUT")
	router.HandleFunc("/platforms/{id:[0-9]+}", platformController.DeletePlatform).Methods("DELETE")

	// Inicia o servidor HTTP
	log.Printf("Iniciando o servidor na porta %s...", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
