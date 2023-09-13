package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/hitokoto-osc/Moe/assets"
	"go.uber.org/zap"
)

// Favicon is a middleware for favicon.ico
func Favicon() fiber.Handler {
	data, err := assets.FaviconIco.ReadFile("favicon.ico")
	if err != nil {
		zap.L().Fatal("failed to load favicon", zap.Error(err))
	}
	return favicon.New(favicon.Config{
		Data: data,
	})
}
