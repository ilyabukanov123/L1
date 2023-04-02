package main

import (
	"fmt"
)

// Наш основной сервис работает с файлами только в Json формате. Но сервис аналитики работает с данными в формате XML. Нам надо применить паттерн адаптер и написать адаптер, который позволит конвертировать json в xml и работать с внешним сервисом аналитики
type JsonDocument struct {
}

// Функция конвертации Json объекта в XML структуру
func (doc *JsonDocument) ConvertToXml() {
	fmt.Println("<xml></xml>")
}

type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument // инжектим все поля структуры в наш адаптер. Таким образом мы можем использовать методы и поля
}

// Здесь мы реализуем интерфейс внешнего сервиса аналитики
func (adapter *JsonDocumentAdapter) SendXmlData() {
	// Конвертируем Json в Xml
	adapter.jsonDocument.ConvertToXml()
	// Отправляем XML в сервис аналитики
	fmt.Println("Отправка Xml данных!")
}

func main() {
	adapter := JsonDocumentAdapter{}
	adapter.SendXmlData()
}
