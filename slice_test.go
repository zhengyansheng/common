package common

import "testing"

func TestSliceTrimSpace(t *testing.T) {
	t.Run("trim space in slice", func(t *testing.T) {
		var tests = []string{"", "", "zhengyansheng"}
		result := SliceTrimSpace(tests)
		t.Logf("result: %v", result)
	})

	t.Run("set trim space in slice", func(t *testing.T) {
		var tests = []string{"", "", "zhengyansheng", "zhengyansheng"}
		result := SetTrimSpace(tests)
		t.Logf("result: %v", result)
	})
}
