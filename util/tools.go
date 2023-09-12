package util

import (
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

// GetCurrentTimeStampMS 用于获取当前的毫秒级时间戳
func GetCurrentTimeStampMS() int64 {
	return time.Now().UnixNano() / 1e6
}

// MustGetExecDir 用于获取当前执行文件的目录
func MustGetExecDir() string {
	path, err := os.Executable()
	if err != nil {
		zap.L().Fatal("无法获得程序执行路径", zap.Error(err))
	}
	dir := filepath.Dir(path)
	return dir
}
