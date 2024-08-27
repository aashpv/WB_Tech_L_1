package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить
тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel int")
	case chan string:
		fmt.Println("channel string")
	case chan bool:
		fmt.Println("channel bool")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	i := 10
	s := "hello"
	b := true
	chInt := make(chan int)
	chStr := make(chan string)
	chBool := make(chan bool)

	// Проверяем тип переменных с использованием функции checkType
	checkType(i)
	checkType(s)
	checkType(b)
	checkType(chInt)
	checkType(chStr)
	checkType(chBool)
	checkType(3.14) // float64

	fmt.Println("-------------------")

	// Определение типа переменной с использованием fmt.Sprintf и спецификатора %T
	fmt.Println(fmt.Sprintf("%T", i))
	fmt.Println(fmt.Sprintf("%T", s))
	fmt.Println(fmt.Sprintf("%T", b))
	fmt.Println(fmt.Sprintf("%T", chInt))
	fmt.Println(fmt.Sprintf("%T", chStr))
	fmt.Println(fmt.Sprintf("%T", chBool))
	fmt.Println(fmt.Sprintf("%T", 3.14))

	fmt.Println("-------------------")

	// Определение типа переменной с использованием reflect.TypeOf
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(chInt))
	fmt.Println(reflect.TypeOf(chStr))
	fmt.Println(reflect.TypeOf(chBool))
	fmt.Println(reflect.TypeOf(3.14))
}
