package main

import "fmt"

func main() {
	// создаем массив и мапу, которая будет хранить группы
	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)
	for _, v := range nums {
		// вычисляем из числа его десяток
		group := (int(v) / 10) * 10
		// в вычисленный десяток(группу) добавляем число
		groups[group] = append(groups[group], v)

	}
	for i, v := range groups {
		fmt.Println(i, v)
	}
}
