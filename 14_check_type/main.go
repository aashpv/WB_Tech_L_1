package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить
тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	case chan int:
		fmt.Println("channel int", v)
	case chan string:
		fmt.Println("channel string", v)
	case chan bool:
		fmt.Println("channel bool", v)
	default:
		fmt.Println("unknown type", v)
	}
}

func main() {
	i := 10
	s := "hello"
	b := true
	chInt := make(chan int)
	chStr := make(chan string)
	chBool := make(chan bool)

	checkType(i)
	checkType(s)
	checkType(b)
	checkType(chInt)
	checkType(chStr)
	checkType(chBool)
	checkType(3.14)
}
