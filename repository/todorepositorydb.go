package repository

import (
	"context"
	"errors"
	"gofiberbasic/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks/repository/mockTodoRepository.go -package=repository gofiberbasic/repository TodoRepository
type todoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Create(todo models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	GetByID(id string) (models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
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

func (t todoRepositoryDB) GetByID(id string) (models.Todo, error) {
	var todo models.Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := t.TodoCollection.FindOne(ctx, bson.M{"title": "Learn Go"}).Decode(&todo); err != nil {
		return todo, err
	}

	return todo, nil
}

func (t todoRepositoryDB) Update(todo models.Todo) (models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := t.TodoCollection.UpdateOne(ctx, bson.M{"_id": todo.ID}, bson.M{"$set": bson.M{"completed": todo.Completed, "title": todo.Title, "description": todo.Description}})
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (t todoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {

		return false, err
	}

	if result.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}
