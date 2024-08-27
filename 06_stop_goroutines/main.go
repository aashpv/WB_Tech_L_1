package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// Пример остановки горутины через контекст
func stopCtx() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done() // Уменьшаем счетчик WaitGroup

		for {
			select {
			case <-ctx.Done(): // Проверка на завершение контекста
				fmt.Println("Горутина завершена через контекст.")
				return // Выход из горутины
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond) // Пауза между итерациями
			}
		}
	}()

	// Подождем 2 секунды перед тем как отменить контекст
	time.Sleep(2 * time.Second)
	cancel()  // Отменяем контекст
	wg.Wait() // Ожидаем завершения горутины
}

// Пример остановки горутины через таймаут в контексте
func stopCtxTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Отменяем контекст при завершении функции

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done() // Уменьшаем счетчик WaitGroup

		for {
			select {
			case <-ctx.Done(): // Проверка на завершение контекста из-за таймаута
				fmt.Println("Горутина завершена через Context таймаут.")
				return // Выход из горутины
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond) // Пауза между итерациями
			}
		}
	}()
	wg.Wait() // Ожидаем завершения горутины
}

// Пример остановки горутины через таймер
func stopTimer() {
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop() // Останавливаем таймер при завершении функции

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done() // Уменьшаем счетчик WaitGroup

		for {
			select {
			case <-timer.C: // Проверка на срабатывание таймера
				fmt.Println("Горутина завершена через таймаут.")
				return // Выход из горутины
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond) // Пауза между итерациями
			}
		}
	}()
	wg.Wait() // Ожидаем завершения горутины
}

// Пример остановки горутины через закрытие канала
func stopChannel() {
	stopCh := make(chan struct{}) // Создаем канал для сигнала остановки
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done() // Уменьшаем счетчик WaitGroup

		for {
			select {
			case <-stopCh: // Проверка на закрытие канала
				fmt.Println("Горутина завершена через закрытие канала.")
				return // Выход из горутины
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond) // Пауза между итерациями
			}
		}
	}()

	// Подождем 2 секунды перед тем как закрыть канал
	time.Sleep(2 * time.Second)
	close(stopCh) // Закрываем канал
	wg.Wait()     // Ожидаем завершения горутины
}

func main() {
	stopCtx()
	fmt.Println("-------------------")
	stopCtxTimeout()
	fmt.Println("-------------------")
	stopTimer()
	fmt.Println("-------------------")
	stopChannel()
}
