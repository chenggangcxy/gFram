package models

import (
	"time"
	//"github.com/golang/protobuf/ptypes/timestamp"
)

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

//获取年月日时间
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

//获取年月日时分秒
func GetDate() string {
	template := "2006-01-02 13:04:05"
	return time.Now().Format(template)
}

//获取当前时间戳
func GetUnix() int64 {

	return time.Now().Unix()
}
