package cache

import (
	"github.com/bytedance/sonic"
	"github.com/cockroachdb/errors"
	"github.com/hitokoto-osc/Moe/logging"
	"time"

	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/task/status/types"
)

// MustStoreStatusData 存储统计结果
func MustStoreStatusData(data any) {
	buff, err := sonic.Marshal(data)
	if err != nil {
		panic(errors.Wrap(err, "无法序列化缓存数据"))
	}
	Collection.SetWithExpire("status_data", buff, 30*time.Minute)
}

// MustGetStatusData 获得缓存中的统计结果
func MustGetStatusData() (*types.GeneratedData, bool) {
	buff, ok := Collection.Get("status_data")
	if !ok {
		return nil, ok
	}
	var data types.GeneratedData
	err := sonic.Unmarshal(buff, &data)
	if err != nil {
		panic(errors.Wrap(err, "无法反序列化缓存数据"))
	}
	return &data, ok
}

// MustGetAPIList 用于获取 API 记录
// 此为快捷方法，如果缓存中为空会拉取数据库再写缓存
func MustGetAPIList() []database.APIRecord {
	var (
		data []database.APIRecord
		err  error
	)
	logger := logging.GetLogger()
	defer logger.Sync()
	buff, ok := Collection.Get("hitokoto_api_server_list")
	if !ok {
		data, err = database.GetHitokotoAPIHostList()
		if err != nil {
			panic(errors.Wrap(err, "无法获取 API 列表"))
		}
		buff, err = sonic.Marshal(data)
		if err != nil {
			panic(errors.Wrap(err, "无法序列化 API 列表"))
		}
		Collection.Set("hitokoto_api_server_list", buff)
	} else {
		err = sonic.Unmarshal(buff, &data)
		if err != nil {
			panic(errors.Wrap(err, "无法反序列化 API 列表"))
		}
	}
	return data
}
