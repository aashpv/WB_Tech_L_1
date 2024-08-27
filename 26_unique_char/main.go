package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func isUnique(str string) bool {
	str = strings.ToLower(str)

	chars := make(map[rune]bool)

	for _, char := range str {
		if chars[char] {
			return false
		}
		chars[char] = true
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("%s — %t\n", str1, isUnique(str1)) // true
	fmt.Printf("%s — %t\n", str2, isUnique(str2)) // false
	fmt.Printf("%s — %t\n", str3, isUnique(str3)) // false
}
