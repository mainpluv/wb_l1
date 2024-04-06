package main

import (
	"fmt"
	"time"
)

// отправляем каждую секунду в канал
func send(ch chan<- int) {
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

// при наличии числа в канале читаем
func receive(ch <-chan int) {
	for {
		val := <-ch
		fmt.Println(val)
	}
}

func main() {

	ch := make(chan int)

	go send(ch)
	go receive(ch)
	// ждем 5 сек, потом после sleepa в send main завершится
	time.Sleep(5 * time.Second)
}
