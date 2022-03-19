package main

import (
	"context"
	"fmt"
	"time"
)

var time_N time.Duration = 10 * time.Second

func Fibonacci(c chan int) {
	x, y := 0, 1
	for {
		c <- x
		x, y = y, x+y

		// Раз в секунду отправляем число Фибоначчи в канал
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c := make(chan int)

	// Создаем контекст тайм-аута
	ctx, cancel := context.WithTimeout(context.Background(), time_N)
	// Отмена контекста после завершения программы
	defer cancel()

	go Fibonacci(c)

	for {
		select {
		// По истечению тайм-аута сработает контекст
		case <-ctx.Done():
			fmt.Println("Завершение")
			return
		case f_number := <-c:
			fmt.Println(f_number)
		}
	}
}

// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// Завершение
