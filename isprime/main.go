package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}

func IsPrime(nb int) bool {
	if nb == 0 || nb == 1 {
		return false
	}
	if nb == 2 || nb == 3 {
		return true
	}
	// check if nb is divisible by 2 or 3
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	for i := 5; i*i <= nb; i = i + 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}

	return true
}
