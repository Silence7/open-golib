package convert

import (
	"fmt"
	"strconv"
	"time"
)

func StringToUint(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if nil != err {
		return 0
	}

	return i
}

func StringToInt(s string) int64 {
	i, err := strconv.Atoi(s)
	if nil != err {
		return 0
	}

	return int64(i)
}

func UintToString(i uint64) string {
	return fmt.Sprintf("%d", i)
}

func IntToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

// fmt = "2006-01-02 15:04:05"
// 秒级
func TimestampFmtString(t int64, fmt string) string {
	return time.Unix(t, 0).Format(fmt)
}

// fmt = "2006-01-02 15:04:05"
func GetTimeFmtString(fmt string) string {
	return time.Now().Format(fmt)
}

func TimeStringFmtInt(t string) int64 {
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	if nil != err {
		return 0
	}

	return tm.Unix()
}