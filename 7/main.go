package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
	Реализовать конкурентную запись данных в map.
*/

type MapMutex struct {
	data map[int]int
	mu   sync.Mutex
	mg   sync.WaitGroup
}

// Метод для записи данных в map
func (m *MapMutex) SetMap(key int, value int) {
	m.mu.Lock()         // Блокируем область в памяти
	defer m.mu.Unlock() // Разблокируем область в памяти
	defer m.mg.Done()
	m.data[key] = value
}

func main() {
	// Объявляем и инициализируем хранилище для map
	mapMutex := MapMutex{data: make(map[int]int)}
	for i := 0; i < 1000; i++ {
		mapMutex.mg.Add(1)
		go mapMutex.SetMap(rand.Intn(10), rand.Intn(100))
	}
	mapMutex.mg.Wait()
	fmt.Println(mapMutex.data)

}
