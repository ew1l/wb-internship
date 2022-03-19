package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Структура-счетчик
type Counter struct {
	mu    sync.Mutex
	value int
}

// Функция, инициализирующая счетчик
func NewCounter() *Counter {
	return &Counter{
		value: 0,
	}
}

// Функция, инкрементирующая значение счетчика
func (c *Counter) Inc() {
	defer wg.Done()
	// Блокируем инкремент для остальных горутин
	c.mu.Lock()
	c.value++
	// Разблокировка
	c.mu.Unlock()
}

// Функция, возвращающая значение счетчика
func (c *Counter) Value() int {
	return c.value
}

func main() {
	// Инициализируем счетчик
	counter := NewCounter()

	// Запускаем горутины
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter.Inc()
	}

	// Дожидаемся выполнения горутин
	wg.Wait()

	fmt.Println(counter.Value())
}

// 1000
