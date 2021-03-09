package lib

func Contains(arrayToCheck []string, str string) bool {
	for _, value := range arrayToCheck {
		if value == str {
			return true
		}
	}

	return false
}
