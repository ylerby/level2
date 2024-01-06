#### 2. Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.
```go
package main

import (
   "fmt"
)

func test() (x int) {
   defer func() {
       x++ 
   }()
   x=1
   return 
}
   
func anotherTest() int {
   var x int
   defer func() {
       x++
   }() 
   x=1 
   return x
}

func main() {
   fmt.Println(test())
   fmt.Println(anotherTest())
}
```

Будет выведено:
```go
2
1
```
defer позволяет отложить вызов функции до момента выхода из текущей функции

В функции test в defer будет инкрементирована возвращаемая переменная x,
а в функции anotherTest в defer будет инкрементирована внутренняя переменная x
