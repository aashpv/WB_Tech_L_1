package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

// Counter представляет счетчик с использованием sync.Mutex.
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment инкрементирует счетчик на единицу в защищенной блокировкой области.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

// Value возвращает текущее значение счетчика в защищенной блокировкой области.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

func main() {
	var wg sync.WaitGroup // Создаем WaitGroup
	counter := &Counter{} // Инициализируем экземпляр структуры Counter

	// Запускаем 5 горутин, каждая из которых выполняет инкремент счетчика 100 раз
	for i := 0; i < 5; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go func() {
			defer wg.Done() // Уменьшаем счетчик WaitGroup
			for j := 0; j < 100; j++ {
				counter.Increment() // Инкрементируем счетчик
			}
		}()
	}

	wg.Wait() // Ожидаем завершения всех горутин

	// Выводим итоговое значение счетчика
	fmt.Println(counter.Value())
}
