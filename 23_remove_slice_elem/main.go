package main

import (
	"fmt"
)

/*
Удалить i-ый элемент из слайса.
*/

func removeAppend(arr []int, i int) []int {
	// Проверяем, чтобы индекс находился в пределах допустимых значений
	if i < 0 || i >= len(arr) {
		fmt.Println("Index out of bounds")
		return arr // Если индекс некорректен, возвращаем исходный слайс
	}

	// Слайс arr[:i] включает все элементы до i, а arr[i+1:] — все элементы после i.
	// Функция append объединяет эти два среза, исключая i-й элемент.
	return append(arr[:i], arr[i+1:]...)
}

func removeCopy(arr []int, i int) []int {
	// Проверяем, чтобы индекс находился в пределах допустимых значений
	if i < 0 || i >= len(arr) {
		fmt.Println("Index out of bounds")
		return arr // Если индекс некорректен, возвращаем исходный слайс
	}

	// copy(arr[i:], arr[i+1:]) копирует элементы после i-го элемента на его место.
	// Таким образом, элемент на позиции i замещается следующим элементом, а все последующие смещаются влево.
	// В итоге, последний элемент становится дублированным, и мы просто уменьшаем длину слайса на 1.
	copy(arr[i:], arr[i+1:])

	// Возвращаем слайс без последнего элемента
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
