package main

import (
	"fmt"
	"os"
)

func main() {
	flags, args := parseFlags(os.Args[1:])

	var isHelp = len(os.Args) == 1
	var shouldOrder = false
	var insert string

	for _, flag := range flags {
		if flag[:2] == "-i" || (len(flag) > 7 && flag[:8] == "--insert") {
			_, insert = parseFlagValue(flag)
		} else if flag[:2] == "-o" || (len(flag) > 6 && flag[:7] == "--order") {
			shouldOrder = true
		} else if flag[:2] == "-h" || (len(flag) > 5 && flag[:6] == "--help") {
			isHelp = true
		}
	}

	if isHelp {
		printHelp()
		return
	}

	var str = args[0]

	if insert != "" {
		str = str + insert
	}
	if shouldOrder {
		str = order(str)
	}

	fmt.Println(str)
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--open")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func order(s string) string {
	str := []rune(s)
	for range str {
		for j := 0; j < len(str)-1; j++ {
			if str[j] > str[j+1] {
				str[j], str[j+1] = str[j+1], str[j]
			}
		}
	}
	return string(str)
}

// return the flag and the value
func parseFlagValue(flag string) (string, string) {
	var name string
	var value string

	for i, char := range flag {
		if char == '=' {
			name = flag[:i]
			value = flag[i+1:]
		}
	}

	return name, value
}

// return the flags, and the remaining args
func parseFlags(args []string) ([]string, []string) {
	var flags []string
	var remainingArgs []string

	for _, arg := range args {
		if arg[0] == '-' && arg[1] == '-' { // long flag
			flags = append(flags, arg)
		} else if arg[0] == '-' { // short flag
			flags = append(flags, arg)
		} else {
			remainingArgs = append(remainingArgs, arg)
		}
	}

	return flags, remainingArgs
}

// check if a slice contains a value
func contains(arr []string, e string) bool {
	for _, v := range arr {
		if v == e {
			return true
		}
	}
	return false
}
