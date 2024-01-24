package piscine

func Any(f func(string) bool, arr []string) bool {
	for _, e := range arr {
		if f(e) {
			return true
		}
	}
	return false
}
