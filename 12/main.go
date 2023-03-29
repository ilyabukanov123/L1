package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

// Создаем именованный тип для множества т.к все ключи являются уникальными
type Set map[string]struct{}

func main() {
	// Создаем слайст строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// создаем множество из map
	set := make(Set)

	// Перебираем все значения из слайса. Т.К в map все ключи являются уникальными дублирующие строки не будут добавлятся в map
	for _, value := range strings {
		set[value] = struct{}{}
	}

	fmt.Println(set)

}
