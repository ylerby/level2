package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Запуск сервера...")
	server, _ := net.Listen("tcp", ":443")

	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			os.Exit(1)
		}
	}(server)

	for {
		client, err := server.Accept()
		fmt.Println("Подключен клиент")
		scanner := bufio.NewScanner(client)
		var read string
		for scanner.Scan() {
			read = scanner.Text()
			fmt.Print("Считано: ")
			fmt.Println(read)
		}

		err = client.Close()
		if err != nil {
			os.Exit(1)
		}
	}
}
