package main

import "fmt"

func setBit(a int64, pos uint, bitVal bool) int64 {
	// создаем маску, устанавливая бит в заданную позицию
	mask := int64(1 << pos)
	if bitVal {
		// устанавливаем бит в 1, побитовое исключающее ИЛИ (^)
		return a ^ mask
	} else {
		// устанавливаем бит в 0, побитовое И-НЕТ (&^)
		return a &^ mask
	}
}

func main() {
	var a int64 = 42 //пример
	fmt.Printf("%064b\n", a)
	pos := uint(1) // позиция бита, который мы хотим установить
	bitVal := true // устанавливаем бит в 1 или 0 (true - 1, false - 0)

	res := setBit(a, pos, bitVal)
	fmt.Printf("%064b\n", res)
}
