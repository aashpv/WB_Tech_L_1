package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(binarySearch(arr, 2))
	fmt.Println(binarySearch(arr, 9))
	fmt.Println(binarySearch(arr, 0))
}
