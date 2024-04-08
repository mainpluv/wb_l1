package main

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	// меняем местами проивоположные относительно центра слайса элементы
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	in := "главрыба"
	res := reverse(in)
	fmt.Println(res)
}
