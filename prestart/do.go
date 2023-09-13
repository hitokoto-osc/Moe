package prestart

import (
	"encoding/gob"
	"github.com/hitokoto-osc/Moe/cache"
	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/task"
	"github.com/hitokoto-osc/Moe/task/status"
	"github.com/hitokoto-osc/Moe/task/status/types"
)

// Do is a func will be called at init, registering the drivers of program
func Do() {
	// TODO: 用更好的方法修复缓存读写问题
	gob.Register([]database.APIRecord{})
	gob.Register(types.GeneratedData{})
	gob.Register(status.TDownServerList{})

	cache.LoadFromDisk()
	initConfigDriver()
	database.InitDB()
	task.Run()
}
