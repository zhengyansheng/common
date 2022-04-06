package common

import (
	"fmt"
	"time"
)

// Now 获取当前日期
func Now() string {
	return time.Now().Format(TimeFormat)
}

// SubTime 获取时间差
func SubTime(startTime, endTime string) string {
	s1, _ := time.ParseInLocation(TimeFormat, startTime, time.Local)
	s2, _ := time.ParseInLocation(TimeFormat, endTime, time.Local)
	return s2.Sub(s1).String()
}

// RuntimeAge 获取Pod运行时长
func RuntimeAge(second int64) string {
	minSeconds := int64(60)
	hoursSeconds := minSeconds * 60
	daysSeconds := hoursSeconds * 24

	switch {
	case second <= minSeconds:
		return fmt.Sprintf("%vs", second)
	case second > minSeconds && second <= hoursSeconds:
		return fmt.Sprintf("%vm%vs", second/minSeconds, second%minSeconds)
	case second > hoursSeconds && second <= daysSeconds:
		h := second / hoursSeconds
		ms := second % hoursSeconds
		m := ms / minSeconds
		return fmt.Sprintf("%vh%vm", h, m)
	default:
		d := second / daysSeconds
		hs := second % daysSeconds
		h := hs / hoursSeconds
		return fmt.Sprintf("%vd%vh", d, h)
	}
}
