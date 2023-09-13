package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hitokoto-osc/Moe/util/web"
)

// Ping is a controller func that impl a Pong response,
// intended to notify the client that server is ok
func Ping(c *fiber.Ctx) error {
	return web.Success(c, map[string]interface{}{})
}
