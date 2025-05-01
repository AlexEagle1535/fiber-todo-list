package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func tableExists(conn *pgx.Conn) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE table_schema = 'public' 
		AND table_name = 'tasks'
	);`
	err := conn.QueryRow(context.Background(), query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if table exists: %w", err)
	}
	return exists, nil
}

func RunMigrations(conn *pgx.Conn) error {
	err := createDatabase(conn, os.Getenv("DB_NAME"))
	if err != nil {
		return err
	}
	exists, err := tableExists(conn)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("Table 'tasks' already exists, migration skipped.")
		return nil
	}

	sqlBytes, err := os.ReadFile("db/migration.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	_, err = conn.Exec(context.Background(), string(sqlBytes))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	fmt.Println("Migration completed successfully!")
	return nil
}

func createDatabase(conn *pgx.Conn, dbName string) error {
	var exists bool
	query := `SELECT EXISTS (
		SELECT FROM pg_database
		WHERE datname = $1
	);`
	err := conn.QueryRow(context.Background(), query, dbName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if exists {
		fmt.Printf("Database %s already exists.\n", dbName)
		return nil
	}

	createDbQuery := fmt.Sprintf(`CREATE DATABASE "%s";`, dbName)
	_, err = conn.Exec(context.Background(), createDbQuery)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	fmt.Printf("Database %s successfully created!\n", dbName)
	return nil
}
