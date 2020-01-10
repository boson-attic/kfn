package util

func AnyOf(toCompare interface{}, comparing ...interface{}) bool {
	for _, c := range comparing {
		if toCompare == c {
			return true
		}
	}
	return false
}
