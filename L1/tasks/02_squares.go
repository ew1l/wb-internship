package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// С помощью WaitGroup дожидаемся выполнения всех горутин
	wg := new(sync.WaitGroup)

	for _, number := range numbers {
		// Увеличиваем счетчик горутин на единицу
		wg.Add(1)
		// Горутина, рассчитывающая квадрат числа number
		go func(num int) {
			// После выполнения горутины уменьшаем счетчик на единицу
			defer wg.Done()
			fmt.Printf("%d^2 = %d\n", num, num*num)
		}(number)
	}

	// Wait блокируется пока счетчик горутин не станет равным нулю
	wg.Wait()
}
