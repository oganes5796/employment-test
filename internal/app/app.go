package app

import (
	"context"
	"os"

	_ "github.com/oganes5796/employment-test/docs"
	"github.com/oganes5796/employment-test/internal/api"
	"github.com/oganes5796/employment-test/internal/client/db"
	"github.com/oganes5796/employment-test/internal/config"
	repoTodo "github.com/oganes5796/employment-test/internal/repository/todo"
	servTodo "github.com/oganes5796/employment-test/internal/service/todo"
	"github.com/oganes5796/employment-test/pkg/logger"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8080
// @BasePath /

func Run(ctx context.Context) error {
	// TODO: init config
	err := config.Load(".env")
	if err != nil {
		return err
	}

	cfg := config.DefaultLoggerConfig()
	logger.InitLogger(cfg)
	defer logger.Sync()

	pool, err := db.NewPostgresPool(os.Getenv("PG_DSN"))
	if err != nil {
		return err
	}
	defer pool.Close()

	repo := repoTodo.NewRepository(pool)

	service := servTodo.NewService(repo)

	handler := api.NewImplementation(service)

	app := handler.InitRoutes()

	logger.Info("Server is running on port" + os.Getenv("LOCAL_PORT"))
	err = app.Listen(":" + os.Getenv("LOCAL_PORT"))
	if err != nil {
		return err
	}

	return nil
}
