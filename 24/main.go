package main

import (
	"fmt"
	"math"
)

// структура точки
type Point struct {
	x float64
	y float64
}

// создаем новую точку
func newPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// вычисляем расстояние, используя дельты
func Distance(p1, p2 *Point) float64 {
	dX := p2.x - p1.x
	dY := p2.y - p1.y
	return math.Sqrt(dX*dX + dY*dY)
}

func main() {
	p1 := newPoint(3, 5)
	p2 := newPoint(6, 10)

	dist := Distance(p1, p2)
	fmt.Println(dist)
}
