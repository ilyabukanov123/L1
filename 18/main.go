package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

// Cоздаем именованный тип счетчика
type Counter struct {
	i  int
	mu sync.Mutex
	wg sync.WaitGroup
}

// Создаем метод для инкреминтации значений
func (c *Counter) inc() {
	defer c.wg.Done() // Завершаем работу текущей задачи
	//Блокируем доступ к участку памяти для всех горутин, кроме текущей
	c.mu.Lock()

	// Инкрементируем значение счетчика
	c.i++

	//Разрешаем доступ к участку памяти для всех горутин, кроме текущей
	c.mu.Unlock()
}

func main() {

	counter := Counter{i: 0}
	for i := 0; i < 1000; i++ {
		counter.wg.Add(1) // Добавляем задачу
		go func() {
			counter.inc()
		}()
	}
	counter.wg.Wait() // Ждем пока все горутины завершат свою работу
	fmt.Printf("Счетчик равен: %v", counter.i)
}