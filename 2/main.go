package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	// создаем канал для результатов
	resCh := make(chan int)
	// создаем ожидание
	var wg sync.WaitGroup
	// для каждого числа запускаем горутину
	for _, num := range arr {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			num *= num
			resCh <- num
		}(num)
	}
	// в горутине ждем завершения горутин и закрываем канал
	go func() {
		defer close(resCh)
		wg.Wait()
	}()
	// выводим, горутина main приостанавливается, тк канал пустой
	for res := range resCh {
		fmt.Println(res)
	}
}
