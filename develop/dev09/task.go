package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком
*/

func wget(outputFileName, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("ошибка при получении ответа - %s", err)
	}
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла - %s", err)
	}

	outputWriter := bufio.NewWriter(outputFile)
	_, err = io.Copy(outputWriter, response.Body)
	if err != nil {
		return fmt.Errorf("ошибка при записи - %s", err)
	}

	err = outputFile.Close()
	if err != nil {
		return fmt.Errorf("ошибка при закрытии файла - %s", err)
	}
	return nil
}

func main() {
	var (
		url            = "https://metanit.com/go/tutorial/9.6.php"
		outputFileName = "output.txt"
	)

	err := wget(outputFileName, url)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}
