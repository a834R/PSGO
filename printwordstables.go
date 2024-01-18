package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, w := range table {
		for _, c := range w {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
