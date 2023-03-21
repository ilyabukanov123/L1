package main

import "fmt"

/*
	1.	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской
	структуры Human (аналог наследования).
*/

// Объявляем именованный тип Human с 3 поля - фамилия, имя, отчество
type Human struct {
	lastName   string
	firstName  string
	patronymic string
}

// Вариант №1
// Объявляет именнованный тип Action и встраиваем в него тип Human. Тем самым в Action будут доступны поля и методы Human
type ActionOne struct {
	Human // анонимное поле структуры
}

// Вариант №2. Указатель на структуру позволяет сэкономить память т.к именованный тип Human не копируется в памяти при вызове метода
// Объявляет именнованный тип Action и встраиваем в него тип Human. Тем самым в Action будут доступны поля и методы Human
type ActionTwo struct {
	*Human // анонимное поле структуры
}

// Объявляем методы, который будет прикреплен с именованному типу Human
func (h Human) work() {
	fmt.Println("Человек может работать")
}

func (h Human) sleep() {
	fmt.Println("Человек может спать")
}

func main() {
	// Cоздаем переменной с типой структур и заполняем структуру Human
	var actionOne ActionOne = ActionOne{Human{lastName: "Иванов", firstName: "Иван", patronymic: "Иванович"}}
	// var actionTwo ActionTwo = ActionTwo{&Human{lastName: "Иванов", firstName: "Иван", patronymic: "Иванович"}}

	fmt.Printf("Фамилия человека: %v\n", actionOne.lastName)
	fmt.Printf("Имя человека: %v\n", actionOne.firstName)
	fmt.Printf("Отчество человека: %v\n", actionOne.patronymic)
	actionOne.work()
	actionOne.sleep()
}
