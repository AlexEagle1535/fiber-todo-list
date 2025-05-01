package main

import (
	"context"
	"log"

	"github.com/AlexEagle1535/fiber-todo-list/db"
	"github.com/AlexEagle1535/fiber-todo-list/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}
	app := fiber.New()
	conn := db.Run()
	defer conn.Close(context.Background())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	router.Run(app, conn)
	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
