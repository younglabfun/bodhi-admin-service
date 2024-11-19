package utils

import (
	"time"
)

const CSTLayout string = "2006-01-02 15:04:05"

func UnixToStr(timeUnix int64) string {
	if timeUnix == 0 {
		return ""
	}
	timeStr := time.Unix(timeUnix, 0).Format(CSTLayout)
	return timeStr
}

func GetDateTime() string {
	return time.Now().Format(CSTLayout)
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
