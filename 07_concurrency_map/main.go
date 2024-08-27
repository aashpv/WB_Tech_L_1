package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func writeSyncMap() {
	var myMap sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			myMap.Store(i, i*10)
		}()
	}
	wg.Wait()

	myMap.Range(func(key, value any) bool {
		fmt.Printf("%d: %d\n", key, value)
		return true
	})
}

func writeMutex() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	myMap := make(map[int]int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			myMap[i] = i * 10
			mu.Unlock()
		}()
	}
	wg.Wait()

	for key, value := range myMap {
		fmt.Printf("%d: %d\n", key, value)
	}
}

func main() {
	writeSyncMap()
	fmt.Println("-------------------")
	writeMutex()
}
