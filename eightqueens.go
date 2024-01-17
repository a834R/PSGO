package piscine

import "fmt"

const size int = 8

var board [size][size]bool

func EightQueens() {
	placeX(0)
}

func printQueens() {
	x := 0
	for x < size {
		y := 0
		for y < size {
			if board[x][y] {
				// We have found a queen, let's print her y
				fmt.Print(y + 1)
			}
			y++
		}
		x++
	}
	fmt.Println('\n')
}

func placeX(x int) {
	for y := 0; y < size; y++ {
		if canBePlaced(x, y) {
			board[y][x] = true
			if x == size-1 {
				printQueens()
			} else {
				placeX(x + 1)
			}
			board[y][x] = false
		}
	}
}

func canBePlaced(x, y int) bool {
	for i := 0; i < size; i++ {
		// check row
		if board[y][i] {
			return false
		}
		// check column
		if board[i][x] {
			return false
		}
		// check diagonal
		py := y + i
		px := x + i
		my := y - i
		mx := x - i
		if px < size && py < size && board[py][px] {
			return false
		}
		if mx >= 0 && my >= 0 && board[my][mx] {
			return false
		}
		if mx >= 0 && py < size && board[py][mx] {
			return false
		}
		if px < size && my >= 0 && board[my][px] {
			return false
		}
	}
	return true
}
