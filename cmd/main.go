package main

import (
	"context"

	"github.com/oganes5796/employment-test/internal/app"
	"github.com/oganes5796/employment-test/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	err := app.Run(ctx)
	if err != nil {
		logger.Error("failed to run app", zap.Error(err))
	}
}
