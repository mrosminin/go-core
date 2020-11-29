package main

import (
	"bufio"
	"fmt"
	"go-core-own/homework-14/pkg/gosearch"
	"log"
	"net"
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
	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handler(conn, gs)
	}
}

func handler(conn net.Conn, gs *gosearch.Service) {
	//conn.SetDeadline(time.Now().Add(time.Second * 10))
	r := bufio.NewReader(conn)
	for {
		query, _, err := r.ReadLine()
		if err != nil {
			return
		}

		//conn.SetDeadline(time.Now().Add(time.Second * 10))
		for i, d := range gs.Engine.Find(string(query)) {
			_, _ = fmt.Fprintf(conn, "%d %v\n", i+1, d)
		}
	}
}
