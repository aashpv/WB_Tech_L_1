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
	defer close(dataCh)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			dataCh <- rand.Intn(100)
			time.Sleep(500 * time.Millisecond) // пауза между отправками
		}
	}
}

func worker(ctx context.Context, id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер: %d завершил работу.\n", id)
			return
		case data, ok := <-dataCh:
			if !ok {
				return
			}
			fmt.Printf("Воркер: %d получил данные: %d\n", id, data)
		}
	}
}

func main() {
	var workerCount int
	fmt.Print("Введите количество воркеров: ")
	_, err := fmt.Scan(&workerCount)
	if err != nil {
		return
	}

	dataCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	go dataGenerator(ctx, dataCh)

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, dataCh, &wg)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	// Ожидаем сигнал
	<-stopChan
	fmt.Println("\nПолучен сигнал завершения, остановка программы...")

	// Отменяем контекст
	cancel()

	// Ожидаем завершения всех воркеров
	wg.Wait()
	fmt.Println("Программа завершена.")
}
