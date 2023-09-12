package web

import (
	"encoding/base64"
	"github.com/hitokoto-osc/Moe/logging"
	"go.uber.org/zap"
	"strings"

	"github.com/gin-gonic/gin"
)

// ParseBearerTokenFromHeader is a func that parse bearer token in authorization header
func ParseBearerTokenFromHeader(ctx *gin.Context) (token string, ok bool) {
	logger := logging.WithContext(ctx)
	defer logger.Sync()
	authorization := ctx.GetHeader("Authorization")
	// parse authorization
	if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
		ok = false
		return
	}
	token = strings.ReplaceAll(authorization, "Bearer ", "")
	// base64 decode
	buffer, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		logger.Error("无法解析 Authorization 头部的 Bearer Token", zap.Error(err))
		ok = false
		return
	}
	ok = true
	token = string(buffer[:])
	return
}
