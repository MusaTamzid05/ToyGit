package util

func StringContains(dataSlice []string, data string) bool {

	for _, currentData := range dataSlice {
		if currentData == data {
			return true
		}
	}
	return false
}
