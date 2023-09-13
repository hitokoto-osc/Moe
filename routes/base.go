package routes

import (
	"github.com/bytedance/sonic"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	recoverMiddleware "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hitokoto-osc/Moe/middlewares"
	"github.com/hitokoto-osc/Moe/util/web"
)

// InitWebServer is a web server register, implemented by gin
func InitWebServer() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "Moe",
		AppName:      "@hitokoto/Moe",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			var e *fiber.Error
			code := fiber.StatusInternalServerError
			if errors.As(err, &e) {
				code = e.Code
			}
			err = web.Fail(ctx, nil, code)
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			return nil
		},
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	// middleware
	app.Use(middlewares.Tracing())
	app.Use(middlewares.Logger())
	app.Use(recoverMiddleware.New())
	app.Use(middlewares.Cors())
	app.Use(middlewares.Favicon())

	// setup routes
	setupRoutes(app)
	return app
}
