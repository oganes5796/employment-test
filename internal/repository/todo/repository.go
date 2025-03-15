package todo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oganes5796/employment-test/internal/domain"
	"github.com/oganes5796/employment-test/internal/repository"
)

const (
	tableName = "tasks"

	idColumn          = "id"
	titleColumn       = "title"
	descriptionColumn = "description"
	statusColumn      = "status"
	createdColumn     = "created"
	updatedColumn     = "updated"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.TodoRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, task *domain.Todo) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description, status) VALUES ($1, $2, $3) RETURNING %s", tableName, idColumn)
	err := r.db.QueryRow(ctx, query, task.Title, task.Description, task.Status).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *repo) GetAll(ctx context.Context) ([]domain.Todo, error) {
	var tasks []domain.Todo
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var task domain.Todo
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *repo) Update(ctx context.Context, id int, todo *domain.Todo) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if todo.Title != "" {
		setValues = append(setValues, fmt.Sprintf("%s = $%d", titleColumn, argId))
		args = append(args, todo.Title)
		argId++
	}
	if todo.Description != "" {
		setValues = append(setValues, fmt.Sprintf("%s = $%d", descriptionColumn, argId))
		args = append(args, todo.Description)
		argId++
	}
	if todo.Status != "" {
		setValues = append(setValues, fmt.Sprintf("%s = $%d", statusColumn, argId))
		args = append(args, todo.Status)
		argId++
	}

	if len(setValues) == 0 {
		return nil
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = $%d", tableName, strings.Join(setValues, ", "), idColumn, len(args))

	_, err := r.db.Exec(ctx, query, args...)
	return err
}

func (r *repo) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", tableName, idColumn)
	_, err := r.db.Exec(ctx, query, id)
	return err
}
