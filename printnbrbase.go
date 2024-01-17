package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')       // let's print the minus sign
		PrintNbrBase(-nbr, base) // then recall the function with the positive value
		return
	}

	if nbr >= len(base) { // if the number is bigger than the base length
		PrintNbrBase(nbr/len(base), base) // call the function with the number divided by the base length
	}

	z01.PrintRune(rune(base[nbr%len(base)])) // print the rune at the index of the remainder of the number divided by the base length
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	checked := ""

	for _, b := range base {
		if b == '-' || b == '+' {
			return false
		}
		for _, c := range checked {
			if c == b {
				return false
			}
		}
		checked += string(b)
	}
	return true
}
