package util

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

// GetCurrentTimeStampMS 用于获取当前的毫秒级时间戳
func GetCurrentTimeStampMS() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetExecDir 用于获取当前执行文件的目录
func GetExecDir() string {
	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(path)
	return dir
}
