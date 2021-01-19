package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Collection *cache.Cache = cache.New(5*time.Minute, 10*time.Minute)

func getAPIList () {
	data, ok := Collection.Get("hitokoto_api_server_list")
	if !ok {
		data =
	}
}
