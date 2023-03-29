package main

import "fmt"

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

// Пустой интерфейс может содержать значения любого типа
func Type(value interface{}) {
	fmt.Printf("Type = %T\n", value)
}

func main() {
	Type(true)
	Type(110)
	Type(make(chan string))
	Type("Hello world")
}
