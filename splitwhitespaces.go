package piscine

func SplitWhiteSpaces(str string) []string {
	var arr []string

	current := ""

	for _, c := range str {
		if c == '\n' || c == '\t' || c == ' ' {
			if current != "" {
				arr = append(arr, current)
				current = ""
			}
		} else {
			current += string(c)
		}
	}
	// the last word
	if current != "" {
		arr = append(arr, current)
	}

	return arr
}
