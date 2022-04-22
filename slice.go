package common

// Contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// SliceTrimSpace 删除 slice 中的空元素
func SliceTrimSpace(s []string) (opt []string) {
	for _, t := range s {
		if t == "" {
			continue
		}
		opt = append(opt, t)
	}
	return
}

// SetTrimSpace 删除 slice 中的空元素 和 重复的元素
func SetTrimSpace(s []string) (opt []string) {
	tmpSlice := SliceTrimSpace(s)
	for _, s1 := range tmpSlice {
		if ok := Contains(opt, s1); !ok {
			opt = append(opt, s1)
		}
	}
	return
}
