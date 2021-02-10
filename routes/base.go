package routes

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/hitokoto-osc/Moe/config"
	"github.com/hitokoto-osc/Moe/middlewares"
	"github.com/hitokoto-osc/Moe/util/web"
)

// InitWebServer is a web server register, implemented by gin
func InitWebServer() *gin.Engine {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// load middleware
	r.Use(requestid.New())
	r.Use(middlewares.Cors())

	// 404
	r.NoRoute(func(context *gin.Context) {
		context.Status(404)
		web.Fail(context, nil, 404)
		return
	})

	// 405
	r.NoMethod(func(context *gin.Context) {
		context.Status(405)
		web.Fail(context, nil, 405)
		return
	})

	// setup routes
	setupRoutes(r)
	return r
}
