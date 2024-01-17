package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false
	for _, arg := range args {
		if arg == "--upper" {
			upper = true
			continue
		}
		nmb, _ := strconv.Atoi(arg)
		if nmb <= 25 {
			if upper {
				z01.PrintRune(rune(nmb + 64))
			} else {
				z01.PrintRune(rune(nmb + 96))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
}
