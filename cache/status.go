package cache

import (
	"time"

	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/task/status/types"
	log "github.com/sirupsen/logrus"
)

// StoreStatusData 存储统计结果
// TODO: 替换 interface{}
func StoreStatusData(data interface{}) {
	Collection.Set("status_data", data, 30*time.Minute)
}

// GetStatusData 获得缓存中的统计结果
func GetStatusData() (*types.GeneratedData, bool) {
	data, ok := Collection.Get("status_data")
	if !ok {
		return nil, ok
	}
	r, ok := data.(types.GeneratedData)
	return &r, ok
}

// GetAPIList 用于获取 API 记录
// 此为快捷方法，如果缓存中为空会拉取数据库再写缓存
func GetAPIList() []database.APIRecord {
	var tmp interface{}
	tmp, ok := Collection.Get("hitokoto_api_server_list")
	if !ok {
		var err error
		if tmp, err = database.GetHitokotoAPIHostList(); err != nil {
			log.Fatal(err)
		}
		if err = Collection.Add("hitokoto_api_server_list", tmp, 3*time.Minute); err != nil {
			log.Fatal(err)
		}
	}
	return tmp.([]database.APIRecord)
}
