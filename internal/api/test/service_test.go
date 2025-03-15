package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/oganes5796/employment-test/internal/api"
	"github.com/oganes5796/employment-test/internal/domain"
	"github.com/oganes5796/employment-test/internal/repository/todo/test"
	"github.com/oganes5796/employment-test/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockRepo := new(test.MockRepository)
	service := service.TodoService(mockRepo)
	handler := api.NewImplementation(service)

	app := fiber.New()
	app.Post("/tasks", handler.CreateTodo)

	task := domain.Todo{
		Title:       "Test Task",
		Description: "Description",
		Status:      "new",
	}

	mockRepo.On("Create", mock.Anything, &task).Return(nil)

	body, _ := json.Marshal(task)
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	mockRepo.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	mockRepo := new(test.MockRepository)
	service := service.TodoService(mockRepo)
	handler := api.NewImplementation(service)

	app := fiber.New()
	app.Get("/tasks", handler.GetTodo)

	tasks := []domain.Todo{
		{ID: 1, Title: "Task 1", Status: "new"},
		{ID: 2, Title: "Task 2", Status: "done"},
	}

	mockRepo.On("GetAll", mock.Anything).Return(tasks, nil)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockRepo.AssertExpectations(t)
}
