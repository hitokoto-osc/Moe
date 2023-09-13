package middlewares

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hitokoto-osc/Moe/consts"
)

// Tracing is a middleware func that set tracing id to request header
func Tracing() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tracingID := uuid.NewString()
		ctx.Set("X-Request-ID", tracingID)
		ctx.Locals("request_id", tracingID)
		// set tracing id to user context
		c := context.WithValue(ctx.UserContext(), consts.ContextKeyRequestID, tracingID)
		ctx.SetUserContext(c)
		return ctx.Next()
	}
}
