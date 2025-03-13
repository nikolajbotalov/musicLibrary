package app

import (
	"go.uber.org/zap"
	"music/internal/config"
	"music/internal/logger"
)

type App struct {
	Logger *zap.Logger
	Config *config.Config
	Server *Server
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

	// инициализация сервера
	server := NewServer(cfg, zapLogger)

	return &App{
		Logger: zapLogger,
		Config: cfg,
		Server: server,
	}, nil
}

func (a *App) Close() {
	a.Logger.Info("Closing application")
}
