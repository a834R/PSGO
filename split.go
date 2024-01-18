package piscine

func Split(str, charset string) []string {
	set := []rune(charset)

	var arr []string

	var currentWord = ""
	var currentSep = ""

	for _, c := range str {
		if currentSep == charset {
			arr = append(arr, string(currentWord[:len(currentWord)-len(charset)]))
			currentSep = ""
			currentWord = ""
		}
		if c == set[len(currentSep)] {
			currentSep += string(c)
		}
		currentWord += string(c)
	}

	arr = append(arr, currentWord)

	return arr
}
