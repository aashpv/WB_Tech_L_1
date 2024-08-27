package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// устанавливает i-й бит в 1
func setBit(num int64, i int) int64 {
	return num | (1 << i)
}

// устанавливает i-й бит в 0
func clearBit(num int64, i int) int64 {
	return num &^ (1 << i)
}

func main() {
	var num int64 = 4 // 0100

	i := 1 // 0001

	num = setBit(num, i) // 0110
	fmt.Printf("%b\n", num)

	num = clearBit(num, i) // 0100
	fmt.Printf("%b\n", num)
}
