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
	fmt.Println("-----------------------")
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// first let's convern nbr to base 10
	b10 := 0
	flen := len(baseFrom)
	for _, c := range nbr {
		b10 = b10*flen + findCharIndex(baseFrom, c)
	}

	// then convert from base 10 to baseTo
	res := ""
	tlen := len(baseTo)
	for b10 > 0 {
		res = string(baseTo[b10%tlen]) + res
		b10 /= tlen
	}

	return res
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
