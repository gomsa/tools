package uitl

import (
	"strings"
	"time"
)

// TimeLayout 转换字符
const TimeLayout = "2006-01-02 15:04:05"

// Where 条件查询
func Where(value string) (param string, val interface{}) {
	split := strings.Split(value, "|")
	param = split[0] + split[1]
	switch split[2] {
	case `string`:
		val = split[3]
	case `time`:
		val, _ = time.ParseInLocation(TimeLayout, split[3], time.Local)
	}
	return param, val
}
