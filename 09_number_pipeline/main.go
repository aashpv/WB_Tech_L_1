package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func writer(nums []int, ch1 chan<- int) {
	for _, num := range nums {
		ch1 <- num
	}
	close(ch1)
}

func multiplier(ch1 <-chan int, ch2 chan<- int) {
	for num := range ch1 {
		ch2 <- num * 2
	}
	close(ch2)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go writer(nums, ch1)
	go multiplier(ch1, ch2)

	for num := range ch2 {
		fmt.Println(num)
	}
}
