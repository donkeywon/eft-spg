package util

func ArrayContainsString(arr []string, d string) bool {
	for _, i := range arr {
		if i == d {
			return true
		}
	}

	return false
}
