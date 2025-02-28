package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/Bladforceone/go_hw_otus/hw15_go_sql/cmd/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	log.Println("Database connected")

	return pool, nil
}
