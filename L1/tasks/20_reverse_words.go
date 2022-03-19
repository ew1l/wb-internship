package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	// Делим строку на слова и кладем в слайс data
	data := strings.Split(str, " ")

	// Переворачиваем слайс к в 19 задании
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	// Объединяем слова через пробел в строку
	return strings.Join(data, " ")
}

func main() {
	str := "snow dog sun"

	reverse_str := ReverseWords(str)

	fmt.Printf("%s - %s", str, reverse_str)
}

// snow dog sun - sun dog snow
