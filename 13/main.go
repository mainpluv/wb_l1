package main

import "fmt"

func main() {
	a := 10
	b := 20
	// меняем местами числа без создания временной переменной
	a, b = b, a
	fmt.Println(a, b)
}
