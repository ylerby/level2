package main

import (
	"dev11/internal/app"
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func main() {
	err := os.Chdir("..")
	if err != nil {
		os.Exit(1)
	}

	file, err := os.Open("config/config.json")
	if err != nil {
		os.Exit(1)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			os.Exit(1)
		}
	}(file)

	res, err := io.ReadAll(file)
	if err != nil {
		os.Exit(1)
	}

	currentConfig := &Config{}

	err = json.Unmarshal(res, currentConfig)
	if err != nil {
		os.Exit(1)
	}

	application := app.New(currentConfig.Port)
	application.Run()
}
