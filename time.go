package common

import (
	"errors"
	"fmt"
	"time"
)

// Now 获取当前日期
func Now() string {
	return time.Now().Format(SecLocalTimeFormat)
}

// SubTime 获取时间差
func SubTime(startTime, endTime string) string {
	s1, _ := time.ParseInLocation(SecLocalTimeFormat, startTime, time.Local)
	s2, _ := time.ParseInLocation(SecLocalTimeFormat, endTime, time.Local)
	return s2.Sub(s1).String()
}

// SubTimeInterval 求2个时间差
func SubTimeInterval(startTime, endTime string) (m int64, err error) {
	st1, err := time.ParseInLocation(SecLocalTimeFormat, startTime, time.Local)
	if err != nil {
		return
	}
	et2, err := time.ParseInLocation(SecLocalTimeFormat, endTime, time.Local)
	if err != nil {
		return
	}
	if st1.Before(et2) {
		m = (et2.Unix() - st1.Unix()) / 60
		return
	} else {
		err = errors.New("startTime不能大于endTime")
		return
	}
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

// ParseMilliTimestamp 毫秒转日期
func ParseMilliTimestamp(tm int64) string {
	sec := tm / 1000
	msec := tm % 1000
	return time.Unix(sec, msec*int64(time.Millisecond)).Format(SecLocalTimeFormat)
}