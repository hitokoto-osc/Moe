package cache

import (
	"github.com/bytedance/sonic"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
	"zgo.at/zcache/v2"

	"github.com/hitokoto-osc/Moe/util"
)

// Collection 存储缓存实例
var Collection *zcache.Cache[string, []byte]

// DataFilePath 是缓存记录文件地址
var DataFilePath = filepath.Join(util.MustGetExecDir(), "cache.data")

// LoadFromDisk 用于初始化缓存驱动
func LoadFromDisk() {
	defer zap.L().Sync()
	Collection = zcache.New[string, []byte](5*time.Minute, 10*time.Minute)
	zap.L().Debug("[cache] 加载缓存文件...", zap.String("path", DataFilePath))
	buff, err := os.ReadFile(DataFilePath)
	if err != nil {
		zap.L().Error("无法加载缓存文件。", zap.Error(err))
	}
	var items map[string]zcache.Item[[]byte]
	if err = sonic.Unmarshal(buff, &items); err != nil {
		zap.L().Error("无法解析缓存文件。", zap.Error(err))
	} else {
		Collection = zcache.NewFrom(5*time.Minute, 10*time.Minute, items)
	}
	go func() {
		for {
			time.Sleep(5 * time.Second)
			zap.L().Debug("[cache.loop] 保存缓存文件到本地...", zap.Any("data", Collection.Items()))
			items := Collection.Items()
			buff, err := sonic.Marshal(items)
			if err != nil {
				zap.L().Error("[cache.loop] 无法序列化缓存数据", zap.Error(err))
				continue
			}
			if err = os.WriteFile(DataFilePath, buff, 0700); err != nil {
				zap.L().Error("[cache.loop] 保存缓存文件到本地时发生错误", zap.Error(err))
			}
			zap.L().Sync()
		}
	}()
}
