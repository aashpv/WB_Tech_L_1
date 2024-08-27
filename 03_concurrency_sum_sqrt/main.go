package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

func unbufferedChannel(nums []int) int {
	sum := 0
	ch := make(chan int)

	for _, num := range nums {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	for range nums {
		sum += <-ch
	}
	close(ch)

	return sum
}

func bufferedChannel(nums []int) int {
	sum := 0
	ch := make(chan int, len(nums))

	for _, num := range nums {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	for range nums {
		sum += <-ch
	}
	close(ch)

	return sum
}

func waitGroup(nums []int) int {
	sum := 0
	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()
			sq := num * num
			sum += sq
		}(num)

		wg.Wait()
	}

	return sum
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(unbufferedChannel(nums))
	fmt.Println("-------------------")
	fmt.Println(bufferedChannel(nums))
	fmt.Println("--------------------")
	fmt.Println(waitGroup(nums))
}
