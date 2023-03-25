package app

import (
	"gofiberbasic/models"
	"gofiberbasic/services"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	TodoService services.TodoService
}

type TodoHandlerInterface interface {
	CreateTodo() error
	GetAllTodos() error
}

func NewTodoHandler(todoService services.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: todoService}
}

func (h *TodoHandler) CreateTodo(ctx *fiber.Ctx) error {
	todo := models.Todo{}
	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	response, err := h.TodoService.Create(todo)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (h *TodoHandler) GetAllTodos(ctx *fiber.Ctx) error {
	todos, err := h.TodoService.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(todos)
}
