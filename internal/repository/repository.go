package repository

import (
	"context"

	"github.com/oganes5796/employment-test/internal/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, task *domain.Todo) (int, error)
	GetAll(ctx context.Context) ([]domain.Todo, error)
	Update(ctx context.Context, id int, task *domain.Todo) error
	Delete(ctx context.Context, id int) error
}
