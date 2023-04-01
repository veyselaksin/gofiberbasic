package app

import (
	"gofiberbasic/models"
	"gofiberbasic/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	TodoService services.TodoService
}

type TodoHandlerInterface interface {
	CreateTodo() error
	GetAllTodos() error
	GetTodoByID() error
	UpdateTodo() error
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(todos)
}

func (h *TodoHandler) GetTodoByID(ctx *fiber.Ctx) error {
	// get id from url
	var id string = ctx.Params("id")

	todo, err := h.TodoService.GetByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(todo)
}

func (h *TodoHandler) UpdateTodo(ctx *fiber.Ctx) error {
	// get id from url
	var id string = ctx.Params("id")

	todo := models.Todo{}
	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_id, err := primitive.ObjectIDFromHex(id)
	todo.ID = _id
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	result, err := h.TodoService.Update(todo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (h *TodoHandler) DeleteTodo(ctx *fiber.Ctx) error {
	// get id from url
	var id string = ctx.Params("id")

	result, err := h.TodoService.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
