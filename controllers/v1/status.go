package v1

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/hitokoto-osc/Moe/cache"
	"github.com/hitokoto-osc/Moe/util/web"
	"time"
)

func GetStatus(c *gin.Context) {
	if data, found := cache.Collection.Get("hitokoto_api:status"); !found {
		web.Fail(c, map[string]interface{}{}, 503)
	} else {
		c.JSON(200, map[string]interface{}{
			"code":       200,
			"message":    "ok",
			"data":       data,
			"request_id": requestid.Get(c),
			"ts":         time.Now().UnixNano() / 1e6,
		})
	}
}
