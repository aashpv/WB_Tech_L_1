package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func intersect(set1, set2 []int) []int {
	// Создаем карту для хранения элементов первого множества
	elementMap := make(map[int]bool)

	// Проходим по первому множеству и добавляем каждый элемент в карту.
	// Значение true означает, что элемент присутствует в первом множестве.
	for _, elem := range set1 {
		elementMap[elem] = true
	}

	// Создаем срез для хранения пересечения множеств
	var intersection []int

	// Проходим по второму множеству
	for _, elem := range set2 {
		// Если элемент из второго множества также присутствует в первом множестве,
		// то добавляем его в пересечение и удаляем из карты, чтобы избежать дублирования
		if elementMap[elem] {
			intersection = append(intersection, elem)
			delete(elementMap, elem) // Удаление элемента из карты, чтобы избежать повторов в пересечении
		}
	}
	// Возвращаем результат пересечения множеств
	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := intersect(set1, set2)
	fmt.Println(result) // [3 4 5]
}
