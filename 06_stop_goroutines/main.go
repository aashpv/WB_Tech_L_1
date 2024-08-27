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

func stopCtx() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена через контекст.")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}

func stopCtxTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена через Context таймаут.")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}

	}()
	wg.Wait()
}

func stopTimer() {
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-timer.C:
				fmt.Println("Горутина завершена через таймаут.")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}

func stopChannel() {
	stopCh := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-stopCh:
				fmt.Println("Горутина завершена через закрытие канала.")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(stopCh)
	wg.Wait()
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
