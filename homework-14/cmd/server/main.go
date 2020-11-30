package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-core-own/homework-14/pkg/gosearch"
	"log"
	"net"
	"strings"
	"time"
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
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Second * 20))

	for {
		// Будем прослушивать все сообщения разделенные \n
		query, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Fprintf(conn, "Ошибка: %v\n", err)
			fmt.Printf("Ошибка: %v\n", err)
			break
		}
		query = strings.TrimSuffix(query, "\n")
		//query = strings.Replace(query, "\n", "", -1)
		//query = strings.Replace(query, "\r", "", -1)

		docs := gs.Engine.Find(query)
		fmt.Printf("По запросу \"%s\" найдено %d документов\n", query, len(docs))
		if len(docs) == 0 {
			fmt.Fprintf(conn, "По запросу \"%s\" ничего не найдено\n", query)
			continue
		}
		jsonData, err := json.Marshal(docs)
		if err != nil {
			fmt.Fprintf(conn, "Ошибка: %v\n", err)
			continue
		}
		fmt.Fprintf(conn, "%s\n", jsonData)
		conn.SetDeadline(time.Now().Add(time.Second * 20))
	}
}
