package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Cors is a middleware func that solve the issue of CORS response
func Cors() fiber.Handler {
	config := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS",
		AllowHeaders:     "Origin, Content-Length, Content-Type, Accept, Cookie",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
	return cors.New(config)
}
