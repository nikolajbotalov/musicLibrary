package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"music/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
	logger     *zap.Logger
}

// NewServer Создает и возвращает новый HTTP сервер
func NewServer(cfg *config.Config, logger *zap.Logger) *Server {
	router := gin.New()

	router.Use(ginZapLogger(logger))
	router.Use(getRecoveryWithLogging(logger))

	// создание HTTP сервера
	address := fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	return &Server{
		httpServer: &http.Server{
			Addr:    address,
			Handler: router,
		},
		logger: logger,
	}
}

// Run запускает сервер с graceful shutdown
func (s *Server) Run() {
	s.logger.Info("Starting server", zap.String("address", s.httpServer.Addr))

	// канал для graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("Server error", zap.Error(err))
		}
	}()

	// ожидание сигнала для graceful shutdown
	<-done
	s.logger.Info("Shutting down server")

	// завершение работы с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("Server shutdown error", zap.Error(err))
	}

	s.logger.Info("Application stopped")
}

// возвращает middleware для Gin, который использует Zap для логирования
func ginZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logger.Info("Request handled",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			zap.Duration("duration", time.Since(start)),
			zap.Int("response_size", c.Writer.Size()),
		)
	}
}

// возвращает middleware для обработки panic с логгированием
func getRecoveryWithLogging(logger *zap.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		logger.Error("Recovered from panic", zap.Any("error", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	})
}
