package types

// DownServerData 是程序内部存储宕机记录的类型
type DownServerData struct {
	ID      string `json:"id"`
	StartTS int64  `json:"start_timestamp"`
}

// StatusData 是统计中状态 status 键的结构
type StatusData struct {
	Load     [3]float64 `json:"load"`
	Memory   float64    `json:"memory"`
	Hitokoto struct {
		Total       int      `json:"total"`
		Category    []string `json:"category"`
		LastUpdated int64    `json:"last_updated"`
	} `json:"hitokoto"`
	ChildStatus []ChildStatus `json:"child_status"`
}

// ChildStatus 是统计中 `child_status` 键的结构
type ChildStatus struct {
	Memory struct {
		Total float64 `json:"total"`
		Free  float64 `json:"free"`
		Usage float64 `json:"usage"`
	} `json:"memory"`
	Load     [3]float64 `json:"load"`
	Hitokoto struct {
		Category    []string `json:"category"`
		Total       int      `json:"total"`
		LastUpdated int64    `json:"last_updated"`
	} `json:"hitokoto"`
}

// RequestStruct 是统计中 requests 键的结构
type RequestStruct struct {
	All struct {
		Total          int64   `json:"total"`
		PastMinute     int     `json:"past_minute"`
		PastHour       int     `json:"past_hour"`
		PastDay        int     `json:"past_day"`
		DayMap         [24]int `json:"day_map"`
		FiveMinutesMap [5]int  `json:"five_minutes_map"`
	} `json:"all"`
	Hosts map[string]HostData `json:"hosts"`
}

// GeneratedData 是统计的生成结构
type GeneratedData struct {
	Version     string           `json:"version"`
	Children    []string         `json:"children"`
	DownServer  []DownServerData `json:"down_server"`
	Status      StatusData       `json:"status"`
	Requests    RequestStruct    `json:"requests"`
	LastUpdated int64            `json:"last_updated"`
}

// HostData 定义了 host 请求数结构
type HostData struct {
	Total      int64   `json:"total"`
	PastMinute int     `json:"past_minute"`
	PastHour   int     `json:"past_hour"`
	PastDay    int     `json:"past_day"`
	DayMap     [24]int `json:"day_map"`
}

// APIStatusResponseData 定义了 v1 接口 `host/status` 的响应结构
type APIStatusResponseData struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Message      string `json:"message"`
	Website      string `json:"website"`
	ServerID     string `json:"server_id"`
	ServerStatus struct {
		Memory struct {
			Total float64 `json:"total"`
			Free  float64 `json:"free"`
			Usage float64 `json:"usage"`
		} `json:"memory"`
		Load     [3]float64 `json:"load"`
		Hitokoto struct {
			Category    []string `json:"category"`
			Total       int      `json:"total"`
			LastUpdated int64    `json:"last_updated"`
		} `json:"hitokoto"`
	} `json:"server_status"`
	Requests struct {
		All struct {
			Total          int64   `json:"total"`
			PastMinute     int     `json:"past_minute"`
			PastHour       int     `json:"past_hour"`
			PastDay        int     `json:"past_day"`
			DayMap         [24]int `json:"day_map"`
			FiveMinutesMap [5]int  `json:"five_minutes_map"`
		} `json:"all"`
		Hosts map[string]HostData `json:"hosts"`
	} `json:"requests"`
	Feedback struct {
		Kuertianshi string `json:"Kuertianshi"`
		Freejishu   string `json:"freejishu"`
		A632079     string `json:"a632079"`
	} `json:"feedback"`
	Copyright string `json:"copyright"`
	Now       string `json:"now"`
	Ts        int64  `json:"ts"`
}
