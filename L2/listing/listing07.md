Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()

	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
3
4
5
6
7
8
0
0
0
0
0
...

В функции main range считывает данные из канала c, которые отправляются из горутины в функции merge. После того как горутина в merge считает все данные из каналов a и b, каналы a и b закроются. Так как каналы a и b закрыты, а горутина в merge не завершит свою работу из-за бесконечного for, в канал с будут отправляться нулевые значения (чтение из закрытого канала возвращает нулевое значение в зависимости от типа данных канала).

Решением будет проверять в горутине функции merge закрыты ли каналы a и b. Если да, то выходим из for и закрываем канал c (закрыть канал c нужно для завершения цикла range).
```