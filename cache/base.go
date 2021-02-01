package cache

import (
	"github.com/hitokoto-osc/Moe/util"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)
// Collection 存储缓存实例
var Collection *cache.Cache
// DataFilePath 是缓存记录文件地址
var DataFilePath = filepath.Join(util.GetExecDir(), "cache.data")

// Init 用于初始化缓存驱动
func Init() {
	Collection = cache.New(5*time.Minute, 10*time.Minute)
	log.Debug("[cache] 加载缓存文件...")
	if err := Collection.LoadFile(DataFilePath); err != nil {
		log.Error(err)
	}
	go func() {
		for {
			time.Sleep(5 * time.Second)
			log.Debug("[cache.loop] 保存缓存文件到本地...")
			if err := Collection.SaveFile(DataFilePath); err != nil {
				log.Error("[cache.loop] 保存缓存文件到本地时发生错误：" + err.Error())
			}
		}
	}()
}
