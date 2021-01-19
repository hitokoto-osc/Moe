package task

import (
	"encoding/json"
	"github.com/blang/semver/v4"
	"github.com/go-resty/resty/v2"
	"time"
)

func RunTask () {
	// 获取 API 列表
}


func requestServerAPI(url string) (data APIStatusResponseData, err error) {
	client := resty.New()
	client.
		// 设置重试逻辑
		SetRetryCount(3). // 重试次数
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(3 * time.Second).
		SetTimeout(2 * time.Second)
	responseData, e := client.
		R().
		EnableTrace().
		Get(url)
	if e != nil {
		err = &GenStatusRequestFailureError{
			Detail:       e.Error(),
			ResponseData: responseData,
		}
		return
	} else if responseData.StatusCode() != 200 {
		err = &GenStatusRequestFailureError{
			Code:         responseData.StatusCode(),
			Detail:       "HTTP Code is not eq 200",
			ResponseData: responseData,
		}
		return
	}
	// status is ok
	var buffer map[string]interface{}
	if e := json.Unmarshal(responseData.Body(), &buffer); e != nil {
		err = &GenStatusRequestFailureError{
			Code:         responseData.StatusCode(),
			Detail:       "JSON parsed failed, detail:" + e.Error(),
			ResponseData: responseData,
		}
		return
	}
	if _, ok := buffer["status"]; ok { // API Throw Error
		err = &GenStatusRequestFailureError{
			Code:         buffer["status"].(int),
			Detail:       buffer["message"].(string),
			ResponseData: responseData,
		}
		return
	}
	// 正常数据
	var authenticResponseData APIStatusResponseData
	if e := json.Unmarshal(responseData.Body(), &authenticResponseData); e != nil {
		err = &GenStatusRequestFailureError{
			Code:         responseData.StatusCode(),
			Detail:       "authentic status JSON parsed failed, detail:" + e.Error(),
			ResponseData: responseData,
		}
		return
	}
	data = authenticResponseData
	return
}

type DownServer struct {
	URL         string
	ID          string
	Active      bool
	UpdatedTime int64
	CreatedTime int64
}

type DownServerData struct {
	ID            string `json:"id"`
	StartTS       int64  `json:"start_ts"`
	LastTime      int64  `json:"last_time"`
	StatusMessage struct {
		IsError bool   `json:"is_error"` // is Error
		ID      string `json:"id"`       // Server_id
		Code    int    `json:"code"`     // StatusCode
		Message string `json:"message"`  // error msg
		Stack   string `json:"stack"`    // error stack
		TS      int64  `json:"ts"`       // current timestamp
	} `json:"status_message"`
}

func genStatusData(inputData []APIStatusResponseData, downServerList []DownServer) (data GeneratedData) {
	var baseHitokotoAPIVersion semver.Version
	data.DownServer = []DownServerData{}
	if downServerList != nil && len(downServerList) > 0 {

	}
	if inputData == nil || len(inputData) == 0 {
		// 抛出异常
		panic("没有节点正常工作.")
	}
	// 合并主体记录
	for i, v := range inputData {
		if i == 0 {
			// 解析版本
			baseHitokotoAPIVersion = semver.MustParse(v.Version)
			initGenData(&data, &v)
		} else {
			// 累加记录
			data.Children = append(data.Children, v.ServerID)
			compareAndUpdateGenDataVersion(&data, &v, &baseHitokotoAPIVersion)
			mergeStatusRecord(&data, &v)
			mergeRequestsRecord(&data, &v)
		}
	}
	data.LastUpdated = time.Now().UnixNano() / 1e6
	return
}
