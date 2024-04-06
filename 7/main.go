package main

import (
	"fmt"
	"sync"
	"time"
)

var mp = make(map[int]int)

// пишет в канал 50 значений
func writer(in chan<- int, done chan<- struct{}) {
	for i := 100; i <= 150; i++ {
		in <- i
		time.Sleep(100 * time.Millisecond)
	} //закрывает канал и отправляет сигнал
	close(in)
	done <- struct{}{}
}

// добавляем элементы в мапу(ключ берем из канала, значение - номер воркера)
func worker(id int, in <-chan int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for {
		v, ok := <-in // получаем значение и проверяем закрытие канала
		if !ok {
			fmt.Printf("worker %d shot down\n", id)
			return
		}
		//блокируем на запись
		mu.Lock()
		mp[v] = id
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	in := make(chan int)
	done := make(chan struct{})
	wg.Add(5)
	go writer(in, done)
	for i := 1; i <= 5; i++ {
		go worker(i, in, &wg, &mu)
	}
	// ждем сигнала о закрытии канала и let's go
	<-done
	wg.Wait()

	fmt.Println("Итоговое состояние:")
	for i, v := range mp {
		fmt.Printf("key: %d value: %d\n", i, v)
	}
}
