package romanNumeralsDecoder

import (
	"fmt"
	"strings"
)

func Decode(roman string) int {

	var total int
	prevValue := 0
	for _, letterByte := range roman {
		letter := strings.ToUpper(string(letterByte))
		var currentValue int
		switch letter {
		case "I":
			currentValue = 1
		case "V":
			currentValue = 5
		case "X":
			currentValue = 10
		case "L":
			currentValue = 50
		case "C":
			currentValue = 100
		case "D":
			currentValue = 500
		case "M":
			currentValue = 1000
		default:
			fmt.Printf("Invalid character: %v\n", letter)
			return 0
		}

		if prevValue < currentValue {
			currentValue -= prevValue * 2
		}

		total += currentValue
		prevValue = currentValue
	}

	return total
}
