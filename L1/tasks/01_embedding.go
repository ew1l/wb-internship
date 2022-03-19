package main

import "fmt"

// Родительская структура Human
type Human struct {
	Name string
	Age  int
}

// Метод родительской структуры SayHi
func (h Human) SayHi() {
	fmt.Printf("Hi! My name's %s, i'm %d\n", h.Name, h.Age)
}

// Встраивание метода SayHi в структуру Action
type Action struct {
	Human
}

func main() {
	h := Human{"Emil", 21}
	a := Action{h}

	a.SayHi() // Hi! My name's Emil, I'm 21
}
