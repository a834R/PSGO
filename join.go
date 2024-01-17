package piscine

func Join(strs []string, sep string) string {
	str := ""

	for i, s := range strs {
		if i != 0 {
			str += sep
		}
		str += s
	}

	return str
}
