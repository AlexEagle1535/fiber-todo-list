package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func Run(migration bool) *pgx.Conn {
	if migration {
		pgURL := os.Getenv("PG_URL")
		if pgURL == "" {
			log.Fatal("PG_URL not set in environment")
		}
		conn, err := pgx.Connect(context.Background(), pgURL)
		if err != nil {
			log.Fatalf("Failed to connect to PG_URL: %v", err)
		}
		err = RunMigrations(conn)
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		return conn
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to DATABASE_URL: %v", err)
	}
	return conn
}
