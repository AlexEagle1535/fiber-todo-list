package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Run() *pgx.Conn {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:password@localhost:5432/todo-list"
	}
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	return conn
}
