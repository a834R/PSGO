package piscine

func ToLower(s string) string {
	r := []rune(s)
	for i := range r {
		if r[i] <= 'A' && r[i] >= 'Z' {
			r[i] = r[i] + 32
		}
	}

	return string(r)
}
