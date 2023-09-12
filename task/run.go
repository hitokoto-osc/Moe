package task

import (
	"reflect"
	"time"

	"github.com/hitokoto-osc/Moe/task/status"
	log "github.com/sirupsen/logrus"
)

// Run 定义了 task 的启动入口
func Run() {
	status.DownServerList.Recover()
	go taskLoop(6*time.Second, status.RunTask)
}

func taskLoop(t time.Duration, task func()) {
	for {
		log.Debugf("[taskLoop] 等待 %v 后执行 %s...", t, reflect.TypeOf(task).Name())
		time.Sleep(t)
		task()
	}
}
