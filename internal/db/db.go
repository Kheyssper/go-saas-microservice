package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

// Connect cria uma conexão com o banco de dados PostgreSQL
func Connect(databaseURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	dbPool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	log.Println("Conectado ao banco de dados com sucesso!")
	return dbPool, nil
}
