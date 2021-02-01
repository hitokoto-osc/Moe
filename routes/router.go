package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hitokoto-osc/Moe/config"
	apiV1 "github.com/hitokoto-osc/Moe/controllers/v1"
	"github.com/hitokoto-osc/Moe/middlewares"
	"github.com/hitokoto-osc/Moe/util/web"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime"
)

func setupRoutes(r *gin.Engine) {
	if !viper.IsSet("server.secret") {
		log.Fatal("[web] can't start server because of the secret is not set.")
	}
	r.Use(middlewares.Session(viper.GetString("server.secret")))

	// Setup router
	r.GET("/", func(context *gin.Context) {
		web.Success(context, map[string]interface{}{
			"copyright": "Moe, a lightweight hitokoto status merge tool, authored by MoeTeam. Built with love.",
			"env": map[string]interface{}{
				"go":    runtime.Version(),
				"os":    runtime.GOOS,
				"debug": config.Debug,
			},
			"version": config.Version,
			"build_info": map[string]interface{}{
				"make": config.MakeVersion,
				"hash": config.BuildTag,
				"time": config.BuildTime,
			},
		})
	})

	v1 := r.Group("/v1")
	{
		// protected routes
		// protected := r.Group("/api/v1", middlewares.AuthByMasterKey())
		// {
		// }
		// common routes
		v1.GET("/ping", apiV1.Ping)
		v1.GET("/statistic", apiV1.Statistic)
	}
}
