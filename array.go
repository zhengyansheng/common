package common

// SliceTrimSpace 删除 slice 中的空元素
func SliceTrimSpace(s []string) []string {
	var s []string
	for _, t := range s {
		if len(t) == "" {
			continue
		}
		s = append(s, t)
	}
	return s
}
