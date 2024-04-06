package main

import "fmt"

//записываем данные в канал
func writer(nums []int, out chan<- int) {
	for _, v := range nums {
		out <- v
	}
	close(out)
}

// удваиваем числа во входном канале и записываем в выходной
func doubler(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2
	}
	close(out)

}

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 8}
	numsCh := make(chan int)
	doublesCh := make(chan int)
	go writer(nums, numsCh)
	go doubler(numsCh, doublesCh)
	for v := range doublesCh {
		fmt.Println(v)
	}
}
