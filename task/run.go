package task

import (
	"github.com/hitokoto-osc/Moe/logging"
	"reflect"
	"time"

	"github.com/hitokoto-osc/Moe/task/status"
)

// Run 定义了 task 的启动入口
func Run() {
	status.DownServerList.Recover()
	go taskLoop(6*time.Second, status.RunTask)
}

func taskLoop(t time.Duration, task func()) {
	logger := logging.GetLogger()
	for {
		logger.Sugar().Debug("[taskLoop] 等待 %v 后执行 %s...", t, reflect.TypeOf(task).Name())
		logger.Sync()
		time.Sleep(t)
		task()
	}
}
