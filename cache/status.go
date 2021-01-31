package cache

import (
	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/task/status/types"
	log "github.com/sirupsen/logrus"
	"time"
)

// TODO: 替换 interface{}
func StoreStatusData(data interface{}) {
	Collection.Set("status_data", data, 30*time.Minute)
}

func GetStatusData() (*types.GeneratedData, bool) {
	data, ok := Collection.Get("status_data")
	if !ok {
		return nil, ok
	}
	r, ok := data.(types.GeneratedData)
	return &r, ok
}

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
