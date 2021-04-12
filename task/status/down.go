package status

import (
	"github.com/hitokoto-osc/Moe/cache"
	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/task/status/types"
	"time"
)

var isTest = false
var apiList = func() []database.APIRecord {
	return cache.GetAPIList()
}

// SDownServer 是 []DownServer（请求异常产生的类型） 的抽象方法，提供了一系列的便捷操作
type SDownServer []DownServer

// Exist 用于检索切片中是否存在指定 key 的内容
func (p *SDownServer) Exist(key string) bool {
	for _, v := range *p {
		if v.ID == key {
			return true
		}
	}
	return false
}

// Exclude 用于去除提供 keys 的元素
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

// Convert 可快捷转换类型为 []types.DownServerData
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

// TDownServerList 是 []types.DownServerData 的操作类型
type TDownServerList []types.DownServerData

// DownServerList 是用于外部调用的保存宕机记录的变量
var DownServerList = &TDownServerList{}

func sliceExist(slice []database.APIRecord, key string) bool {
	for _, v := range slice {
		if v.Name == key {
			return true
		}
	}
	return false
}

// Diff 与数据库中的接口记录对比，筛除失效 ID
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

// Merge 用于和新宕机集合对比，更新现有记录，返回新记录
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
		// fmt.Print(isTest)
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

// Exist 查询切片中是否存在给定 id 的记录
func (p *TDownServerList) Exist(id string) bool {
	for _, v := range *p {
		if v.ID == id {
			return true
		}
	}
	return false
}

//  Find 取得给定 ID 的记录
func (p *TDownServerList) Find(id string) (*types.DownServerData, bool) {
	for _, v := range *p {
		if v.ID == id {
			return &v, true
		}
	}
	return nil, false
}

// Remove 移除给定 ID 的记录
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

// Add 向集合添加记录
func (p *TDownServerList) Add(data ...types.DownServerData) {
	*p = append(*p, data...)
	if !isTest {
		p.Save() // 手动触发保存
	}
}

// Unique 集合去重
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
