package util

import "time"

func GetCurrentTimeStampMS() int64 {
	return time.Now().UnixNano() / 1e6
}
