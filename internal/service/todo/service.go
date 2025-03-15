package todo

import (
	"context"
	"strconv"

	"github.com/oganes5796/employment-test/internal/domain"
	"github.com/oganes5796/employment-test/internal/repository"
	"github.com/oganes5796/employment-test/internal/service"
)

type serv struct {
	todoRepository repository.TodoRepository
}

func NewService(todoRepository repository.TodoRepository) service.TodoService {
	return &serv{todoRepository: todoRepository}
}

func (s *serv) Create(ctx context.Context, todo *domain.Todo) (int, error) {
	return s.todoRepository.Create(ctx, todo)
}

func (s *serv) GetAll(ctx context.Context) ([]domain.Todo, error) {
	return s.todoRepository.GetAll(ctx)
}

func (s *serv) Update(ctx context.Context, id string, todo *domain.Todo) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return s.todoRepository.Update(ctx, idInt, todo)
}

func (s *serv) Delete(ctx context.Context, id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return s.todoRepository.Delete(ctx, idInt)
}
