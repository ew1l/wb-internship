package main

import (
	"fmt"
	"sort"
)

func main() {
	// Последовательность, в которой будем искать x
	// Слайс отсортирован
	sequence := []int{12, 21, 34, 51, 68, 72, 76, 85, 97, 121}

	// Искомый элемент
	x := 97

	// Размер слайса
	n := len(sequence)

	// Будем использовать стандартную функцию поиска sort.Search()
	// Данная функция возвращает индекс найденного элемента
	// Если индекс не найден, функция возвращает n
	if index := sort.Search(n,
		func(i int) bool { return x <= sequence[i] }); sequence[index] == x && index < len(sequence) {
		fmt.Printf("Индекс элемента %v: %d\n", x, index)
	} else {
		fmt.Printf("Индекс элемента %v не найден!", x)
	}
}

// Индекс элемента 97: 8
