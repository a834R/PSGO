package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		return
	}

	val, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	val2, err := strconv.Atoi(args[2])
	if err != nil {
		return
	}

	ope := args[1]

	switch ope {
	case "+":
		fmt.Println(val + val2)
	case "-":
		fmt.Println(val - val2)
	case "*":
		fmt.Println(val * val2)
	case "/":
		if val2 == 0 {
			fmt.Println("No division by 0")
			return
		}
		fmt.Println(val / val2)
	case "%":
		if val2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		fmt.Println(val % val2)
	default:
		fmt.Println("0")
	}
}
