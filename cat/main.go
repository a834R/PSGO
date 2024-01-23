package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	output := ""

	if len(args) == 0 {
		// take user input
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			output += err.Error() + "\n"
		} else {
			output += string(input) + "\n"
		}
	} else {
		// read files
		for _, arg := range args {
			file, err := os.ReadFile(arg)
			if err != nil {
				output += err.Error() + "\n"
			} else {
				output += string(file) + "\n"
			}
		}
	}

	fmt.Print(output)
}
