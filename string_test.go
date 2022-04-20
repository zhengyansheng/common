package common

import "testing"

func TestSplitEmail(t *testing.T) {
	t.Run("split email", func(t *testing.T) {
		var s = "郑彦生(zhengyansheng@xxx.com)"
		info, err := SplitEmail(s)
		if err != nil {
			t.Fatalf("err: %+v", err)
			return
		}
		t.Logf("result: %+v", info)
	})
}
