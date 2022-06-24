package common

import (
	"fmt"
	"testing"
)

func TestSubTimeIntervalFormat(t *testing.T) {
	//startTime := "2022-01-01 13:00:00"
	endTime := "2022-06-23 17:05:00"
	interval := NewTimeInterval(endTime)
	fmt.Println(interval.Sub())
}
