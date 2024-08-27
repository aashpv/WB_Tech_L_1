package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

// Пример конкурентной записи в map с использованием sync.Map
func writeSyncMap() {
	var myMap sync.Map // sync.Map является потокобезопасным и предоставляет методы для конкурентного доступа
	var wg sync.WaitGroup

	// Запускаем несколько горутин для записи в myMap
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) { // Передаем i в замыкание, чтобы избежать проблем с доступом к переменной i
			defer wg.Done()      // Уменьшаем счетчик WaitGroup
			myMap.Store(i, i*10) // Используем метод Store для записи в sync.Map
		}(i)
	}
	wg.Wait() // Ожидаем завершения всех горутин

	// Перебираем и выводим все значения из sync.Map
	myMap.Range(func(key, value any) bool {
		fmt.Printf("%d: %d\n", key, value)
		return true // Продолжаем итерацию
	})
}

// Пример конкурентной записи в map с использованием мьютекса
func writeMutex() {
	var mu sync.Mutex // Мьютекс для синхронизации доступа к обычной map
	var wg sync.WaitGroup

	myMap := make(map[int]int) // Обычная map без потокобезопасности

	// Запускаем несколько горутин для записи в myMap
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) { // Передаем i в замыкание, чтобы избежать проблем с доступом к переменной i
			defer wg.Done()   // Уменьшаем счетчик WaitGroup после завершения горутины
			mu.Lock()         // Захватываем мьютекс для обеспечения безопасного доступа к myMap
			myMap[i] = i * 10 // Записываем данные в map
			mu.Unlock()       // Освобождаем мьютекс
		}(i)
	}
	wg.Wait() // Ожидаем завершения всех горутин

	// Перебираем и выводим все значения из myMap
	for key, value := range myMap {
		fmt.Printf("%d: %d\n", key, value)
	}
}

func main() {
	writeSyncMap()
	fmt.Println("-------------------")
	writeMutex()
}
