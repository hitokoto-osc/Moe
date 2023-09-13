package routes

import (
	"github.com/gofiber/fiber/v2"
	"runtime"

	"github.com/hitokoto-osc/Moe/config"
	apiV1 "github.com/hitokoto-osc/Moe/controllers/v1"
	"github.com/hitokoto-osc/Moe/util/web"
)

var osInfo = runtime.GOOS + " " + runtime.GOARCH

func setupRoutes(app *fiber.App) {
	//logger := logging.GetLogger()

	//if !viper.IsSet("server.secret") {
	//	logger.Fatal("[web] can't start server because of the secret is not set.")
	//}
	//r.Use(middlewares.Session(viper.GetString("server.secret")))

	// Setup router
	app.Get("/", func(ctx *fiber.Ctx) error {
		return web.Success(ctx, map[string]interface{}{
			"build_info": map[string]interface{}{
				"version":      config.Version,
				"commit_hash":  config.BuildTag,
				"commit_time":  config.BuildTime,
				"generated_by": runtime.Version(),
				"os":           osInfo,
				"debug":        config.Debug,
			},
			"donate":    "Love us? donate at https://hitokoto.cn/donate",
			"copyright": "Moe, a lightweight hitokoto status merge tool, authored by MoeTeam. Built with love.",
			"feedback": map[string]interface{}{
				"tips": "if you find anything that not works as expected, you can contact us directly. Thanks.",
				"email": []string{
					"i@loli.oneline",
					"i@freejishu.com",
					"i@a632079.me",
				},
				"qq": map[string]int{
					"group":  33542648,
					"person": 442971704,
				},
			},
		})
	})

	v1 := app.Group("/v1")
	{
		// protected routes
		// protected := r.Group("/api/v1", middlewares.AuthByMasterKey())
		// {
		// }
		// common routes
		v1.Get("/ping", apiV1.Ping)
		v1.Get("/statistic", apiV1.Statistic)
	}
}
