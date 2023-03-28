package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
5.	Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
func main() {
	// создаем таймер на 1 минуту
	time := time.NewTimer(time.Minute)

	// Создаем канал в который будем писать данные
	values := make(chan int)

	go func() {
		for {
			select {
			// Если таймрер завершил свою работу
			case <-time.C:
				fmt.Println("1 минут прошла и таймер завершен")
				close(values) // закрываем канал
				return
			default:
				values <- rand.Intn(10000) // пишем данные в канал
			}
		}
	}()

	// Перебираем значение из канала
	for value := range values {
		fmt.Println(value)
	}
}
