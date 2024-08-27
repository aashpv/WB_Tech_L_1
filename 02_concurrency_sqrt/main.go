package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел,
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func unbufferedChannel(nums []int) {
	ch := make(chan int)

	for _, num := range nums {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	for range nums {
		fmt.Println(<-ch)
	}

	close(ch)
}

func bufferedChannel(nums []int) {
	ch := make(chan int, len(nums))

	for _, num := range nums {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	for range nums {
		fmt.Println(<-ch)
	}

	close(ch)
}

func waitGroup(nums []int) {
	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()
			sq := num * num
			fmt.Println(sq)
		}(num)

		wg.Wait()
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	unbufferedChannel(nums)
	fmt.Println("-------------------")
	bufferedChannel(nums)
	fmt.Println("--------------------")
	waitGroup(nums)
}
