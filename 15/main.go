package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	//заполняем строку символами a и b
	return "a" + strings.Repeat("b", n-1)
}

func someFunc() {
	v := createHugeString(1 << 10)
	// проблема утечки памяти решена
	justString = string(v[0:100])
	fmt.Print(justString)
}

func main() {
	someFunc()
}
