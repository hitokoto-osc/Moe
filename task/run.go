package task

import (
	"github.com/hitokoto-osc/Moe/task/status"
	log "github.com/sirupsen/logrus"
	"reflect"
	"time"
)

func Run() {
	status.DownServerList.Recover()
	go taskLoop(6*time.Second, status.RunTask)
}

type CallFunc func()

func taskLoop(t time.Duration, task CallFunc) {
	for {
		log.Debugf("[taskLoop] 等待 %v 后执行 %s...", t, reflect.TypeOf(task).Name())
		time.Sleep(t)
		task()
	}
}
