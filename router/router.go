package router

import (
	"github.com/AlexEagle1535/fiber-todo-list/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func Run(app *fiber.App, db *pgx.Conn) {
	app.Post("/tasks", handlers.CreateTask(db))
	app.Get("/tasks", handlers.GetTasks(db))
	app.Put("/tasks/:id", handlers.UpdateTask(db))
	app.Delete("/tasks/:id", handlers.DeleteTask(db))
}
