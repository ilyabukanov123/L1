package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

func reverse(s string) string {
	//Fields разбивает строку s вокруг каждого экземпляра одного или нескольких последовательных символов пробела, как определено в unicode.IsSpace, возвращая часть подстрок строки s или пустую часть, если s содержит только пробелы.

	temp := strings.Fields(s)

	// i - 0 объекта массива temp
	// j - последний объект массива temp
	// i, j = i+1, j-1; 1 итерация цикла - первый и последний объект массива temp; 2 итерация цикла - второй и предпоследний объект массива temp
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {

		// поменяйте местами объекты массива temp,
		// например, первый на последний и так далее.
		temp[i], temp[j] = temp[j], temp[i]
	}
	// fmt.Printf("%T", temp)
	justString := strings.Join(temp, " ")
	return justString
}

func main() {
	var str string
	fmt.Print("Введите строку:")
	fmt.Scanf("%q", &str)
	fmt.Println("Введенная строка:", str)
	str = reverse(str)
	fmt.Println("Перевернутая строка:", str)
}
