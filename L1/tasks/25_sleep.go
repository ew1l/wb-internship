package main

import "time"

func Sleep(d time.Duration) {
	// Используем функцию After, которая ожидает истечения продолжительности d и возвращает текущее время по каналу
	<-time.After(d)
}

func main() {
	// Засыпаем на определенное время
	Sleep(2 * time.Second)
}
