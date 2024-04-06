package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	//создаем для ожидания
	var wg sync.WaitGroup
	// запускаем горутины для каждого числа
	for _, v := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// пишем в канал
			ch <- v * v
		}()
	}
	go func() {
		// ждем и закрывааем
		defer close(ch)
		wg.Wait()
	}()

	var sum int
	// читаем
	for num := range ch {
		sum += num
	}
	fmt.Println(sum)
}
