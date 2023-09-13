package middlewares

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/hitokoto-osc/Moe/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a middleware that logs the request and response
func Logger() fiber.Handler {
	logger := logging.GetLogger().WithOptions(zap.AddCallerSkip(2))
	cfg := fiberzap.Config{
		Logger: logger,
		Fields: []string{"latency", "status", "method", "url"},
		FieldsFunc: func(c *fiber.Ctx) []zap.Field {
			return []zap.Field{
				zap.String("request_id", c.Locals("request_id").(string)),
			}
		},
		Levels: []zapcore.Level{zapcore.ErrorLevel, zapcore.WarnLevel, zapcore.DebugLevel},
	}
	fn := fiberzap.New(cfg)
	log.SetLogger(fiberzap.NewLogger(fiberzap.LoggerConfig{
		SetLogger: logger,
		ExtraKeys: []string{"request_id"},
	}))
	return fn
}
