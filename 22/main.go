package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	// Создаем big int, который равен 2 в 20 степени
	var a = big.NewInt(2 << 20)
	// Создаем big int, который равен 2 в 25 степени
	var b = big.NewInt(2 << 25)

	// результат работы с а и b
	var bigInt big.Int
	fmt.Printf("A = %v, B = %v\n", a, b)
	fmt.Printf("A * B = %v\n", bigInt.Mul(a, b))
	fmt.Printf("A / B = %v\n", bigInt.Div(a, b))
	fmt.Printf("A + B = %v\n", bigInt.Add(a, b))
	fmt.Printf("A - B = %v\n", bigInt.Sub(a, b))
}
