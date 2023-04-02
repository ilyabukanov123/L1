package dataservice

import (
	"fmt"
)

// Паттерн адаптер - это паттерн, который позволяет объектам какого либо пакета несовместимого друг с другом работать в одном ключе

// data-service это внешний сервис, который будет рабоать и возвращать аналитику в формате xml

// Данный интерфейс работает с аналитикой
type AnalitycalDataService interface {
	SendXmlData()
}

// Данная структура предназначена для отправки документа в формате Xml
type XmlDocument struct {
}

func (doc *XmlDocument) SendXmlData() {
	fmt.Println("Отправка xml документа!")
}
