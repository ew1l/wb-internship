package main

import "fmt"

// Функция, определяющая тип переменной. В качестве аргумента передаем interface{}
func TypeFor(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan string:
		return "chan string"
	default:
		return "unknown"
	}
}

func main() {
	a := 42
	b := "test"
	c := false
	d := make(chan string)

	fmt.Println(TypeFor(a))
	fmt.Println(TypeFor(b))
	fmt.Println(TypeFor(c))
	fmt.Println(TypeFor(d))
}

// int
// string
// bool
// chan string
