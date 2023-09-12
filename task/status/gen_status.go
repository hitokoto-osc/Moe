package status

import (
	"errors"
	"github.com/hitokoto-osc/Moe/logging"
	"go.uber.org/zap"
	"time"

	"github.com/blang/semver/v4"
	"github.com/hitokoto-osc/Moe/cache"
	"github.com/hitokoto-osc/Moe/task/status/types"
)

// LimitedHost 以编码模式定义了应统计的 API 主机地址
var LimitedHost = []string{
	"v1.hitokoto.cn",
	"international.v1.hitokoto.cn",
	"api.a632079.me",
}

// RunTask 用于运行统计流程
func RunTask() {
	logger := logging.GetLogger()
	defer logger.Sync()
	// 获取 API 列表
	logger.Debug("[task.GenStatus] 开始执行合并任务...")
	logger.Debug("[task.GenStatus] 取得 API 列表...")
	apiList := cache.MustGetAPIList()
	// log.Debug(apiList)
	logger.Debug("[task.GenStatus] 发起请求...")
	dataList, downServerList := performRequest(apiList)
	// log.Debug(dataList, downServerList)
	logger.Debug("[task.GenStatus] 合并请求记录...")
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
	logger := logging.GetLogger()
	defer logger.Sync()
	// TODO: recover 捕获错误
	var baseHitokotoAPIVersion semver.Version
	data = &types.GeneratedData{}
	data.DownServer = []types.DownServerData{}
	if downServerList != nil && len(downServerList) > 0 {
		logger.Debug("[task.GenStatus] 存在宕机节点.")
		for _, downServer := range downServerList {
			if downServer.IsGenStatusRequestFailureError {
				var e *GenStatusRequestFailureError
				errors.As(downServer.Error, &e)
				// log.Error(e.ResponseData)
				logger.Error("[task.genStatus] 获取统计信息时出错：请求异常。",
					zap.String("server_id", downServer.ID),
					zap.String("cause", downServer.Cause),
					zap.Int64("occurred_at", downServer.StartTS),
					zap.Any("detail", e),
				)
			} else {
				logger.Error("[task.genStatus] 获取统计信息时出错：未知错误。",
					zap.String("server_id", downServer.ID),
					zap.String("cause", downServer.Cause),
					zap.Int64("occurred_at", downServer.StartTS),
					zap.Error(downServer.Error),
				)
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
		logger.Warn("[task.GenStatus] 没有节点正常工作.")
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
