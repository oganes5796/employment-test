package test

import (
	"context"

	"github.com/oganes5796/employment-test/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx context.Context, todo *domain.Todo) (int, error) {
	args := m.Called(ctx, todo)
	return 0, args.Error(0)
}

func (m *MockRepository) GetAll(ctx context.Context) ([]domain.Todo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Todo), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, id string, todo *domain.Todo) error {
	args := m.Called(ctx, id, todo)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
