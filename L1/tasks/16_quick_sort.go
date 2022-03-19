package main

import (
	"fmt"
	"sort"
)

type User struct {
	id   int
	name string
}

type ByName []User

// Реализуем интерфейс sort.Interface
// Пример: сортировка по имени
// Функция Len() возвращает количество элементов в слайсе []User
func (u ByName) Len() int {
	return len(u)
}

// Функция Less определяет, должен ли элемент с индексом i быть отсортированным перед элементом с индексом j
func (u ByName) Less(i, j int) bool {
	return u[i].name < u[j].name
}

// Функция Swap меняет местами элементы с индексами i и j
func (u ByName) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	a := []int{6, 2, 8, 3, 9, 0, 4, 1, 5, 7}
	b := []string{"ba", "aa", "bb", "ab"}

	users := []User{
		{id: 3, name: "Bob"},
		{id: 1, name: "Alice"},
		{id: 2, name: "Eve"},
	}

	// Сортировка слайса целых чисел
	sort.Ints(a)

	fmt.Println(a)

	// Сортировка слайса строк
	sort.Strings(b)

	fmt.Println(b)

	// Сортировка пользовательских структур данных

	// 1 способ. Через sort.SliceStable()
	// sort.SliceStable(), в отличие от sort.Slice(), гарантирует стабильность сортировки (т.е. сохраняет равные элементы в их исходном порядке)
	// Пример: сортировка по id
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].id < users[j].id
	})

	fmt.Println(users)

	// 2 способ. Через sort.Interface
	// sort.Sort(ByName(users)) // Не сохраняет порядок равных элементов
	sort.Stable(ByName(users))

	fmt.Println(users)
}

// [0 1 2 3 4 5 6 7 8 9]
// [aa ab ba bb]
// [{1 Alice} {2 Eve} {3 Bob}]
// [{1 Alice} {3 Bob} {2 Eve}]
