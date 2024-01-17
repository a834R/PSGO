package main

import (
	"fmt"
	"os"
)

func main() {
	args := bubbleSort(os.Args[1:])
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func bubbleSort(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
