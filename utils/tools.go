package utils

import "time"

/**
格式化时间
*/
func GetDay() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

/**
纳秒
*/
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

/**
当前时间戳
*/
func GetUnix() int64 {
	return time.Now().Unix()
}
