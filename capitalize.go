package piscine

func Capitalize(s string) string {
	r := []rune(ToLower(s))

	if IsAlpha(string(r[0])) {
		r[0] = []rune(ToUpper(string(r[0])))[0]
	}

	for i := 1; i < len(r); i++ {
		if !IsAlpha(string(r[i-1])) {
			r[i] = []rune(ToUpper(string(r[i])))[0]
		}
	}

	return string(r)
}
