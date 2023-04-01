package main

import (
	"gofiberbasic/app"
	"gofiberbasic/config"
	"gofiberbasic/repository"
	"gofiberbasic/services"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	appRoute := fiber.New()

	config.ConnectMongoDB()                    // Connect to MongoDB
	client := config.CreateCollection("todos") // Create a new collection

	todoRepository := repository.NewTodoRepository(client)
	todoService := services.NewTodoService(todoRepository)
	todoHandler := app.NewTodoHandler(todoService)

	// create a new todo
	// POST http://localhost:3000/api/todo
	// {
	// 	"title": "Learn Go",
	// 	"description": "Learn Go and build a REST API"
	// }
	appRoute.Post("/api/todo", todoHandler.CreateTodo)
	appRoute.Get("/api/todos", todoHandler.GetAllTodos)
	appRoute.Get("/api/todo/:id", todoHandler.GetTodoByID)
	appRoute.Put("/api/todo/:id", todoHandler.UpdateTodo)
	appRoute.Delete("/api/todo/:id", todoHandler.DeleteTodo)

	appRoute.Listen(":3000")
}
