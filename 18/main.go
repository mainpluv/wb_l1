package main

import (
	"fmt"
	"sync"
)

// структура
type Counter struct {
	mu    sync.Mutex
	count int
	wg    sync.WaitGroup
}

func newCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment(id int) {
	defer c.wg.Done()
	for i := 1; i <= 100; i++ {
		// увеличиваем после лока
		c.mu.Lock()
		c.count++
		fmt.Printf("worker %d count to %d\n", id, c.count)
		c.mu.Unlock()
	}
}

func main() {
	counter := newCounter()
	workers := 10
	counter.wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go counter.Increment(i)
	}
	counter.wg.Wait()

}
