package middlewares

import (
	"github.com/sqlc_test/config"
	"go.uber.org/zap"
)

type Middleware struct {
	config config.AppConfig
	logger *zap.Logger
}

func NewMiddleware(cfg config.AppConfig, logger *zap.Logger) Middleware {
	return Middleware{
		config: cfg,
		logger: logger,
	}
}
