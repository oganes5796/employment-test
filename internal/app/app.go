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

// @title Todo API
// @version 1.0
// @description This is a sample server for a todo application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func Run(ctx context.Context) error {
	// TODO: init config
	err := config.Load(".env")
	if err != nil {
		return err
	}

	// TODO: init logger
	cfg := config.DefaultLoggerConfig()
	logger.InitLogger(cfg)
	defer logger.Sync()

	// TODO: init db
	pool, err := db.NewPostgresPool(os.Getenv("PG_DSN"))
	if err != nil {
		return err
	}
	defer pool.Close()

	// TODO: init repository
	repo := repoTodo.NewRepository(pool)

	// TODO: init service
	service := servTodo.NewService(repo)

	// TODO: init handler
	handler := api.NewImplementation(service)

	// TODO: init router
	app := handler.InitRoutes()

	logger.Info("Server is running on port" + os.Getenv("LOCAL_PORT"))
	err = app.Listen(":" + os.Getenv("LOCAL_PORT"))
	if err != nil {
		return err
	}

	return nil
}
