package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort
(man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками,
на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:
-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке -u — не выводить повторяющиеся строки

*/

type SortUtil struct {
	sortedColumn int
	sortingByNumber,
	reversedSorting,
	uniqueFields bool
	inputFile,
	outputFile string
	data [][]string
}

func main() {
	sortUtil := &SortUtil{}

	flag.IntVar(&sortUtil.sortedColumn, "k", 0, "sorted column")
	flag.BoolVar(&sortUtil.sortingByNumber, "n", false, "sorting by number")
	flag.BoolVar(&sortUtil.reversedSorting, "r", false, "reversed sorting")
	flag.BoolVar(&sortUtil.uniqueFields, "u", false, "unique fields")
	flag.Parse()
	filesArgs := flag.Args()

	switch len(filesArgs) {
	case 1:
		sortUtil.inputFile = filesArgs[0]
	case 2:
		sortUtil.inputFile = filesArgs[0]
		sortUtil.outputFile = filesArgs[1]
	default:
		os.Exit(1)
	}

	err := sortUtil.sort()
	if err != nil {
		fmt.Printf("ошибка при сортировке - %s", err)
		os.Exit(5)
	}
}

func (s *SortUtil) sort() error {
	err := s.readData()
	if err != nil {
		return fmt.Errorf("ошибка при чтении из файла - %s", err)
	}

	for _, value := range s.data {
		if s.sortedColumn > len(value) {
			os.Exit(3)
		}
	}

	sort.Slice(s.data, func(i, j int) bool {
		if !s.sortingByNumber {
			if !s.reversedSorting {
				return s.data[i][s.sortedColumn-1] < s.data[j][s.sortedColumn-1]
			}

			return s.data[i][s.sortedColumn-1] > s.data[j][s.sortedColumn-1]
		}

		firstValue, err := strconv.Atoi(s.data[i][s.sortedColumn-1])
		if err != nil {
			return false
		}

		secondValue, err := strconv.Atoi(s.data[j][s.sortedColumn-1])
		if err != nil {
			return false
		}

		if !s.reversedSorting {
			return firstValue < secondValue
		}

		return firstValue > secondValue
	})

	if s.outputFile != "" {
		err = s.writeFile()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SortUtil) readData() error {
	s.data = make([][]string, 0)
	file, err := os.Open(s.inputFile)
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла - %s", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			os.Exit(2)
		}
	}(file)

	uniqueFieldMap := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if s.uniqueFields {
			if !uniqueFieldMap[line] {
				uniqueFieldMap[line] = true
				s.data = append(s.data, strings.Split(line, " "))
			} else {
				continue
			}
		} else {
			s.data = append(s.data, strings.Split(line, " "))
		}
	}

	return nil
}

func (s *SortUtil) writeFile() error {
	file, err := os.Create(s.outputFile)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			os.Exit(4)
		}
	}(file)

	for _, val := range s.data {
		line := strings.Join(val, " ") + "\n"
		_, err = file.Write([]byte(line))
		if err != nil {
			return err
		}
	}

	return nil
}
