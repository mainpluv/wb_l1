package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg *sync.WaitGroup

// проверка закрыт ли канал
func stopIfClosed(in <-chan int) {
	defer wg.Done()
	for {
		if _, ok := <-in; !ok {
			fmt.Println("StopIfClosed is shutting down")
			return
		}
		fmt.Printf("StopIfClosed got: %d\n", <-in)
	}
}

// проверка пуст ли канал - если да, то метод остановится сам
func stopEmpty(in <-chan int) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("stopEmpty got: &d\n", v)
	}
	fmt.Println("stopEmpty is shutting down")
}

// второй канал для остановки горутины
func stopChannel(in <-chan int, stopCh <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case v := <-in:
			fmt.Printf("stopChannel got: %d\n", v)
		case <-stopCh:
			fmt.Println("stopChannel is shutting down")
			return
		}
	}
}

// контекст для остановки горутины
func stopCtx(ctx context.Context, in <-chan int) {
	defer wg.Done()
	for {
		select {
		case v := <-in:
			fmt.Printf("stopCtx got: %d\n", v)
		case <-ctx.Done():
			fmt.Println("stopCtx is shutting down")
			return
		}

	}
}

// писатель
func writer(ctx context.Context, in chan<- int) {
	defer wg.Done()
	timer := time.NewTicker(2 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			in <- rand.Intn(100) + 1
		case <-ctx.Done():
			fmt.Println("writer finished")
			return
		}
	}
}

func main() {
	wg = &sync.WaitGroup{}
	in := make(chan int)
	// проверки идут по порядку
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go writer(ctx, in)
	go stopIfClosed(in)
	time.Sleep(3 * time.Second)
	close(in)
	cancel()
	wg.Wait()

	ctx, cancel = context.WithCancel(context.Background())
	in = make(chan int)
	wg.Add(2)
	go writer(ctx, in)
	go stopEmpty(in)
	time.Sleep(3 * time.Second)
	close(in)
	cancel()
	wg.Wait()

	ctx, cancel = context.WithCancel(context.Background())
	in = make(chan int)
	st := make(chan struct{})
	wg.Add(2)
	go writer(ctx, in)
	go stopChannel(in, st)
	time.Sleep(3 * time.Second)
	st <- struct{}{}
	close(in)
	close(st)
	cancel()
	wg.Wait()

	ctx, cancel = context.WithCancel(context.Background())
	in = make(chan int)
	wg.Add(2)
	go writer(ctx, in)
	go stopCtx(ctx, in)
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}
