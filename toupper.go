package piscine

func ToUpper(s string) string {
	r := []rune(s)
	for i := range r {
		if r[i] >= 'a' && r[i] <= 'z' {
			r[i] = r[i] - 32
		}
	}

	return string(r)
}
