package main

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	temp := make(map[rune]struct{})

	// Переводим все в нижний регистр
	for _, rn := range strings.ToLower(str) {
		// Если rn есть в temp, возвращаем false
		if _, ok := temp[rn]; ok {
			return false
		}

		temp[rn] = struct{}{}
	}

	return true
}

func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd", "тест", "привет", "abBcd"}

	for _, str := range strs {
		fmt.Printf("%s - %t\n", str, IsUnique(str))
	}
}

// abcd - true
// abCdefAaf - false
// aabcd - false
// тест - false
// привет - true
// abBcd - false
