Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}
```

Ответ:
```
[77 78 79]

С помощью операции среза создается новый слайс на 3 элемента (с 1ого индекса по 4ый не включительно).
```