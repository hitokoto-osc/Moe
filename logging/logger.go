package logging

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const loggerKey = "logger"

var logger *zap.Logger

// GetLogger returns a global logger
func GetLogger() *zap.Logger {
	return logger
}

func setZapGlobalLogger() {
	zap.ReplaceGlobals(logger)
}

// NewContext add some fields to logger and set it to gin.Context
func NewContext(ctx context.Context, fields ...zap.Field) {
	if ctx == nil {
		logger.Panic("context is nil")
	}
	gctx, ok := ctx.(*gin.Context)
	if !ok {
		logger.Panic("context is not *gin.Context")
	}
	gctx.Set(loggerKey, WithContext(ctx).With(fields...))
}

// WithContext return a logger from gin.Context or global logger
func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	gctx, ok := ctx.(*gin.Context)
	if !ok {
		return logger
	}
	l, ok := gctx.Get(loggerKey)
	if !ok {
		return logger
	}
	ctxLogger, ok := l.(*zap.Logger)
	if !ok {
		return logger
	}
	return ctxLogger
}
