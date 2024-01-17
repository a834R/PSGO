package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for _, v := range bubbleSort(intToDigits(n)) {
		z01.PrintRune(rune(v + '0'))
	}
}

func intToDigits(n int) []int {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return digits
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
