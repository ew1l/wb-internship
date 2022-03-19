package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)

	// Создаем буферизованный канал (без блокировки горутины) со значением буфера равным длине numbers
	c := make(chan int, len(numbers))
	sum_sq := 0

	for _, number := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			// Отправляем квадрат числа number в канал
			c <- num * num
		}(number)
	}

	wg.Wait()

	// Закрываем канал
	close(c)

	// Читаем данные с помощью for range
	for num := range c {
		sum_sq += num
	}

	// Аналог:
	// for {
	// 	if num, ok := <-c; ok {
	// 		sum_sq += num
	// 	} else {
	// 		break
	// 	}
	// }

	fmt.Println(sum_sq)
}
