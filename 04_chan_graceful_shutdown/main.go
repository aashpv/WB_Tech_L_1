package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func dataGenerator(ctx context.Context, dataCh chan<- int) {
	defer close(dataCh) // Закрыть канал при завершении функции

	for {
		select {
		case <-ctx.Done(): // Проверка, был ли вызван cancel для контекста
			return
		default:
			dataCh <- rand.Intn(100)           // Генерация случайного числа и отправка в канал
			time.Sleep(500 * time.Millisecond) // Пауза между отправками данных
		}
	}
}

func worker(ctx context.Context, id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счетчика WaitGroup

	for {
		select {
		case <-ctx.Done(): // Проверка, был ли вызван cancel для контекста
			fmt.Printf("Воркер: %d завершил работу.\n", id)
			return
		case data, ok := <-dataCh: // Получение данных из канала
			if !ok { // Проверка, был ли канал закрыт
				return
			}
			fmt.Printf("Воркер: %d получил данные: %d\n", id, data)
		}
	}
}

func main() {
	var workerCount int
	fmt.Print("Введите количество воркеров: ")
	_, err := fmt.Scan(&workerCount) // Получение количества воркеров
	if err != nil {
		return
	}

	dataCh := make(chan int)                                // Создание канала для передачи данных
	ctx, cancel := context.WithCancel(context.Background()) // Создание контекста с возможностью отмены
	defer cancel()                                          // Отмена контекста при завершении main функции

	var wg sync.WaitGroup // Создание WaitGroup

	go dataGenerator(ctx, dataCh) // Запуск генератора данных

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)                      // Увеличение счетчика WaitGroup
		go worker(ctx, i, dataCh, &wg) // Запуск воркера
	}

	stopChan := make(chan os.Signal, 1)   // Создание канала для получения сигналов
	signal.Notify(stopChan, os.Interrupt) // Подписка на сигнал прерывания (Ctrl+C)

	// Ожидание сигнала прерывания
	<-stopChan
	fmt.Println("\nПолучен сигнал завершения, остановка программы...")

	// Отменяем контекст
	cancel()

	// Ожидаем завершения всех воркеров
	wg.Wait()
	fmt.Println("Программа завершена.")
}
