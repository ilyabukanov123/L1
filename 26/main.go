package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
//
//	aabcd — false

func Unique(s string) bool {
	// Переводим строку в нижний регистр
	s = strings.ToLower(s)

	for _, value := range s {
		// Count подсчитывает количество неперекрывающихся экземпляров substr в s. Если substr - пустая строка, Count возвращает 1 + количество кодовых точек Unicode в s.

		// Переводим байт в строку и сравниваем новый байт в каждой итерации цикла
		// https://golang-blog.blogspot.com/2020/09/strings-count-fields.html?ysclid=lfzpqsjaey683514275

		if strings.Count(s, string(value)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	// var str string
	// fmt.Print("Введите строку:")
	// fmt.Scanf("%q", &str)
	// Unique(str)

	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(str1, Unique(str1))
	fmt.Println(str2, Unique(str2))
	fmt.Println(str3, Unique(str3))
}
