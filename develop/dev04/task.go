package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству, 'листок', 'слиток' и 'столик' - другому.
Требования:
1. Входные данные для функции: ссылка на массив, каждый
элемент которого - слово на русском языке в кодировке
utf8
2. Выходные данные: ссылка на мапу множеств анаграмм
3. Ключ - первое встретившееся в словаре слово из
множества. Значение - ссылка на массив, каждый элемент
которого,
слово из множества.
4. Массив должен быть отсортирован по возрастанию.
5. Множества из одного элемента не должны попасть в
результат.
6. Все слова должны быть приведены к нижнему регистру.
7. В результате каждое слово должно встречаться только один
раз.
*/

func main() {
	// инициализируем слайс строк
	stringValues := &[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "тест", "естт",
		"слово", "волос", "ноль", "один", "два", "птаяк", "тапяк", "капят", "апктя"}
	result := getAnagramGroups(stringValues)
	fmt.Printf("Результат: %v\n", result)
}

func getAnagramGroups(stringValues *[]string) map[string][]string {
	// инициализируем map`у для создания групп анаграмм
	anagramMap := make(map[string][]string)
	// итерируемся по слайсу
	for _, value := range *stringValues {
		// приводим значение к нижнему регистру
		value = strings.ToLower(value)

		isEmpty := true

		// итерируемся по ключам map`ы
		for key := range anagramMap {
			// если ключ в map`е и значение из слайса являются анаграммами
			if isAnagramGroup(key, value) {
				// то добавляем значение (строку) в слайс по ключу
				anagramMap[key] = append(anagramMap[key], value)
				isEmpty = false
			}
		}

		// создаем новое значение для строки
		if isEmpty {
			// инициализируем значение пустым слайсом
			anagramMap[value] = []string{}
		}
	}

	for key, value := range anagramMap {
		// если для ключа значением является пустой слайс, то удаляем значение
		if len(anagramMap[key]) == 0 {
			delete(anagramMap, key)
		}

		// сортируем значения в map`е
		sort.Strings(value)
	}

	return anagramMap
}

func isAnagramGroup(groupKey, groupValue string) bool {
	isAnagramMap := make(map[rune]int)

	// итерируемся по строке и инкрементируем значение
	for _, letter := range groupKey {
		isAnagramMap[letter]++
	}

	// итерируемся по второй строке
	for _, letter := range groupValue {
		_, isOk := isAnagramMap[letter]
		// если значение из второй строке не найдено в map`е, то строки не являются анаграммами
		if !isOk {
			return false
		}

		// декрементируем значение
		isAnagramMap[letter]--

		// если значение меньше нуля, то число повторяющихся символов в строках не равно
		if isAnagramMap[letter] < 0 {
			return false
		}
	}

	// итерируемся по map`е и если значение не равно нулю, то слова не являются анаграммами
	for _, value := range isAnagramMap {
		if value != 0 {
			return false
		}
	}

	return true
}
