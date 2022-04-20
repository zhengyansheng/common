package common

import "testing"

func TestRegex(t *testing.T) {
	t.Run("test domain", func(t *testing.T) {
		s := "parent-api.huohua.cn 5xx 错误一分钟大于30"
		expr := "[a-zA-Z-]*.huohua.c[omn]{1,2}"
		result, err := Regex(s, expr)
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}
		t.Logf("result: %+v", result)
	})
}
