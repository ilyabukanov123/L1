package main

import "fmt"

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func OneBitToNull(num *int64, position uint) {
	//Используем оператор деления по модулю 2 (XOR)
	*num ^= 1 << position
}

func main() {
	var value int64 = 8
	// %.8b выводит 8 бит
	fmt.Printf("%.8b\n", value)
	OneBitToNull(&value, 3)
	fmt.Printf("%.8b\n", value)
}
