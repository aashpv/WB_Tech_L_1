package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func dataGenerator(ctx context.Context, dataChan chan<- int, wg *sync.WaitGroup) {
	defer close(dataChan) // Закрытие канала при завершении работы функции
	defer wg.Done()       // Уменьшение счетчика WaitGroup

	for i := 1; ; i++ {
		select {
		case <-ctx.Done(): // Проверка на истечение времени
			fmt.Println("dataGenerator остановлен.")
			return
		case dataChan <- i: // Отправка текущего значения в канал
			time.Sleep(500 * time.Millisecond) // Пауза между отправками данных
		}
	}
}

func worker(ctx context.Context, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счетчика WaitGroup

	for {
		select {
		case data, ok := <-dataChan: // Чтение данных из канала
			if !ok { // Проверка на закрытие канала
				return
			}
			fmt.Println("Получены данные:", data)
		case <-ctx.Done(): // Проверка на истечение времени
			fmt.Println("Время истекло. Завершение программы.")
			return
		}
	}
}

func main() {
	const duration = 3 * time.Second // Время работы программы

	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel() // Отмена контекста при завершении main функции

	dataChan := make(chan int) // Создание канала для передачи данных
	var wg sync.WaitGroup      // Создание WaitGroup для синхронизации горутин

	wg.Add(2) // Установка счетчика WaitGroup для двух горутин

	go dataGenerator(ctx, dataChan, &wg)
	go worker(ctx, dataChan, &wg)

	wg.Wait() // Ожидание завершения всех горутин
	fmt.Println("Программа завершена.")
}
