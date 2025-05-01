package handlers

import (
	"strconv"

	"github.com/AlexEagle1535/fiber-todo-list/db"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func GetTasks(conn *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasks, err := db.GetAllTasks(conn)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch tasks"})
		}
		return c.JSON(tasks)
	}
}

func CreateTask(conn *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var t db.Task
		if err := c.BodyParser(&t); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
		if t.Status == "" {
			t.Status = "new"
		}
		if err := db.CreateTask(conn, &t); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create task"})
		}
		return c.Status(201).JSON(t)
	}
}

func UpdateTask(conn *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}
		var t db.Task
		if err := c.BodyParser(&t); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
		if err := db.UpdateTask(conn, id, &t); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update task"})
		}
		return c.SendStatus(204)
	}
}

func DeleteTask(conn *pgx.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}
		if err := db.DeleteTask(conn, id); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete task"})
		}
		return c.SendStatus(204)
	}
}
