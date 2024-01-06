package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type CutUtil struct {
	fields    []int
	delimiter string
	separated bool
}

func main() {
	cutUtil := CutUtil{}
	field := flag.String("f", "", "fields")
	flag.StringVar(&cutUtil.delimiter, "d", "\t", "another delimiter")
	flag.BoolVar(&cutUtil.separated, "s", false, "separated")
	flag.Parse()

	if *field == "" {
		fmt.Println("некорректное значение для поля")
		os.Exit(1)
	}

	var err error
	cutUtil.fields, err = cutUtil.parseColumns(*field)
	if err != nil {
		fmt.Printf("Ошибка при разборе флага -f: %v\n", err)
		os.Exit(1)
	}

	var result string

	if flag.NArg() < 1 {
		f := os.Stdin
		result, err = cutUtil.Cut(f)
	} else {
		filename := flag.Arg(0)
		var f *os.File
		f, err = os.Open(filename)
		if err != nil {
			_, err = fmt.Fprint(os.Stderr, "error:\t", err)
			if err != nil {
				os.Exit(1)
			}
			return
		}
		result, err = cutUtil.Cut(f)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func (c *CutUtil) Cut(f *os.File) (string, error) {
	scanner := bufio.NewScanner(f)
	var columns []string
	var resultColumns []string
	var result string

	for scanner.Scan() {
		str := scanner.Text()
		columns = strings.Split(str, c.delimiter)
		if c.separated && !strings.Contains(str, c.delimiter) {
			continue
		}
		for i := range c.fields {
			if c.fields[i] > len(columns) {
				errMessage := fmt.Sprintf("отсутствует колонка номер %d", c.fields[i])
				return "", fmt.Errorf(errMessage)
			}
			resultColumns = append(resultColumns, columns[c.fields[i]-1])
		}
		result += strings.Join(resultColumns, c.delimiter)
		resultColumns = resultColumns[:0]
		result += "\n"
	}
	return result, nil
}

func (c *CutUtil) parseColumns(input string) ([]int, error) {
	var columns []int
	ranges := strings.Split(input, ",")

	for _, r := range ranges {
		if strings.Contains(r, "-") {
			bounds := strings.Split(r, "-")
			if len(bounds) != 2 {
				return nil, fmt.Errorf("неверный формат диапазона: %s", r)
			}

			start, err := strconv.Atoi(bounds[0])
			if err != nil {
				return nil, fmt.Errorf("неверный формат начала диапазона: %s", bounds[0])
			}

			end, err := strconv.Atoi(bounds[1])
			if err != nil {
				return nil, fmt.Errorf("неверный формат конца диапазона: %s", bounds[1])
			}

			for i := start; i <= end; i++ {
				columns = append(columns, i)
			}
		} else {
			col, err := strconv.Atoi(r)
			if err != nil {
				return nil, fmt.Errorf("неверный формат колонки: %s", r)
			}
			columns = append(columns, col)
		}
	}
	return columns, nil
}
