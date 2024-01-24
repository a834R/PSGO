package piscine

func ForEach(f func(int), arr []int) {
	for _, e := range arr {
		f(e)
	}
}
