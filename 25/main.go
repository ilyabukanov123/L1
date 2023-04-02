package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
// https://golang-blog.blogspot.com/2020/07/package-time-in-golang-duration.html?ysclid=lfzozep0sv314013966

// Duration представляет прошедшее время между двумя моментами в виде подсчета int64 наносекунд. Представление ограничивает наибольшую представляемую продолжительность приблизительно 290 годами.

func Sleep(duration time.Duration) {
	// Устанавливаем время работы таймера
	timer := time.NewTimer(duration)
	// Пишем сообщение в канал как только таймер завершен
	<-timer.C
	fmt.Println("Работа таймера завершена")
}

func main() {
	// Задаем время работы таймера
	Sleep(10 * time.Second)
}
