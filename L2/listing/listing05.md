Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
    msg string
}

func (e *customError) Error() string {
    return e.msg
}

func test() *customError {
    {
        // do something
    }
    return nil
}

func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```

Ответ:
```
error

Создаем переменную err интерфейса error равную nil и помещаем в нее результат выполнения функции test, которая возвращает *customError (customError удовлетворяет интерфейсу error, так как реализована функция Error). Теперь у переменной интерфейса err тип (itable) - *customError, а значение (data) - nil.
```