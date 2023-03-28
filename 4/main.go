package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте. Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	// cоздаем контекст, который возвращает произвольный контекст и функцию отмены
	ctx, cancel := context.WithCancel(context.Background())

	// инициализируем WaitGroup
	var wg sync.WaitGroup

	// Cоздаем переменную под количество воркеров
	var numberWorkers int

	fmt.Println("Введите количество воркеров: ")
	fmt.Scanf("%d\n", &numberWorkers) // считываем данные из консоли

	// Создаем канал в который воркеры будут писать значения
	values := make(chan int)

	// results := make(chan int, numberWorkers)

	// Запускаем горутины на работу воркеров
	for i := 0; i < numberWorkers; i++ {
		wg.Add(1)                        // создаем одну задачу
		go worker(ctx, i+1, values, &wg) // запускам горутину
	}

	// Обработка сигнала завершения программы
	sigint := make(chan os.Signal, 1)

	//Подписываемся на уведомления от ОС, был ли применен сигнал Sigint(ctrl + c) и кладем это все в канал sigint
	signal.Notify(sigint, syscall.SIGINT, os.Interrupt)

	for {
		select {
		// Закрываем контекст и завершаем работу если пришел Sigint сигнал
		case <-sigint:
			close(values) // закрываем канал
			cancel()      // вызываем функцию на закрытие контекста
			wg.Wait()
			return
		default:
			values <- rand.Intn(10000) // пишем данные в канал
		}
	}
}

// Принимаем контекст, id, канал из которого сможем только читать, указатель на WaitGroup
func worker(ctx context.Context, id int, values <-chan int, wg *sync.WaitGroup) {

	for {
		select {
		// Если поступил сигнал в контекст о завершении завершаем работу воркера
		case <-ctx.Done():
			fmt.Printf("Worker №%d завершил работу\n", id)
			wg.Done() // Сообщаем о завершении задачи
			return    // выходим из функции
		// Если воркер не завершил работу считываем значение из канала
		default:
			fmt.Printf("Worker №%d прочитал информацию из канала - %d\n", id, <-values)
		}
	}
}
