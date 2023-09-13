package cache

import (
	"go.uber.org/zap"
	"path/filepath"
	"time"

	"github.com/hitokoto-osc/Moe/util"
	"github.com/patrickmn/go-cache"
)

// Collection 存储缓存实例
var Collection *cache.Cache

// DataFilePath 是缓存记录文件地址
var DataFilePath = filepath.Join(util.MustGetExecDir(), "cache.data")

// LoadFromDisk 用于初始化缓存驱动
func LoadFromDisk() {
	defer zap.L().Sync()

	Collection = cache.New(5*time.Minute, 10*time.Minute)
	zap.L().Debug("[cache] 加载缓存文件...")
	if err := Collection.LoadFile(DataFilePath); err != nil {
		zap.L().Error("无法加载缓存文件。", zap.Error(err))
	}
	go func() {
		for {
			time.Sleep(5 * time.Second)
			zap.L().Debug("[cache.loop] 保存缓存文件到本地...")
			if err := Collection.SaveFile(DataFilePath); err != nil {
				zap.L().Error("[cache.loop] 保存缓存文件到本地时发生错误", zap.Error(err))
			}
		}
	}()
}
