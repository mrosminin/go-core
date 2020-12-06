package netsrv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-core-own/homework-14/pkg/engine"
	"net"
	"strings"
	"time"
)

func Serve(engine *engine.Service, network, addr string) error {
	// регистрация сетевой службы
	listener, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	fmt.Println("Сервер слушает порт 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn, engine)
	}
}

// обработка подключения
func handleConnection(conn net.Conn, engine *engine.Service) {
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
		// для mac этого достаточно, но не уверен насчет других (менее продвинутых) ОС
		// query = strings.TrimSuffix(query, "\n")
		query = strings.Replace(query, "\n", "", -1)
		query = strings.Replace(query, "\r", "", -1)

		docs := engine.Find(query)
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
