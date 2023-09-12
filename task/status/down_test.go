package status

import (
	"testing"

	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/task/status/types"
	"github.com/stretchr/testify/assert"
)

func TestSDownServerExist(t *testing.T) {
	isTest = true
	s := SDownServer{
		DownServer{
			ID:                             "test_id_1",
			StartTS:                        1612091723908,
			Cause:                          "Test",
			Error:                          nil,
			IsGenStatusRequestFailureError: false,
		},
	}
	assert.True(t, s.Exist("test_id_1"), "`test_id_1` 应该存在")
	assert.False(t, s.Exist("test_id_2"), "`test_id_2` 应该不存在")
}

func TestSDownServerExclude(t *testing.T) {
	isTest = true
	s := SDownServer{
		DownServer{
			ID:                             "test_id_1",
			StartTS:                        1612091723908,
			Cause:                          "Test",
			Error:                          nil,
			IsGenStatusRequestFailureError: false,
		},
	}
	s.Exclude([]string{"test_id_1"})
	assert.Len(t, s, 0, "移除 `test_id_1` 后切片长度应为 0")
}

func TestSDownServerConvert(t *testing.T) {
	isTest = true
	s := SDownServer{
		DownServer{
			ID:                             "test_id_1",
			StartTS:                        1612091723908,
			Cause:                          "Test",
			Error:                          nil,
			IsGenStatusRequestFailureError: false,
		},
	}
	target := s.Convert()
	assert.IsType(t, []types.DownServerData{}, target, "转换后类型应为 []types.DownServerData")
	assert.Len(t, target, 1, "转换后目标数组长度应为 1")
	assert.Equal(t, target[0].StartTS, int64(1612091723908), "转换后的切片引索 0 的 StartTS 值应为 1612091723908")
	assert.Equal(t, target[0].ID, "test_id_1", "转换后的切片引索 0 的 ID 值应为 `test_id_1`")
}

func TestTDownServerListDiff(t *testing.T) {
	isTest = true
	s := TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612091723908,
		},
		types.DownServerData{
			ID:      "test_id_2",
			StartTS: 1612091723908,
		},
	}
	c := []database.APIRecord{
		database.APIRecord{
			ID:   0,
			Name: "test_id_1",
			URL:  "",
			Desc: "",
		},
		database.APIRecord{
			ID:   1,
			Name: "test_id_3",
			URL:  "",
			Desc: "",
		},
	}
	s.Diff(c)
	assert.ElementsMatch(t, s, TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612091723908,
		},
	}, "Diff 结果应该与示例相同")
}

func TestTDownServerListMerge(t *testing.T) {
	isTest = true
	apiList = func() []database.APIRecord {
		return []database.APIRecord{
			database.APIRecord{
				Name: "test_id_1",
			},
			database.APIRecord{
				Name: "test_id_2",
			},
			database.APIRecord{
				Name: "test_id_3",
			},
		}
	}
	s := TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
		types.DownServerData{
			ID:      "test_id_2",
			StartTS: 1612101075609,
		},
	}
	c := SDownServer{
		DownServer{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
		DownServer{
			ID:      "test_id_3",
			StartTS: 1612101075609,
		},
	}
	s.Merge(c)
	assert.ElementsMatch(t, s, TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
		types.DownServerData{
			ID:      "test_id_3",
			StartTS: 1612101075609,
		},
	}, "执行合并操作后，切片应该相同")
}

func TestTDownServerListExist(t *testing.T) {
	isTest = true
	s := TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
	}
	assert.True(t, s.Exist("test_id_1"), "`test_id_1` 存在")
	assert.False(t, s.Exist("test_id_2"), "`test_id_2` 不存在")
}

func TestTDownServerListFind(t *testing.T) {
	isTest = true
	s := TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
	}
	data, ok := s.Find("test_id_1")
	assert.True(t, ok, "`test_id_1` 存在")
	assert.Equal(t, data.ID, "test_id_1", "结构体内容应相同")
}

func TestSDownServerRemove(t *testing.T) {
	isTest = true
	s := TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
	}
	s.Remove("test_id_1")
	assert.Len(t, s, 0, "删除记录后切片长度应为 0")
}

func TestSDownServerAdd(t *testing.T) {
	isTest = true
	s := TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
	}
	s.Add(types.DownServerData{
		ID:      "test_id_2",
		StartTS: 1612101075609,
	})
	assert.ElementsMatch(t, s, TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
		types.DownServerData{
			ID:      "test_id_2",
			StartTS: 1612101075609,
		},
	}, "添加元素后，内容应和预期一致")
}

func TestSDownServerUnique(t *testing.T) {
	isTest = true
	s := TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
	}
	s.Unique()
	assert.ElementsMatch(t, s, TDownServerList{
		types.DownServerData{
			ID:      "test_id_1",
			StartTS: 1612101075609,
		},
	}, "去重后，内容应和预期一致")
}
