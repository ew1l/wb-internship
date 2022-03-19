package main

import (
	"fmt"
	"sync"
)

func main() {
	squares := make(map[int]int, 0)
	wg := new(sync.WaitGroup)

	// Создаем мьютекс
	mu := new(sync.Mutex)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Блокируем доступ для остальных горутин
			mu.Lock()
			squares[i] = i * i
			// Разблокировка
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(squares)
}

// map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64 9:81 10:100]
