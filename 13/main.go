package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	one, two := 10, 20
	one, two = two, one

	fmt.Println(one, two)
}
