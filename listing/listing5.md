#### 5. Что выведет программа? Объяснить вывод программы.
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

Выводом будет: 
```go
error
```

Тип error является интерфейсным типом, а значит содержит 
два поля (data, tab). 
В данном случае возвращаемая ошибка не равна nil, 
потому что в переменной err хранится информация о типе. 
Поле data интерефейса не равно nil.

