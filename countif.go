package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, e := range tab {
		if f(e) {
			count++
		}
	}
	return count
}
