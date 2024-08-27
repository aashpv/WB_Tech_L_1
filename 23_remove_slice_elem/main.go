package main

import (
	"fmt"
)

/*
Удалить i-ый элемент из слайса.
*/

func removeAppend(arr []int, i int) []int {
	if i < 0 || i >= len(arr) {
		fmt.Println("Index out of bounds")
		return arr
	}

	return append(arr[:i], arr[i+1:]...)
}

func removeCopy(arr []int, i int) []int {
	if i < 0 || i >= len(arr) {
		fmt.Println("Index out of bounds")
		return arr
	}

	copy(arr[i:], arr[i+1:])

	return arr[:len(arr)-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println("-------------------")
	arr = removeAppend(arr, 4)
	fmt.Println(arr) // 1, 2, 3, 4
	fmt.Println("-------------------")
	arr = removeCopy(arr, 0)
	fmt.Println(arr) // 2, 3, 4
}
