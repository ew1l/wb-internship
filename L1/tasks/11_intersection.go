package main

import "fmt"

func main() {
	set_1 := []int{1, 2, 3, 4}
	set_2 := []int{3, 4, 5, 6}
	intersection := make([]int, 0)

	for _, value_1 := range set_1 {
		for _, value_2 := range set_2 {
			// Проверка на совпадение значения из первого множества со значением из второго множества
			// Если совпали, кладем значение в intersection
			if value_1 == value_2 {
				intersection = append(intersection, value_1)
			}
		}
	}

	// Пересечение двух множеств
	fmt.Printf("%v & %v => %v\n", set_1, set_2, intersection)
}

// [1 2 3 4] & [3 4 5 6] => [3 4]
