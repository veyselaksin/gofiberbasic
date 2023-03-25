package repository

import (
	"context"
	"errors"
	"gofiberbasic/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Create(todo models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	// GetByID(id string) (models.Todo, error)
	// Update(todo models.Todo) error
	// Delete(id string) (bool, error)
}

func NewTodoRepository(todoCollection *mongo.Collection) TodoRepository {
	return &todoRepositoryDB{TodoCollection: todoCollection}
}

func (t todoRepositoryDB) Create(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.InsertOne(ctx, todo)
	if err != nil || result.InsertedID == nil {
		errors.New("Error while creating todo")
		return false, err
	}

	return true, nil
}

func (t todoRepositoryDB) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	var todo models.Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})
	if err != nil {
		return todos, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatal(err)
			return []models.Todo{}, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
