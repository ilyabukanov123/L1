package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.
// Пример реализован с помощью грокаем алгоритмы

func binary_search(list []int, item int) (position int, error bool) {
	// В переменной low и high хранятся границы той части списка, в которой выполнятся поиск
	low := 0
	high := len(list) - 1

	// Пока граница не скоратится до одного элемента
	for low <= high {
		// Проверяем средний элемент
		mid := low + high
		// Получаем средний элемент
		guess := list[mid]

		// Если значение найдено
		if guess == item {
			fmt.Println(mid)
			return mid, true
		}

		// слишком большой
		if guess > item {
			high = mid - 1
			// слишком маленький
		} else {
			low = mid + 1
		}
	}
	return 0, false
}

func main() {
	my_list := []int{1, 3, 5, 7, 9}
	position, ok := binary_search(my_list, 3)
	//Если ошибка, то значение было не найдено
	if ok {
		fmt.Printf("Позиция 3 элемента в слайсе: %v", position)
	} else {
		fmt.Print("Значение 3 в слайсе не найдено")
	}
}
