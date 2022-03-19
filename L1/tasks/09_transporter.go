package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	// Создаем канал, который принимает числа из numbers
	n := make(chan int)

	// Создаем канал, который принимает число из n, умноженное на 2
	r := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6}
	wg := new(sync.WaitGroup)

	// Горутина считывает число из numbers и отправляет в канал n
	wg.Add(1)
	go func() {
		defer wg.Done()
		// После считывания всех чисел из numbers закрываем канал
		defer close(n)
		for _, number := range numbers {
			n <- number
		}
	}()

	// Горутина принимает число из канала n и отправляет число, умноженное на 2, в канал r
	wg.Add(1)
	go func() {
		defer wg.Done()
		// После считывания всех чисел из канал n и отправки результата в канал r закрываем канал r
		defer close(r)
		for number := range n {
			r <- number * 2
		}
	}()

	// Горутина выводит полученное число из канала r в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range r {
			fmt.Fprintln(os.Stdout, result)
		}
	}()

	wg.Wait()
}
