package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// ValidTokenByContext is a func that verify the authorization by bearer token
func ValidTokenByContext(ctx *fiber.Ctx) (ok bool) {
	ok = false
	masterKey := viper.GetString("server.auth.master_key")
	token, isOk := ParseBearerTokenFromHeader(ctx)
	if !isOk && token == masterKey { // fallback: support token that was not encoded by base64
		ok = true
	} else if token == masterKey {
		ok = true
	}
	return
}
