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
	defer close(dataChan)
	defer wg.Done()

	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("dataGenerator остановлен.")
			return
		case dataChan <- i:
			time.Sleep(500 * time.Millisecond) // Пауза между отправками
		}
	}
}

func worker(ctx context.Context, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				return
			}
			fmt.Println("Получены данные:", data)
		case <-ctx.Done():
			fmt.Println("Время истекло. Завершение программы.")
			return
		}
	}
}

func main() {
	const duration = 3 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	dataChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go dataGenerator(ctx, dataChan, &wg)
	go worker(ctx, dataChan, &wg)

	wg.Wait()
	fmt.Println("Программа завершена.")
}
