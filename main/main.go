package main

import (
	"fmt"
	"piscine"
)

func main() {
	str := "HellobHAhowbbHAareHAyou?HA"
	fmt.Println(piscine.Split(str, "HA"))
}
