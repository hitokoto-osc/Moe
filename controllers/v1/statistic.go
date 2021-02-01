package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hitokoto-osc/Moe/cache"
	"github.com/hitokoto-osc/Moe/task/status/types"
	"github.com/hitokoto-osc/Moe/util/web"
)

// Statistic 是用于返回统计分析结果的控制器
func Statistic(c *gin.Context) {
	data, ok := cache.GetStatusData()
	if !ok {
		web.Fail(c, map[string]interface{}{}, -1)
		return
	}
	if data.DownServer == nil {
		data.DownServer = make([]types.DownServerData, 0)
	}
	web.Success(c, data)
}
