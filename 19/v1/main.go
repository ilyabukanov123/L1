package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

// https://translated.turbopages.org/proxy_u/en-ru.ru.5b60f6b1-6427fc62-46eb3e6e-74722d776562/https/www.geeksforgeeks.org/how-to-reverse-a-string-in-golang/

// https://golang-blog.blogspot.com/2019/09/string-byte-rune-character-golang.html?ysclid=lfxshbyfza688837183

// В Go строка в действительности является срезом (slice) байтов, доступным только для чтения.

// 1 вариант
// Функция переварота строка
func reverse(s string) string {
	// Руна распознает байты как код юникода в int32
	rns := []rune(s) // convert to rune
	fmt.Println(rns)

	// i - 0
	// j - последний элемент строки
	// i, j = i+1, j-1; 1 итерация цикла - первый и последний элемент руны; 2 итерация цикла - второй и предпоследний элемент руны
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// поменяйте местами буквы строки,
		// например, первую на последнюю и так далее.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// возвращаем руну конвертированную в строку
	return string(rns)
}

func main() {
	fmt.Println(reverse("Hello"))

}
