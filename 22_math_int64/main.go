package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a, b, значение которых > 2^20.
*/

func main() {
	a := big.NewInt(888888888888888888) // Переменная a
	b := big.NewInt(111111111111111111) // Переменная b

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("-------------------")
	// new(big.Int).Add(a, b) создает новое big.Int для хранения результата сложения
	fmt.Println("a + b =", new(big.Int).Add(a, b))
	// new(big.Int).Sub(a, b) создает новое big.Int для хранения результата вычитания
	fmt.Println("a - b =", new(big.Int).Sub(a, b))
	// new(big.Int).Mul(a, b) создает новое big.Int для хранения результата умножения
	fmt.Println("a * b =", new(big.Int).Mul(a, b))
	// new(big.Int).Div(a, b) создает новое big.Int для хранения результата деления
	fmt.Println("a / b =", new(big.Int).Div(a, b))
}
