package main

import (
	"fmt"
	"math"
)

func main() {
	// Присваиваем значение, большее чем 2^20
	a := math.Pow(2, 20) + 1
	b := math.Pow(2, 20) + 2

	// умножение
	multiplication := a * b
	fmt.Printf("Умножение: %.0f * %.0f = %.0f\n", a, b, multiplication)

	// деление
	division := a / b
	fmt.Printf("Деление: %.0f / %.0f = %.4f\n", a, b, division)

	// сложение
	addition := a + b
	fmt.Printf("Сложение: %.0f + %.0f = %.0f\n", a, b, addition)

	// вычитание
	subtraction := a - b
	fmt.Printf("Вычитание: %.0f - %.0f = %.0f\n", a, b, subtraction)
}
