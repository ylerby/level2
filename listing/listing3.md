#### 3. Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go   
package main

import (
   "fmt"
   "os" 
)

func Foo() error {
    var err *os.PathError = nil
    return err
}

func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}
```

Выводом будет:
```go
<nil>
false
```

В строке:
```go
fmt.Println(err == nil)
```
производится сравнение интерфейсов, а именно их динамическое значение и динамический тип, 
у них различаются динамические типы, поэтому выводом будет:
```go
false 
```