package web

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

var errorMessageMap = map[int]string{
	fiber.StatusBadRequest:          "Bad request.",
	fiber.StatusUnauthorized:        "Unauthorized.",
	fiber.StatusNotFound:            "Not found specific route/file.",
	fiber.StatusForbidden:           "Forbidden.",
	fiber.StatusMethodNotAllowed:    "Method not permitted.",
	fiber.StatusInternalServerError: "Server error. If occurs frequently, please contact the author.",
	fiber.StatusServiceUnavailable:  "service Unavailable.",
	// Custom Error
	-1: "The Status Data is not found. Please try again later.",
}

// Success is a func that do the common situation of responding successful formation
func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"code":       200,
		"message":    "ok",
		"data":       data,
		"request_id": ctx.Locals("request_id"),
		"ts":         time.Now().UnixNano() / 1e6,
	})
}

// Fail is a func that do the common situation of responding failed formation
func Fail(ctx *fiber.Ctx, data interface{}, code int) error {
	var status int
	if code <= 0 {
		status = fiber.StatusOK
	} else {
		status = code
	}
	msg, ok := errorMessageMap[code]
	if !ok {
		msg = "Unknown status code, please contact author."
	}
	return ctx.Status(status).JSON(map[string]interface{}{
		"code":       code,
		"message":    msg,
		"data":       data,
		"request_id": ctx.Locals("request_id"),
		"ts":         time.Now().UnixNano() / 1e6,
	})
}
