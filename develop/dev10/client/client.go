package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

/*
Утилита telnet
Реализовать простейший telnet-клиент.
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123
Требования:
1. Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
2. Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
3. При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

type TelnetUtil struct {
	port    int
	host    string
	timeout time.Duration
}

func main() {
	telnetUtil := &TelnetUtil{}
	var timeout int

	flag.StringVar(&telnetUtil.host, "host", "", "host")
	flag.IntVar(&telnetUtil.port, "port", 0, "port")
	flag.IntVar(&timeout, "timeout", 0, "timeout")
	flag.Parse()

	if telnetUtil.host == "" {
		fmt.Println("не указан хост")
		os.Exit(1)
	}

	if telnetUtil.port == 0 {
		fmt.Println("не указан порт")
		os.Exit(1)
	}
	conn, err := net.DialTimeout(
		"tcp",
		net.JoinHostPort(telnetUtil.host, strconv.Itoa(telnetUtil.port)), time.Second*time.Duration(timeout))

	if err != nil {
		fmt.Printf("ошибка подключения - %s", err)
		os.Exit(1)
	}

	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	consoleReader(conn)
}

func consoleReader(c net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			err = c.Close()
			if err != nil {
				os.Exit(1)
			}
			fmt.Printf("ошибка при чтении - %s", err)
			os.Exit(1)
		}
		_, err = c.Write([]byte(data))
		if err != nil {
			err := c.Close()
			if err != nil {
				os.Exit(1)
			}
			fmt.Printf("ошибка при записи - %s", err)
			os.Exit(1)
		}
	}
}
