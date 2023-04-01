package services

import (
	"gofiberbasic/mocks/repository"
	"gofiberbasic/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockRepository *repository.MockTodoRepository
var service TodoService

var fakeData = []models.Todo{
	{
		ID:          primitive.NewObjectID(),
		Title:       "Title 1",
		Description: "Description 1",
		Completed:   false,
	},
	{
		ID:          primitive.NewObjectID(),
		Title:       "Title 2",
		Description: "Description 2",
		Completed:   false,
	},
	{
		ID:          primitive.NewObjectID(),
		Title:       "Title 3",
		Description: "Description 3",
		Completed:   true,
	},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepository = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepository)

	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestGetAll(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	mockRepository.EXPECT().GetAll().Return(fakeData, nil)

	todos, err := service.GetAll()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.NotEmpty(t, todos)
}

func TestGetByID(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	mockRepository.EXPECT().GetByID(gomock.Any()).Return(fakeData[0], nil)

	todo, err := service.GetByID("5f9f1b5b9c9c1b1b8c1c1c1c")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.NotEmpty(t, todo)
}

func TestCreate(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	var fakeData_ = fakeData[0]
	mockRepository.EXPECT().Create(gomock.Any()).Return(true, nil)

	todo, err := service.Create(fakeData_)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, todo.Status, true)
}

func TestUpdate(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	var fakeData_ = fakeData[0]

	mockRepository.EXPECT().Update(gomock.Any()).Return(fakeData_, nil)

	todo, err := service.Update(fakeData_)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, todo, fakeData_)
	assert.NotEmpty(t, todo)
}

func TestDelete(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	mockRepository.EXPECT().Delete(gomock.Any()).Return(true, nil)

	todo, err := service.Delete("5f9f1b5b9c9c1b1b8c1c1c1c")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, todo.Status, true)
}
