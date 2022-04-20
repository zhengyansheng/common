package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// SplitEmail 郑彦生(zhengyansheng@xxx.com)
func SplitEmail(s string) (info map[string]string, err error) {
	info = make(map[string]string)
	if !strings.Contains(s, "(") {
		return
	}
	emails := strings.Split(s, "(")
	info["alias"] = emails[0]
	info["email"] = strings.Trim(emails[1], ")")
	return
}

// CustomUnitGi 转换 cpu/memory 到 uint
func CustomUnitGi(s string) (float64, error) {
	if ok := strings.HasSuffix(s, "m"); ok {
		// cpu
		v, err := strconv.ParseFloat(s[:len(s)-1], 64)
		if err != nil {
			return 0, errors.New("must end in m, convert failed ")
		}
		return v / 1000, nil
	} else if ok := strings.HasSuffix(s, "Gi"); ok {
		// memory
		v, err := strconv.ParseFloat(s[:len(s)-2], 64)
		if err != nil {
			return 0, errors.New("must end in Gi, convert failed")
		}
		return v * 1024, nil
	} else {
		// cpu
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, errors.New("convert cpu failed")
		}
		return v, nil
	}
}

// CustomUnitGiInter 转换 任意类型 到 float64
func CustomUnitGiInter(s interface{}) (float64, error) {
	if v, ok := s.(string); ok {
		return CustomUnitGi(v)
	} else {
		return 0, errors.New(fmt.Sprintf("args interface value %v not string", s))
	}
}
