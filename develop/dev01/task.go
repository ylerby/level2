package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

/*
Создать программу печатающую точное время с использованием NTP -библиотеки.
Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
Требования:
1. Программа должна быть оформлена как go module
2. Программа должна корректно обрабатывать ошибки
библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS

*/

func main() {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "ошибка при получении текущего времени: %s", err.Error())
		if err != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}
	fmt.Printf("Текущее время - %v\n", currentTime)
}
