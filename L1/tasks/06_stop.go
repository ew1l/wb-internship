package main

import "fmt"

// ############################################

// 1 способ. Отдельный канал
// func main() {
// 	stop := make(chan struct{})

// 	go func() {
// 		fmt.Println("Выполнение горутины")

// 		select {
// 		case stop <- struct{}{}:
// 			fmt.Println("Завершение горутины")
// 			return
// 		}
// 	}()

// 	<-stop
// }

// ############################################

// 2 способ. Основной канал
func main() {
	c := make(chan int)

	go func() {
		counter := 0
		for {
			select {
			case c <- counter:
				counter++
			case <-c:
				// При попадании данных в канал выходим из горутины
				return
			}
		}
	}()

	fmt.Println(<-c) // 0
	fmt.Println(<-c) // 1
	fmt.Println(<-c) // 2

	c <- 0

	// fmt.Println(<-c) // fatal error: all goroutines are asleep - deadlock!
}

// ############################################
