package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

// Конструктор реализуем самостоятельно путем создания метода с типом Point. Каждая точка имеет координаты по x и y
type Point struct {
	x float64
	y float64
}

func (p *Point) Point(x, y float64) {
	p.x = x
	p.y = y
}

// Вычисление расстояния между взятыми на плоскости двумя точками А(х А ; у А ) и В(х В ; у В ), выполняется по формуле d = √((хА – хВ)2 + (уА – уВ)2)

func (pointOne *Point) distanceBetweenTwoPoints(pointTwo *Point) float64 {
	return math.Sqrt((math.Pow((pointTwo.x-pointOne.x), 2) + math.Pow((pointTwo.y-pointOne.y), 2)))
}

func main() {
	//Объявляем и инициализируем первую точку
	var pointOne Point
	pointOne.Point(5, 1)

	//Объявляем и инициализируем вторую точку
	var pointTwo Point
	pointTwo.Point(15, 10)

	// +v показывает поля структуры
	fmt.Printf("Первая точка на отрезке = %+v\n", pointOne)
	fmt.Printf("Вторая точка на отрезке = %+v\n", pointTwo)
	//Вычисляем расстояние между точкой А и B
	fmt.Printf("Расстояние между двумя точками на отрезке = %v", pointOne.distanceBetweenTwoPoints(&pointTwo))
}
