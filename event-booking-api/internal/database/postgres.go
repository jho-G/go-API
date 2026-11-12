package database

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"

	"event-booking-api/config"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	// Safely encode the database username and password.
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, cfg.DBPassword),
		Host:   cfg.DBHost + ":" + cfg.DBPort,
		Path:   cfg.DBName,
	}

	query := u.Query()
	query.Set("sslmode", cfg.DBSSLMode)
	u.RawQuery = query.Encode()

	// Create a connection pool.
	pool, err := pgxpool.New(context.Background(), u.String())
	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}

	// Verify the database is reachable.
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	return pool, nil
}
