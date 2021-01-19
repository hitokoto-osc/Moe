package task

var limitedHost = []string{
	"v1.hitokoto.cn",
	"international.hitokoto.cn",
	"api.a632079.me",
}

type statusData struct {
	Load     [3]float64 `json:"load"`
	Memory   float64    `json:"memory"`
	Hitokoto struct {
		Total       int      `json:"total"`
		Category    []string `json:"category"`
		LastUpdated int64    `json:"last_updated"`
	} `json:"hitokoto"`
	ChildStatus []childStatus `json:"child_status"`
}

type childStatus struct {
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

type requestStruct struct {
	All struct {
		Total          int64   `json:"total"`
		PastMinute     int     `json:"past_minute"`
		PastHour       int     `json:"past_hour"`
		PastDay        int     `json:"past_day"`
		DayMap         [24]int `json:"day_map"`
		FiveMinutesMap [5]int  `json:"five_minutes_map"`
	} `json:"all"`
	Hosts map[string]hostData `json:"hosts"`
}

type GeneratedData struct {
	Version         string           `json:"version"`
	Children        []string         `json:"children"`
	DownServer      []DownServerData `json:"down_server"`
	Status          statusData       `json:"status"`
	Requests        requestStruct    `json:"requests"`
	LastUpdated     int64            `json:"last_updated"`
	ServerTimestamp int64            `json:"Server_timestamp"`
}

type hostData struct {
	Total      int64   `json:"total"`
	PastMinute int     `json:"past_minute"`
	PastHour   int     `json:"past_hour"`
	PastDay    int     `json:"past_day"`
	DayMap     [24]int `json:"day_map"`
}

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
		Hosts map[string]hostData `json:"hosts"`
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
