package util

func ArrayContains(arr []interface{}, d interface{}) bool {
	for _, i := range arr {
		if i == d {
			return true
		}
	}

	return false
}
