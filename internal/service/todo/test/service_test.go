package test

import (
	"context"
	"errors"
	"testing"

	"github.com/oganes5796/employment-test/internal/domain"
	"github.com/oganes5796/employment-test/internal/repository/todo/test"
	"github.com/oganes5796/employment-test/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTask(t *testing.T) {
	mockRepo := new(test.MockRepository)
	service := service.TodoService(mockRepo)

	task := &domain.Todo{
		Title:       "Test Task",
		Description: "Description",
		Status:      "new",
	}

	mockRepo.On("Create", mock.Anything, task).Return(nil)

	_, err := service.Create(context.Background(), task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	mockRepo := new(test.MockRepository)
	service := service.TodoService(mockRepo)

	tasks := []domain.Todo{
		{ID: 1, Title: "Task 1", Status: "new"},
		{ID: 2, Title: "Task 2", Status: "done"},
	}

	mockRepo.On("GetAll", mock.Anything).Return(tasks, nil)

	result, err := service.GetAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Task 1", result[0].Title)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(test.MockRepository)
	service := service.TodoService(mockRepo)

	taskID := "1"
	updatedTask := &domain.Todo{
		Title:       "Обновленная задача",
		Description: "Описание обновлено",
		Status:      "in_progress",
	}

	mockRepo.On("Update", mock.Anything, taskID, updatedTask).Return(nil)

	err := service.Update(context.Background(), taskID, updatedTask)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_NotFound(t *testing.T) {
	mockRepo := new(test.MockRepository)
	service := service.TodoService(mockRepo)

	taskID := "999" // Несуществующий ID
	updatedTask := &domain.Todo{
		Title:       "Несуществующая задача",
		Description: "Описание",
		Status:      "done",
	}

	mockRepo.On("Update", mock.Anything, taskID, updatedTask).Return(errors.New("task not found"))

	err := service.Update(context.Background(), taskID, updatedTask)

	assert.Error(t, err)
	assert.Equal(t, "task not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(test.MockRepository)
	service := service.TodoService(mockRepo)

	taskID := "1"

	mockRepo.On("Delete", mock.Anything, taskID).Return(nil)

	err := service.Delete(context.Background(), taskID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_NotFound(t *testing.T) {
	mockRepo := new(test.MockRepository)
	service := service.TodoService(mockRepo)

	taskID := "999" // Несуществующий ID

	mockRepo.On("Delete", mock.Anything, taskID).Return(errors.New("task not found"))

	err := service.Delete(context.Background(), taskID)

	assert.Error(t, err)
	assert.Equal(t, "task not found", err.Error())
	mockRepo.AssertExpectations(t)
}
