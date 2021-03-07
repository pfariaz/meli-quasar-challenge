package services

import (
	"strings"
)

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string) {
	if len(messages) == 0 {
		return
	}

	var words_by_position = map[int]string{}
	for _, message := range messages {
		for i := 0; i < len(message); i++ {
			if len(message[i]) > 0 {
				words_by_position[i] = message[i]
			}
		}
	}
	words := []string{}
	for i := 0; i < len(words_by_position); i++ {
		words = append(words, words_by_position[i])
	}
	msg = strings.Join(words, " ")
	return
}
