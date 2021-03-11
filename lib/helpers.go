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

func Equal(firstArrayToCompare, secondArrayToCompare []string) bool {
	if len(firstArrayToCompare) != len(secondArrayToCompare) {
		return false
	}
	sort.Strings(firstArrayToCompare)
	sort.Strings(secondArrayToCompare)

	for index, value := range firstArrayToCompare {
		if value != secondArrayToCompare[index] {
			return false
		}
	}
	return true
}

func GetBaseURL() string {
	environment := os.Getenv("GIN_MODE")
	host := os.Getenv("HOST")
	if environment != "release" {
		host = fmt.Sprintf("%s:%s", host, os.Getenv("PORT"))
	}
	return host
}
