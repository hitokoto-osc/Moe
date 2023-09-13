package web

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"strings"
)

// ParseBearerTokenFromHeader is a func that parse bearer token in authorization header
func ParseBearerTokenFromHeader(ctx *fiber.Ctx) (token string, ok bool) {
	authorization := ctx.Get("Authorization")
	// parse authorization
	if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
		ok = false
		return
	}
	token = strings.ReplaceAll(authorization, "Bearer ", "")
	// base64 decode
	buffer, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Error("无法解析 Authorization 头部的 Bearer Token", zap.Error(err))
		ok = false
		return
	}
	ok = true
	token = string(buffer[:])
	return
}
