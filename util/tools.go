package util

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

func GetCurrentTimeStampMS() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetExecDir() string {
	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(path)
	return dir
}
