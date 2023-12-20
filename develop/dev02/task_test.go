package main

import "testing"

func Test_StringUnpacking(t *testing.T) {
	testCases := []struct {
		testCase,
		str,
		expected string
	}{
		{"пустая строка", "", ""},
		{"есть повторяющиеся элементы", "a4bc2d5e", "aaaabccddddde"},
		{"нет повторяющихся элементов", "abcd", "abcd"},
		{"некорректная строка", "45", ""},
		{"есть escape символы", `qwe\4\5`, "qwe45"},
		{"есть escape символы 2", `qwe\45`, "qwe44444"},
		{"есть повторяющиеся escape символы", `qwe\\5`, "qwe\\\\\\\\\\"},
	}

	for index, value := range testCases {
		result, err := StringUnpacking(value.str)
		if result != value.expected {
			t.Errorf("тест №%d (%s) не пройден. Ожидалось: %s, получено: %s, "+
				"ошибка (при возникновении - %s\n",
				index, value.testCase, value.expected, result, err)
		}
	}
}
