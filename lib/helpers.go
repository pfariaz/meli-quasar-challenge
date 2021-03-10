package lib

import (
	"fmt"
	"os"
	"sort"
)

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

func GetBaseURL() string {
	environment := os.Getenv("GIN_MODE")
	host := os.Getenv("HOST")
	if environment == "release" {
		host = fmt.Sprintf("%s:%s", host, os.Getenv("PORT"))
	}
	return host
}
