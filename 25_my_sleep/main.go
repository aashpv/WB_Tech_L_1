package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func sleepSince(d time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= d {
			break
		}
	}
}

func sleepTimer(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func sleepAfter(d time.Duration) {
	<-time.After(d)
}

func sleepCtx(d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	<-ctx.Done()
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
