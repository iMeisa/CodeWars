package spinWords

import (
	"strings"
)

func SpinWords(str string) string {
	words := strings.Fields(str)

	for index, word := range words {
		if len(word) >= 5 {
			var reversedWord string
			for i := len(word) - 1; i >= 0; i-- {
				reversedWord += string(word[i])
			}

			words[index] = reversedWord
		}
	}

	str = strings.Join(words, " ")
	return str
}
