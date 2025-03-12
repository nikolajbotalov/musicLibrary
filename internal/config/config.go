package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"sync"
)

type Config struct {
	Listen Listen
}

type Listen struct {
	BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
	Port   string `env:"PORT" env-default:"8080"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *zap.Logger) *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			logger.Error("Error reading config", zap.Error(err))
		}
	})

	return instance
}
