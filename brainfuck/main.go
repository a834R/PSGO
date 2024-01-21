package main

import (
	"fmt"
)

var code = `
++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.
`

func main() {
	i := Interpreter{code: code}
	i.Run()
}

type Interpreter struct {
	code        string     // The code to be interpreted
	memory      [30000]int // The memory array
	pointer     int        // The pointer to the current memory cell
	codePointer int        // The pointer to the current code cell
}

func (i *Interpreter) Run() {
	for i.codePointer < len(i.code) {
		switch i.code[i.codePointer] {
		case '>':
			i.pointer++
		case '<':
			i.pointer--
		case '+':
			i.memory[i.pointer]++
		case '-':
			i.memory[i.pointer]--
		case '.':
			fmt.Printf("%c", i.memory[i.pointer])
		case ',':
			fmt.Scanf("%c", &i.memory[i.pointer])
		case '[':
			if i.memory[i.pointer] == 0 {
				i.codePointer = i.findMatchingBracket()
			}
		case ']':
			if i.memory[i.pointer] != 0 {
				i.codePointer = i.findMatchingBracket()
			}
		}
		i.codePointer++
	}
}

func (i *Interpreter) findMatchingBracket() int {
	if i.code[i.codePointer] == '[' {
		for i.codePointer < len(i.code) {
			i.codePointer++
			if i.code[i.codePointer] == ']' {
				return i.codePointer
			}
		}
	}
	if i.code[i.codePointer] == ']' {
		for i.codePointer > 0 {
			i.codePointer--
			if i.code[i.codePointer] == '[' {
				return i.codePointer
			}
		}
	}
	return -1
}
