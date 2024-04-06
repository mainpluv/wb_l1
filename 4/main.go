package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func writer(ch chan<- int) {
	//пишем в канал
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

func worker(workerId int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v, ok := <-ch // читаем из канала
		if !ok {
			fmt.Printf("Worker %d has shut down\n", workerId)
			return // если канал закрыт, завершаем работу
		}
		fmt.Printf("Worker %d received: %d\n", workerId, v)
	}
}

func main() {
	var numWorkers int
	// считываем число воркеров
	fmt.Println("enter the number of workers:")
	fmt.Scan(&numWorkers)
	ch := make(chan int)
	var wg sync.WaitGroup

	go writer(ch)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}
	// канал, ожидающий сигнала
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	close(ch)
	wg.Wait()
}
