package common

func MergeJSONMaps(maps ...map[string]interface{}) (result map[string]interface{}) {
	result = make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
