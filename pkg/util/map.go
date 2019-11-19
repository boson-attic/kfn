package util

func GetNestedMap(in map[string]interface{}, path ...string) map[string]interface{} {
	result := in
	for _, p := range path {
		var next interface{}
		var ok bool
		if next, ok = result[p]; !ok {
			next = make(map[string]interface{})
			result[p] = next
		}
		result = next.(map[string]interface{})
	}
	return result
}

func WriteNestedEntry(in map[string]interface{}, key string, value interface{}, path ...string) {
	GetNestedMap(in, path...)[key] = value
}
