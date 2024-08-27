package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов (2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

func bufferedChannel(nums []int) int {
	sum := 0
	ch := make(chan int, len(nums)) // Создание буферизированного канала

	for _, num := range nums {
		go func(num int) {
			ch <- num * num // Отправка квадрата числа в канал
		}(num)
	}

	for range nums {
		sum += <-ch // Получение значений из канала и суммирование их
	}
	close(ch) // Закрытие канала

	return sum // Возврат общей суммы
}

func waitGroup(nums []int) int {
	sum := 0
	var mu sync.Mutex     // Создание мьютекса для синхронизации доступа к переменной
	var wg sync.WaitGroup // Создание экземпляра WaitGroup

	for _, num := range nums {
		wg.Add(1) // Увеличение счетчика WaitGroup

		go func(num int) {
			defer wg.Done() // Уменьшение счетчика WaitGroup
			sq := num * num // Вычисление квадрата числа
			mu.Lock()       // Захват мьютекса перед изменением переменной
			sum += sq
			mu.Unlock() // Освобождение мьютекса
		}(num)
	}

	wg.Wait() // Ожидание завершения всех горутин

	return sum // Возврат общей суммы
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	fmt.Println(bufferedChannel(nums))
	fmt.Println("-------------------")
	fmt.Println(waitGroup(nums))
}
