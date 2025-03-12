package app

import (
	"go.uber.org/zap"
	"music/internal/config"
	"music/internal/logger"
)

type App struct {
	Logger *zap.Logger
	Config *config.Config
}

func NewApp() (*App, error) {
	// инициализация логгера
	zapLogger, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	// инициализация конфига
	cfg := config.GetConfig(zapLogger)
	zapLogger.Debug("Config loaded")

	return &App{
		Logger: zapLogger,
		Config: cfg,
	}, nil
}
