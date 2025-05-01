package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetAllTasks(db *pgx.Conn) ([]Task, error) {
	rows, err := db.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func CreateTask(db *pgx.Conn, t *Task) error {
	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	return db.QueryRow(context.Background(), query, t.Title, t.Description, t.Status).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func UpdateTask(db *pgx.Conn, id int, t *Task) error {
	query := `UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=$4 WHERE id=$5`
	_, err := db.Exec(context.Background(), query, t.Title, t.Description, t.Status, time.Now(), id)
	return err
}

func DeleteTask(db *pgx.Conn, id int) error {
	_, err := db.Exec(context.Background(), "DELETE FROM tasks WHERE id=$1", id)
	return err
}
