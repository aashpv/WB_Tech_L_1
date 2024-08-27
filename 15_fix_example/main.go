package main

import "fmt"

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	justString = subStr(v, 100)
	fmt.Println(justString)
}

func subStr(s string, length int) string {
	if len(s) <= length {
		return s
	}

	runes := []rune(s)
	if len(runes) < length {
		return s
	}

	return string(runes[:length])
}

func createHugeString(size int) string {
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = 'A'
	}
	return string(bytes)
}

func main() {
	someFunc()
}
