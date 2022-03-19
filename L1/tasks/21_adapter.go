package main

import (
	"log"
	"time"
)

// Реализация структурного паттерна "адаптер"
// Карта памяти не может взаимодействовать с компьютером, т.к. у них нет общего интерфейса. У компьютеров есть разъем USB, но мы не можем вставить нашу карту памяти туда.
// Картридер является нашим адаптером, т.к. у него есть USB кабель и его можно вставить в компьютер.

// Структура карты памяти
// Адаптируемый объект
type MemoryCard struct {
	SerialNumber string
}

// Функции вставки и копирования данных с карты памяти
func (mc *MemoryCard) Insert() {
	log.Printf("Memory card %s inserted successfully!\n", mc.SerialNumber)
}

func (mc *MemoryCard) CopyData() {
	time.Sleep(2 * time.Second) // Для визуализации
	log.Println("The data has been copied to the computer!")
}

// "Интерфейс компьютера и картридера"
type USB interface {
	ConnectWithUSB()
}

// Структура картридера
// Адаптер
type CardReader struct {
	// Чтобы наш картридер мог взаимодействовать с нашей картой памяти, передадим ему поле типа MemoryCard
	MemoryCard *MemoryCard
}

// Реализуем интерфейс USB
func (cr *CardReader) ConnectWithUSB() {
	cr.MemoryCard.Insert()
	cr.MemoryCard.CopyData()
}

func main() {
	// Наша карта памяти
	memoryCard := &MemoryCard{SerialNumber: "LXW00K236912743"}

	// Передадим нашу карту памяти в картридер
	cardReader := &CardReader{MemoryCard: memoryCard}

	// "Вставляем картридер в компьютер"
	cardReader.ConnectWithUSB()
}

// 2022/02/25 16:04:56 Memory card LXW00K236912743 inserted successfully!
// 2022/02/25 16:04:58 The data has been copied to the computer!
