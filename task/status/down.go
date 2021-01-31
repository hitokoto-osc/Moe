package status

import (
	"fmt"
	"github.com/hitokoto-osc/Moe/cache"
	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/task/status/types"
	"time"
)

var apiList = func () []database.APIRecord {
	return cache.GetAPIList()
}
var isTest bool = false

type SDownServer []DownServer

func (p *SDownServer) Exist(key string) bool {
	for _, v := range *p {
		if v.ID == key {
			return true
		}
	}
	return false
}

func (p *SDownServer) Exclude(keys []string) {
	for _, key := range keys {
		t := 0
		for _, v := range *p {
			if v.ID != key {
				(*p)[t] = v
				t++
			}
		}
		*p = (*p)[:t]
	}
}

func (p *SDownServer) Convert() []types.DownServerData {
	var tmp []types.DownServerData
	for _, v := range *p {
		tmp = append(tmp, types.DownServerData{
			ID:      v.ID,
			StartTS: v.StartTS,
		})
	}
	return tmp
}

type TDownServerList []types.DownServerData

var DownServerList = &TDownServerList{}

func sliceExist(slice []database.APIRecord, key string) bool {
	for _, v := range slice {
		if v.Name == key {
			return true
		}
	}
	return false
}

func (p *TDownServerList) Diff(list []database.APIRecord) {
	t := 0
	for _, v := range *p {
		if sliceExist(list, v.ID) {
			(*p)[t] = v
			t++
		}
	}
	*p = (*p)[:t]
}

func (p *TDownServerList) Merge(newCollection SDownServer) []types.DownServerData {
	p.Diff(apiList()) // 清除失效的节点
	// 保留仍然宕机的节点
	t := 0
	var s []string
	for _, v := range *p {
		if newCollection.Exist(v.ID) {
			(*p)[t] = v
			t++
			s = append(s, v.ID)
		}
	}
	*p = (*p)[:t]            // 移除恢复的节点
	newCollection.Exclude(s) // 移除已存在的故障节点
	// 合并
	for _, v := range newCollection {
		p.Add(types.DownServerData{
			ID:      v.ID,
			StartTS: v.StartTS,
		})
	}
	if !isTest {
		fmt.Print(isTest)
		p.Save() // 手动触发保存
	}
	return *p
}

// Save 用于保存到缓存
func (p *TDownServerList) Save() {
	cache.Collection.Set("down", *p, 24*time.Hour)
}

// Recover 函数用于从缓存种恢复
func (p *TDownServerList) Recover() {
	if data, ok := cache.Collection.Get("down"); ok {
		*p = data.(TDownServerList)
	}
}

func (p *TDownServerList) Exist(id string) bool {
	for _, v := range *p {
		if v.ID == id {
			return true
		}
	}
	return false
}

func (p *TDownServerList) Find(id string) (*types.DownServerData, bool) {
	for _, v := range *p {
		if v.ID == id {
			return &v, true
		}
	}
	return nil, false
}

func (p *TDownServerList) Remove(id string) {
	t := 0
	for _, v := range *p {
		if v.ID != id {
			(*p)[t] = v
			t++
		}
	}
	*p = (*p)[:t]
	if !isTest {
		p.Save() // 手动触发保存
	}
}

func (p *TDownServerList) Add(data ...types.DownServerData) {
	*p = append(*p, data...)
	if !isTest {
		p.Save() // 手动触发保存
	}

}

func (p *TDownServerList) Unique() {
	if len(*p) < 1024 {
		*p = uniqueByLoop(*p)
	} else {
		*p = uniqueByMap(*p)
	}
	if !isTest {
		p.Save() // 手动触发保存
	}
}

func uniqueByLoop(slc TDownServerList) TDownServerList {
	result := TDownServerList{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为 false
				break
			}
		}
		if flag { // 标识为 false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

func uniqueByMap(slc TDownServerList) TDownServerList {
	result := TDownServerList{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e.ID] = 0
		if len(tempMap) != l { // 加入 map 后，map 长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
