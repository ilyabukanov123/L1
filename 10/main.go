package main

import (
	"fmt"
	"math"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func getGroup(value float64) int {
	// Функция Trunc () — это встроенная функция пакета math , которая используется для получения целочисленного значения заданного значения.

	// Пример: 32,5 / 10 - > 3,25 и с помощью функции Trunc возвращаем целую часть
	group := math.Trunc(value / 10)

	// Домножаем на 10 чтоб получить группу
	group *= 10

	return int(group)
	// Домнажаем значение
}

// Создаем именованный тип подмножества
// ключи будут отвечать за параметры температуры, а значения будут пустой структурой
type subset map[float64]struct{}

func main() {

	temperature := [...]float64{25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	manyTemperatures := make(map[int]subset)

	for _, value := range temperature {
		group := getGroup(value)
		//Если группа не существует создаем данную группу
		if _, exists := manyTemperatures[group]; !exists {
			manyTemperatures[group] = make(subset)
		}
		// Получаем map, который отвечает за группу
		subset := manyTemperatures[group]

		// Заполняем подмножество группы
		subset[value] = struct{}{}
		manyTemperatures[group] = subset
	}
	fmt.Println(manyTemperatures)

}
