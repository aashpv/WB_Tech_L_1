package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на вход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func reverseString(s string) string {
	// Преобразуем строку в срез рун, чтобы корректно работать с Unicode-символами.
	runes := []rune(s)

	// Используем два указателя: i (слева) и j (справа).
	// Обмениваем элементы до тех пор, пока i не станет больше или равен j.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// Меняем местами руны (символы Unicode) в срезе.
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем срез рун обратно в строку и возвращаем ее.
	return string(runes)
}

func main() {
	input := "главрыба&∧"
	fmt.Println(input)
	fmt.Println(reverseString(input))
}
