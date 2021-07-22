package status

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/hitokoto-osc/Moe/database"
	"github.com/hitokoto-osc/Moe/task/status/types"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
	"sync"
	"time"
)

func performRequest(records []database.APIRecord) (data []types.APIStatusResponseData, downList []DownServer) {
	wg := sync.WaitGroup{} // 使用 waitGroup 控制并发
	wg.Add(len(records))
	downListMutex := sync.Mutex{}
	dataMutex := sync.Mutex{}
	for _, record := range records {
		record := record // 此行不能去除
		go func() {
			id := record.Name
			url := record.URL + "/status"
			result, err := requestServerAPI(url)
			if err != nil {
				if e, ok := err.(*GenStatusRequestFailureError); ok {
					// 添加到 DownServer 列表
					downListMutex.Lock()
					downList = append(downList, DownServer{
						ID:                             id,
						StartTS:                        time.Now().UnixNano() / 1e6,
						Cause:                          e.Detail,
						Error:                          e,
						IsGenStatusRequestFailureError: true,
					})
					downListMutex.Unlock()
				} else {
					// 错误类型应该可预测；此处错误断言失败，遵守 let it crash 哲学，在遇到崩溃的情况下未来增加预测分支
					log.Fatal(e) // TODO: 处理此种无奈的情况
				}
			} else {
				// 正常记录
				dataMutex.Lock()
				data = append(data, result)
				dataMutex.Unlock()
			}
			wg.Done()
		}()
	}
	return
}

func requestServerAPI(url string) (data types.APIStatusResponseData, err error) {
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
			Code:         -1,
			Detail:       e.Error(),
			ResponseData: responseData,
			Stack:        debug.Stack(),
		}
		return
	} else if responseData.StatusCode() != 200 {
		err = &GenStatusRequestFailureError{
			Code:         responseData.StatusCode(),
			Detail:       "HTTP Code is not 200",
			ResponseData: responseData,
			Stack:        debug.Stack(),
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
			Stack:        debug.Stack(),
		}
		return
	}
	if _, ok := buffer["status"]; ok { // API Throw Error
		err = &GenStatusRequestFailureError{
			Code:         buffer["status"].(int),
			Detail:       buffer["message"].(string),
			ResponseData: responseData,
			Stack:        debug.Stack(),
		}
		return
	}
	// 正常数据
	var authenticResponseData types.APIStatusResponseData
	if e := json.Unmarshal(responseData.Body(), &authenticResponseData); e != nil {
		err = &GenStatusRequestFailureError{
			Code:         responseData.StatusCode(),
			Detail:       "authentic status JSON parsed failed, detail:" + e.Error(),
			ResponseData: responseData,
			Stack:        debug.Stack(),
		}
		return
	}
	data = authenticResponseData
	return
}
