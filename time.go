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

// SubTimeSeconds 获取时间秒数
func SubTimeSeconds(startTime, endTime string) int64 {
	s1, _ := time.ParseInLocation(SecLocalTimeFormat, startTime, time.Local)
	s2, _ := time.ParseInLocation(SecLocalTimeFormat, endTime, time.Local)
	return int64(s2.Sub(s1).Seconds())
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
		unixSeconds := et2.Unix() - st1.Unix()
		fmt.Println(unixSeconds, 24*3600)
		// day 24 * 3600
		if unixSeconds > 24*3600 {
			m = unixSeconds / (24 * 3600)
		} else {
			// hour
			m = unixSeconds / 60
		}
		return

	} else {
		err = errors.New(fmt.Sprintf("startTime: %v不能大于endTime: %v", startTime, endTime))
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

type TimeInterval struct {
	StartTime string `json:"start_time"`
	Result    string `json:"result"`
}

func NewTimeInterval(startTime string) *TimeInterval {
	return &TimeInterval{
		StartTime: startTime,
		Result:    "",
	}
}

func (t *TimeInterval) Sub() string {
	startTime, _ := time.ParseInLocation(SecLocalTimeFormat, t.StartTime, time.Local)
	seconds := int64(time.Now().Sub(startTime).Seconds())
	return t.interval(seconds)
}

func (t *TimeInterval) interval(seconds int64) string {
	if seconds > 86400 {
		days := int(seconds / 86400)
		remainSeconds := seconds % 86400
		t.Result += fmt.Sprintf("%vd", days)
		t.interval(remainSeconds)
	} else if seconds > 3600 {
		hours := int(seconds / 3600)
		remainSeconds := seconds % 3600
		t.Result += fmt.Sprintf("%vh", hours)
		t.interval(remainSeconds)
	} else if seconds > 60 {
		mins := int(seconds / 60)
		remainSeconds := seconds % 60
		t.Result += fmt.Sprintf("%vm", mins)
		t.interval(remainSeconds)
	} else {
		t.Result += fmt.Sprintf("%vs", seconds)
	}
	return t.Result
}
