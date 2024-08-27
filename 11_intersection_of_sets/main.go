package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func intersect(set1, set2 []int) []int {
	elementMap := make(map[int]bool)
	for _, elem := range set1 {
		elementMap[elem] = true
	}

	var intersection []int
	for _, elem := range set2 {
		if elementMap[elem] {
			intersection = append(intersection, elem)
			delete(elementMap, elem)
		}
	}
	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := intersect(set1, set2)

	fmt.Println(result)
}
