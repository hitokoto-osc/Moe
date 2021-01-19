package task

import (
	"github.com/blang/semver/v4"
)

func compareAndUpdateGenDataVersion(data *GeneratedData, t *APIStatusResponseData, baseVersion *semver.Version) {
	v := semver.MustParse(t.Version)
	if v.GT(*baseVersion) {
		*baseVersion = v
		data.Version = t.Version
	}
}

func mergeStatusRecord(data *GeneratedData, v *APIStatusResponseData) {
	// Load
	for i := range data.Status.Load {
		data.Status.Load[i] += v.ServerStatus.Load[i]
	}
	// 内存占用
	data.Status.Memory += v.ServerStatus.Memory.Usage
	// 一言部分
	if v.ServerStatus.Hitokoto.LastUpdated > data.Status.Hitokoto.LastUpdated {
		data.Status.Hitokoto.LastUpdated = v.ServerStatus.Hitokoto.LastUpdated
		data.Status.Hitokoto.Total = v.ServerStatus.Hitokoto.Total
		data.Status.Hitokoto.Category = v.ServerStatus.Hitokoto.Category
	}
	// 加入节点统计信息
	data.Status.ChildStatus = append(data.Status.ChildStatus, v.ServerStatus)
}

func mergeRequestsRecord(data *GeneratedData, v *APIStatusResponseData) {
	// ALL
	data.Requests.All.Total += v.Requests.All.Total
	data.Requests.All.PastMinute += v.Requests.All.PastMinute
	data.Requests.All.PastHour += v.Requests.All.PastHour
	data.Requests.All.PastDay += v.Requests.All.PastDay
	for i := range v.Requests.All.DayMap {
		data.Requests.All.DayMap[i] = v.Requests.All.DayMap[i]
	}
	for i := range v.Requests.All.FiveMinutesMap {
		data.Requests.All.FiveMinutesMap[i] += v.Requests.All.FiveMinutesMap[i]
	}

	// 合并 HOST 请求数
	for _, host := range limitedHost {
		if hostData, ok := v.Requests.Hosts[host]; ok {
			t, o := data.Requests.Hosts[host]
			if !o {
				t = hostData
			}
			t.Total += hostData.Total
			t.PastMinute += hostData.PastMinute
			t.PastHour += hostData.PastHour
			t.PastDay += hostData.PastDay
			for i := range hostData.DayMap {
				t.DayMap[i] += hostData.DayMap[i]
			}
		}
	}
}

func initGenData(data *GeneratedData, v *APIStatusResponseData) {
	*data = GeneratedData{
		Version:  v.Version,
		Children: []string{v.ServerID},
		Status: statusData{
			Load:   v.ServerStatus.Load,
			Memory: v.ServerStatus.Memory.Usage,
			Hitokoto: struct {
				Total       int      `json:"total"`
				Category    []string `json:"category"`
				LastUpdated int64    `json:"last_updated"`
			}{
				Total:       v.ServerStatus.Hitokoto.Total,
				Category:    v.ServerStatus.Hitokoto.Category,
				LastUpdated: v.ServerStatus.Hitokoto.LastUpdated,
			},
			ChildStatus: []childStatus{
				{
					Memory:   v.ServerStatus.Memory,
					Load:     v.ServerStatus.Load,
					Hitokoto: v.ServerStatus.Hitokoto,
				},
			},
		},
		Requests: struct {
			All struct {
				Total          int64   `json:"total"`
				PastMinute     int     `json:"past_minute"`
				PastHour       int     `json:"past_hour"`
				PastDay        int     `json:"past_day"`
				DayMap         [24]int `json:"day_map"`
				FiveMinutesMap [5]int  `json:"five_minutes_map"`
			} `json:"all"`
			Hosts map[string]hostData `json:"hosts"`
		}{
			All:   v.Requests.All,
			Hosts: v.Requests.Hosts,
		},
		LastUpdated:     0,
		ServerTimestamp: 0,
	}
}
