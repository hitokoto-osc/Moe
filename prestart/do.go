package prestart

import (
	"github.com/hitokoto-osc/Moe/cache"
	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/task"
)

// Do is a func will be called at init, registering the drivers of program
func Do() {
	cache.LoadFromDisk()
	initConfigDriver()
	database.InitDB()
	task.Run()
}
