package main

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/hitokoto-osc/Moe/config"
	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/flag"
	"github.com/hitokoto-osc/Moe/prestart"
	"github.com/hitokoto-osc/Moe/routes"
	"github.com/hitokoto-osc/Moe/task/status/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime"
)

var (
	// BuildTag is a commit hash that will be injected in release mode
	BuildTag = "Unknown"
	// BuildTime is a time, when it build, that will be injected in release mode
	BuildTime = "Unknown"
	// Version is the version of this program, will be injected in release mode
	Version = "development"
)

var r *gin.Engine

func init() {
	// TODO: 用更好的方法修复缓存读写问题
	gob.Register([]database.APIRecord{})
	gob.Register(types.GeneratedData{})

	// Global set build information
	config.BuildTag = BuildTag
	config.BuildTime = BuildTime
	config.GoVersion = runtime.Version()
	config.Version = Version

	// Parse Flag
	flag.Parse()

	// Init Drivers
	prestart.Do()

	if config.Debug {
		log.Info("[debug] 已启用调试模式")
	}

	// init Web Server
	r = routes.InitWebServer()
}

func main() {
	// start Server
	if err := r.Run(":" + viper.GetString("server.port")); err != nil {
		log.Fatal(err)
	}
}
