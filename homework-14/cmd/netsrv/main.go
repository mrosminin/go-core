package main

import (
	"fmt"
	"go-core-own/homework-14/pkg/gosearch"
	"log"
	"net"
	"strings"
)

var sites = []string{
	"https://go.dev",
	"http://www.transflow.ru",
	"https://www.newsru.com",
	"https://www.gov-murman.ru/",
	"https://www.anekdot.ru/",
	"https://en.wikipedia.org/wiki/Main_Page",
	"https://www.prj-exp.ru/gost-34",
	"https://habr.com/ru/",
}
var depth = 2

func main() {
	gs := gosearch.New(sites, depth)
	gs.Init()

	// регистрация сетевой службы
	listener, err := net.Listen("tcp", ":8080")
	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Сервер слушает порт 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn, gs)
	}
}

// обработка подключения
func handleConnection(conn net.Conn, gs *gosearch.Service) {
	// Читаем из соединение через буфер, т.к. ReadString('\n'), ReadLine, ReadAll не работают - нет EOF
	input := make([]byte, 1024*4)
	n, err := conn.Read(input)
	if n == 0 || err != nil {
		fmt.Println("Read error:", err)
		return
	}
	query := string(input[0:n])
	query = strings.Replace(query, "\n", "", -1)
	query = strings.Replace(query, "\r", "", -1)

	docs := gs.Engine.Find(query)
	for i, d := range docs {
		_, _ = fmt.Fprintf(conn, "%d %v\n", i+1, d)
	}
	fmt.Printf("По запросу \"%s\" найдено %d документов\n", query, len(docs))
	conn.Close()
}
