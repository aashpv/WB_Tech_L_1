package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

// sleepSince реализует задержку, используя цикл и проверку времени.
// Эта функция создаёт бесконечный цикл, который продолжается до тех пор,
// пока текущее время не превысит время начала на заданный интервал d.
func sleepSince(d time.Duration) {
	start := time.Now() // Запоминаем текущее время
	for {
		if time.Since(start) >= d { // Проверяем, прошло ли указанное время
			break
		}
	}
}

// sleepTimer использует таймер для реализации задержки.
// Он создаёт новый таймер с заданным интервалом d и ожидает,
// пока таймер не завершится, получая сигнал от канала C таймера.
func sleepTimer(d time.Duration) {
	timer := time.NewTimer(d) // Создаём новый таймер
	<-timer.C                 // Ожидаем завершения таймера
}

// sleepAfter использует функцию time.After для реализации задержки.
// time.After возвращает канал, который получает значение по истечении заданного интервала d.
func sleepAfter(d time.Duration) {
	<-time.After(d) // Ожидаем получения значения из канала
}

// sleepCtx использует контекст с тайм-аутом для реализации задержки.
// Он создаёт контекст с заданным тайм-аутом и ожидает завершения контекста.
func sleepCtx(d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d) // Создаём контекст с тайм-аутом
	defer cancel()                                              // Отменяем контекст при завершении функции

	<-ctx.Done() // Ожидаем завершения контекста
}

func main() {
	fmt.Println("Sleep for 2 seconds...")
	sleepSince(2 * time.Second)
	fmt.Println("Awake!")

	fmt.Println("Sleep for 2 seconds...")
	sleepTimer(2 * time.Second)
	fmt.Println("Awake!")

	fmt.Println("Sleep for 2 seconds...")
	sleepAfter(2 * time.Second)
	fmt.Println("Awake!")

	fmt.Println("Sleep for 2 seconds...")
	sleepCtx(2 * time.Second)
	fmt.Println("Awake!")
}
