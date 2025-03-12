package app

import (
	"go.uber.org/zap"
	"music/internal/logger"
)

type App struct {
	Logger *zap.Logger
}

func NewApp() (*App, error) {
	zapLogger, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	return &App{
		Logger: zapLogger,
	}, nil
}
