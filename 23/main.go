package main

import "fmt"

// Функция сдвига элементов слайса
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// func remove[T any](slice []T, s int) []T {
// 	return append(slice[:s], slice[s+1:]...)
// }

// Удалить i-ый элемент из слайса.
// Элементы из Slice нельзя удалять. Мы должны сдвинуть все элементы справа от индекса удаления на один
// Можно использовать T как дженерик для любого типа(как в java)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("Cлайс: ", slice)
	slice = remove(slice, 3)
	fmt.Println("Обновленный слайс с удаленным элементом под индексом 3", slice)
}
