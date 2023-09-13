package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hitokoto-osc/Moe/cache"
	"github.com/hitokoto-osc/Moe/task/status/types"
	"github.com/hitokoto-osc/Moe/util/web"
)

// Statistic 是用于返回统计分析结果的控制器
func Statistic(c *fiber.Ctx) error {
	data, ok := cache.GetStatusData()
	if !ok {
		return web.Fail(c, map[string]interface{}{}, -1)
	}
	if data.DownServer == nil {
		data.DownServer = make([]types.DownServerData, 0)
	}
	return web.Success(c, data)
}
