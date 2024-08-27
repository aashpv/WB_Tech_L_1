package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел,
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func square(num int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счетчика WaitGroup
	fmt.Println(num * num)
}

func waitGroup(nums []int) {
	var wg sync.WaitGroup // Создание экземпляра WaitGroup

	for _, num := range nums {
		wg.Add(1)           // Увеличение счетчика WaitGroup
		go square(num, &wg) // Запуск горутины для вычисления квадрата
	}
	wg.Wait() // Ожидание завершения всех горутин
}

func bufferedChannel(nums []int) {
	ch := make(chan int, len(nums)) // Создание буферизированного канала

	for _, num := range nums {
		go func(num int) {
			ch <- num * num // Отправка квадрата числа в канал
		}(num)
	}

	for range nums {
		fmt.Println(<-ch) // Получение значения из канала и вывод его на экран
	}
	close(ch) // Закрытие канала
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	bufferedChannel(nums)
	fmt.Println("-------------------")
	waitGroup(nums)
}
