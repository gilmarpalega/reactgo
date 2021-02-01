package main

import (
	"reactgo/handlers"
	"reactgo/models"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/gofiber/fiber/middleware/logger"
)

func init() {
	models.Setup()
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/tasks/:id", handlers.GetTask)
	app.Get("/tasks", handlers.GetTasks)
	app.Post("/tasks", handlers.AddTask)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	app.Listen("localhost:5000")

}
