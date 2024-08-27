package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func reverseWords(s string) string {
	// Используем strings.Fields для разделения строки на срез слов.
	// В отличие от strings.Split, strings.Fields автоматически обрабатывает несколько пробелов как один разделитель.
	words := strings.Fields(s)

	// Алгоритм разворота слов в срезе аналогичен развороту символов в строке:
	// используем два указателя i (слева) и j (справа) и меняем местами элементы.
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		// Меняем местами слова.
		words[i], words[j] = words[j], words[i]
	}

	// Соединяем слова обратно в строку с помощью strings.Join, используя пробел в качестве разделителя.
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(s)
	fmt.Println(reverseWords(s))
}
