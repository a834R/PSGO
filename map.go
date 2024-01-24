package piscine

func Map(f func(int) bool, arr []int) []bool {
	res := make([]bool, len(arr))

	for _, e := range arr {
		res = append(res, f(e))
	}

	return res
}
