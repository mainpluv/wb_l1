package main

import (
	"fmt"
	"time"
)

func Sleep(tim time.Duration) {
	// программа не продолжит выполнение, пока не придет сигнал (через 5 сек)
	<-time.After(tim)
}

func main() {
	start := time.Now()
	Sleep(5 * time.Second)
	fmt.Println(time.Since(start))
}
