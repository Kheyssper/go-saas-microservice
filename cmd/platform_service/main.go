package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kheyssper/go-saas-microservice/internal/config"
	"github.com/kheyssper/go-saas-microservice/internal/db"
	"github.com/kheyssper/go-saas-microservice/internal/middleware"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/controllers"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/repositories"
	"github.com/kheyssper/go-saas-microservice/pkg/platform_service/services"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load configuration
	cfg, err := config.LoadConfig() // Ensure your LoadConfig method reads from environment variables
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Connect to the database using the SetupDB method
	dbPool, err := db.SetupDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbPool.Close()

	// Initialize repositories, services, and controllers
	platformRepo := repositories.NewPlatformRepository(dbPool)
	platformService := services.NewPlatformService(platformRepo)
	platformController := controllers.NewPlatformController(platformService)

	// Create router and define routes
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/platforms", platformController.CreatePlatformHandler).Methods("POST")
	router.HandleFunc("/platforms", platformController.ListPlatformsHandler).Methods("GET")
	router.HandleFunc("/platforms/{id:[0-9]+}", platformController.GetPlatformByID).Methods("GET")
	router.HandleFunc("/platforms/{id:[0-9]+}/status", platformController.RunOrStopPlatformHandler).Methods("PUT")
	router.HandleFunc("/platforms/{id:[0-9]+}", platformController.DeletePlatform).Methods("DELETE")

	// Start HTTP server
	log.Printf("Starting server on port %s...", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
