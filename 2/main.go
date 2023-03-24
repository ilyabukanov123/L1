package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение
квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {

	var wg sync.WaitGroup           // объявляем wait group
	array := [5]int{2, 4, 6, 8, 10} // объявляем массив
	for _, value := range array {   // перебираем элементы массива
		wg.Add(1)           // добавляем количество задач на выполнение
		go sqrt(value, &wg) // запускам горутину на выполнение функции
	}

	wg.Wait() // ожидаем завершение всех горутин из группы wg
}

func sqrt(number int, wg *sync.WaitGroup) {
	fmt.Printf("%v в степени 2 : %v\n", number, math.Pow(float64(number), 2))
	wg.Done() // указываем на завершение запущенной задачи
}
