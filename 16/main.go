package main

import "fmt"

func main() {
	array := []int{6, 5, 1, 3, 8, 4, 7, 9, 2}
	arraySort := quicksort(array)
	fmt.Println(arraySort)
}

func quicksort(array []int) []int {
	// На примере грокаем алгоритмы

	// Если длинна массива состоит из 2 элементов(0,1) то его уже не нужно сортировать
	if len(array) < 2 {
		return array
	} else {
		// Устанавливаем последний элемент в качестве опорного
		pivot := array[0]

		// Числа меньше опорного
		var les []int

		// Числа больше опорного
		var more []int

		// Начинаем перебор с 1 элемента массива т.к. 0 элемент мы взяли в качестве опорного
		for _, value := range array[1:] {
			if value <= pivot {
				les = append(les, value)
			}
		}

		// Начинаем перебор с 1 элемента массива т.к. 0 элемент мы взяли в качестве опорного
		for _, value := range array[1:] {
			if value > pivot {
				more = append(more, value)
			}
		}

		// Далее нужно взять опорный элемент из les/more и проделать те же действия
		var middle []int
		middle = append(middle, pivot)

		// Объединяем 3 слайса путем вызова рекурсивной функции
		array = append(quicksort(les), append(middle, quicksort(more)...)...)
		return array
	}
}
