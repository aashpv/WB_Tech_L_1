package main

import (
	"fmt"
)

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, str := range sequence {
		set[str] = struct{}{}
	}

	for word := range set {
		fmt.Println(word)
	}
}
