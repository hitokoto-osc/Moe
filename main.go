package main

import (
	"encoding/gob"
	"github.com/hitokoto-osc/Moe/logging"
	"go.uber.org/zap"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/hitokoto-osc/Moe/config"
	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/flag"
	"github.com/hitokoto-osc/Moe/prestart"
	"github.com/hitokoto-osc/Moe/routes"
	"github.com/hitokoto-osc/Moe/task/status"
	"github.com/hitokoto-osc/Moe/task/status/types"
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

var r *gin.Engine

func init() {
	// TODO: 用更好的方法修复缓存读写问题
	gob.Register([]database.APIRecord{})
	gob.Register(types.GeneratedData{})
	gob.Register(status.TDownServerList{})

	// Global set build information
	config.BuildTag = BuildTag
	config.BuildTime = BuildTime
	config.GoVersion = runtime.Version()
	config.Version = Version

	// Parse Flag
	flag.Parse()

	if config.Debug {
		logging.GetLogger().Info("Debug mode enabled.")
	}

	// Init Drivers
	prestart.Do()

	// init Web Server
	r = routes.InitWebServer()
}

func main() {
	defer zap.L().Sync()
	// start Server
	if err := r.Run(":" + viper.GetString("server.port")); err != nil {
		zap.L().Fatal("无法启动服务器", zap.Error(err))
	}
}
