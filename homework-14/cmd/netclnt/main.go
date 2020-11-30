package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	for {
		var query string
		fmt.Print("Введите запрос: ")
		_, err := fmt.Scanln(&query)
		if err != nil {
			fmt.Println("Некорректный ввод: ", err)
			continue
		}
		conn, err := net.Dial("tcp4", "localhost:8080")
		if err != nil {
			log.Fatal(err)
		}
		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(query)); n == 0 || err != nil {
			log.Fatal(err)
		}
		// получем ответ
		data, err := ioutil.ReadAll(conn)
		if err != nil {
			fmt.Println("Ошибка: ", err)
			conn.Close()
			continue
		}
		fmt.Print(string(data))
		fmt.Println()
		conn.Close()
	}
}
