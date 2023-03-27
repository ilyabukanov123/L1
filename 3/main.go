package main

import (
	"fmt"
	"math"
)

/*
 3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
    с использованием конкурентных вычислений.
*/
func main() {

	msg := make(chan float64) // создаем канал
	array := [5]int{2, 4, 6, 8, 10}
	go func() { // запускаем 1 горутину
		for _, value := range array {
			msg <- (sqrt(value)) // пишем в канал то, что вернет функция sqrt
		}
		close(msg)
	}()

	sumSqrt := 0
	for value := range msg { // перебираем значение из канала msg
		fmt.Printf("Квадратный корень: %v\n", value)
		sumSqrt += int(value)
	}
	fmt.Printf("Сумма квадратов: %v", sumSqrt)

}

func sqrt(number int) float64 {
	return math.Pow(float64(number), 2) // возводим number во 2 степень
}
