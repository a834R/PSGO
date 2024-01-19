package main

import (
	"fmt"
)

func main() {
	PrintNbrBase(125, "0123456789")
	PrintNbrBase(-125, "01")
	PrintNbrBase(125, "0123456789ABCDEF")
	PrintNbrBase(-125, "choumi")
	PrintNbrBase(125, "aa")
	fmt.Println("-----------------------")
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	if !IsValidBase(base) {
		return 0
	}
	blen := len(base)
	res := 0
	for _, c := range s {
		res = res*blen + findCharIndex(base, c)
	}

	return res
}

func findCharIndex(s string, c rune) int {
	for i, v := range s {
		if v == c {
			return i
		}
	}
	return -1
}

func PrintNbrBase(nb int, base string) {
	if !IsValidBase(base) {
		fmt.Println("NV")
		return
	}
	sign := 1
	if nb < 0 {
		sign = -1
		nb = -nb
	}
	res := ""
	blen := len(base)

	for nb > 0 {
		res = string(base[nb%blen]) + res
		nb /= blen
	}

	if sign == -1 {
		fmt.Print("-")
	}
	fmt.Print(res)
	fmt.Println()
}

func IsValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	uniq := ""
	for _, c := range base {
		if c == '+' || c == '-' {
			return false
		}
		for _, u := range uniq {
			if c == u {
				return false
			}
		}
		uniq += string(c)
	}
	return true
}
