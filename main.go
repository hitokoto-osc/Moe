package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hitokoto-osc/Moe/config"
	"github.com/hitokoto-osc/Moe/flag"
	"github.com/hitokoto-osc/Moe/logging"
	"github.com/hitokoto-osc/Moe/prestart"
	"github.com/hitokoto-osc/Moe/routes"
	"go.uber.org/zap"
	"runtime"

	"github.com/spf13/viper"
)

var (
	// BuildTag is a commit hash that will be injected in release mode
	BuildTag = "Unknown"
	// BuildTime is a time, when it build, that will be injected in release mode
	BuildTime = "Unknown"
	// CommitTime is a time, when it is committed, that will be injected in release mode
	CommitTime = "Unknown"
	// Version is the version of this program, will be injected in release mode
	Version = "development"
)

var app *fiber.App

func init() {
	// Global set build information
	config.BuildTag = BuildTag
	config.BuildTime = BuildTime
	config.GoVersion = runtime.Version()
	config.Version = Version

	// Parse Flag
	flag.Do()

	if config.Debug {
		logging.GetLogger().Info("Debug mode enabled.")
	}
}

func main() {
	defer zap.L().Sync()
	// Init Drivers
	prestart.Do()

	// init Web Server
	app = routes.InitWebServer()
	// start Server
	if err := app.Listen(":" + viper.GetString("server.port")); err != nil {
		zap.L().Fatal("无法启动服务器", zap.Error(err))
	}
}
