package lib

import "sort"

func Contains(arrayToCheck []string, str string) bool {
	for _, value := range arrayToCheck {
		if value == str {
			return true
		}
	}

	return false
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
