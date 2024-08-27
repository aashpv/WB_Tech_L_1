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
	a := big.NewInt(888888888888888888)
	b := big.NewInt(111111111111111111)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("-------------------")
	fmt.Println("a + b =", new(big.Int).Add(a, b))
	fmt.Println("a - b =", new(big.Int).Sub(a, b))
	fmt.Println("a * b =", new(big.Int).Mul(a, b))
	fmt.Println("a / b =", new(big.Int).Div(a, b))
}
