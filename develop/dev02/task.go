package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки,
содержащую повторяющиеся символы/руны, например:

  "a4bc2d5e"=>"aaaabccddddde"
  "abcd"=>"abcd"
  "45"=>""(некорректнаястрока)
  ""=>""

Дополнительно
Реализовать поддержку escape-последовательностей. Например:

  "qwe\4\5" => "qwe45"
  "qwe\45"  => "qwe44444"
  "qwe\\5"  => "qwe\\\\\"

В случае если была передана некорректная строка, функция должна возвращать ошибку.
Написать unit-тесты.
*/

// константа для ASCII-кода escape-символа
const escapeASCII = 92

func StringUnpacking(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	// преобразуем строку в слайс рун и создаем builder для создания результирующей строки
	var (
		runeSlice = []rune(str)
		builder   strings.Builder
	)

	// проверяем, является ли первый элемент числом
	if unicode.IsDigit(runeSlice[0]) {
		return "", fmt.Errorf("некорректная строка")
	}

	// итерируемся по элементам слайсу
	for i := 0; i < len(runeSlice); i++ {
		// если символ является escape-символом, то прибавляем i++ (переходим к следующему элементу)
		if runeSlice[i] == escapeASCII {
			i++
		}

		// проверяем, что строка не закончилась
		if i+1 != len(runeSlice) {
			// проверяем, является ли следующий элемент числом
			if unicode.IsDigit(runeSlice[i+1]) {
				// если следующий элемент - число, то конвертируем его сначала в строку, а потом в число
				num, err := strconv.Atoi(string(runeSlice[i+1]))
				if err != nil {
					return "", fmt.Errorf("некорректная строка")
				}
				// добавляем в результирующую строку num элементов
				for j := 0; j < num; j++ {
					builder.WriteRune(runeSlice[i])
				}
				// переходим к следующему элементу, чтобы число не было добавлено в результирующую строку
				i++
			} else {
				// добавляем элемент в строку
				builder.WriteRune(runeSlice[i])
			}
		} else {
			// добавляем элемент в строку
			builder.WriteRune(runeSlice[i])
		}
	}
	// возвращаем полученную строку
	return builder.String(), nil
}
