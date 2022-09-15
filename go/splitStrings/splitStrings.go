package splitStrings

func Solution(str string) []string {
	var splitString []string
	for i := 0; i < len(str); i += 2 {
		if len(str)-i < 2 {
			splitString = append(splitString, string(str[i])+"_")
			continue
		}
		splitString = append(splitString, str[i:i+2])
	}

	return splitString
}
