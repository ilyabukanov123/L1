package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

// Создаем именованный тип для множества т.к все ключи являются уникальными
type Set map[int]struct{}

func main() {
	// Создаем два слайса под множества
	setOne := []int{10, 8, 6, 4, 2}
	setTwo := []int{9, 8, 7, 5, 4, 1}

	// Множество пересечений
	setIntersection := []int{}

	// Создаем мэп т.к у него ключи уникальны
	setMap := make(Set)

	// Заполняем мэп на основе первого множества
	for _, value := range setOne {
		setMap[value] = struct{}{}
	}

	for _, value := range setTwo {
		// Если элемент второго множества существует в мэпе, который создан на основе второго множества
		if _, exists := setMap[value]; exists {
			setIntersection = append(setIntersection, value)
		}
	}

	fmt.Println(setIntersection)

}
