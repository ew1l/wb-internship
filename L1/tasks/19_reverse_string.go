package main

import "fmt"

func ReverseString(str string) string {
	// Конвертируем в rune, так как символы могут быть unicode
	data := []rune(str)

	// Меняем местами первый - последний элемент, второй - предпоследний и тд
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	// Возвращаем перевернутую строку
	return string(data)
}

func main() {
	strs := []string{"some text", "главрыба"}

	for _, str := range strs {
		fmt.Printf("%s - %s\n", str, ReverseString(str))
	}
}

// some text - txet emos
// главрыба - абырвалг
