package main

import "fmt"

// v2
func reverse(str *string) (result string) {
	//buffer := []rune(*str)
	// В Go присутствует синтаксический сахар при обходе строки.
	// Если использовать конструкцию for range,
	// строка автоматически будет преобразована в []rune,
	//  то есть обход будет по Юникод символам,
	// поэтому нужно преобразовывать результат в string
	for _, symbol := range *str {
		result = string(symbol) + result
	}
	return
}

func main() {
	var str string
	fmt.Print("Введите строку:")
	fmt.Scanf("%q", &str)
	fmt.Println("Введенная строка:", str)
	str = reverse(&str)
	fmt.Println("Перевернутая строка:", str)
}
