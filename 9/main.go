package main

import (
	"fmt"
)

/*
Разработать конвейер чисел. Даны два канала:
в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	// Канал для целый чисел из массива
	writeNumberArray := make(chan int)
	// Канал для домножения целый числе из массива на 2
	writeSqrtNumberArray := make(chan int)
	array := [5]int{2, 4, 6, 8, 10}

	// Горутина на запись данных из массива в канал
	go func() {
		for _, value := range array {
			writeNumberArray <- value
		}
		close(writeNumberArray)
	}()

	// Горутина на чтение данных из канала, домножение данных на 2, запись во 2 канал
	go func() {
		for value := range writeNumberArray {
			writeSqrtNumberArray <- value * 2
		}
		close(writeSqrtNumberArray)
	}()

	// Чтение данных из 2 канала
	for value := range writeSqrtNumberArray {
		fmt.Println(value)
	}
}
