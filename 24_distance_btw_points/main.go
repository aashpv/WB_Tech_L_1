package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x, y и конструктором.
*/

// Структура Point представляет точку на плоскости с координатами x и y
// Поля инкапсулированы, т.к. в нижнем регистре
type Point struct {
	x float64
	y float64
}

// Конструктор NewPoint создает и возвращает указатель на новый экземпляр структуры Point.
func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

// Формула для вычисления расстояния между двумя точками в пространстве: sqrt((x2 - x1)^2 + (y2 - y1)^2).
func Distance(p1 *Point, p2 *Point) float64 {
	// Разница по оси x
	dx := p1.x - p2.x
	// Разница по оси y
	dy := p1.y - p2.y
	// Вычисляем расстояние с использованием теоремы Пифагора
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(0, 0)
	distance := Distance(p1, p2)
	fmt.Println(distance) // Ожидаемый результат: 5
}
