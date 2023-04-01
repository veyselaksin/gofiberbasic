package services

import (
	"errors"
	"gofiberbasic/dto"
	"gofiberbasic/models"
	"gofiberbasic/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=../mocks/services/mockTodoService.go -package=services gofiberbasic/services TodoService
type todoService struct {
	TodoRepository repository.TodoRepository
}

type TodoService interface {
	Create(todo models.Todo) (*dto.TodoResponse, error)
	GetAll() ([]models.Todo, error)
	GetByID(id string) (models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(id string) (*dto.TodoResponse, error)
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoService{TodoRepository: todoRepository}
}

func (t todoService) Create(todo models.Todo) (*dto.TodoResponse, error) {

	var response = &dto.TodoResponse{}
	var err error
	if len(todo.Title) < 3 {
		response.Status = false
		return response, errors.New("Title must be at least 3 characters long")
	}

	if len(todo.Description) < 10 {
		response.Status = false
		return response, errors.New("Description must be at least 10 characters long")
	}

	todo.ID = primitive.NewObjectID()
	todo.Completed = false

	created, err := t.TodoRepository.Create(todo)
	if err != nil {
		response.Status = false
		return response, err
	}

	response.Status = created
	return response, nil
}

func (t todoService) GetAll() ([]models.Todo, error) {
	todos, err := t.TodoRepository.GetAll()
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (t todoService) GetByID(id string) (models.Todo, error) {
	todo, err := t.TodoRepository.GetByID(id)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (t todoService) Update(todo models.Todo) (models.Todo, error) {
	todo, err := t.TodoRepository.Update(todo)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (t todoService) Delete(id string) (*dto.TodoResponse, error) {
	var response = &dto.TodoResponse{}
	_id, err := primitive.ObjectIDFromHex(id)
	deleted, err := t.TodoRepository.Delete(_id)
	if err != nil {
		return response, err
	}

	response.Status = deleted

	return response, nil
}
