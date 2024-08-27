package main

import (
	"fmt"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := arr[right]

	for i := range arr {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{4, 1, 2, 5, 3}
	fmt.Println(arr)
	fmt.Println(quickSort(arr))
}
