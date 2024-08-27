package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func swapArithmetic(a, b int) {
	a = a + b // a теперь 15
	b = a - b // b теперь 5 (исходное значение a)
	a = a - b // a теперь 10 (исходное значение b)

	fmt.Println(a, b)
}

func swapXOR(a, b int) {
	a = a ^ b // a теперь 15 (1010 ^ 0101 = 1111)
	b = a ^ b // b теперь 5 (1111 ^ 1010 = 0101)
	a = a ^ b // a теперь 10 (1111 ^ 0101 = 1010)

	fmt.Println(a, b)
}

func main() {
	a, b := 5, 10
	fmt.Println(a, b)
	fmt.Println("-------------------")
	swapArithmetic(a, b)
	fmt.Println("-------------------")
	swapXOR(a, b)
}
