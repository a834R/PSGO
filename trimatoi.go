package piscine

func TrimAtoi(s string) int {
	var nbr int = 0
	sign := 1
	for _, v := range []rune(s) {
		if nbr < 1 && v == '-' {
			sign = -1
		}
		if v >= '0' && v <= '9' {
			nbr = (nbr * 10) + int(v-48)
		}
	}
	return sign * nbr
}
