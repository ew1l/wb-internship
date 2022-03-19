package main

import "fmt"

func ToSet(sequence []string) []string {
	// Создаем слайс, который будет нашим итоговым множеством
	set := make([]string, 0)

	// Создаем промежуточную map для хранения уникальных значений sequence
	// В качестве значения в map будем использовать пустую структуру, так она не занимает места в памяти
	temp := make(map[string]struct{})

	for _, value := range sequence {
		// Проверяем, есть ли value в map
		if _, ok := temp[value]; !ok {
			// Если нет, кладем в map
			// Также кладем наше value в set
			temp[value] = struct{}{}

			// fmt.Println(unsafe.Sizeof(temp[value])) // 0
			set = append(set, value)
		}
	}

	// Возвращаем итоговое множество
	return set
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Конвертируем во множество
	set := ToSet(sequence)

	fmt.Println(set)
}

// [cat dog tree]
