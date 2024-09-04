package utils

import (
	"time"
)

func UnixToStr(timeUnix int64) string {
	if timeUnix == 0 {
		return ""
	}
	CSTLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(timeUnix, 0).Format(CSTLayout)
	return timeStr
}
func GetTimestamp() int64 {
	now := time.Now()
	timestamp := now.Unix()
	return timestamp
}

func GetMilliTimestamp() int64 {
	now := time.Now()
	timestamp := now.UnixMilli()
	return timestamp
}
