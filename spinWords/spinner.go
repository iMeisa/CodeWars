package spinWords

import (
	"fmt"
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
			fmt.Println(word)
		}
	}

	str = strings.Join(words, " ")
	return str
}
