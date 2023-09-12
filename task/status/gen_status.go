package status

import (
	"time"

	"github.com/blang/semver/v4"
	"github.com/hitokoto-osc/Moe/cache"
	"github.com/hitokoto-osc/Moe/task/status/types"
	log "github.com/sirupsen/logrus"
)

// LimitedHost 以编码模式定义了应统计的 API 主机地址
var LimitedHost = []string{
	"v1.hitokoto.cn",
	"international.v1.hitokoto.cn",
	"api.a632079.me",
}

// RunTask 用于运行统计流程
func RunTask() {
	// 获取 API 列表
	log.Debug("[task.GenStatus] 开始执行合并任务...")
	log.Debug("[task.GenStatus] 取得 API 列表...")
	apiList := cache.GetAPIList()
	// log.Debug(apiList)
	log.Debug("[task.GenStatus] 发起请求...")
	dataList, downServerList := performRequest(apiList)
	// log.Debug(dataList, downServerList)
	log.Debug("[task.GenStatus] 合并请求记录...")
	data := genStatusData(dataList, downServerList)
	// log.Debug(data)
	cache.StoreStatusData(*data)
}

// DownServer 定义了请求中出现异常时应递交给合并器的类型
type DownServer struct {
	ID                             string
	StartTS                        int64
	Cause                          string
	Error                          error
	IsGenStatusRequestFailureError bool
}

func genStatusData(inputData []types.APIStatusResponseData, downServerList []DownServer) (data *types.GeneratedData) {
	// TODO: recover 捕获错误
	var baseHitokotoAPIVersion semver.Version
	data = &types.GeneratedData{}
	data.DownServer = []types.DownServerData{}
	if downServerList != nil && len(downServerList) > 0 {
		log.Debug("[task.GenStatus] 存在宕机节点.")
		for _, downServer := range downServerList {
			log.Errorf("[task.genStatus] 服务标识：%s, 获取统计信息时出错，错误原因：%s. 触发时间：%v", downServer.ID, downServer.Cause, downServer.StartTS)
			if downServer.IsGenStatusRequestFailureError {
				e := downServer.Error.(*GenStatusRequestFailureError)
				// log.Error(e.ResponseData)
				log.Error("[task.genStatus] 调用堆栈：" + string(e.Stack))
			} else {
				log.Error(downServer.Error)
			}
		}
		// log.Debug(downServerList)
		result := DownServerList.Merge(downServerList)
		// log.Debug(result)
		data.DownServer = append(data.DownServer, result...)
		// log.Debug(data.DownServer)
	}
	if inputData == nil || len(inputData) == 0 {
		// 抛出异常
		data.LastUpdated = time.Now().UnixNano() / 1e6
		log.Warn("[task.GenStatus] 没有节点正常工作.")
		return
	}
	// 合并主体记录
	for i, v := range inputData {
		if i == 0 {
			// 解析版本
			baseHitokotoAPIVersion = semver.MustParse(v.Version)
			initGenData(data, &v)
		} else {
			// 累加记录
			data.Children = append(data.Children, v.ServerID)
			compareAndUpdateGenDataVersion(data, &v, &baseHitokotoAPIVersion)
			mergeStatusRecord(data, &v)
			mergeRequestsRecord(data, &v)
		}
	}
	data.LastUpdated = time.Now().UnixNano() / 1e6
	// log.Debug(data.DownServer)
	return
}
