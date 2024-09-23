package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// SetupDB initializes the database connection using environment variables
func SetupDB() (*pgxpool.Pool, error) {
	// Get environment variables
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST") // assuming you might also have a HOST variable
	port := os.Getenv("POSTGRES_PORT") // if you need port number as well

	// Construct the connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbName)

	// Initialize connection pool
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully.")
	return pool, nil
}
