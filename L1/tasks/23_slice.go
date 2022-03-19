package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	// Индекс удаляемого элемента
	i := 3

	// Сдвигаем все элементы после удаляемого влево на один индекс
	copy(slice[i:], slice[i+1:]) // [1 2 3 5 5]

	// Обнуляем последний элемент
	slice[len(slice)-1] = 0 // [1 2 3 5 0]

	// Урезаем до последнего элемента
	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}

// [1 2 3 5]
