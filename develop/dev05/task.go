package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"flag"
)

/*
	Реализовать утилиту фильтрации (man grep)

	Поддержать флаги:
	-A - "after" печатать +N строк после совпадения
	-B - "before" печатать +N строк до совпадения
	-C - "context" (A+B) печатать ±N строк вокруг совпадения
	-c - "count" (количество строк)
	-i - "ignore-case" (игнорировать регистр)
	-v - "invert" (вместо совпадения, исключать)
	-F - "fixed", точное совпадение со строкой, не паттерн
	-n - "line num", печатать номер строки

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type GrepUtil struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func main() {
	grepUtil := GrepUtil{}
	lines := make([]string, 0)
	flag.IntVar(&grepUtil.after, "A", 0, "+N after")
	flag.IntVar(&grepUtil.before, "B", 0, "+N before")
	flag.IntVar(&grepUtil.context, "C", 0, "+-N")
	flag.BoolVar(&grepUtil.count, "c", false, "count")
	flag.BoolVar(&grepUtil.ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&grepUtil.invert, "v", false, "invert")
	flag.BoolVar(&grepUtil.fixed, "F", false, "fixed")
	flag.BoolVar(&grepUtil.lineNum, "n", false, "line number")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("")
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	var re *regexp.Regexp
	var err error
	if grepUtil.fixed {
		re, err = regexp.Compile(regexp.QuoteMeta(pattern))
	} else {
		if grepUtil.ignoreCase {
			pattern = "(?i)" + pattern
		}
		re, err = regexp.Compile(pattern)
	}
	if err != nil {
		fmt.Printf("ошибка при использовании regex - %s", err)
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("ошибка при открытии файла - %s", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			os.Exit(1)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	grepUtil.grep(grepUtil, lines, re)

}

func (g *GrepUtil) grep(grepUtil GrepUtil, lines []string, pattern *regexp.Regexp) []string {
	resultLines := make([]string, 0)
	count := 0
	lastBeforeIndex := -1

	for index := 0; index < len(lines); index++ {
		match := (pattern.MatchString(lines[index]) && !grepUtil.invert) || (!pattern.MatchString(lines[index]) && grepUtil.invert)
		if grepUtil.count && match {
			count++
		} else if match {
			if grepUtil.context > 0 {
				start := max(0, index-grepUtil.context)
				if start <= lastBeforeIndex {
					start = lastBeforeIndex + 1
				}
				for i := start; i < min(len(lines), index+grepUtil.context+1); i++ {
					resultLines = append(resultLines, grepUtil.addToResult(grepUtil.lineNum, lines[i], i))
				}
				lastBeforeIndex = index + grepUtil.context
			} else if grepUtil.before > 0 {
				start := max(0, index-grepUtil.before)
				if start <= lastBeforeIndex {
					start = lastBeforeIndex + 1
				}
				for i := start; i <= index; i++ {
					resultLines = append(resultLines, grepUtil.addToResult(grepUtil.lineNum, lines[i], i))
				}
				lastBeforeIndex = index
			} else if grepUtil.after > 0 {
				i := index
				for ; i < min(len(lines), index+grepUtil.after+1); i++ {
					resultLines = append(resultLines, grepUtil.addToResult(grepUtil.lineNum, lines[i], i))
				}
				index = i - 1
			} else {
				resultLines = append(resultLines, grepUtil.addToResult(grepUtil.lineNum, lines[index], index))
			}

		}
	}

	if grepUtil.count {
		resultLines = append(resultLines, strconv.Itoa(count))
	}

	for i := 0; i < len(resultLines); i++ {
		fmt.Println(resultLines[i])
	}

	return resultLines
}

func (g *GrepUtil) addToResult(lineNum bool, line string, index int) string {
	if lineNum {
		return strconv.Itoa(index+1) + ":" + line
	}
	return line
}
