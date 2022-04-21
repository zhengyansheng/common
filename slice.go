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
func SliceTrimSpace(s []string) []string {
	var opt []string
	for _, t := range s {
		if t == "" {
			continue
		}
		opt = append(opt, t)
	}
	return opt
}
